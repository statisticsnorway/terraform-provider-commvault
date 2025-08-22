# UserGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**GUID** | **string** |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Enabled** | **bool** | Returns if the user group is enabled or disabled | [optional] [default to null]
**EnforceFSQuota** | **bool** |  | [optional] [default to null]
**QuotaLimitInGB** | **int32** |  | [optional] [default to null]
**ServiceType** | [***ServiceTypes**](ServiceTypes.md) |  | [optional] [default to null]
**Email** | **string** | Returns if there is a email ID associated with the usergroup. | [optional] [default to null]
**EnableTwoFactorAuthentication** | [***TfaOptions**](TFAOptions.md) |  | [optional] [default to null]
**EnableLocalAuthentication** | [***TfaOptions**](TFAOptions.md) |  | [optional] [default to null]
**AssociatedExternalGroups** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**LaptopAdmins** | **bool** | When set to true, users in this group cannot activate or be set as server owner | [optional] [default to null]
**AssociatedLocalGroups** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Users** | [**[]IdName**](IdName.md) | Returns list of users that are associated with this userGroup | [optional] [default to null]
**AssociatedEntities** | [**[]AssocEntity**](AssocEntity.md) |  | [optional] [default to null]
**EligibleToAllowMultipleCompanyMembers** | **bool** | Read only property. Denotes if the group is eligible for allowMultipleCompanyMembers | [optional] [default to null]
**AllowMultipleCompanyMembers** | **bool** | This property denotes that addition of users/groups from child companies is allowed. Only applicable for commcell and reseller company group. | [optional] [default to false]
**DoNotInheritRestrictConsoleTypes** | **bool** | RestrictConsoleTypes are inherited from the parent or not. | [optional] [default to null]
**RestrictedConsoleTypes** | [**[]RestrictedConsoleTypes**](RestrictedConsoleTypes.md) |  | [optional] [default to null]
**AzureGUID** | **string** | Azure Object ID used to link this user group to Azure AD group and manage group membership of the user during SAML login | [optional] [default to null]
**ShowAzureGuidOption** | **bool** | Read only property. Denotes if the group is eligible to have Azure Object ID property | [optional] [default to null]
**GlobalConfigInfo** | [***GlobalConfigInfo**](GlobalConfigInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

