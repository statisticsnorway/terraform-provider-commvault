# RecoveryEntityRpDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityRecoveryPoint** | **int32** | Timestamp for entity restore in case of disaster, default value is 0 as latest recovery point | [optional] [default to 0]
**EntityRecoveryPointCategory** | **string** | Type of recovery point category | [optional] [default to ENTITY_RECOVERY_POINT_CATEGORY.LATEST]
**InheritedFrom** | **string** | Describes from where the recovery point is inherited from. If not provided, recovery point will be used as if it is set at recovery group level | [optional] [default to INHERITED_FROM.GROUP]
**TimeZoneId** | **int32** | TimeZone Id of the CS | [optional] [default to 0]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

