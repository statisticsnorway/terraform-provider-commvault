# KubernetesApplicationItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NameSpace** | **string** | Specifies the Namespace of the Application | [optional] [default to null]
**Kind** | **string** | Specifies the resource Kind of the Application | [optional] [default to null]
**ApplicationGroup** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Backupset** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Commcell** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Cluster** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Instance** | [***IdName**](IdName.md) |  | [optional] [default to null]
**UsedSize** | **int32** | Describes the Total used space of all volumes of the Application | [optional] [default to null]
**ProvisionedSize** | **int32** | Describes the Total provisioned size of the volumes of the Application | [optional] [default to null]
**ApplicationSize** | **int32** | Describes the application size of the Application | [optional] [default to null]
**LastBackup** | [***LastBackup**](lastBackup.md) |  | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Tags** | [**[]IdNameValue**](IdNameValue.md) | List of Tags associated to the Kubernetes Application | [optional] [default to null]
**SLA** | [***KubernetesSlaDetails**](KubernetesSLADetails.md) |  | [optional] [default to null]
**GUID** | **string** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

