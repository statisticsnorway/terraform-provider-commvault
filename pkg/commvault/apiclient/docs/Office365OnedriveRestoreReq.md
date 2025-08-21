# Office365OnedriveRestoreReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]NameGuid**](NameGUID.md) |  | [optional] [default to null]
**Groups** | [**[]NameGuid**](NameGUID.md) |  | [optional] [default to null]
**IfFileExists** | **string** |  | [optional] [default to IF_FILE_EXISTS.SKIP]
**DeletedItems** | **bool** | Include deleted items on Restore | [optional] [default to null]
**MatchEmailAddress** | **bool** | Match destination user based on the email address | [optional] [default to null]
**SkipFilePermissions** | **bool** | Skip Onedrive file permissions during restore | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

