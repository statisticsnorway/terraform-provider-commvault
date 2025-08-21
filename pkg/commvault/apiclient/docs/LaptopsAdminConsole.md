# LaptopsAdminConsole

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id of the laptop | [optional] [default to null]
**Name** | **string** | name of the laptop | [optional] [default to null]
**Owners** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**Username** | **string** | comma seperated value for username of the laptop | [optional] [default to null]
**Email** | **string** | comma seperated value for email of the laptop | [optional] [default to null]
**Configured** | **bool** | defines if the laptop is configured or not | [optional] [default to null]
**LastBackupJobInfo** | [***LastBackupJobInfo**](LastBackupJobInfo.md) |  | [optional] [default to null]
**TotalBackupSize** | **int32** | application size (in bytes) for the laptop | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Sla** | [***SlaDetails**](SLADetails.md) |  | [optional] [default to null]
**SlaReason** | **string** | reason for the current sla status | [optional] [default to null]
**LastSuccessfulBackup** | **int32** | time (in unix timestamp) for the last successful backup | [optional] [default to null]
**Tags** | [**[]IdNameType**](IdNameType.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

