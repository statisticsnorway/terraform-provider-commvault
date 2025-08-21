# CreateGlobalConfigInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [optional] [default to null]
**Commcells** | [**[]GlobalConfigCommcellInfo**](GlobalConfigCommcellInfo.md) | List of Service CommCells where the global configuration should be applied | [optional] [default to null]
**Scope** | **string** | The entity level at which the config has to be applied. | [optional] [default to null]
**ScopeFilterQuery** | **string** | CommCellEntityCache filter query string using for filtering the scope | [optional] [default to null]
**ApplyOnAllCommCells** | **bool** | Decides whether the global configuration should be applied to all the Service commcells, including the newly created ones | [optional] [default to null]
**ActionOnLocalEntity** | **string** | Action that will be taken on the local entity that has the same name as the global entity that needs to be created | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

