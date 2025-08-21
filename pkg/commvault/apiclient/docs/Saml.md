# Saml

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | SAML name. | [optional] [default to null]
**AppKey** | **string** | Unique key for the SAML app | [optional] [default to null]
**Description** | **string** | SAML description | [optional] [default to null]
**Enabled** | **bool** | Boolean to indicate whether SAML is enabled. | [optional] [default to null]
**AutoCreateUser** | **bool** | This auto-creates non-existing user if the user detail match with the identity rule. | [optional] [default to null]
**UserGroups** | [**[]CompanyWithUserGroupAssocDetails**](CompanyWithUserGroupAssocDetails.md) | By default, auto-created users will be associated to the Tenant Users group of the company. Add mapping to override this behaviour for a company. | [optional] [default to null]
**NameIDAttribute** | **string** | nameID in SAML assertion subject is used to perform login. | [optional] [default to null]
**UserGroup** | [***IdName**](IdName.md) |  | [optional] [default to null]
**CreatedForCompany** | [***IdName**](IdName.md) |  | [optional] [default to null]
**AttributeMappings** | [**[]SamlAttributes**](SAMLAttributes.md) | attribute mapping details | [optional] [default to null]
**IdentityProviderMetaData** | [***SamlidpMetaDataResp**](SAMLIDPMetaDataResp.md) |  | [optional] [default to null]
**ServiceProviderMetaData** | [***SamlspMetaDataResp**](SAMLSPMetaDataResp.md) |  | [optional] [default to null]
**Associations** | [***SamlAssociations**](SAMLAssociations.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

