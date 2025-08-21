# HpeCatalystBucketContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudType** | **string** | Name of cloud vendor | [default to null]
**MediaAgent** | [***IdName**](IdName.md) |  | [default to null]
**StoreOnceHost** | **string** | IP address or COFC identifier in case of Fibre Channel associated with the HPE StoreOnce device | [optional] [default to null]
**Store** | **string** | Name of the store that is created on the StoreOnce management console | [optional] [default to null]
**Username** | **string** | Username used to access StoreOnce management console | [optional] [default to null]
**Password** | **string** | Password used to access StoreOnce management console (Should be in Base64 format) | [optional] [default to null]
**Access** | [***AccessType**](AccessType.md) |  | [optional] [default to null]
**Enable** | **bool** | Enable/Disable access of bucket to a media Agent | [optional] [default to null]
**Configuration** | [***CloudBucketConfiguration**](CloudBucketConfiguration.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

