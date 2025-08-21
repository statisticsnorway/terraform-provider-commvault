# AzureCredentialInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **string** | Aunthentication type applicable only for Cloud Accounts with Microsoft Azure as vendor. | [default to null]
**AccountName** | **string** | Account name of Credential, applicable only if authType is Access Key | [default to null]
**AccessKeyId** | **string** | Access key ID of Credential, applicable only if authType is Access Secret Key and must be in base64 encoded format. | [default to null]
**Description** | **string** | Description of Credential | [optional] [default to null]
**Security** | [***CredentialSecurity**](CredentialSecurity.md) |  | [optional] [default to null]
**AccountType** | [***CredentialAccountType**](CredentialAccountType.md) |  | [default to null]
**VendorType** | [***CloudVendorType**](CloudVendorType.md) |  | [optional] [default to null]
**Name** | **string** | Name of Credential | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

