# GetEmailServerDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SMTPServerName** | **string** | The name of the mail server that sends alerts, scheduled reports, log files, and additional information. | [optional] [default to null]
**SMTPPort** | **int32** | The port number that connects to the mail server. | [optional] [default to null]
**SenderEmail** | **string** | The sender email address used for emails sent from the software. | [optional] [default to null]
**SenderName** | **string** | The sender name used for emails sent from the software. | [optional] [default to null]
**EncryptionAlgorithm** | [***SmtpServerEncryptionAlgorithm**](SMTPServerEncryptionAlgorithm.md) |  | [optional] [default to null]
**UseAuthentication** | **bool** | The option to use authentication on the mail server. This is typically used in external or internet mail server configurations. | [optional] [default to null]
**Username** | **string** | only when useAuthentication is true, username would be populated | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

