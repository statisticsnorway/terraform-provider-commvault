# LdapAttributes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserGroupFilter** | **string** | Custom attribute string of LDAP query paramters | [optional] [default to (objectClass=group)]
**UserFilter** | **string** | Custom attribute string of LDAP query paramters | [optional] [default to (&(objectCategory=User)(sAMAccountName=*))]
**UniqueIdentifier** | **string** | Custom attribute string of LDAP query paramters | [optional] [default to sAMAccountName]
**BaseDN** | **string** | Custom attribute string of LDAP query paramters | [optional] [default to base DN]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

