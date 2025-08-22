# SamlUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | SAML description | [optional] [default to null]
**Enabled** | **bool** | Boolean to indicate whether SAML is enabled. | [optional] [default to null]
**AutoCreateUser** | **bool** | This auto-creates non-existing user if the user detail match with the identity rule. | [optional] [default to null]
**UserGroups** | [**[]CompanyWithUserGroupAssocDetails**](CompanyWithUserGroupAssocDetails.md) | By default, auto-created users will be associated to the Tenant Users group of the company. Add mapping to override this behaviour for a company. | [optional] [default to null]
**NameIDAttribute** | **string** | nameID in SAML assertion subject is used to perform login. | [optional] [default to null]
**AttributeMappings** | [**[]SamlAttributes**](SAMLAttributes.md) | attribute mapping details | [optional] [default to null]
**IdentityProviderMetaData** | [***SamlidpMetaDataReq**](SAMLIDPMetaDataReq.md) |  | [optional] [default to null]
**ServiceProviderMetaData** | [***SamlspMetaData**](SAMLSPMetaData.md) |  | [optional] [default to null]
**Associations** | [***SamlAssociations**](SAMLAssociations.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

