# AzureCredentialContentWithTenantId

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | **string** | Aunthentication type applicable only for Cloud Accounts with Microsoft Azure as vendor. | [optional] [default to null]
**NewName** | **string** | Updated name of credential | [optional] [default to null]
**TenantId** | **string** | Unique Azure active directory ID | [default to null]
**ApplicationId** | **string** | Unique Azure application ID | [default to null]
**NewApplicationSecret** | **string** | Application secret of Credential and must be in base64 encoded format. | [default to null]
**Environment** | **string** | Azure cloud deployed region | [default to null]
**Endpoints** | [***AzureEndpoints**](AzureEndpoints.md) |  | [default to null]
**Description** | **string** | Updated description of Credential | [optional] [default to null]
**Security** | [***CredentialSecurity**](CredentialSecurity.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

