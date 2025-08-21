# VmDiskFilter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overwrite** | **bool** | True if disk filter in vmgroup has to be overwritten, by default it will append the content | [optional] [default to false]
**FilterType** | [***VmDiskFilterType**](vmDiskFilterType.md) |  | [optional] [default to null]
**Name** | **string** | The string to be filtered | [optional] [default to null]
**Value** | **string** | The value string to be filtered, in case of disk tag , value of tag to be filtered | [optional] [default to null]
**Condition** | [***RuleOperationType**](RuleOperationType.md) |  | [optional] [default to null]
**VmName** | **string** | VM Name of the Virtual Machine whose disk has to be filtered . This is optional. if not given, all disks of name and type from all Vms added in content will be filtered | [optional] [default to null]
**VmGuid** | **string** | VM Guid of the Virtual Machine whose disk has to be filtered . This is optional. if not given, all disks of name and type from all Vms added in content will be filtered | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

