# InstanceSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Gives the id of the instance. | [optional] [default to null]
**Name** | **string** | Gives the name of the instance. | [optional] [default to null]
**Server** | **string** | Gives the server to which the instance belongs to. | [optional] [default to null]
**Client** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Vendor** | [***Vendors**](Vendors.md) |  | [optional] [default to null]
**DatabaseEngine** | [***DatabaseEngine**](DatabaseEngine.md) |  | [optional] [default to null]
**Status** | [***InstanceStatus**](InstanceStatus.md) |  | [optional] [default to null]
**NotReadyReason** | **string** | If the instance isn&#x27;t ready, this provides the reason as to why the instance isn&#x27;t ready. | [optional] [default to null]
**LastBackup** | [***LastBackupJobInfo**](LastBackupJobInfo.md) |  | [optional] [default to null]
**ApplicationSize** | **int32** | Gives the application size of the instance. It is returned in bytes. | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**SLAStatus** | **string** |  | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Commcell** | [***CommcellInfo**](CommcellInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

