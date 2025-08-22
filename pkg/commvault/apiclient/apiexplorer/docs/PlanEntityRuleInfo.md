# PlanEntityRuleInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | [***IdName**](IdName.md) |  | [default to null]
**Plan** | [***IdNameGuid**](IdNameGUID.md) |  | [default to null]
**Workloads** | [**[]IdName**](IdName.md) | This will include list of apptypes that should be evaluated against workload for plan association | [optional] [default to null]
**Regions** | [**[]IdNameGuid**](IdNameGUID.md) | This will include list of regions that should be evaluated against workload region for plan association | [optional] [default to null]
**Tags** | [**[]PlanEntityRuleTag**](PlanEntityRuleTag.md) | This will include list of tags that should be evaluated against workload tags (if any) for plan association | [optional] [default to null]
**ServerGroups** | [**[]IdNameGuid**](IdNameGUID.md) | This will include list of Server groups that should be evaluated against workload server group for plan association | [optional] [default to null]
**Rank** | **int32** | This will suggest rank/priority of the plan rule. | [optional] [default to null]
**Solutions** | [**[]IdName**](IdName.md) | This will include list of solutions that should be evaluated against workload for plan association. | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**GlobalConfigInfo** | [***GlobalConfigInfo**](GlobalConfigInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

