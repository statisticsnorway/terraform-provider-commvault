# DcPlanDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**ThreatAnalysis** | **bool** | Whether Threat Analysis is enabled for plan | [optional] [default to null]
**TargetApplication** | [***PlanTargetApp**](PlanTargetApp.md) |  | [optional] [default to null]
**ContentAnalyzer** | [**[]IdName**](IdName.md) | Content Analyzer details for Data Classification plan | [optional] [default to null]
**IndexServer** | [***IdName**](IdName.md) |  | [optional] [default to null]
**EntityDetection** | [***DcPlanEePolicy**](DCPlanEEPolicy.md) |  | [optional] [default to null]
**ContentIndexing** | [***GetDcPlanCiPolicy**](GetDCPlanCIPolicy.md) |  | [optional] [default to null]
**Permissions** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**AssociatedEntities** | [**[]IdNameCount**](IdNameCount.md) |  | [optional] [default to null]
**Schedule** | [***DcPlanJobSchedule**](DCPlanJobSchedule.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

