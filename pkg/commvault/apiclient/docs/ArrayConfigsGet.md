# ArrayConfigsGet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterConfigId** | **int32** | This is the masterConfigId, which is available for each vendors configs | [optional] [default to null]
**Name** | **string** | This is the name of the config which is displayed on the Command Center Console | [optional] [default to null]
**Type_** | **int32** | Type of the config value. type can accept 7 values [1,2,8,10,12,13,14], 1: boolean, 2: integer, 8: text, 10: array[string] are common configs, the rest three are special keys, 12 is for password type key for NetApp E-Series and HPE Nimble, 13 is Private Key for GCP, 14 is a config to select type of Disk for GCP | [optional] [default to null]
**Value** | **string** | Values to be set for all types except 10 and 14 | [optional] [default to null]
**Values** | [**[]IdName**](IdName.md) | Values to be set for type 10 and 14. Incase of HPE 3PAR StoreServ, name field will have the remote snap MA Id. | [optional] [default to null]
**Flags** | **int32** | Flag regarding placement of config in the CC page | [optional] [default to null]
**Description** | **string** | Description about the config that tells the user what it is for and the range of values it accepts | [optional] [default to null]
**IsEnabled** | **bool** | Whether the config is enabled or not | [optional] [default to null]
**CheckRange** | **bool** | Whether to check if the value is in range or not of accepted values | [optional] [default to null]
**MinValue** | **int64** | This is the lower limit value to which the config can be set for numeric types | [optional] [default to null]
**MaxValue** | **int64** | This is the higher limit value to which the config can be set for numeric types | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

