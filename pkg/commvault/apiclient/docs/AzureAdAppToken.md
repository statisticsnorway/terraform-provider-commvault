# AzureAdAppToken

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefreshToken** | **string** | User based refresh token generated with corresponding app credentials. Can be left blank if the azure app is protected API approved. | [optional] [default to null]
**ApplicationId** | **string** | Azure Application ID | [default to null]
**ApplicationSecret** | **string** | Azure app secret key, needs to be base64 encoded | [default to null]
**CertificatePassword** | **string** | Password of the azure application certificate and must be in base64 encoded format. Only applicable for SharePoint Online | [optional] [default to null]
**Certificate** | **string** | Content of the azure application certificate file and must be in base64 encoded format. Accepted fileTypes are pfx, p12. Only applicable for SharePoint Online | [optional] [default to null]
**CertificateThumbprint** | **string** | Thumbprint of the azure application certificate file. Only applicable for SharePoint Online | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

