# UpdateCompany

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [***CompanyStatus**](CompanyStatus.md) |  | [optional] [default to null]
**NewName** | **string** | Used to change the name of a company | [optional] [default to null]
**General** | [***General**](General.md) |  | [optional] [default to null]
**Security** | [**[]UpdateSecurityAssoc**](UpdateSecurityAssoc.md) |  | [optional] [default to null]
**EmailSettings** | [***EmailSettings**](EmailSettings.md) |  | [optional] [default to null]
**Sites** | [***Sites**](Sites.md) |  | [optional] [default to null]
**TenantOperators** | [**[]TenantOperator**](TenantOperator.md) |  | [optional] [default to null]
**Plans** | [**[]IdName**](IdName.md) | Provide a list of data protection plans to use for the company. The plans that are provided are the plans that the tenant administrator can choose from. | [optional] [default to null]
**DefaultPlans** | [**[]DefaultPlan**](DefaultPlan.md) | Refers to default data protection plans to use for the company. | [optional] [default to null]
**FileExceptions** | [***FileExceptions**](FileExceptions.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

