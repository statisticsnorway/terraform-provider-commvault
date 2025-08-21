# PlanBackupDestinationGeneralInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Storage** | [***StoragePool**](StoragePool.md) |  | [optional] [default to null]
**StorageType** | [***StorageType**](StorageType.md) |  | [optional] [default to null]
**Source** | [***IdName**](IdName.md) |  | [optional] [default to null]
**NetAppCloudTarget** | **bool** | Only for snap copy. Tells if the snap copy supports SVM Mapping to NetApp cloud targets only. | [optional] [default to null]
**IsImmutableSnapCopy** | **bool** | Only for snap copy. Tells if the snap copy has immutable snap option enabled. | [optional] [default to null]
**IsDefault** | **bool** | Is this a default backup destination? | [optional] [default to null]
**IsActive** | **bool** | Is this an active backup destination? | [optional] [default to null]
**IsCopyInMaintenanceMode** | **bool** | Is this copy in Maintenance mode? | [optional] [default to null]
**IsCopyPromotionRequestSubmitted** | **bool** | Is copy promotion request submitted for this copy? | [optional] [default to null]
**ComplianceLock** | **bool** | Is compliance lock enabled on backup destination? | [optional] [default to null]
**IsSnapCopy** | **bool** | Is this a snap copy? | [optional] [default to null]
**IsMirrorCopy** | **bool** | Is this a mirror copy? | [optional] [default to null]
**IsAirgapCopy** | **bool** | Set to true if this copy uses Airgap storage | [optional] [default to null]
**CopyType** | [***PlanCopyType**](PlanCopyType.md) |  | [optional] [default to null]
**SnapCopyType** | [***PlanCopyTypeName**](PlanCopyTypeName.md) |  | [optional] [default to null]
**CopyPrecedence** | **int32** | Order of backup destinaion copy created in storage policy | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

