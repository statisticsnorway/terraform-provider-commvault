# CreateUserGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | To create an active directory usergroup, the domain name should be mentioned along with the usergroup name (domainName\\\\usergroupName) and localUserGroup value must be given. | [default to null]
**Description** | **string** |  | [optional] [default to null]
**EnforceFSQuota** | **bool** | Used to determine if a backup data limit will be set for the user group being created | [optional] [default to null]
**QuotaLimitInGB** | **int32** | if enforceFSQuota is set to true, the quota limit can be set in GBs | [optional] [default to null]
**LocalUserGroups** | [**[]IdName**](IdName.md) | This option is for AD user groups being created. Local user groups can be added to the active directory user groups. | [optional] [default to null]
**GlobalConfigInfo** | [***CreateGlobalConfigInfo**](CreateGlobalConfigInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

