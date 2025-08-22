# CreateUser

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | **string** | Used to provide a name to the new user. | [optional] [default to null]
**Name** | **string** | Used to provide the new user with a username. This username can be used for logging in the user instead of email-id when duplicate email-ids are present. For external user, it is necessary to provide the domain name along with the username (domainName\\\\username). To create a company user, the company id or name needs to be provided in the company entity. | [optional] [default to null]
**ExternalProviderId** | **int32** | Please provide this entity only when creating an AD user. Not needed for local and company users. | [optional] [default to null]
**Email** | **string** | Used to provide an email-id to the new user. This email-id is used for logging in the user. Please note that email ids are compulsory for company and local users and optional for external users. | [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**UserGroups** | [**[]IdName**](IdName.md) | Provide a list of userGroups that the user should be a part of. Either id or name or both can be provided. | [optional] [default to null]
**UseSystemGeneratePassword** | **bool** | Choose to provide a system generated password to the user instead of providing your own password. An email will be sent to the user to reset the password. If it is set to true, password tag need not be provided. If it is set to false, password needs to be provided in the password tag in Base64 format. | [optional] [default to null]
**Password** | **string** | Used to provide a password to the user being created. This will be accepted when the useSystemGeneratePassword tag is false. The password has to be provided in Base64 format. | [optional] [default to null]
**InviteUser** | **bool** | User will receive an email to install backup software package on their device if this is set to true. | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

