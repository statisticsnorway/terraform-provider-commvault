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
	IsVirtualServerDiscoveredClient          bool   `json:"isVirtualServerDiscoveredClient"`
	ClientStatus                             int    `json:"clientStatus"`
	IsMA                                     bool   `json:"isMA"`
	IsIndexingV2                             bool   `json:"isIndexingV2"`
	ISCSIPort                                bool   `json:"iSCSIPort"`
	IsInfrastructure                         bool   `json:"isInfrastructure"`
	SpecialClientFlags                       int    `json:"specialClientFlags"`
	IsCommCellPackageAbsent                  bool   `json:"isCommCellPackageAbsent"`
	ClusterType                              int    `json:"clusterType"`
	IsIndexingV2VSA                          bool   `json:"isIndexingV2VSA"`
	GrcMigratedClient                        bool   `json:"grcMigratedClient"`
	IsDeletedClient                          bool   `json:"IsDeletedClient"`
	IsUserCentricClient                      bool   `json:"isUserCentricClient"`
	IsVirtualClient                          bool   `json:"IsVirtualClient"`
	VendorType                               int    `json:"vendorType,omitempty"`
	SLAStatus                                int    `json:"slaStatus"`
	CipherType                               int    `json:"CipherType"`
	RetryFrequencyInSecOnNetworkError        int    `json:"RetryFrequencyInSecOnNetworkError"`
	RestoreAccessFlag                        int    `json:"RestoreAccessFlag"`
	IsRemoteCacheEnabled                     bool   `json:"isRemoteCacheEnabled"`
	SnapBackupMAInstalled                    bool   `json:"SnapBackupMAInstalled"`
	PathToExchangeMiningTool                 string `json:"PathToExchangeMiningTool"`
	RestoreAccessSynFull                     bool   `json:"RestoreAccessSynFull"`
	ClientInterface                          string `json:"ClientInterface"`
	OptimizeDataForSearch                    bool   `json:"OptimizeDataForSearch"`
	SLACategoryDescription                   string `json:"slaCategoryDescription"`
	RetireClientPhase                        int    `json:"retireClientPhase"`
	ContentIndexEngineInstalled              bool   `json:"ContentIndexEngineInstalled"`
	RetryCountOnNetworkError                 int    `json:"RetryCountOnNetworkError"`
	DevsnapServerType                        string `json:"devsnapServerType"`
	IsPersonalWorkstation                    bool   `json:"isPersonalWorkstation"`
	IsIndexingV2NAS                          bool   `json:"isIndexingV2NAS"`
	EnableSnapBackups                        bool   `json:"EnableSnapBackups"`
	StorePatchLocally                        bool   `json:"StorePatchLocally"`
	IsFastSearchSupported                    bool   `json:"isFastSearchSupported"`
	SLACategory                              int    `json:"slaCategory"`
	CopyManagementLicense                    bool   `json:"copyManagementLicense"`
	IsDescriptionInHTMLFormat                bool   `json:"isDescriptionInHTMLFormat"`
	EnableRetryOnNetworkError                bool   `json:"EnableRetryOnNetworkError"`
	QueueJobsIfOtherConflictingJobsAreActive bool   `json:"queueJobsIfOtherConflictingJobsAreActive"`
	JobPriority                              int    `json:"JobPriority"`
	OverrideGlobalDesktopGuiProperties       bool   `json:"overrideGlobalDesktopGuiProperties"`
	IsPseudoDBCluster                        bool   `json:"isPseudoDBCluster"`
	CDRLogFileLocation                       string `json:"CDRLogFileLocation"`
	IsSecureInstallEnabled                   bool   `json:"isSecureInstallEnabled"`
	IsFarmClient                             bool   `json:"isFarmClient"`
	AutoUpdateLocation                       string `json:"AutoUpdateLocation"`
	JobResultsDiskCapThreshold               int    `json:"JobResultsDiskCapThreshold"`
	MaxStreams                               int    `json:"maxStreams"`
	ForceClientSideDownload                  int    `json:"forceClientSideDownload"`
	Is1TouchClient                           bool   `json:"is1TouchClient"`
	IsMARoleSet                              bool   `json:"isMARoleSet"`
	AltCachePartitionForQSnap                string `json:"AltCachePartitionForQSnap"`
	CacheSource                              int    `json:"cacheSource"`
	IsAdvancedFSPackageInstalled             bool   `json:"isAdvancedFSPackageInstalled"`
	IndexingV2EnabledMask                    int    `json:"indexingV2EnabledMask"`
	EncryptKeyLength                         int    `json:"EncryptKeyLength"`
	IsQsnapUNIX                              bool   `json:"IsQsnapUNIX"`
	Dm2CacheDiskCapThreshold                 int    `json:"dm2CacheDiskCapThreshold"`
	DevsnapTargetPort                        int    `json:"devsnapTargetPort"`
	EnableContentIndexing                    bool   `json:"EnableContentIndexing"`
	RecallService                            string `json:"recallService"`
	CenteraResourceFilename                  string `json:"CenteraResourceFilename"`
	IsQsnapInstalled                         int    `json:"isQsnapInstalled"`
	Dm2CacheRetentionDays                    int    `json:"dm2CacheRetentionDays"`
	OverrideGlobalEncryptionSettings         int    `json:"overrideGlobalEncryptionSettings"`
	IsDCInstalled                            bool   `json:"IsDCInstalled"`
	EncryptionSettings                       int    `json:"encryptionSettings"`
	IsExchangeOnePassClient                  bool   `json:"IsExchangeOnePassClient"`
	IsExchangePseudoClient                   bool   `json:"isExchangePseudoClient"`
	DirectMediaAccessFlag                    int    `json:"DirectMediaAccessFlag"`
	IsMAPackageInstalled                     bool   `json:"isMAPackageInstalled"`
	JobResultsThresholdMB                    int    `json:"JobResultsThresholdMB"`
	JobResultsRetentionDays                  int    `json:"JobResultsRetentionDays"`
	BlockLevelCacheDir                       string `json:"BlockLevelCacheDir"`
	IsCommServer                             bool   `json:"IsCommServer"`
	EnableCollectDelegateInfo                bool   `json:"EnableCollectDelegateInfo"`
	FileLevelAnalyticsLicense                bool   `json:"fileLevelAnalyticsLicense"`
	Dm2CacheThresholdMB                      int    `json:"dm2CacheThresholdMB"`
	DevsnapInitiatorIQN                      string `json:"devsnapInitiatorIQN"`
	EnableAccessControl                      bool   `json:"enableAccessControl"`
	ClientNoPhysicalMachine                  bool   `json:"ClientNoPhysicalMachine"`
	RestoreAccessPubKey                      bool   `json:"RestoreAccessPubKey"`

	Attributes []Attribute            `json:"attributes,omitempty"`
	Plan       *ClientGetResponsePlan `json:"plan,omitempty"`

	NetworkReadiness                 *NetworkReadiness                 `json:"networkReadiness,omitempty"`
	DlpPropertise                    *DlpPropertise                    `json:"dlpPropertise,omitempty"`
	FirewallConfiguration            *FirewallConfiguration            `json:"firewallConfiguration,omitempty"`
	ClientRegionInfo                 *ClientRegionInfo                 `json:"clientRegionInfo,omitempty"`
	DeDuplicationProperties          *DeDuplicationProperties          `json:"deDuplicationProperties,omitempty"`
	SpWebServerProperties            *SpWebServerProperties            `json:"spWebServerProperties,omitempty"`
	Dm2CacheDir                      *Dm2CacheDir                      `json:"dm2CacheDir,omitempty"`
	RightManagementServiceProperties *RightManagementServiceProperties `json:"rightManagementServiceProperties,omitempty"`
	Company                          *Company                          `json:"company,omitempty"`
	NetworkThrottle                  *NetworkThrottle                  `json:"networkThrottle,omitempty"`
	SecurityAssociations             *SecurityAssociations             `json:"securityAssociations,omitempty"`
	ClientSecurity                   *ClientSecurity                   `json:"clientSecurity,omitempty"`
	ActivityControl                  *ActivityControl                  `json:"activityControl,omitempty"`
	AdvancedClientProperties         *AdvancedClientProperties         `json:"advancedClientProperties,omitempty"`
	ClientActivityControl            *ClientActivityControl            `json:"clientActivityControl,omitempty"`
	JobThrottleSettings              *JobThrottleSettings              `json:"jobThrottleSettings,omitempty"`
	EmailWebAppProperties            *EmailWebAppProperties            `json:"emailWebAppProperties,omitempty"`
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
	Size                   int          `json:"size"`
	IsSLAMet               bool         `json:"isSLAMet"`
	LastBackupTime         int64        `json:"lastBackupTime,omitempty"`
	DefaultBackupset       BackupsetRef `json:"defaultBackupset"`
	SLAStatus              int          `json:"slaStatus"`
	SLACategoryDescription string       `json:"slaCategoryDescription"`
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
	HostName     string `json:"hostName"`
	ClientId     int    `json:"clientId"`
	ClientName   string `json:"clientName"`
	DisplayName  string `json:"displayName"`
	CommCellName string `json:"commCellName"`
	ClientGUID   string `json:"clientGUID"`
}
type VersionInfo struct {
	UpdateStatus  int           `json:"UpdateStatus"`
	Version       string        `json:"version"`
	GalaxyRelease GalaxyRelease `json:"GalaxyRelease"`
}
type GalaxyRelease struct {
	ReleaseID     int    `json:"ReleaseID"`
	ReleaseString string `json:"ReleaseString"`
}

