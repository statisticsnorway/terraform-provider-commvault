# LaptopPlanGeneralInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | **int32** | Number of users associated with this plan | [optional] [default to null]
**Laptops** | **int32** | Number of laptops associated with this plan | [optional] [default to null]
**OptimizedForCloudBackups** | **bool** | This feature allows laptops to write backup directly to the cloud storage. It helps to optimize scale by reducing server dependency and extra data hops. Once the feature is enabled, the existing and the newly-added laptops use optimized backups. | [optional] [default to false]
**StorageResourcePoolMap** | [**[]StorageResourcePoolMap**](StorageResourcePoolMap.md) | Optimization for cloud backups can only be used when resource pool are configured for atleast one storage. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

