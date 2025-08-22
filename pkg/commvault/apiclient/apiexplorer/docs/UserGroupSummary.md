# UserGroupSummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | User group id | [optional] [default to null]
**Name** | **string** | user group name | [optional] [default to null]
**GUID** | **string** | globally unique identifier for user group | [optional] [default to null]
**ServiceType** | [***ServiceTypes**](ServiceTypes.md) |  | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**Plan** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Provider** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Enabled** | **bool** | returns if the user group is enabled or disabled | [optional] [default to null]
**AllowMultipleCompanyMembers** | **bool** | This property denotes that addition of users/groups from child companies is allowed. Only applicable for commcell and reseller company group. | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Commcell** | [***CommcellInfo**](CommcellInfo.md) |  | [optional] [default to null]
**Users** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**GlobalConfigInfo** | [***GlobalConfigInfo**](GlobalConfigInfo.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

