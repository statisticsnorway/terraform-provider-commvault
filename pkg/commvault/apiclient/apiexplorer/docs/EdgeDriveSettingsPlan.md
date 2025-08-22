# EdgeDriveSettingsPlan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateNewIndexServer** | **bool** | If set to false, uses existing edge index server and requires IndexServer IdName to be passed. If set to true, requires client IdName and indexCachePath to create a new Index Server. | [default to null]
**IndexServer** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Client** | [***IdName**](IdName.md) |  | [optional] [default to null]
**IndexCachePath** | **string** |  | [optional] [default to null]
**AuditDriveOperations** | **bool** | Enable auditing which logs the activities based on user interaction like creating, editing, moving, renaming, downloading or deleting files. | [optional] [default to true]
**NotificationsForShares** | **bool** | Enables alert notification feature which allows the share user or share owner to subscribe for share notifications when any activities are performed on the Edge Drive or the Collaborative share. The user can receive the notifications on the Web Console or as an email notification. | [optional] [default to true]
**EdgeDriveQuota** | **int32** | Maximum number of gigabytes that you can store in the Edge Drive. Giving value as -1 means no quota. | [optional] [default to -1]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

