# SnapConfigsOverrideEdit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterConfigId** | **int32** | This is the masterConfigId, which is available for each vendors configs | [default to null]
**Id** | **int32** | This is the id of the config which is required during config override | [optional] [default to null]
**Name** | **string** | This is the name of the config which is displayed on the Command Center Console | [default to null]
**Type_** | **int32** | Type of the config value. type can accept 7 values [1,2,8,10,12,13,14], 1: boolean, 2: integer, 8: text, 10: array[string] are common configs, the rest three are special keys, 12 is for password type key for NetApp E-Series and HPE Nimble, 13 is Private Key for GCP, 14 is a hidden config to select type of Disk for GCP | [default to null]
**Value** | **string** | Value of the config | [optional] [default to null]
**Values** | [**[]IdName**](IdName.md) | Holds a single value for types except 10, for 10 it holds one or more values. For HPE 3PAR StoreServ, pass MA Id in name field to configure that MA as remote snap MA. | [optional] [default to null]
**Flags** | **int32** | Flag regarding placement of config in the CC page | [default to null]
**IsUpdated** | **bool** | Whether the config is updated or not | [optional] [default to null]
**IsOverridden** | **bool** | whether the config is overridden or not | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

