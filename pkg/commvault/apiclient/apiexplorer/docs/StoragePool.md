# StoragePool

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id of Storage Pool | [optional] [default to null]
**Name** | **string** | Name of Storage Pool | [optional] [default to null]
**Type_** | [***StoragePoolType**](StoragePoolType.md) |  | [optional] [default to null]
**RetentionPeriodDays** | **int32** | Retention period of pool in days | [optional] [default to null]
**WormStoragePoolFlag** | **int32** | Bit flag indicating WORM configuration of storage pool. 0 means no lock, 1 means compliance lock, 2 means worm storage lock, 4 means Object worm lock and 8 means bucket worm lock. | [optional] [default to null]
**DeviceType** | [***DeviceType**](DeviceType.md) |  | [optional] [default to null]
**StorageClass** | **string** | Storage class of the Storage Pool | [optional] [default to null]
**IsArchiveStorage** | **bool** | Flag indicating whether the storage tier is archive. | [optional] [default to null]
**Region** | [***IdNameDisplayName**](IdNameDisplayName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

