# ModifyAdditionalSetting

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the setting to be modified. This cannot be modified but is required to identify the setting for which the new value is to be set. | [default to null]
**Category** | **string** | Category of the setting to be modified. This cannot be modified but is required to identify the setting for which the new value is to be set. | [default to null]
**Type_** | **string** | Type of the setting to be modified. This cannot be modified but is required to identify the setting for which the new value is to be set. | [default to null]
**NewValue** | **string** | New value which will be set for the specified setting, there can be fixed acceptables values for some settings. To get more details about what are acceptable value for a setting, use GET GlobalSettings or GET EntitySettings. | [optional] [default to null]
**Comment** | **string** | Comment to specify why this value was set for the specified setting | [optional] [default to null]
**Reset** | **bool** | To reset the value of any already modified setting, set this to true to reset the value and regain default behaviour | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

