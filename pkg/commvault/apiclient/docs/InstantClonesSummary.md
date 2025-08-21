# InstantClonesSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneJobId** | **int32** | Job id for clone operation | [optional] [default to null]
**SqlRecoveryId** | **int32** | Id of sql recovery point for SQL Agent | [optional] [default to null]
**SourceClient** | [***IdName**](IdName.md) |  | [optional] [default to null]
**SourceInstance** | [***IdName**](IdName.md) |  | [optional] [default to null]
**MountedHost** | [***IdName**](IdName.md) |  | [optional] [default to null]
**TargetInstance** | [***IdName**](IdName.md) |  | [optional] [default to null]
**DatabaseAsOfTime** | **int32** | The Point in time up to which the database has been recovered using the backups. The time is provided in unix time format. | [optional] [default to null]
**CreationTime** | **int32** | The time when the database clone was created. The time is provided in unix time format. | [optional] [default to null]
**ExpirationDate** | **int32** | The time till which the clone is kept active and will get cleaned up automatically after that. The time is provided in unix time format. | [optional] [default to null]
**Status** | **string** | Status of instant clone | [optional] [default to null]
**Commcell** | [***CommcellInfo**](CommcellInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

