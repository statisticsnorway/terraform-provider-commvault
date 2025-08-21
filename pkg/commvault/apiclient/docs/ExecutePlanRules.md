# ExecutePlanRules

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subclients** | [**[]IdName**](IdName.md) | List of subclients against which plan rules need to be  executed and evaluated. | [optional] [default to null]
**Backupsets** | [**[]IdName**](IdName.md) | List of backupsets against which plan rules need to be  executed and evaluated. | [optional] [default to null]
**Instances** | [**[]IdName**](IdName.md) | List of instances against which plan rules need to be  executed and evaluated. | [optional] [default to null]
**Clients** | [**[]IdName**](IdName.md) | List of clients against which plan rules need to be  executed and evaluated. | [optional] [default to null]
**IgnorePreviousPlanAssociation** | **bool** | Boolean to indicate if we want to evaluate rule for entities associated to plan. Default is false. Only entities with no plan associated will be evaluated. | [optional] [default to null]
**IsPreviewOnly** | **bool** | Boolean to indicate if request is to preview list of subclients that will be associated via plan rules | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

