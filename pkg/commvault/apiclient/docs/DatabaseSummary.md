# DatabaseSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **int32** | Application Id of database | [optional] [default to null]
**Name** | **string** | This gives the name of the database. | [optional] [default to null]
**Instance** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Server** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Vendor** | [***Vendors**](Vendors.md) |  | [optional] [default to null]
**DatabaseEngine** | [***DatabaseEngine**](DatabaseEngine.md) |  | [optional] [default to null]
**RecoveryModel** | **string** | SQL recovery model is a database configuration option that determines the type of backup that one could perform, and provides the ability to restore the data or recover it from a failure. | [optional] [default to null]
**SLAStatus** | **string** | SLA status for last backup of database | [optional] [default to null]
**LastBackupSize** | **int32** | Size of last backup of database in bytes | [optional] [default to null]
**LastBackupTime** | **int32** | Time of last backup of database. It is given in UNIX time. | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Commcell** | [***CommcellInfo**](CommcellInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

