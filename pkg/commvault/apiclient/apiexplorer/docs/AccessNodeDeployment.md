# AccessNodeDeployment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Client name for vm, if not provided it will be same as vmName | [optional] [default to null]
**HostName** | **string** | Host name for vm, if not provided it will be same as vmName | [optional] [default to null]
**Os** | **string** |  | [optional] [default to null]
**DeploymentType** | **int32** | 0: Access node deployment, 1: File recovery enabler for linux deployment | [optional] [default to DEPLOYMENT_TYPE.0_]
**CommCell** | [***CommCellNameHostName**](CommCellNameHostName.md) |  | [optional] [default to null]
**VmName** | **string** |  | [default to null]
**UserCredentials** | [***UserCredentials**](userCredentials.md) |  | [default to null]
**VCenter** | **string** | vCenter server instance | [default to null]
**VmLocation** | [***VmLocation**](vmLocation.md) |  | [default to null]
**HardDiskType** | **int32** | 1: Thin Provision, 2: Thick provision lazy zeroed, 3: Thick provision eager zeroed | [optional] [default to HARD_DISK_TYPE.2_]
**Network** | [***NetworkDetails**](NetworkDetails.md) |  | [default to null]
**ClientGroup** | **string** |  | [optional] [default to null]
**NotifyUserOnJobCompletion** | **bool** |  | [optional] [default to false]
**AutomaticOSUpdates** | **bool** |  | [optional] [default to false]
**Timezone** | **string** | Linux OS qualified timezones | [optional] [default to America/New_York]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

