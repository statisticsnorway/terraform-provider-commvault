# AlertDefinitionsDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**AlertSummary** | [***AlertDetailsSummary**](AlertDetailsSummary.md) |  | [optional] [default to null]
**SendIndividualNotifications** | **bool** | true if individual notifications are on for the alert definitions | [optional] [default to null]
**ProcessDiscoveredVMs** | **bool** | True if individual notifications for each discovered virtual machine are enabled for the alert definition | [optional] [default to null]
**SubscriptionBasedAlert** | **bool** | Flag to indicate whether its a subscription based alert. | [optional] [default to null]
**Associations** | [***[]AlertAssociationIdNameType**](array.md) |  | [optional] [default to null]
**AlertTarget** | [***AlertDefinitionsTarget**](AlertDefinitionsTarget.md) |  | [optional] [default to null]
**Templates** | [***AlertDefinitionsTemplate**](AlertDefinitionsTemplate.md) |  | [optional] [default to null]
**AdditionalProperties** | [***AlertAdditionalProperties**](AlertAdditionalProperties.md) |  | [optional] [default to null]
**Tokens** | [***TokenRuleGroups**](TokenRuleGroups.md) |  | [optional] [default to null]
**CustomQueryDetails** | [***CustomQueryDetails**](CustomQueryDetails.md) |  | [optional] [default to null]
**EventCriteriaDetails** | [***EventCriteriaDetails**](EventCriteriaDetails.md) |  | [optional] [default to null]
**GlobalConfigInfo** | [***GlobalConfigInfo**](GlobalConfigInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

