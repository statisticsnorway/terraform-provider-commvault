# User

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**GUID** | **string** |  | [optional] [default to null]
**FullName** | **string** |  | [optional] [default to null]
**Email** | **string** |  | [optional] [default to null]
**Enabled** | **bool** |  | [optional] [default to null]
**LockInfo** | [***LockProperties**](LockProperties.md) |  | [optional] [default to null]
**Description** | **string** | Returns the description of the user which was specified at the time of user creation or modification. | [optional] [default to null]
**ServiceType** | [***ServiceTypes**](ServiceTypes.md) |  | [optional] [default to null]
**LastLoggedIn** | **int64** | Returns the most recent time the user was logged in. It is provided in unix format. | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**AuthenticationMethod** | **string** | Specifies authentication method used by user. Default value is Commcell. | [optional] [default to null]
**AuthenticationMethodServer** | [***IdNameCompany**](IdNameCompany.md) |  | [optional] [default to null]
**AssociatedUserGroups** | [**[]IdNameProvider**](IdNameProvider.md) |  | [optional] [default to null]
**UserPrincipalName** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

