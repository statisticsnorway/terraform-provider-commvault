# BackupLocationDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id of the backup location | [optional] [default to null]
**Name** | **string** | Name of the backup location | [optional] [default to null]
**BackupLocation** | **string** | backup location path | [optional] [default to null]
**MediaAgent** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Configuration** | [***DiskStorageConfiguration**](DiskStorageConfiguration.md) |  | [optional] [default to null]
**TotalCapacity** | **int32** | total capacity of the backup location | [optional] [default to null]
**FreeSpace** | **int32** | total free space on the backup location | [optional] [default to null]
**DiskAccessPaths** | [**[]AccessPathDetails**](AccessPathDetails.md) | access paths available on the backup location | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**Access** | [***AccessType**](AccessType.md) |  | [optional] [default to null]
**Credentials** | [***CredentialUserName**](CredentialUserName.md) |  | [optional] [default to null]
**SavedCredentials** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

