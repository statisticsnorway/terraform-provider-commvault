# ResourcePoolSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Resource Pool id | [optional] [default to null]
**Name** | **string** | Resource Pool name. | [optional] [default to null]
**SolutionType** | [***SolutionTypes**](SolutionTypes.md) |  | [optional] [default to null]
**Storage** | [***IdName**](IdName.md) |  | [optional] [default to null]
**IndexServer** | [***IdName**](IdName.md) |  | [optional] [default to null]
**AccessNodes** | [**[]AccessNodes**](AccessNodes.md) | List of access nodes associated to the resource pool | [optional] [default to null]
**AssociatedPlans** | **int32** | Refers to the number of plans associated to the resource pool. | [optional] [default to null]
**IsIndexingEnabledOnAnyPlan** | **bool** | Denotes if the resource pool is associated to any plan with indexing enabled | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

