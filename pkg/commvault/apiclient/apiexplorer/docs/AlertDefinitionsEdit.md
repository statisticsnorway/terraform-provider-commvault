# AlertDefinitionsEdit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | **string** | The new name of the alert definition | [optional] [default to null]
**CriteriaList** | [**[]AlertCriteriaIdName**](AlertCriteriaIdName.md) | List of alert criteria | [optional] [default to null]
**SendIndividualNotifications** | **bool** | Flag indicating the functionality to send individual notifications | [optional] [default to null]
**SubscriptionBasedAlert** | **bool** | Flag to indicate whether its a subscription based alert. | [optional] [default to null]
**ProcessDiscoveredVMs** | **bool** | Flag indicating the functionality to send individual notifications for each discovered virtual machine | [optional] [default to null]
**AssociationsOperationType** | **string** | Allows adding to, overwriting and deleting existing alert associations. default is adding to existing alert associations | [optional] [default to ASSOCIATIONS_OPERATION_TYPE.ADD]
**Associations** | [***[]AlertAssociationIdNameType**](array.md) |  | [optional] [default to null]
**AlertTarget** | [***AlertDefinitionsTarget**](AlertDefinitionsTarget.md) |  | [optional] [default to null]
**Templates** | [***AlertDefinitionsTemplate**](AlertDefinitionsTemplate.md) |  | [optional] [default to null]
**Tokens** | [***TokenRuleGroups**](TokenRuleGroups.md) |  | [optional] [default to null]
**EventCriteriaDetails** | [***EventCriteriaDetails**](EventCriteriaDetails.md) |  | [optional] [default to null]
**GlobalConfigInfo** | [***GlobalConfigInfo**](GlobalConfigInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

