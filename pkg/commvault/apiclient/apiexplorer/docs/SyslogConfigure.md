# SyslogConfigure

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | **string** | The hostname or the IP address of the syslog server. | [default to null]
**Port** | **int32** | The port on which the syslog server accepts the logs. | [optional] [default to null]
**Enabled** | **bool** | Specifies if the syslog server is enabled or disabled | [optional] [default to null]
**SecureMessaging** | **bool** | When false UDP port will be used, when true TLS encryption will be used to connect to Syslog Server. To upload Certificate Authority file it should be enabled. | [optional] [default to null]
**CertificateAuthorityName** | **string** | File name of the uploaded certificate authority file | [optional] [default to null]
**CertificateAuthorityContent** | **string** | Content of the uploaded certificate authority file. It should be base64 encoded. Accepted fileTypes are key,crt,pem. | [optional] [default to null]
**ForwardToSyslog** | [***SyslogPolicies**](SyslogPolicies.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

