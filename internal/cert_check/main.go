package certcheck

import (
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	types "gitlab.com/slavditore/check-domain/v2/internal/types"
)

func CheckDomains(domains []types.DomainStatus) []types.DomainStatus {
	var checkResults []types.DomainStatus
	var domainCheckThreadPool sync.WaitGroup
	for i := 0; i < len(domains); i++ {
		domainCheckThreadPool.Add(1)
		go func(wg *sync.WaitGroup, domain *types.DomainStatus) {
			defer wg.Done()
			connection, err := tls.Dial("tcp", fmt.Sprintf("%s:443", domain.Name), nil)
			if err != nil {
				domain.Note = "Domain does not support TLS"
			}
			err = connection.VerifyHostname(domain.Name)
			if err != nil {
				domain.Note = "Hostname does not match with certificate"
			}
			expiry := connection.ConnectionState().PeerCertificates[0].NotAfter
			domain.Issuer = connection.ConnectionState().PeerCertificates[0].Issuer.String()
			domain.ValidBefore = expiry.Format(time.RFC850)
		}(&domainCheckThreadPool, &domains[i])
	}
	domainCheckThreadPool.Wait()
	checkResults = append(checkResults, domains...)
	return checkResults
}
