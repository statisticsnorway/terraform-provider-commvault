# ApplicationGroupSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Status** | [***Vmslastatus**](VMSLASTATUS.md) |  | [optional] [default to null]
**ClusterType** | [***VsVendor**](VSVendor.md) |  | [optional] [default to null]
**ApplicationGroup** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Cluster** | [***IdNameDisplayName**](IdNameDisplayName.md) |  | [optional] [default to null]
**Instance** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**LastBackup** | [***LastBackupJobInfo**](LastBackupJobInfo.md) |  | [optional] [default to null]
**Company** | [***CompanyInfo**](companyInfo.md) |  | [optional] [default to null]
**IsDefaultApplicationGroup** | **bool** | True if subclient is default subclient | [optional] [default to false]
**StoragePolicyRetentionExtension** | **bool** | True if extend storage policy Retention was set in VMGroup for deleted files | [optional] [default to false]
**IndexingStatus** | [***IndexingStatusType**](IndexingStatusType.md) |  | [optional] [default to null]
**SnapBackupEnabled** | **bool** |  true if snap backup enabled | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

