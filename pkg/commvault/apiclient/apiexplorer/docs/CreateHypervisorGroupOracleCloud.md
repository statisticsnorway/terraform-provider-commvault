# CreateHypervisorGroupOracleCloud

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorType** | **string** |  | [default to null]
**TenancyOCId** | **string** | OCID for the tenant. | [default to null]
**UserOCId** | **string** | OCID for the admin user for the hypervisor | [default to null]
**FingerPrint** | **string** | Finger print for the private key | [default to null]
**PrivateKeyFileName** | **string** | File Name for the private key | [default to null]
**PrivateKeyPassword** | **string** | password for the private key.This is the passphrase that was used to encrypt the private key. | [optional] [default to null]
**RegionName** | **string** | home region for the tenant | [default to null]
**Name** | **string** | The name of the hypervisor group being created | [default to null]
**SkipCredentialValidation** | **bool** | if credential validation has to be skipped. | [optional] [default to false]
**AccessNodes** | [**[]AccessNodeModel**](accessNodeModel.md) |  | [optional] [default to null]
**Credentials** | [***IdName**](IdName.md) |  | [optional] [default to null]
**EtcdProtection** | [***EtcdProtectionItem**](EtcdProtectionItem.md) |  | [optional] [default to null]
**PlanEntity** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

