# UpdateUserGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | **string** | gives a new name to the user group. | [optional] [default to null]
**NewDescription** | **string** | gives a new description to the user group. | [optional] [default to null]
**LaptopPlan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**PlanOperationType** | [***Operations**](Operations.md) |  | [optional] [default to null]
**Enabled** | **bool** | allows the enabling/disabling of the user group. | [optional] [default to null]
**EnforceFSQuota** | **bool** | determines if a data limit will be set for the user group. | [optional] [default to null]
**QuotaLimitInGB** | **int32** | if enforceFSquota is enabled, the quota limit can be provided in GBs | [optional] [default to null]
**EnableTwoFactorAuthentication** | [***TfaOptions**](TFAOptions.md) |  | [optional] [default to null]
**EnableLocalAuthentication** | [***TfaOptions**](TFAOptions.md) |  | [optional] [default to null]
**AssociatedExternalGroups** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**ExternalUserGroupsOperationType** | **string** | Allows adding, deleting or overwriting associated external user groups of a user group. Default is adding associated external user groups | [optional] [default to EXTERNAL_USER_GROUPS_OPERATION_TYPE.ADD]
**LaptopAdmins** | **bool** | When set to true, users in this group cannot activate or be set as server owner | [optional] [default to null]
**AssociatedLocalGroups** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**LocalUserGroupsOperationType** | **string** | Allows adding, deleting or overwriting associated local user groups of a user group. Default is adding associated local user groups | [optional] [default to LOCAL_USER_GROUPS_OPERATION_TYPE.ADD]
**Users** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**UserOperationType** | [***Operations**](Operations.md) |  | [optional] [default to null]
**AllowMultipleCompanyMembers** | **bool** | This property can be used to allow addition of users/groups from child companies. Only applicable for commcell and reseller company group. | [optional] [default to false]
**DoNotInheritRestrictConsoleTypes** | **bool** | Option to not inherit the RestrictConsoleTypes from the parent. By default the value is false, parent RestrictConsoleTypes will be inherited. | [optional] [default to null]
**ConsoleTypeOperationType** | **string** |  | [optional] [default to null]
**RestrictConsoleTypes** | [***RestrictConsoleTypes**](RestrictConsoleTypes.md) |  | [optional] [default to null]
**AzureGUID** | **string** | Azure Object ID used to link this user group to Azure AD group and manage group membership of the user during SAML login | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

