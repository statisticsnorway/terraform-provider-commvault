package apiclient

type LoginResponseCustomerManaged struct {
	AliasName           string `json:"aliasName"`
	UserGUID            string `json:"userGUID"`
	LoginAttempts       int    `json:"loginAttempts"`
	RemainingLockTime   int    `json:"remainingLockTime"`
	SmtpAddress         string `json:"smtpAddress"`
	UserName            string `json:"userName"`
	ProviderType        int    `json:"providerType"`
	Ccn                 int    `json:"ccn"`
	Token               string `json:"token"`
	Capability          int    `json:"capability"`
	ForcePasswordChange bool   `json:"forcePasswordChange"`
	IsAccountLocked     bool   `json:"isAccountLocked"`
	OwnerOrganization   struct {
		GUID               string `json:"GUID"`
		ProviderId         int    `json:"providerId"`
		ProviderDomainName string `json:"providerDomainName"`
	} `json:"ownerOrganization"`
	AdditionalResp struct {
		NameValues []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"nameValues"`
	} `json:"additionalResp"`
	ProviderOrganization struct {
		GUID               string `json:"GUID"`
		ProviderId         int    `json:"providerId"`
		ProviderDomainName string `json:"providerDomainName"`
	} `json:"providerOrganization"`
	ErrList []struct {
		Company struct {
			ProviderId         int    `json:"providerId"`
			ProviderDomainName string `json:"providerDomainName"`
		} `json:"company"`
	} `json:"errList"`
}
