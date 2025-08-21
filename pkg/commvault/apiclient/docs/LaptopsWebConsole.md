# LaptopsWebConsole

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | id of the laptop client | [optional] [default to null]
**Name** | **string** | name of the laptop client | [optional] [default to null]
**LastBackupTime** | **int32** | This specifies the last backup time (in unix timestamp) | [optional] [default to null]
**LastBackupJobInfo** | [***LastBackupJobInfo**](LastBackupJobInfo.md) |  | [optional] [default to null]
**TotalBackupSize** | **int32** | This specifies the total backup size (in bytes) | [optional] [default to null]
**NextBackupTime** | **int32** | This species the timestamp (in unix timestamp) of the next scheduled backup | [optional] [default to null]
**SlaStatus** | [***SlaStatus**](SLAStatus.md) |  | [optional] [default to null]
**SlaReason** | **string** | reason for the current sla status | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