// ---------------------------------------------

type ClientGetByIdResponse struct {
	Processinginstructioninfo Processinginstructioninfo `json:"processinginstructioninfo"`
	ClientProperties          []ClientProperties        `json:"clientProperties"`
}

type Processinginstructioninfo struct {
	Attributes []Attribute `json:"attributes"`
}
type ClusterClientProperties struct {
	ConfigureClusterClient bool `json:"configureClusterClient"`
	ShowAllAgents          bool `json:"showAllAgents"`
}
type VirtualServerCredentialinfo struct {
}
type Azure struct {
	Credentials Credentials `json:"credentials"`
}
type AzureResourceManager struct {
	UseManagedIdentity bool        `json:"useManagedIdentity"`
	Credentials        Credentials `json:"credentials"`
}
type AssociatedClients struct {
	UsePredefinedAccessNodes bool `json:"usePredefinedAccessNodes"`
}
type AmazonInstanceInfo struct {
	UseIamRole bool `json:"useIamRole"`
}
type VirtualServerInstanceInfo struct {
	VsInstanceType              int                         `json:"vsInstanceType"`
	VirtualServerCredentialinfo VirtualServerCredentialinfo `json:"virtualServerCredentialinfo"`
	Azure                       Azure                       `json:"azure"`
	AzureResourceManager        AzureResourceManager        `json:"azureResourceManager"`
	AssociatedClients           AssociatedClients           `json:"associatedClients"`
	AmazonInstanceInfo          AmazonInstanceInfo          `json:"amazonInstanceInfo"`
}
type VMBackupInfo struct {
}
type VirtualServerClientProperties struct {
	ManagementVersion         string                    `json:"managementVersion"`
	VirtualServerInstanceInfo VirtualServerInstanceInfo `json:"virtualServerInstanceInfo"`
	VMBackupInfo              VMBackupInfo              `json:"vmBackupInfo"`
}
type AccessNodes struct {
}

