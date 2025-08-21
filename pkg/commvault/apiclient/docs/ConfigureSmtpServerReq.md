# ConfigureSmtpServerReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SMTPServerName** | **string** | The name of the mail server that sends alerts, scheduled reports, log files, and additional information. | [default to null]
**SMTPPort** | **int32** | The port number that connects to the mail server. | [default to null]
**SenderEmail** | **string** | The sender email address used for emails sent from the software. | [default to null]
**SenderName** | **string** | The sender name used for emails sent from the software. | [default to null]
**EncryptionAlgorithm** | [***SmtpServerEncryptionAlgorithm**](SMTPServerEncryptionAlgorithm.md) |  | [optional] [default to null]
**Username** | **string** | This option is used when authentication on the mail server is required | [optional] [default to null]
**Password** | **string** | Password must be in base64 encoded format. Similar to username, if authentication on the mail server is required, this option is required. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

