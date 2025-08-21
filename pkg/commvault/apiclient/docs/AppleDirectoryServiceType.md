# AppleDirectoryServiceType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryType** | **string** | Apple directory service type | [default to null]
**OSXServerName** | **string** | Domain shortname to create apple directory type app | [default to null]
**Name** | **string** | Domain connect name | [default to null]
**Id** | **int32** | Required when configuring an existing dummy domain as LDAP/AD | [optional] [default to null]
**Username** | **string** | Username to create LDAP/AD app | [default to null]
**Password** | **string** | Password to create LDAP app, it should be base64 encoded | [default to null]
**AccessViaClient** | **bool** | Denotes if the domain is accessed via a proxy | [optional] [default to null]
**Proxies** | [**[]IdName**](IdName.md) | List of proxies used to connect to the domain. Required if accessViaClient is true. | [optional] [default to null]
**DoNotValidateNetBIOSName** | **bool** | If true, the NetBIOS name will not be validated. Providing a custom name without validation may cause problems during Single Sign-On. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

