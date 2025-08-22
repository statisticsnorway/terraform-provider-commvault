# LaptopPlanStorageAndSchedule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryStorage** | [***IdName**](IdName.md) |  | [default to null]
**SecondaryStorage** | [***IdName**](IdName.md) |  | [optional] [default to null]
**BackupFrequency** | **int32** | Recovery Point Objective (RPO) is the maximum amount of time that data can be lost during a service disruption. Your RPO determines the frequency of your backup jobs. Your RPO is met through automatic options. The time specified in minutes here is your RPO if none of the automatic options are met. Default is 480 minutes (8 hours). | [optional] [default to 480]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

