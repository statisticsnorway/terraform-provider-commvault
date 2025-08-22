# ArrayConfigsEdit

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterConfigId** | **int32** | This is the masterConfigId, which is available for each vendor&#x27;s configs | [default to null]
**Name** | **string** | This is the name of the config which is displayed on the Command Center Console | [default to null]
**Type_** | **int32** | Type of the config value. type can accept 7 values [1,2,8,10,12,13,14], 1: boolean, 2: integer, 8: text, 10: array[string] are common configs, the rest three are special keys, 12 is for password type key for NetApp E-Series and HPE Nimble, 13 is Private Key for GCP, 14 is a hidden config to select type of Disk for GCP | [default to null]
**Value** | **string** | Takes a single value for all types and for type 14 it holds the id of the selected value from values | [optional] [default to null]
**Values** | [**[]IdName**](IdName.md) | For type 10 and 14, it can take one or more value of idname type where id by default is 0 for 10. For HPE 3PAR StoreServ, pass MA Id in name field to configure that MA as remote snap MA. | [optional] [default to null]
**Flags** | **int32** | Flag regarding placement of config in the CC page | [default to null]
**IsUpdated** | **bool** | Whether the config is updated/edited or not | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

