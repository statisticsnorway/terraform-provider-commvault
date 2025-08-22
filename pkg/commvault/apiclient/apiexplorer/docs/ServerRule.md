# ServerRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleName** | **string** |  | [optional] [default to null]
**RuleValue** | **string** | Primary value for the rule | [optional] [default to null]
**RuleSecValue** | **string** | Secondary value used for BETWEEN and NOT_BETWEEN matchCondition. For ruleName which have enum values (like OS_TYPE), this is used to store displayName. | [optional] [default to null]
**MatchCondition** | [***MatchesWithEnum**](matchesWithEnum.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

