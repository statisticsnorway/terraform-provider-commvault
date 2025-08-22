# MountPathJobInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | **int64** | The system-generated ID assigned to the job. | [optional] [default to null]
**JobStatus** | [***JobStatus**](JobStatus.md) |  | [optional] [default to null]
**ClientName** | **string** | The name of the client associated with the job. | [optional] [default to null]
**AppName** | **string** | The name of application | [optional] [default to null]
**InstanceName** | **string** | The name of the instance associated with the job. | [optional] [default to null]
**BackupSetName** | **string** | The name of the backup set associated with the job. | [optional] [default to null]
**SubclientName** | **string** | The name of the subclient associated with the job. | [optional] [default to null]
**BackupLevel** | **string** | The level of the backup job | [optional] [default to null]
**CyclesNSeq** | **string** | Cycles and Sequences | [optional] [default to null]
**ArchFileType** | [***ArchFileType**](ArchFileType.md) |  | [optional] [default to null]
**DataStatus** | **string** | Status of the data of job | [optional] [default to null]
**BackupTime** | **int64** | The time when backup is done | [optional] [default to null]
**RetainTill** | **int64** | Retention DateTime for job | [optional] [default to null]
**MediaSideName** | **string** | Media Side Name | [optional] [default to null]
**ArchFileName** | **string** | Archive File Name | [optional] [default to null]
**StoragePolicyNCopy** | **string** | Storage Policy and Copy Name | [optional] [default to null]
**DataWritten** | **int64** | Amount of data written on mountpath during job | [optional] [default to null]
**SizeOfBackup** | **int64** | The amount of application data that was protected during the job | [optional] [default to null]
**SizeOfApplication** | **int64** | The amount of application data that was protected during the job. | [optional] [default to null]
**IsAged** | **bool** | Indicates if job is aged or not. | [optional] [default to null]
**IsValidData** | **bool** | Indicates if job has valid Data | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

