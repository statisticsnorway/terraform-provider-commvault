# UpdateUser

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | Change name for existing user | [optional] [default to null]
**NewName** | **string** | Change user name for existing user. Username can be used for logging-in as an alternate to email-id if duplicate email-ids are present. | [optional] [default to null]
**Email** | **string** | Change email-id for existing user. This email-id can be used for logging-in. | [optional] [default to null]
**UserGroupOperation** | **string** | Allows adding to, overwriting and deleting existing user groups. default is adding to existing userGroups | [optional] [default to USER_GROUP_OPERATION.ADD]
**UserGroups** | [**[]IdName**](IdName.md) | Provide a list of userGroups that the user should be a part of. Note that for external users, user groups cannot be modified. | [optional] [default to null]
**NewPassword** | **string** | Change existing password for user. validationPassword must also be provided when changing password. Password should be in provided in Base64 format. | [optional] [default to null]
**ValidationPassword** | **string** | Provide the old password in Base64 format when updating the password. The new password has to be provided in the password tag. | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Enabled** | **bool** | enable or disable the user. | [optional] [default to null]
**UserPrincipalName** | **string** | Change User Principal Name(UPN) for existing user. This User Principal Name can be used for logging-in. | [optional] [default to null]
**AuthenticationMethod** | **string** | Change the current authentication method of user. SAML user association can be removed using this. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

