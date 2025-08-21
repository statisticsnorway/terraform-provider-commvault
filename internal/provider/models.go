package provider

type createResponse struct {
	Response struct {
		ErrorCode int `json:"errorCode"`
		Entity    struct {
			ClientID   int    `json:"clientId"`
			ClientName string `json:"clientName"`
		} `json:"entity"`
	} `json:"response"`
}
