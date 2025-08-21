# VmAppValidation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidateVMBackups** | **bool** | True if VM Backup validation is enabled | [optional] [default to null]
**RecoveryTarget** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Schedule** | [***ValidationScheduleObject**](ValidationScheduleObject.md) |  | [optional] [default to null]
**UseSourceVmESXToMount** | **bool** | Use Source VM ESX To Mount | [optional] [default to null]
**KeepValidatedVMsRunning** | **bool** | If true then validated VMs will be available until expiration time set on the recovery target | [optional] [default to null]
**MaximumNoOfThreads** | **int32** | Number of backup Validation Threads | [optional] [default to null]
**CustomValidationScript** | [***AppValidationScript**](appValidationScript.md) |  | [optional] [default to null]
**GuestCredentials** | [***GuestCredentialInfo**](guestCredentialInfo.md) |  | [optional] [default to null]
**Copy** | [***AppValidationSourceCopy**](AppValidationSourceCopy.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

