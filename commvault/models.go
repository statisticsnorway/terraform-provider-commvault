package commvault

type CommvaultConfig struct {
	BaseURL  string
	Username string
	Password string
}

type loginResponse struct {
	Token string `xml:"token,attr"`
}

type createResponse struct {
	Response struct {
		ErrorCode int `json:"errorCode"`
		Entity    struct {
			ClientID   int    `json:"clientId"`
			ClientName string `json:"clientName"`
		} `json:"entity"`
	} `json:"response"`
}
