# AdldapDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**DomainName** | **string** | Domain to which the users are associated | [optional] [default to null]
**DirectoryType** | [***AdldapDirectoryTypeForGet**](ADLDAPDirectoryTypeForGet.md) |  | [optional] [default to null]
**Credentials** | **string** | Username of the domain user used to connect to the domain | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Users** | **int32** | Number of users in the domain | [optional] [default to null]
**UserGroups** | **int32** | Number of user groups in the domain | [optional] [default to null]
**BaseDNForCardUsers** | **string** | Base DN for card users. Available only if domainType is ACTIVE_DIRECTORY. | [optional] [default to null]
**UseSecureLDAP** | **bool** | Boolean to indicate if the identity server uses secure LDAP | [optional] [default to null]
**EnableSSO** | **bool** | Denotes if SSO is enabled | [optional] [default to null]
**AccessViaClient** | **bool** | Denotes if the domain is accessed via a proxy | [optional] [default to null]
**Proxies** | [**[]IdName**](IdName.md) | List of proxies used to connect to the domain. Available only if accessViaClient is true. | [optional] [default to null]
**LDAPQueryParameters** | [**[]LdapAttribute**](LDAPAttribute.md) | List of query parameters for the LDAP domain. Available only if the directory type is LDAP_SERVER | [optional] [default to null]
**AttributeMap** | [**[]LdapAttribute**](LDAPAttribute.md) | List of attribute mappings for the LDAP domain. Available only if the directoryType is LDAP_SERVER. | [optional] [default to null]
**GlobalConfigInfo** | [***GlobalConfigInfo**](GlobalConfigInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

