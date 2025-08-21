# SendTestMailReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SMTPServerName** | **string** | The name of the mail server that sends alerts, scheduled reports, log files, and additional information. | [default to null]
**SMTPPort** | **int32** | The port number that connects to the mail server. | [default to null]
**SenderEmail** | **string** | The sender email address used for emails sent from the software. | [default to null]
**SenderName** | **string** | The sender name used for emails sent from the software. | [default to null]
**EncryptionAlgorithm** | [***SmtpServerEncryptionAlgorithm**](SMTPServerEncryptionAlgorithm.md) |  | [optional] [default to null]
**UseAuthentication** | **bool** | The option to use authentication on the mail server. This is typically used in external or internet mail server configurations. | [optional] [default to null]
**Username** | **string** | If username is specified, password field must also be specified. To use previously saved username and password, leave out username and password in payload. | [optional] [default to null]
**Password** | **string** | If password is specified, username must also be specified. password should be a base 64 encoded string. To use previoulsy saved username and password, leave out username and password in payload. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

