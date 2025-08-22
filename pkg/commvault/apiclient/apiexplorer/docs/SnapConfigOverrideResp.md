# SnapConfigOverrideResp

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssocType** | **int64** | Association type, at subclient level it is 9, at client level it is 8, at copy level it is 6 and by default at array level it is 3 | [optional] [default to null]
**ClientId** | **int64** | Client Id | [optional] [default to null]
**SubclientId** | **int64** | Subclient id | [optional] [default to null]
**CopyId** | **int64** | snap copy id | [optional] [default to null]
**SnapConfigurations** | [**[]SnapConfigsOverride**](SnapConfigsOverride.md) | Snap Config options that can be overridden | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

