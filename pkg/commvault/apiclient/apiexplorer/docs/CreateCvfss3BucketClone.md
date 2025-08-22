# CreateCvfss3BucketClone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTime** | **int32** | Clone objects in the bucket created on or after the specified time in unix format. | [optional] [default to null]
**ToTime** | **int32** | Clone objects in the bucket created on or before the specified time in unix format. | [default to null]
**CopyPrecedence** | **int32** | The copy precedence to use for the clone operation | [default to null]
**ReservationTime** | **int32** | Clone reservation time in seconds. The clone will be automatically deleted after the specified duration. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

