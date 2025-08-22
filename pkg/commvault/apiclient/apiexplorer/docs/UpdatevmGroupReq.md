# UpdatevmGroupReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | **string** | subclient name  | [optional] [default to null]
**Content** | [***VmContent**](vmContent.md) |  | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Storage** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Filters** | [***VmContent**](vmContent.md) |  | [optional] [default to null]
**DiskFilters** | [***VmDiskFilterProp**](vmDiskFilterProp.md) |  | [optional] [default to null]
**SecurityAssociations** | [**[]SecurityAssoc**](SecurityAssoc.md) |  | [optional] [default to null]
**Settings** | [***VmGroupSettings**](vmGroupSettings.md) |  | [optional] [default to null]
**SnapshotManagement** | [***SnapCopyInfo**](snapCopyInfo.md) |  | [optional] [default to null]
**ActivityControl** | [***ActivityControlOptions**](ActivityControlOptions.md) |  | [optional] [default to null]
**ApplicationValidation** | [***VmAppValidation**](vmAppValidation.md) |  | [optional] [default to null]
**EnableFileIndexing** | **bool** | True if file indexing needs to be enabled | [optional] [default to false]
**TimeZone** | [***IdName**](IdName.md) |  | [optional] [default to null]
**AccessNode** | [**[]AccessNodeModel**](accessNodeModel.md) |  | [optional] [default to null]
**MeditechSystems** | [***MeditechPropResp**](meditechPropResp.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

