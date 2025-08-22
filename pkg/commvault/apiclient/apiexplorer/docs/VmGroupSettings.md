# VmGroupSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseChangedBlockTrackingOnVM** | **bool** | True if Changed Block Tracking is enabled | [optional] [default to true]
**CustomSnapshotResourceGroup** | **string** | Custom snapshot resource group name for Azure | [optional] [default to null]
**CustomSnapshotTags** | [**[]ResourceTag**](resourceTag.md) | represents custom tags to be set on snapshots | [optional] [default to null]
**RegionalSnapshot** | **bool** | True when snapshot storage location is regional | [optional] [default to true]
**AutoDetectVMOwner** | **bool** | True if auto detect VM Owner enabled | [optional] [default to null]
**AllowEmptySubclient** | **bool** | True if empty subclient is allowed | [optional] [default to null]
**NoOfReaders** | **int32** | Number of readers for backup | [optional] [default to 5]
**IsVMGroupDiskFiltersIncluded** | **bool** | Is VM group disk filters included in VM instance disk filters | [optional] [default to null]
**VmBackupType** | [***VmBackupType**](vmBackupType.md) |  | [optional] [default to null]
**IsApplicationAware** | **bool** | Is the VM App Aware | [optional] [default to null]
**GuestCredentials** | [***GuestCredentialInfo**](guestCredentialInfo.md) |  | [optional] [default to null]
**UseVMCheckpointSetting** | **bool** | True if use VM CheckPoint setting is enabled | [optional] [default to false]
**TransportMode** | [***TransportMode**](transportMode.md) |  | [optional] [default to null]
**DatastoreFreespaceCheck** | **bool** | True if Datastore Free space check is enabled | [optional] [default to true]
**DatastoreFreespaceRequired** | **int32** | precentage of datastore free space check value | [optional] [default to 10]
**CollectFileDetailsforGranularRecovery** | **bool** | True if metadata collection is enabled. Only applicable for Indexing v1 | [optional] [default to false]
**CollectFileDetailsFromSnapshotCopy** | **bool** | True if metadata collection is enabled for intellisnap jobs. Only applicable for Indexing v1 | [optional] [default to false]
**JobStartTime** | **int32** | Start Time for the VM Group Job | [optional] [default to null]
**CrossAccount** | [***AmazonCrossAccount**](AmazonCrossAccount.md) |  | [optional] [default to null]
**QuiesceGuestFileSystemAndApplications** | **bool** | True if VM backup type is File system and application consistent | [optional] [default to null]
**UseBackupsetDiskFilters** | **bool** | True if use of backupset disk filters is allowed | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

