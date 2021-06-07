package output_plain

import (
	"log"

	"gitlab.com/slavditore/check-domain/v2/internal/types"
)

func PlainOutput(overview *[]types.DomainStatus) {
	for i := 0; i < len(*overview); i++ {
		log.Printf("hostname=%s, issuer=%s, valid_until=%s", (*overview)[i].Name, (*overview)[i].Issuer, (*overview)[i].ValidBefore)
	}
}
