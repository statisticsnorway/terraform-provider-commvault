# RecoveryEntity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Id of the recovery entity | [optional] [default to null]
**Name** | **string** | Name of the recovery entity | [optional] [default to null]
**Entity** | [***IdName**](IdName.md) |  | [optional] [default to null]
**RecoveryGroup** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Workload** | [***SolutionTypes**](SolutionTypes.md) |  | [optional] [default to null]
**Client** | [***IdName**](IdName.md) |  | [optional] [default to null]
**BackupSet** | [***IdName**](IdName.md) |  | [optional] [default to null]
**VmGroup** | [***IdName**](IdName.md) |  | [optional] [default to null]
**VirtualMachine** | [***NameGuid**](NameGUID.md) |  | [optional] [default to null]
**RecoveryConfiguration** | [***RecoveryConfiguration**](RecoveryConfiguration.md) |  | [optional] [default to null]
**Instance** | [***IdName**](IdName.md) |  | [optional] [default to null]
**PreRecoveryActions** | [**[]DrOperationScript**](DROperationScript.md) |  | [optional] [default to null]
**PostRecoveryActions** | [**[]DrOperationScript**](DROperationScript.md) |  | [optional] [default to null]
**Target** | [***IdName**](IdName.md) |  | [optional] [default to null]
**RecoveryPointDetails** | [***RecoveryEntityRpDetails**](RecoveryEntityRPDetails.md) |  | [optional] [default to null]
**RecoveryPoint** | **int32** | Timestamp for entity restore in case of disaster, default value is 0 as latest recovery point | [optional] [default to 0]
**ExecutionOrder** | [***RecoveryEntityExecutionInfo**](RecoveryEntityExecutionInfo.md) |  | [optional] [default to null]
**ConnectionDetails** | [***VmConnectionDetails**](VMConnectionDetails.md) |  | [optional] [default to null]
**BlockSysrepForRecovery** | **bool** |  | [optional] [default to null]
**RecoveryStatus** | [***RecoveryStatus**](RecoveryStatus.md) |  | [optional] [default to null]
**ValidationStatus** | [***ValidationStatus**](ValidationStatus.md) |  | [optional] [default to null]
**ValidationResults** | [**[]ValidationResult**](ValidationResult.md) |  | [optional] [default to null]
**RecoveryFailureReason** | **string** | Reason for last recovery job failure | [optional] [default to null]
**ValidationFailureReason** | **string** | Reason for last validation failure | [optional] [default to null]
**RecoveryStatusNotReadyReason** | **string** | Recovery status not ready reason | [optional] [default to null]
**RecoveryStatusNotReadyCategory** | [***RecoveryStatusNotReadyCategory**](RecoveryStatusNotReadyCategory.md) |  | [optional] [default to null]
**LastRecoveryJobId** | **int32** | Last recovery job Id of the entity | [optional] [default to null]
**OsType** | [***OsType**](OSType.md) |  | [optional] [default to null]
**CopyAvailableTime** | **int32** | Timestamp of the latest backup job start time present in the copy that would be used for recovery. Value would be -1 if no valid copy is present for recovery. | [optional] [default to null]
**InstalledWorkloads** | [**[]SolutionTypes**](SolutionTypes.md) |  | [optional] [default to null]
**SourceVendor** | [***VsVendor**](VSVendor.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

