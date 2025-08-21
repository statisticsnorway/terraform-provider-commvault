# Office365SharepointRestoreReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sites** | [**[]Office365SiteBkpRestoreEntity**](Office365SiteBkpRestoreEntity.md) |  | [optional] [default to null]
**Category** | [***Office365SharepointCategoryRestoreEntity**](Office365SharepointCategoryRestoreEntity.md) |  | [optional] [default to null]
**IfDocumentExists** | **string** |  | [optional] [default to IF_DOCUMENT_EXISTS.SKIP]
**DeletedItems** | **bool** | Include deleted items on Restore | [optional] [default to null]
**AdditionalRestoreOption** | **string** |  | [optional] [default to null]
**LatestVersion** | **bool** | Restore latest version only | [optional] [default to true]
**SkipLookupMetadata** | **bool** | Skip lookup metadata on sharepoint restore | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

