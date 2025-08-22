# CompanySummary

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Company Id | [optional] [default to null]
**Name** | **string** | Company name | [optional] [default to null]
**GUID** | **string** | GUID of the company | [optional] [default to null]
**Alias** | **string** | The company domain or NetBIOS name | [optional] [default to null]
**IsReseller** | **bool** | Enable reseller mode. A reseller is a user or user groups who can operate multiple tenant environments. The service provider can assign one or more resellers as the tenant operator for a company. The reseller can switch to any of their assigned tenant environment and operate on the company as a tenant user. Once enabled, the reseller mode cannot be disabled. | [optional] [default to null]
**PrimaryContacts** | [**[]IdName**](IdName.md) |  | [optional] [default to null]
**AssociatedEntitiesCount** | **int32** | Gives the number of entities associated with the company | [optional] [default to null]
**Status** | **string** | Gives the company status | [optional] [default to null]
**Commcell** | [***CommcellNameDisplayNameInfo**](CommcellNameDisplayNameInfo.md) |  | [optional] [default to null]
**Operators** | [**[]CompanyOperator**](CompanyOperator.md) | List of operators configured for the company | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

