# ValidateVmRestoreReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromTime** | **string** | Restore window UTC from time. Valid Formats: yyyy-MM-ddTHH:mm:ss or yyyy-MM-ddTHH:mm:sszzz or yyyy-MM-dd | [optional] [default to null]
**ToTime** | **string** | Restore window UTC to time. Valid Formats: yyyy-MM-ddTHH:mm:ss or yyyy-MM-ddTHH:mm:sszzz or yyyy-MM-dd | [optional] [default to null]
**BackupSource** | **int32** | Backup source information from where you want to restore. 0 (Automatic), 1 (Snap Copy), 2 (Primary), 3 to N (Aux Copy) | [optional] [default to 0]
**AccessNode** | [***IdName**](IdName.md) |  | [optional] [default to null]
**AccessNodeGroup** | [***IdName**](IdName.md) |  | [optional] [default to null]
**NotifyOnJobCompletion** | **bool** | Enable email notificaiton for job status. Defaults to false | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

