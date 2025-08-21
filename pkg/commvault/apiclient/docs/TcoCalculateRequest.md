# TcoCalculateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProvider** | [***TcoCloudProvider**](TCOCloudProvider.md) |  | [default to null]
**NumberOfVirtualMachines** | **int32** | Total number of Virtual Machines that are included for TCO assessment | [default to null]
**TotalVirtualMachineSize** | **int64** | Total size of Virtual Machines in GB that are included for TCO assessment | [default to null]
**TotalRecoveryPoints** | **int32** | Total number of Recovery Points that are included for TCO assessment | [default to null]
**TotalInstantRestoreSnapshot** | **int32** | Total number of Instant Restore Snapshots that are included for TCO assessment | [default to null]
**DailyRateOfChange** | **int32** | This is the predicted rate of change that impacts the snapshot size and backup size | [default to null]
**ReplicatedSnapsPercentage** | **int32** | This is the total percentage of snapshots replicated that is included for TCO assessment | [default to null]
**NativeDiscountPercentage** | **int32** | This is the percentage of the discount that the customer gets with the cloud vendor | [optional] [default to null]
**CommvaultDiscountPercentage** | **int32** | This is the percentage of the discount that the customer gets with Commvault | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

