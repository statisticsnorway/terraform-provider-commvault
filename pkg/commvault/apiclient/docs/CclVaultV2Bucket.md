# CclVaultV2Bucket

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | **string** | Name of cloud vendor | [default to null]
**MediaAgent** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ServiceHost** | **string** | IP address or fully qualified domain name or URL for the cloud library based on cloud vendor | [default to s3-nzl.vault.net.nz]
**Credentials** | [***IdName**](IdName.md) |  | [default to null]
**Bucket** | **string** | Name of bucket | [default to null]
**ProxyAddress** | **string** | If the MediaAgent accesses the mount path using a proxy then proxy server address needs to be provided. If you want to remove proxy information, pass empty string in proxyAddress. | [optional] [default to null]
**Port** | **int32** | Port for proxy configuration | [optional] [default to null]
**Username** | **string** | Username for proxy configuration | [optional] [default to null]
**Password** | **string** | Password for proxy configuration (Should be in Base64 format) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

