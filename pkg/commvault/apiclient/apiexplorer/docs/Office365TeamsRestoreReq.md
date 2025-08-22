# Office365TeamsRestoreReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]NameGuid**](NameGUID.md) | List of Teams Users for Restore | [optional] [default to null]
**DestinationUserInfo** | [***Office365TeamsDestinationUserEntity**](Office365TeamsDestinationUserEntity.md) |  | [optional] [default to null]
**Teams** | [**[]NameGuid**](NameGUID.md) | List of Teams for Restore | [optional] [default to null]
**Category** | [***Office365TeamsCategoryRestoreEntity**](Office365TeamsCategoryRestoreEntity.md) |  | [optional] [default to null]
**IfFileExists** | **string** |  | [optional] [default to IF_FILE_EXISTS.SKIP]
**DeletedItems** | **bool** | Include deleted items on Restore | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

