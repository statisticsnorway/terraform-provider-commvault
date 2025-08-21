# OracleObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | **string** | Name of cloud vendor | [default to null]
**ServiceHost** | **string** | IP address or fully qualified domain name or URL for the cloud library based on cloud vendor | [optional] [default to objectstorage.us-phoenix-1.oraclecloud.com]
**Credentials** | [***IdName**](IdName.md) |  | [default to null]
**CompartmentName** | **string** | OCI compartment name | [optional] [default to null]
**Bucket** | **string** | Name of the bucket | [default to null]
**StorageClass** | **string** | Appropriate storage class for your account | [default to null]
**UseCombinedStorage** | **bool** | Applicable for Archive storage class only | [optional] [default to null]
**CombinedStorageClass** | **string** | Appropriate combined storage class for archive. Applicable only when useCombinedStorage is true. | [optional] [default to null]
**UseDeduplication** | **bool** | Enables or disables deduplication on the storage | [optional] [default to false]
**DeduplicationDBLocation** | [**[]DedupePath**](DedupePath.md) | A list of dedupe locations can be provided for the storage pool being created. This provides an efficient way to save/store data by eliminating duplicate blocks of data during backups | [optional] [default to null]
**Name** | **string** | Name of the cloud storage library | [default to null]
**MediaAgent** | [***IdName**](IdName.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

