package types

type DomainStatus struct {
	Group       string
	Name        string
	Issuer      string
	ValidBefore string
	DaysLeft    uint
	Note        string
}

type ConfigurationVariables struct {
	DaysToExpration uint
}
