# KubernetesApplicationGroupItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID of the Kubernetes Application group | [optional] [default to null]
**Name** | **string** | Name of the Kubernetes Application Group | [optional] [default to null]
**IsDefaultApplicationGroup** | **bool** | Specifies whether the Application Group is a default Application Group | [optional] [default to null]
**SnapBackupEnabled** | **bool** | Specifies if snap backup is enabled | [optional] [default to null]
**Cluster** | [***IdNameDisplayName**](IdNameDisplayName.md) |  | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Company** | [***CompanyInfo**](companyInfo.md) |  | [optional] [default to null]
**Tags** | [**[]IdNameValue**](IdNameValue.md) | List of Tags associated to the Kubernetes Application Group | [optional] [default to null]
**Instance** | [***IdName**](IdName.md) |  | [optional] [default to null]
**LastBackup** | [***LastBackup**](lastBackup.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

