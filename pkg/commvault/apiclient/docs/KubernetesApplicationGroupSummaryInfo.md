# KubernetesApplicationGroupSummaryInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextBackupTime** | **int32** |  | [optional] [default to null]
**IsETCDApplicationGroup** | **bool** | Describes if the Application Group is an ETCD Application Group | [optional] [default to null]
**IsDefaultApplicationGroup** | **bool** | Describes if the Application Group is a default Application Group | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Timezone** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Tags** | [**[]IdNameValue**](IdNameValue.md) |  | [optional] [default to null]
**Content** | [***KubernetesApplicationGroupContentItem**](KubernetesApplicationGroupContentItem.md) |  | [optional] [default to null]
**Filters** | [***KubernetesApplicationGroupFilterItem**](KubernetesApplicationGroupFilterItem.md) |  | [optional] [default to null]
**LastBackup** | [***LastBackupJobInfo**](LastBackupJobInfo.md) |  | [optional] [default to null]
**ActivityControl** | [***ApplicationGroupActivityControl**](ApplicationGroupActivityControl.md) |  | [optional] [default to null]
**AccessNodes** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**Options** | [***ApplicationGroupGetOptions**](ApplicationGroupGetOptions.md) |  | [optional] [default to null]
**Cluster** | [***IdNameDisplayName**](IdNameDisplayName.md) |  | [optional] [default to null]
**Instance** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Backupset** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

