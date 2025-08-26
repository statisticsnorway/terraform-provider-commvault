package apiclient

type ClientGetResponse struct {
	ClientProperties []ClientPropertyItem `json:"clientProperties"`
}

type Overview struct {
	VmsCount            int `json:"vmsCount"`
	UnixServersCount    int `json:"unixServersCount"`
	NasFilersCount      int `json:"nasFilersCount"`
	WindowsServersCount int `json:"windowsServersCount"`
	FsCount             int `json:"fsCount"`
	DatabasesCount      int `json:"databasesCount"`
	CloudAppsCount      int `json:"cloudAppsCount"`
}

type ClientPropertyItem struct {
	ClientProps ClientProps             `json:"clientProps"`
	Client      ClientGetResponseClient `json:"client"`
}

type ClientProps struct {
	IsVirtualServerDiscoveredClient bool                   `json:"isVirtualServerDiscoveredClient"`
	ClientStatus                    int                    `json:"clientStatus"`
	IsMA                            bool                   `json:"isMA"`
	IsIndexingV2                    bool                   `json:"isIndexingV2"`
	IsInfrastructure                bool                   `json:"isInfrastructure"`
	SpecialClientFlags              int                    `json:"specialClientFlags"`
	IsCommCellPackageAbsent         bool                   `json:"isCommCellPackageAbsent"`
	ClusterType                     int                    `json:"clusterType"`
	IsIndexingV2VSA                 bool                   `json:"isIndexingV2VSA"`
	GrcMigratedClient               bool                   `json:"grcMigratedClient"`
	IsDeletedClient                 bool                   `json:"IsDeletedClient"`
	IsUserCentricClient             bool                   `json:"isUserCentricClient"`
	IsVirtualClient                 bool                   `json:"IsVirtualClient"`
	VendorType                      int                    `json:"vendorType,omitempty"`
	Attributes                      []Attribute            `json:"attributes,omitempty"`
	Plan                            *ClientGetResponsePlan `json:"plan,omitempty"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ClientGetResponsePlan struct {
	Type     int    `json:"_type_"`
	PlanName string `json:"planName"`
	PlanId   int    `json:"planId"`
}

type ClientGetResponseClient struct {
	OsInfo       OSInfo       `json:"osInfo"`
	IdaList      []IdaItem    `json:"idaList"`
	ClientEntity ClientEntity `json:"clientEntity"`
	VersionInfo  VersionInfo  `json:"versionInfo"`
}

type OSInfo struct {
	Type          string        `json:"Type"`
	SubType       string        `json:"SubType"`
	OsId          int           `json:"osId"`
	OsDisplayInfo OSDisplayInfo `json:"OsDisplayInfo"`
}

type OSDisplayInfo struct {
	ProcessorType string `json:"ProcessorType,omitempty"`
	OSName        string `json:"OSName"`
}

type IdaItem struct {
	Summary   *Summary  `json:"summary,omitempty"`
	IdaEntity IdaEntity `json:"idaEntity"`
}

type Summary struct {
	Size             int          `json:"size"`
	IsSLAMet         bool         `json:"isSLAMet"`
	LastBackupTime   int64        `json:"lastBackupTime,omitempty"`
	DefaultBackupset BackupsetRef `json:"defaultBackupset"`
}

type IdaEntity struct {
	AppName       string `json:"appName"`
	Type          int    `json:"_type_"`
	ApplicationId int    `json:"applicationId"`
}

type BackupsetRef struct {
	PlanEntity      interface{}     `json:"planEntity"`
	BackupSetEntity BackupSetEntity `json:"backupSetEntity"`
}

type BackupSetEntity struct {
	BackupsetId   int    `json:"backupsetId"`
	BackupsetName string `json:"backupsetName"`
}

type ClientEntity struct {
	HostName    string `json:"hostName"`
	ClientId    int    `json:"clientId"`
	ClientName  string `json:"clientName"`
	DisplayName string `json:"displayName"`
	ClientGUID  string `json:"clientGUID"`
}
type VersionInfo struct {
	UpdateStatus  int           `json:"UpdateStatus"`
	Version       string        `json:"version"`
	GalaxyRelease GalaxyRelease `json:"GalaxyRelease"`
}
type GalaxyRelease struct {
	ReleaseID int `json:"ReleaseID"`
}

// ---------------------------------------------

type ClientCreateRequest struct {
	ClientInfo ClientInfo `json:"clientInfo"`
	Entity     Entity     `json:"entity"`
}

type ClientInfo struct {
	ClientType               int                      `json:"clientType"`
	CloudConnectorProperties CloudConnectorProperties `json:"cloudClonnectorProperties"`
	Plan                     ClientCreateRequestPlan  `json:"plan"`
}

type CloudConnectorProperties struct {
	Instance     Instance `json:"instance"`
	InstanceType string   `json:"instanceType"`
}

type Instance struct {
	CloudAppsInstance CloudAppsInstance `json:"cloudAppsInstance"`
	Instance          InstanceDetails   `json:"instance"`
}

type CloudAppsInstance struct {
	CredentialType          string                 `json:"credentialType"`
	GeneralCloudProperties  GeneralCloudProperties `json:"generalCloudProperties"`
	GoogleCloudInstance     GoogleCloudInstance    `json:"googleCloudInstance"`
	InstanceType            string                 `json:"instanceType"`
	InstanceTypeDisplayName string                 `json:"instanceTypeDisplayName"`
}

type GeneralCloudProperties struct {
	Credentials           Credentials    `json:"credentials"`
	MemberServers         []MemberServer `json:"memberServers"`
	NumberOfBackupStreams int            `json:"numberOfBackupStreams"`
}

type Credentials struct {
	CredentialID int64 `json:"credentialId"`
}

type MemberServer struct {
	Client ClientCreateRequestClient `json:"client"`
}

type ClientCreateRequestClient struct {
	Type     int   `json:"_type_"`
	ClientID int64 `json:"clientId"`
}

type GoogleCloudInstance struct {
	GCPProjectID string `json:"GCPProjectId"`
	ServerName   string `json:"serverName"`
}

type InstanceDetails struct {
	ApplicationID int    `json:"applicationId"`
	CommCellID    int    `json:"commCellId"`
	InstanceName  string `json:"instanceName"`
}

type ClientCreateRequestPlan struct {
	PlanID int64 `json:"planId"`
}

type Entity struct {
	ClientName string `json:"clientName"`
}

// --------------------------------------------------

type ClientCreateResponse struct {
	Response struct {
		ErrorCode int `json:"errorCode"`
		Entity    struct {
			ClientID   int    `json:"clientId"`
			ClientName string `json:"clientName"`
		} `json:"entity"`
	} `json:"response"`
}