type CloudClonnectorProperties struct {
	Instance Instance `json:"instance"`
}
type AdForestPseudoClientProperties struct {
}
type UserProfile struct {
	UserName string `json:"userName"`
}
type IbmiInstallOptions struct {
	JobProirity       int         `json:"jobProirity"`
	JobQueue          string      `json:"jobQueue"`
	CreateJobQueue    bool        `json:"createJobQueue"`
	SubsysDescription string      `json:"subsysDescription"`
	DataFolder        string      `json:"dataFolder"`
	RunPriority       int         `json:"runPriority"`
	UserProfile       UserProfile `json:"userProfile"`
}
type ImpersonateUser struct {
}
type OutlookAddinStoreInfo struct {
}
type EdgeDrivePseudoClientProperties struct {
	SystemDriveType       int                   `json:"systemDriveType"`
	ImpersonateUser       ImpersonateUser       `json:"impersonateUser"`
	OutlookAddinStoreInfo OutlookAddinStoreInfo `json:"outlookAddinStoreInfo"`
}
type ProxyClients struct {
}
type UserAccount struct {
	UserName string `json:"userName"`
}
type OpenVMSProperties struct {
	CvdPort      int          `json:"cvdPort"`
	ProxyClients ProxyClients `json:"proxyClients"`
	UserAccount  UserAccount  `json:"userAccount"`
}
type DbThinClientProperties struct {
	ProxyClients ProxyClients `json:"proxyClients"`
}
type ServiceMeshClientProperties struct {
}
type OverrideHigherLevelSettings struct {
	UserAccount UserAccount `json:"userAccount"`
}
type SQLInstanceCredentials struct {
}
type AwsSQLInfo struct {
	SQLInstanceCredentials SQLInstanceCredentials `json:"sqlInstanceCredentials"`
}
type AzureSQLInstanceCredentials struct {
}
type AzureInfo struct {
	AzureSQLInstanceCredentials AzureSQLInstanceCredentials `json:"azureSqlInstanceCredentials"`
}
type MSSQLCredentialinfo struct {
}
type SQL61Prop struct {
	OverrideHigherLevelSettings OverrideHigherLevelSettings `json:"overrideHigherLevelSettings"`
	AwsSQLInfo                  AwsSQLInfo                  `json:"awsSqlInfo"`
	AzureInfo                   AzureInfo                   `json:"azureInfo"`
	MSSQLCredentialinfo         MSSQLCredentialinfo         `json:"MSSQLCredentialinfo"`
}
type IdaInfo struct {
	SQL61Prop SQL61Prop `json:"sql61Prop"`
}
type PseudoClientInfo struct {
	ClientType                      int                             `json:"clientType"`
	VirtualServerClientProperties   VirtualServerClientProperties   `json:"virtualServerClientProperties"`
	CloudClonnectorProperties       CloudClonnectorProperties       `json:"cloudClonnectorProperties"`
	AdForestPseudoClientProperties  AdForestPseudoClientProperties  `json:"adForestPseudoClientProperties"`
	IbmiInstallOptions              IbmiInstallOptions              `json:"ibmiInstallOptions"`
	EdgeDrivePseudoClientProperties EdgeDrivePseudoClientProperties `json:"edgeDrivePseudoClientProperties"`
	OpenVMSProperties               OpenVMSProperties               `json:"openVMSProperties"`
	DbThinClientProperties          DbThinClientProperties          `json:"dbThinClientProperties"`
	ServiceMeshClientProperties     ServiceMeshClientProperties     `json:"serviceMeshClientProperties"`
	IdaInfo                         IdaInfo                         `json:"idaInfo"`
}
type ClientConfiguration struct {
	IsMergedExchangePackageInstalled bool `json:"isMergedExchangePackageInstalled"`
}
type LastOnlineTime struct {
	Type int `json:"_type_"`
}
type LastUpdatedTime struct {
	Type int `json:"_type_"`
}
type OfflineTime struct {
	Type int `json:"_type_"`
}
type NetworkReadiness struct {
	Status          int             `json:"status"`
	LastOnlineTime  LastOnlineTime  `json:"lastOnlineTime"`
	LastUpdatedTime LastUpdatedTime `json:"lastUpdatedTime"`
	OfflineTime     OfflineTime     `json:"offlineTime"`
}
type DlpRMProperties struct {
	EnableRmDLP      bool `json:"enableRmDLP"`
	DlpRmNow         bool `json:"dlpRmNow"`
	DlpRmOfflineDays int  `json:"dlpRmOfflineDays"`
}
type DlpPropertise struct {
	DlpScanIntervalMins          int             `json:"dlpScanIntervalMins"`
	DlpStolen                    bool            `json:"dlpStolen"`
	DlpEnableAutomaticDecryption bool            `json:"dlpEnableAutomaticDecryption"`
	DlpUnlockMethod              int             `json:"dlpUnlockMethod"`
	EnableDLP                    bool            `json:"enableDLP"`
	DlpEnableClientKeys          bool            `json:"dlpEnableClientKeys"`
	DlpMinFileAgeMins            int             `json:"dlpMinFileAgeMins"`
	DlpRMProperties              DlpRMProperties `json:"dlpRMProperties"`
}
type ReachMeVia struct {
	ReachMeViaSelected bool `json:"reachMeViaSelected"`
}
type ProxyEntities struct {
	ClientGroupName string `json:"clientGroupName"`
	ClientGroupID   int    `json:"clientGroupId"`
}
type ProxyEntity struct {
	DisplayName string `json:"displayName"`
	ClientID    int    `json:"clientId"`
	ClientName  string `json:"clientName"`
}
type FirewallTrivialDetailsCS struct {
	NetworkType    int         `json:"networkType"`
	ConnectionType int         `json:"connectionType"`
	ProxyEntity    ProxyEntity `json:"proxyEntity"`
}
type FirewallTrivialDetailsMA struct {
	NetworkType    int         `json:"networkType"`
	ConnectionType int         `json:"connectionType"`
	ProxyEntity    ProxyEntity `json:"proxyEntity"`
}
type FwTrivialConfig struct {
	FirewallTrivialDetailsCS FirewallTrivialDetailsCS `json:"firewallTrivialDetailsCS"`
	FirewallTrivialDetailsMA FirewallTrivialDetailsMA `json:"firewallTrivialDetailsMA"`
}
type Tppm struct {
}
type FirewallOptions struct {
	IsRoamingClient      bool   `json:"isRoamingClient"`
	ExtendedProperties   string `json:"extendedProperties"`
	TunnelconnectionPort int    `json:"tunnelconnectionPort"`
	ForeceSSL            bool   `json:"foreceSSL"`
	TunnelInitSeconds    int    `json:"tunnelInitSeconds"`
	Lockdown             bool   `json:"lockdown"`
	BindOpenPortsOnly    bool   `json:"bindOpenPortsOnly"`
	IsDMZ                bool   `json:"isDMZ"`
	KeepAliveSeconds     int    `json:"keepAliveSeconds"`
	Tppm                 []Tppm `json:"tppm"`
}
type FirewallConfiguration struct {
	ConfigureFirewallSettings bool            `json:"configureFirewallSettings"`
	IsTrivialConfig           bool            `json:"isTrivialConfig"`
	ReachMeVia                ReachMeVia      `json:"reachMeVia"`
	ProxyEntities             []ProxyEntities `json:"proxyEntities"`
	FwTrivialConfig           FwTrivialConfig `json:"fwTrivialConfig"`
	FirewallOptions           FirewallOptions `json:"firewallOptions"`
}
type ClientRegionInfo struct {
}
type DeDuplicationProperties struct {
	PerformClientSideDeduplication bool `json:"performClientSideDeduplication"`
	EnableHighLatencyOptimization  bool `json:"enableHighLatencyOptimization"`
	CacheBufferSize                int  `json:"cacheBufferSize"`
	MaxCacheDb                     int  `json:"maxCacheDb"`
	ClientSideDeduplication        int  `json:"clientSideDeduplication"`
	EnableVariableContentAlignment bool `json:"enableVariableContentAlignment"`
	EnableClientSideDiskCache      bool `json:"enableClientSideDiskCache"`
}
type SpWebServerUserPassword struct {
}
type SpWebServerProperties struct {
	EnableOnlineSearch      bool                    `json:"enableOnlineSearch"`
	SpWebServerUserPassword SpWebServerUserPassword `json:"spWebServerUserPassword"`
}
type Dm2CacheDir struct {
	Path        string      `json:"path"`
	UserAccount UserAccount `json:"userAccount"`
}
type RmsCredentials struct {
	UserName string `json:"userName"`
}
type RightManagementServiceProperties struct {
	SMTPAddressOfTheRMSSuperUser            string         `json:"SMTPAddressOfTheRMSSuperUser"`
	DecryptRMSDocumentDuringContentIndexing bool           `json:"decryptRMSDocumentDuringContentIndexing"`
	SMTPAddressOfRMSSuperUser               string         `json:"smtpAddressOfRMSSuperUser"`
	RmsCredentials                          RmsCredentials `json:"rmsCredentials"`
}
type ShortName struct {
	DomainName string `json:"domainName"`
	ID         int    `json:"id"`
}
type Company struct {
	ConnectName string    `json:"connectName"`
	ShortName   ShortName `json:"shortName"`
}
type Throttle struct {
}
type NetworkThrottle struct {
	EnableThrottle bool     `json:"enableThrottle"`
	Throttle       Throttle `json:"throttle"`
}
type OwnerAssociations struct {
}
type TagWithCompany struct {
	ProviderID         int    `json:"providerId"`
	ProviderDomainName string `json:"providerDomainName"`
}
type SecurityAssociations struct {
	OwnerAssociations OwnerAssociations `json:"ownerAssociations"`
	TagWithCompany    TagWithCompany    `json:"tagWithCompany"`
}
type AssociatedUserGroups struct {
	UserGroupID   int    `json:"userGroupId"`
	Type          int    `json:"_type_"`
	UserGroupName string `json:"userGroupName"`
}
type ExpirationTime struct {
	Time int `json:"time"`
}
type PasskeySettings struct {
	EnableAuthorizeForRestore bool           `json:"enableAuthorizeForRestore"`
	ExpirationTime            ExpirationTime `json:"expirationTime"`
}
type AdvancedPrivacySettings struct {
	IsPasskeyFeatureEnabled bool            `json:"isPasskeyFeatureEnabled"`
	AuthType                int             `json:"authType"`
	IsPrivacyFeatureEnabled bool            `json:"isPrivacyFeatureEnabled"`
	PasskeySettings         PasskeySettings `json:"passkeySettings"`
}
type OwnerCapabilities struct {
}
type ClientSecurity struct {
	EnableDataSecurity      bool                    `json:"enableDataSecurity"`
	AssociatedUserGroups    []AssociatedUserGroups  `json:"associatedUserGroups"`
	AdvancedPrivacySettings AdvancedPrivacySettings `json:"advancedPrivacySettings"`
	OwnerCapabilities       OwnerCapabilities       `json:"ownerCapabilities"`
}
type ActivityControl struct {
	EnableDataRecovery       bool `json:"EnableDataRecovery"`
	EnableDataManagement     bool `json:"EnableDataManagement"`
	EnableOnlineContentIndex bool `json:"EnableOnlineContentIndex"`
}
type AdvancedClientProperties struct {
	SLACopyFallenBehindDays int  `json:"slaCopyFallenBehindDays"`
	RpoMinutes              int  `json:"rpoMinutes"`
	MailServerRole          int  `json:"mailServerRole"`
	RtoMinutes              int  `json:"rtoMinutes"`
	SLACopyRedundancy       int  `json:"slaCopyRedundancy"`
	ExcludeFromSLA          bool `json:"excludeFromSLA"`
}
type ActivityControlOptions struct {
	ActivityType       int  `json:"activityType"`
	EnableAfterADelay  bool `json:"enableAfterADelay"`
	EnableActivityType bool `json:"enableActivityType"`
}
type ClientActivityControl struct {
	ActivityControlOptions []ActivityControlOptions `json:"activityControlOptions"`
}
type JobThrottleSettings struct {
	SweepThreshold           int `json:"sweepThreshold"`
	IsJobThrottleEnabledAtCS int `json:"isJobThrottleEnabledAtCS"`
	IsJobThrottleEnabled     int `json:"isJobThrottleEnabled"`
	DataThreshold            int `json:"dataThreshold"`
	ExcludeImmidiateJobs     int `json:"excludeImmidiateJobs"`
	LogThreshold             int `json:"logThreshold"`
}
type ExchangeAdministratorAccount struct {
	UserName string `json:"userName"`
}
type EmailWebAppProperties struct {
	EnableEmailWebApp            bool                         `json:"enableEmailWebApp"`
	ExchangeAdminSMTPAddress     string                       `json:"exchangeAdminSmtpAddress"`
	ExchangeAdministratorAccount ExchangeAdministratorAccount `json:"exchangeAdministratorAccount"`
}

