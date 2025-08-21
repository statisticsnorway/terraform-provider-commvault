# FileServerSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | This gives the id of the file server. | [optional] [default to null]
**Name** | **string** | This give the name of the file server. | [optional] [default to null]
**DisplayName** | **string** | This gives the name of the file server as it shown on the admin console or GUI. | [optional] [default to null]
**GUID** | **string** | This returns the Globally Unique Identifier of the file server. | [optional] [default to null]
**Configured** | **bool** | This returns if the file server is configured or deconfigured. | [optional] [default to null]
**Type_** | **string** | This returns the type of the file server. Example: Windows, Qumulo | [optional] [default to null]
**CommcellName** | **string** | This returns the CommCell the file server is connected to. | [optional] [default to null]
**LastBackup** | [***LastBackupJobInfoWithLastSuccessfulBackup**](LastBackupJobInfoWithLastSuccessfulBackup.md) |  | [optional] [default to null]
**ApplicationSize** | **int32** | Provides the application size of the file server. It is provided in bytes. | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**SLA** | [***SlaDetailsWithReasonEntity**](SLADetailsWithReasonEntity.md) |  | [optional] [default to null]
**Status** | [***FileServerStatus**](FileServerStatus.md) |  | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Tags** | [**[]IdNameValue**](IdNameValue.md) | Tags associated to file server | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

