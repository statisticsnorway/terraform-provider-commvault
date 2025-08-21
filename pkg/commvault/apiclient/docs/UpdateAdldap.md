# UpdateAdldap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectoryType** | [***AdldapDirectoryType**](ADLDAPDirectoryType.md) |  | [default to null]
**Host** | **string** | The fully qualified domain name that you use to identify this network resource. Required only if directoryType is LDAP_SERVER | [optional] [default to null]
**OSXServerName** | **string** | The fully qualified domain name that you use to identify this network resource. Required only if directoryType is APPLE_DIRECTORY_SERVICE | [optional] [default to null]
**NETBIOSName** | **string** | The fully qualified domain name that you use to identify this network resource. Required only if directoryType is ACTIVE_DIRECTORY, OPEN_LDAP or ORACLE_DIRECTORY  | [optional] [default to null]
**Name** | **string** | The fully qualified domain name, for example, my.domain.example.com | [default to null]
**Username** | **string** | The username for a user who has at least read permission for the domain | [optional] [default to null]
**Password** | **string** | Password for the domain user. Should be in Base64 encoded format. | [optional] [default to null]
**BaseDNForCardUsers** | **string** | Base DN for card users | [optional] [default to null]
**UseSecureLDAP** | **bool** | Boolean to indicate if the app use secure LDAP. Valid only for directory types - ACTIVE_DIRECTORY, ORACLE_DIRECTORY and LDAP_SERVER. | [optional] [default to null]
**EnableSSO** | **bool** | Denotes if SSO should be enabled for the domain. Valid only for ACTIVE_DIRECTORY. | [optional] [default to null]
**AccessViaClient** | **bool** | Denotes if the domain is to be accessed via a proxy | [optional] [default to null]
**Proxies** | [**[]IdName**](IdName.md) | List of proxies used to connect to the domain. Available only if accessViaClient is true. | [optional] [default to null]
**LDAPQueryParameters** | [**[]LdapAttribute**](LDAPAttribute.md) | List of overridden query parameters for the LDAP domain. Valid only if the directory type is LDAP_SERVER | [optional] [default to null]
**AttributeMap** | [**[]LdapAttribute**](LDAPAttribute.md) | List of overridden attribute mappings for the LDAP domain. Valid only if the directoryType is LDAP_SERVER. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

