# CreateRole

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the new role | [default to null]
**PermissionList** | [**[]Permissions**](Permissions.md) | Used to provide the list of permissions associated with the role. | [default to null]
**Enabled** | **bool** | Used to determine if the role is enabled or disabled. If not provided, role will be enabled by default. | [optional] [default to null]
**VisibleToAll** | **bool** | Determines if the role is visible to everyone. if not provided, it will be set to false by default. | [optional] [default to null]
**GlobalConfigInfo** | [***CreateGlobalConfigInfo**](CreateGlobalConfigInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

