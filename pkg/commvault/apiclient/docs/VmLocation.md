# VmLocation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InventoryPath** | **string** | Folder path where you can locate vm, empty if Datacenter is selected for location. Default is set to Datacenter | [optional] [default to null]
**DataCenterName** | **string** |  | [default to null]
**Host** | **string** | If ESX-Host for resource the host moref or If ESX-Cluster then cluster moref and if resource pool is used for resource then first go for host moref then cluster moref | [default to null]
**ResourcePool** | **string** | Resource Pool moref if resource pool is selected for resource | [optional] [default to null]
**Cluster** | **string** | ESX-Cluster moref if cluster is selected for resource | [optional] [default to null]
**Datastore** | **string** | If cluster is selected for storage then moref of datastore cluster else dataStoreName | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

