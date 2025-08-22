# VmPreviewInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name  of the VM | [optional] [default to null]
**GUID** | **string** | GUID  of the VM | [optional] [default to null]
**PoweredStatus** | **string** | Power status of the VM from the hypervisor Eg- Powered on ,  Running if the VM is powered on . Values might differ based on hypervisor type | [optional] [default to null]
**ToolStatus** | **string** | if Installed then the corresponding tools version of the VM. VMWare - VMWare tools version Hyper-V - Integration services version Values might differ based on hypervisor type status would be not installed if tools is not installed. | [optional] [default to null]
**Host** | **string** | host of the VM | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

