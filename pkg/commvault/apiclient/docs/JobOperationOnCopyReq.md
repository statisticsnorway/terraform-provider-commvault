# JobOperationOnCopyReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoragePolicyId** | **int32** | ID for chosen storage policy | [default to null]
**CopyId** | **int32** | ID for chosen copy | [default to null]
**OpType** | [***JobOperationType**](JobOperationType.md) |  | [default to null]
**JobIds** | **[]int32** | Comma separated Job IDs to run job operation on | [default to null]
**CommcellId** | **int32** | ID for chosen commcell | [default to null]
**RetainUntilTime** | **int64** | Job will be retained till specified time. The time is provided in unix time format. | [optional] [default to null]
**LoadDependentJobs** | **bool** | Defines whether dependent jobs need to be computed. | [optional] [default to false]
**LoadArchiverJobs** | **bool** | Defines whether archiver jobs need to be computed. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

