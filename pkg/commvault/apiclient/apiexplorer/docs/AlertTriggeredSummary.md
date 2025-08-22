# AlertTriggeredSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Triggered Alert Id | [optional] [default to null]
**Severity** | [***AlertSeverity**](AlertSeverity.md) |  | [optional] [default to null]
**DetectedCriterion** | **string** | detection criteria for the triggered alert to be generated | [optional] [default to null]
**Info** | **string** | Name of the Alert Triggered | [optional] [default to null]
**Notes** | **string** | contains any descriptive note written for the alert | [optional] [default to null]
**Type_** | **string** | alert type for which the triggered alert was generated | [optional] [default to null]
**DetectedTime** | **int32** | Unix Epoch Time | [optional] [default to null]
**Client** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ReadStatus** | **bool** | Gives the Alert Read Status. True if Read and False if Unread | [optional] [default to null]
**PinStatus** | **bool** | Gives the Alert Pin Status. True if pinned and False if not pinned. | [optional] [default to null]
**JobId** | **int32** | Job Id by which this Alert was Triggered | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Commcell** | [***CommcellNameDisplayNameInfo**](CommcellNameDisplayNameInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

