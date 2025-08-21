# PlanSourceCopy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDefault** | **bool** | Is this a default backup destination? | [optional] [default to null]
**IsActive** | **bool** | Is this an active backup destination? | [optional] [default to null]
**IsSnapCopy** | **bool** | Is this a snap copy? | [optional] [default to null]
**IsMirrorCopy** | **bool** | Is this a mirror copy? | [optional] [default to null]
**CopyType** | [***PlanCopyType**](PlanCopyType.md) |  | [optional] [default to null]
**SnapCopyType** | [***PlanCopyTypeName**](PlanCopyTypeName.md) |  | [optional] [default to null]
**CopyPrecedence** | **int32** | Order of backup destinaion copy created in storage policy | [optional] [default to null]
**BackupDestination** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ArrayReplicaCopy** | **bool** | Is this an array replica copy? | [optional] [default to null]
**DefaultReplicaCopy** | **bool** | Is this a default replica copy? | [optional] [default to null]
**IsImmutableSnapCopy** | **bool** | Only for snap copy. Tells if the snap copy has immutable option enabled. | [optional] [default to null]
**IsSourceBackupCopy** | **bool** | Is this the source snap copy for backup copy operations? | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

