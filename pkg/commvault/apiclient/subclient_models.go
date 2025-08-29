package apiclient

type SubclientGetResponse struct {
	FilterQueryCount    int                 `json:"filterQueryCount"`
	SubClientProperties []SubClientProperty `json:"subClientProperties"`
}

type SubClientProperty struct {
	UseLocalContent        bool            `json:"useLocalContent"`
	UseLocalArchivalRules  bool            `json:"useLocalArchivalRules"`
	PlanEntity             interface{}     `json:"planEntity"`
	CassandraProps         interface{}     `json:"cassandraProps"`
	TaPlanEntity           interface{}     `json:"taPlanEntity"`
	SplunkProps            interface{}     `json:"splunkProps"`
	BigDataProps           interface{}     `json:"bigDataProps"`
	ProxyClient            interface{}     `json:"proxyClient"`
	DcPlanEntity           interface{}     `json:"dcPlanEntity"`
	SubClientEntity        SubClientEntity `json:"subClientEntity"`
	DfsSubclientProp       interface{}     `json:"dfsSubclientProp"`
	Region                 interface{}     `json:"region"`
	AnalyticsSubclientProp interface{}     `json:"analyticsSubclientProp"`
}

type SubClientEntity struct {
	ClientName    string     `json:"clientName"`
	InstanceName  string     `json:"instanceName"`
	DisplayName   string     `json:"displayName"`
	BackupsetId   int        `json:"backupsetId"`
	Type          int        `json:"_type_"`
	CommCellName  string     `json:"commCellName"`
	InstanceId    int        `json:"instanceId"`
	CsGUID        string     `json:"csGUID"`
	SubclientGUID string     `json:"subclientGUID"`
	SubclientId   int        `json:"subclientId"`
	ClientId      int        `json:"clientId"`
	AppName       string     `json:"appName"`
	BackupsetName string     `json:"backupsetName"`
	ApplicationId int        `json:"applicationId"`
	SubclientName string     `json:"subclientName"`
	EntityInfo    EntityInfo `json:"entityInfo"`
}

type EntityInfo struct {
	CompanyId       int    `json:"companyId"`
	CompanyName     string `json:"companyName"`
	MultiCommcellId int    `json:"multiCommcellId"`
}

// -------------------------------------------------------

type SubclientCreateOrUpdateRequestAndResponse struct {
	SubclientProperties SubclientProperties `json:"subClientProperties"`
}

type SubclientProperties struct {
	SubclientEntity        SubclientUpdateRequestClientEntity `json:"subClientEntity"`
	UseLocalContent        bool                               `json:"useLocalContent"`
	CloudAppsSubClientProp CloudAppsSubClientProp             `json:"cloudAppsSubClientProp"`
}

type SubclientUpdateRequestClientEntity struct {
	SubclientID   int `json:"subclientId,omitempty"`
	ClientName    int `json:"clientName,omitempty"`
	AppName       int `json:"appName,omitempty"`
	SubclientName int `json:"subclientName,omitempty"`
	BackupsetName int `json:"backupsetName,omitempty"`
}

type CloudAppsSubClientProp struct {
	InstanceType           int                    `json:"instanceType"`
	ObjectStorageSubclient ObjectStorageSubclient `json:"objectStorageSubclient"`
}

type ObjectStorageSubclient struct {
	ContentOperationType int       `json:"contentOperationType"`
	Content              []Content `json:"content"`
}

type Content struct {
	GCPContent GCPContent `json:"gcpContent"`
}

type GCPContent struct {
	BucketName  string `json:"bucketName"`
	ProjectName string `json:"projectName"`
}

// --------------------------------------------------
type SubclientDeleteResponse struct {
	Processinginstructioninfo Processinginstructioninfo `json:"processinginstructioninfo"`
	Response                  []struct {
		ErrorCode int `json:"errorCode"`
	} `json:"response"`
}