type OnlineTime struct {
	Time int `json:"time"`
}
type CcrTime struct {
	Time int `json:"time"`
}
type ClientReadiness struct {
	DRStatus        string      `json:"DRStatus"`
	ReadinessStatus string      `json:"readinessStatus"`
	OnlineTime      OnlineTime  `json:"onlineTime"`
	CcrTime         CcrTime     `json:"ccrTime"`
	OfflineTime     OfflineTime `json:"offlineTime"`
}
type JobResulsDir struct {
	Path        string      `json:"path"`
	UserAccount UserAccount `json:"userAccount"`
}
type OrigCommCell struct {
	CommCellID int `json:"commCellId"`
}
type OsDisplayInfo struct {
	ProcessorType string `json:"ProcessorType"`
	OSName        string `json:"OSName"`
}
type OsInfo struct {
	Type          string        `json:"Type"`
	SubType       string        `json:"SubType"`
	OsID          int           `json:"osId"`
	OsDisplayInfo OsDisplayInfo `json:"OsDisplayInfo"`
}

type IdaList struct {
	Summary   Summary   `json:"summary"`
	IdaEntity IdaEntity `json:"idaEntity"`
}

type TimeZone struct {
	TimeZoneName string `json:"TimeZoneName"`
}
type Client struct {
	DisplayName        string       `json:"displayName"`
	ClockSkewInSeconds int          `json:"ClockSkewInSeconds"`
	ClientDescription  string       `json:"clientDescription"`
	InstallDirectory   string       `json:"installDirectory"`
	TimezoneSetByUser  bool         `json:"timezoneSetByUser"`
	EvmgrcPort         int          `json:"evmgrcPort"`
	CvdPort            int          `json:"cvdPort"`
	JobResulsDir       JobResulsDir `json:"jobResulsDir"`
	OrigCommCell       OrigCommCell `json:"origCommCell"`
	OsInfo             OsInfo       `json:"osInfo"`
	IdaList            []IdaList    `json:"idaList"`
	ClientEntity       ClientEntity `json:"clientEntity"`
	VersionInfo        VersionInfo  `json:"versionInfo"`
	TimeZone           TimeZone     `json:"TimeZone"`
}
type ClientProperties struct {
	ClusterClientProperties ClusterClientProperties `json:"clusterClientProperties"`
	PseudoClientInfo        PseudoClientInfo        `json:"pseudoClientInfo"`
	ClientConfiguration     ClientConfiguration     `json:"clientConfiguration"`
	ClientProps             ClientProps             `json:"clientProps"`
	ClientReadiness         ClientReadiness         `json:"clientReadiness"`
	Client                  Client                  `json:"client"`
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
	AccessNodes           AccessNodes    `json:"accessNodes"`
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
