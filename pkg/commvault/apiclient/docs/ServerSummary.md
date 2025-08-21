# ServerSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id of server | [optional] [default to null]
**Name** | **string** | Name of the client | [optional] [default to null]
**DisplayName** | **string** | Name to be displayed in UI | [optional] [default to null]
**HostName** | **string** | Hostname of the client | [optional] [default to null]
**Agents** | [**[]IdName**](IdName.md) | List of agents the server has installed | [optional] [default to null]
**ServerGroups** | [**[]IdName**](IdName.md) | List of server groups for this server | [optional] [default to null]
**Configured** | **bool** | Property to show whether client is in configured state or not | [optional] [default to null]
**Version** | **string** | Version of Commvault software server is running. Version is in the format: Release.SPversion.Hotfixpack. eg: 11.22.5 | [optional] [default to null]
**OS** | **string** | The operating system, for example, Windows Server 2008 R2 Enterprise. | [optional] [default to null]
**UpdateState** | [***UpdateStatus**](UpdateStatus.md) |  | [optional] [default to null]
**IsInfrastructure** | **bool** | By default, a server is classified as infrastructure if has one of the following packages installed: CommServe, Index Store, Web Server, Content Extractor, Virtual Server Agent (VSA), Web Console, Content Analyzer, Exchange, Cloud Apps | [optional] [default to null]
**IsMARoleSet** | **bool** | Used to determine if MA role is set on the client | [optional] [default to null]
**IsMAPackageInstalled** | **bool** | Used to determine if MA package is instlaled on the client | [optional] [default to null]
**NetworkReadiness** | [***FileServerStatus**](FileServerStatus.md) |  | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Tags** | [**[]IdNameValue**](IdNameValue.md) |  | [optional] [default to null]
**Commcell** | [***CommcellInfo**](CommcellInfo.md) |  | [optional] [default to null]
**ClientRoles** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**Region** | [***IdNameDisplayName**](IdNameDisplayName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

