# AlertTypeItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [***AlertCategoryIdNameValue**](AlertCategoryIdNameValue.md) |  | [optional] [default to null]
**CategoryType** | [***AlertCategoryIdName**](AlertCategoryIdName.md) |  | [optional] [default to null]
**Criteria** | [***AlertCriteriaIdNameValue**](AlertCriteriaIdNameValue.md) |  | [optional] [default to null]
**QueryDetail** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ActionsAllowed** | **int32** | bit flag indicating actions allowed on this alert type. 128 means subscription support is present. | [optional] [default to null]
**ProcessDiscoveredVMs** | **bool** | whether process discovered virtual machines enabled or not | [optional] [default to false]
**NotificationOptionsAllowed** | **int32** | Bit flag indicating notification criteria value. 1 means immediate, 2 means repeated and 8 means delayed | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

