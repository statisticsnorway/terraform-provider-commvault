# CreateCompany

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name of the company to be created. | [default to null]
**Email** | **string** | Email address for the tenant administrator. If provided, contactName for the tenant administrator also needs to be provided | [optional] [default to null]
**ContactName** | **string** | Name of the tenant administrator. If provided, email also needs to be provided. | [optional] [default to null]
**Plans** | [**[]IdName**](IdName.md) | Select data protection plans to use for the company. The plans that are selected are the plans that the tenant administrator can choose from. | [optional] [default to null]
**Alias** | **string** | The company domain or NetBIOS name | [default to null]
**EmailSuffix** | **string** | Supported domains for the company | [optional] [default to null]
**ServiceCommcells** | [**[]IdName**](IdName.md) | Used to add service commcells to the master commcell. Either id or name can be provided. If both are provided, id will be taken into consideration. | [optional] [default to null]
**SendWelcomeEmail** | **bool** | send a welcome email on company creation to the tenant administrator. | [optional] [default to null]
**PrimaryDomain** | **string** | The primary domain name of the company being created. Can be added only if an external domain is already present. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

