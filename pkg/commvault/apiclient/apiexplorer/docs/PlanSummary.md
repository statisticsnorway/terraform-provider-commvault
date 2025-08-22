# PlanSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**PlanType** | [***PlanType**](PlanType.md) |  | [optional] [default to null]
**ParentPlan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**TargetApp** | [**[]PlanTargetApp**](PlanTargetApp.md) | This tells the content indexing target apps for data classification plans | [optional] [default to null]
**AssociatedEntities** | **int32** | Count of associated entities for the plan | [optional] [default to null]
**RPO** | **int32** | RPO in minutes for the plan | [optional] [default to null]
**NumberOfCopies** | **int32** | Number of copies for the plan | [optional] [default to null]
**Status** | [***PlanStatus**](PlanStatus.md) |  | [optional] [default to null]
**MissingEntities** | [**[]IdName**](IdName.md) | For plans in incomplete state contains list of entities missing from the plan | [optional] [default to null]
**ResourcePool** | [**[]IdName**](IdName.md) | This lists the various resource Pools of different app types associated with storage pools of plan | [optional] [default to null]
**Tags** | [**[]IdNameValue**](IdNameValue.md) | Tags associated with this plan | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Commcell** | **string** | Tells which commcell this plan belongs to. | [optional] [default to null]
**Derivable** | **bool** | Tells if this plan can be used to derive from and create a new child plan | [optional] [default to null]
**GlobalConfigInfo** | [***GlobalConfigInfo**](GlobalConfigInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

