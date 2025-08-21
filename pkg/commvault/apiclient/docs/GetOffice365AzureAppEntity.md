# GetOffice365AzureAppEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | **string** | Azure application display name. | [optional] [default to null]
**ApplicationStatus** | **string** | Azure application connection status. | [optional] [default to null]
**CredentialId** | **int64** | Credential manager id of the azure application. | [optional] [default to null]
**Flags** | [**[]GetOffice365AzureAppFlags**](GetOffice365AzureAppFlags.md) |  | [optional] [default to null]
**ApplicationId** | **string** | Azure Application ID | [default to null]
**ApplicationSecret** | **string** | Azure app secret key, needs to be base64 encoded | [default to null]
**CertificatePassword** | **string** | Password of the azure application certificate and must be in base64 encoded format. Only applicable for SharePoint Online | [optional] [default to null]
**Certificate** | **string** | Content of the azure application certificate file and must be in base64 encoded format. Accepted fileTypes are pfx, p12. Only applicable for SharePoint Online | [optional] [default to null]
**CertificateThumbprint** | **string** | Thumbprint of the azure application certificate file. Only applicable for SharePoint Online | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

