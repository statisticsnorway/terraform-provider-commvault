# RmRequestSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Request id | [optional] [default to null]
**Name** | **string** | Request name | [optional] [default to null]
**Type_** | [***RmRequestType**](RMRequestType.md) |  | [optional] [default to null]
**Status** | [***RmRequestStatus**](RMRequestStatus.md) |  | [optional] [default to null]
**Requestor** | **string** | Email of the requestor | [optional] [default to null]
**Application** | [***ActivateApplication**](ActivateApplication.md) |  | [optional] [default to null]
**CreatedOn** | **int64** | Request creation time in unix epoch format | [optional] [default to null]
**Owner** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Reviewers** | [**[]IdName**](IdName.md) | List of request reviewers | [optional] [default to null]
**Approvers** | [**[]IdName**](IdName.md) | List of the request approvers | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

