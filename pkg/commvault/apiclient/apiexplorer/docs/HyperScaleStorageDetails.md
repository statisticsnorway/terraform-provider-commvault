# HyperScaleStorageDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**General** | [***HyperScaleStorageGeneralInfo**](HyperScaleStorageGeneralInfo.md) |  | [optional] [default to null]
**Nodes** | [**[]IdNameStatus**](IdNameStatus.md) | List of disks with the Linux MediaAgent | [optional] [default to null]
**Encryption** | [***Encryption**](Encryption.md) |  | [optional] [default to null]
**Security** | [**[]SecurityAssoc**](SecurityAssoc.md) | List of users or user groups each having a specific set of roles that determine the kind of operations they can perform on hyperscale storage | [optional] [default to null]
**AssociatedPlans** | [**[]IdName**](IdName.md) | List of plans associated with this HyperScale storage | [optional] [default to null]
**Id** | **int32** |  | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

