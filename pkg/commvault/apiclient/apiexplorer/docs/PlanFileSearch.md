# PlanFileSearch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Flag for enabling indexing | [optional] [default to null]
**Status** | [***IndexingStatusType**](IndexingStatusType.md) |  | [optional] [default to null]
**StatusMessage** | **string** | Tells what is happening behind the scene, so that user can knows why indexing is not enabled or if its in progress | [optional] [default to null]
**Errors** | [**[]PlanFileSearchSetupError**](PlanFileSearchSetupError.md) | File search was enabled on plan but failed to process some of the storage pool(s) with these errors | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

