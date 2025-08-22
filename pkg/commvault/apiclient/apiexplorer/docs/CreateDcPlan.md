# CreateDcPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of Data Classification Plan | [optional] [default to null]
**ThreatAnalysis** | **bool** | Enables\\Disables Threat Analysis support for DC Plan | [optional] [default to null]
**Application** | [***PlanTargetApp**](PlanTargetApp.md) |  | [optional] [default to null]
**IndexServer** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ContentAnalyzer** | [**[]IdName**](IdName.md) | Content Analyzer Id&#x60;s for Entity Detection | [optional] [default to null]
**ContentIndexing** | [***DcPlanCIpolicy**](DCPlanCIpolicy.md) |  | [optional] [default to null]
**EntityDetection** | [***DcPlanEePolicy**](DCPlanEEPolicy.md) |  | [optional] [default to null]
**Schedule** | [***DcPlanJobSchedule**](DCPlanJobSchedule.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

