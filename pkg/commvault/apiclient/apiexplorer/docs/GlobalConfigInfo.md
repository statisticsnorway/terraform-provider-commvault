# GlobalConfigInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Commcells** | [**[]GlobalConfigCommcellInfo**](GlobalConfigCommcellInfo.md) | List of Service CommCells where the global configuration is applied | [optional] [default to null]
**Scope** | **string** | The entity level at which the config is applied. | [optional] [default to null]
**ScopeFilterQuery** | **string** | CommCellEntityCache filter query string used for filtering the scope | [optional] [default to null]
**ApplyOnAllCommCells** | **bool** | Decides whether the global configuration should be applied to all the Service commcells, including the newly created ones | [optional] [default to null]
**IsMarkedForDeletion** | **bool** | Indicates whether global configuration deletion has been started. | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

