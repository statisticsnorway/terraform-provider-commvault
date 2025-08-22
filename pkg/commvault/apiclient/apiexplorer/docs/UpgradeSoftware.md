# UpgradeSoftware

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RebootIfRequired** | **bool** | Boolean which determines whether to reboot if required. | [optional] [default to null]
**RunDBMaintenance** | **bool** | Boolean which determines whether or not to run a database maintenance. | [optional] [default to null]
**NotifyWhenJobCompletes** | **bool** | Boolean which determines whether to notify when software is upgraded. | [optional] [default to null]
**InstallOSUpdates** | **bool** | Boolean which determines whether to install operating system updates. | [optional] [default to null]
**InstallStorageUpdates** | [***StorageUpdateType**](StorageUpdateType.md) |  | [optional] [default to null]
**WaitForDownloadJobToComplete** | **bool** | Boolean which determines whether to wait for Download software job to complete before software is upgraded. | [optional] [default to null]
**Entities** | [**[]IdNameGuidType**](IdNameGUIDType.md) | List of all the client and client groups on which UpgradeSoftware should be run | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

