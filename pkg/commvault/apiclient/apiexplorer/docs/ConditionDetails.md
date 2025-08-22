# ConditionDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the condition column | [default to null]
**Id** | **int32** | Unique identifier for the condition column | [default to null]
**Operation** | [***ConditionOperationType**](ConditionOperationType.md) |  | [default to null]
**Value** | **string** | Value associated with the column | [optional] [default to null]
**FromValue** | **string** | First value associated with the column when the operation type is BETWEEN | [optional] [default to null]
**ToValue** | **string** | Second value associated with the column when the operation type is BETWEEN | [optional] [default to null]
**AdvancedCriteria** | [**[]TokenFormat**](TokenFormat.md) | List of advanced criteria options | [optional] [default to null]
**ValuesList** | **[]int32** | List of the values associated with the column in case of multiple values selection | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

