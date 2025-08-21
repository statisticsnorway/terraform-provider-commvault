# UpdateRole

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewName** | **string** | new name of the role | [optional] [default to null]
**PermissionList** | [**[]Permissions**](Permissions.md) | Used to update the list of permissions associated with the role. | [optional] [default to null]
**PermissionOperationType** | **string** | Type of operation to be performed on the permissionList. Default is OVERWRITE. | [optional] [default to null]
**Enabled** | **bool** | Used to determine if the role is enabled or disabled. | [optional] [default to null]
**VisibleToAll** | **bool** | Determines if the role is visible to everyone. | [optional] [default to null]
**Security** | [**[]UpdateSecurityAssoc**](UpdateSecurityAssoc.md) | Used to update the security association for the role | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

