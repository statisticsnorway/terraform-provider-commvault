# TwoFactorAuth

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | enable or disable two factor authentication. if enabled, all or userGroups value should be provided. | [optional] [default to null]
**All** | **bool** | enable two factor authentication for every entity. if set to false when Two factor authentication is enabled, provide userGroupId or userGroupName. If both are provided, userGroupId is taken | [optional] [default to null]
**UserGroups** | [**[]IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

