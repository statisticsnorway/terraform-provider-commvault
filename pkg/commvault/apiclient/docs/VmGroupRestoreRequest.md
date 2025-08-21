# VmGroupRestoreRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PowerOnVmAfterRestore** | **bool** | Turn ON virtual machine after restore. Defaults to false. | [optional] [default to false]
**OverwriteVM** | **bool** | Unconditionally overwrite VM during restore. Defaults to false. | [optional] [default to false]
**InPlaceRestore** | **bool** | Flag to set if restore is in-place or out-of-place. Defaults to false. | [default to false]
**BackupSource** | **int32** | Backup source information from where you want to restore. 0 (Automatic), 1 (Snap Copy), 2 (Primary), 3 to N (Aux Copy) | [optional] [default to 0]
**Destination** | [***IdName**](IdName.md) |  | [optional] [default to null]
**AccessNode** | [***IdName**](IdName.md) |  | [optional] [default to null]
**AccessNodeGroup** | [***IdName**](IdName.md) |  | [optional] [default to null]
**VmDestinationInfo** | [***VmDestinationInfo**](VMDestinationInfo.md) |  | [default to null]
**TransportMode** | [***VmWareTransportMode**](VmWareTransportMode.md) |  | [optional] [default to null]
**RestoreFlavour** | [***RestoreFlavour**](RestoreFlavour.md) |  | [optional] [default to null]
**NotifyOnJobCompletion** | **bool** | Enable email notificaiton for job status. Defaults to false. | [optional] [default to false]
**ReuseVMClient** | **bool** | Reuse the existing VM client instance of creating new one after restore. Defaults to true for inplace restore and false for out of place restore. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

