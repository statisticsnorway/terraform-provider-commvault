# AzureCredentialInfoWithCertificate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **string** | Aunthentication type applicable only for Cloud Accounts with Microsoft Azure as vendor. | [default to null]
**TenantId** | **string** | Unique Azure active directory ID | [default to null]
**ApplicationId** | **string** | Unique Azure application ID | [default to null]
**ApplicationSecret** | **string** | Application secret of Credential and must be in base64 encoded format. | [default to null]
**ApplicationCertificatePassword** | **string** | Password of the certificate and must be in base64 encoded format. | [default to null]
**Certificate** | **string** | Content of the certificate file. It should be base64 encoded. Accepted fileTypes are pfx, p12. | [default to null]
**CertificateThumbprint** | **string** | Thumbprint of the certificate file. | [default to null]
**Environment** | **string** | Azure cloud deployed region | [default to null]
**Endpoints** | [***AzureEndpoints**](AzureEndpoints.md) |  | [optional] [default to null]
**Description** | **string** | Description of Credential | [optional] [default to null]
**Security** | [***CredentialSecurity**](CredentialSecurity.md) |  | [optional] [default to null]
**AccountType** | [***CredentialAccountType**](CredentialAccountType.md) |  | [default to null]
**VendorType** | [***CloudVendorType**](CloudVendorType.md) |  | [optional] [default to null]
**Name** | **string** | Name of Credential | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

