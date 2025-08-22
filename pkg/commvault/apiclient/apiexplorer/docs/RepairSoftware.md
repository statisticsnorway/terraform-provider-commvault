# RepairSoftware

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RebootIfRequired** | **bool** | Boolean which determines whether to reboot if required. | [optional] [default to false]
**NotifyWhenJobCompletes** | **bool** | Boolean which determines whether to notify when software is repaired. | [optional] [default to false]
**Username** | **string** | If using system credentials, provide the user name | [optional] [default to null]
**Password** | **string** | Corresponding password of the user | [optional] [default to null]
**Entities** | [**[]IdNameGuidType**](IdNameGUIDType.md) | List of all the client and client groups on which repair software should be run | [default to null]
**InvokedFrom** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

