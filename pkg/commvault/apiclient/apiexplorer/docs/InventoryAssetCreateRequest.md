# InventoryAssetCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the file server | [default to null]
**HostName** | **string** | Fully qualified domain name of the server | [default to null]
**IpAddress** | **string** | Ip Address of the server | [default to null]
**OperatingSystem** | **string** | Operating system of the server | [default to null]
**CountryCode** | **string** | Country code to which the server belongs | [optional] [default to null]
**StorageName** | **string** | Name of the object storage | [default to null]
**HostURL** | **string** | Service account URL | [default to null]
**Credentials** | [***IdName**](IdName.md) |  | [default to null]
**AccessNode** | [***IdName**](IdName.md) |  | [default to null]
**StorageType** | [***ObjectStorageAssetType**](ObjectStorageAssetType.md) |  | [default to null]
**AssetType** | [***InventoryAssetType**](InventoryAssetType.md) |  | [optional] [default to null]
**IdentityServers** | **[]string** | List of identity server names | [default to null]
**StartDataCollection** | **bool** | Flag to signify if data collection job needs to be invoked after addition of identity server | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

