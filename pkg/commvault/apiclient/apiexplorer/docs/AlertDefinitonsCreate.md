# AlertDefinitonsCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [optional] [default to null]
**AlertType** | [***AlertDefinitionsCreateAlertType**](AlertDefinitionsCreateAlertType.md) |  | [optional] [default to null]
**SendIndividualNotifications** | **bool** | Flag indicating the functionality to send individual notifications | [optional] [default to false]
**SubscriptionBasedAlert** | **bool** | Flag to indicate whether its a subscription based alert. | [optional] [default to false]
**ProcessDiscoveredVMs** | **bool** | Flag indicating the functionality to send individual notifications for each discovered virtual machine | [optional] [default to false]
**Associations** | [***[]AlertAssociationIdNameType**](array.md) |  | [optional] [default to null]
**AlertTarget** | [***AlertDefinitionsTarget**](AlertDefinitionsTarget.md) |  | [optional] [default to null]
**Templates** | [***AlertDefinitionsTemplate**](AlertDefinitionsTemplate.md) |  | [optional] [default to null]
**Tokens** | [***TokenRuleGroups**](TokenRuleGroups.md) |  | [optional] [default to null]
**EventCriteriaDetails** | [***EventCriteriaDetails**](EventCriteriaDetails.md) |  | [optional] [default to null]
**GlobalConfigInfo** | [***CreateGlobalConfigInfo**](CreateGlobalConfigInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

