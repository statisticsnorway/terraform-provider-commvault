# GlobalSettingsItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the setting | [optional] [default to null]
**DisplayLabel** | **string** | Display Label of the setting | [optional] [default to null]
**Description** | **string** | Description of the setting | [optional] [default to null]
**AcceptableValues** | **[]string** | Specifies values which are acceptable when modifying the setting value. This will only be returned if there are specific set of values which can be accepted. | [optional] [default to null]
**DefaultValue** | **string** | Value used to get default system behaviour | [optional] [default to null]
**Value** | **string** | Value of the setting if the setting was modified. | [optional] [default to null]
**Comment** | **string** | Comment set by the user while modifying the setting | [optional] [default to null]
**IsModified** | **bool** | Specifies if the setting is already modified by the user | [optional] [default to null]
**Type_** | **string** | Type of the setting | [optional] [default to null]
**IsRestartRequired** | **bool** | Specifies if it is required to restart the services for any changes to take effect | [optional] [default to null]
**MinValue** | **int32** | Specifies minimum value that can be set to modify specified global setting. It is only returned if the setting type is Integer | [optional] [default to null]
**MaxValue** | **int32** | Specifies maximum value that can be set to modify specified global setting. It is only returned if the setting type is Integer | [optional] [default to null]
**Category** | **string** | Category to which setting belongs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

