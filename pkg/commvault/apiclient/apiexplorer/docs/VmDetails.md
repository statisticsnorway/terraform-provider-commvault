# VmDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of virtual machine | [optional] [default to null]
**State** | [***VmState**](VMState.md) |  | [optional] [default to null]
**IpAddress** | **string** | IP address of virtual machine | [optional] [default to null]
**OperatingSystem** | **string** | Operating system of virtual machine | [optional] [default to null]
**Vendor** | [***VsVendor**](VSVendor.md) |  | [optional] [default to null]
**ExpirationTime** | **int32** | Expiration time of virtual machine | [optional] [default to null]
**Creator** | **string** | Creator of virtual machine | [optional] [default to null]
**Description** | **string** | Description of virtual machine | [optional] [default to null]
**Owners** | **[]string** | List of owners for virtual machine | [optional] [default to null]
**NumberOfCPUs** | **int32** | Number of CPUs in virtual machine | [optional] [default to null]
**Memory** | **int64** | Memory size of virtual machine | [optional] [default to null]
**NumberOfNICs** | **int32** | Number of network interface cards | [optional] [default to null]
**Networks** | **[]string** | List of networks associated with virtual machine | [optional] [default to null]
**Disks** | [**[]VmDisk**](VMDisk.md) | List of the disks associated with virtual machine | [optional] [default to null]
**IsSnapshotSupported** | **bool** | Whether the live mount is enabled or disabled | [optional] [default to null]
**Lab** | [***IdName**](IdName.md) |  | [optional] [default to null]
**VmGUID** | **string** | GUID of virtual machine | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

