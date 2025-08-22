# ChinaMobileOnest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | **string** | Name of cloud vendor | [default to null]
**ServiceHost** | **string** | IP address or fully qualified domain name or URL for the cloud library based on cloud vendor | [default to null]
**Credentials** | [***IdName**](IdName.md) |  | [default to null]
**Container** | **string** | Name of container | [default to null]
**UseDeduplication** | **bool** | Enables or disables deduplication on the storage | [optional] [default to false]
**DeduplicationDBLocation** | [**[]DedupePath**](DedupePath.md) | A list of dedupe locations can be provided for the storage pool being created. This provides an efficient way to save/store data by eliminating duplicate blocks of data during backups | [optional] [default to null]
**Name** | **string** | Name of the cloud storage library | [default to null]
**MediaAgent** | [***IdName**](IdName.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

