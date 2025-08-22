# UpdateHypervisorGroupOracleCloud

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorType** | **string** |  | [default to null]
**TenancyOCId** | **string** | OCID for the tenant. | [optional] [default to null]
**UserOCId** | **string** | OCID for the admin user for the hypervisor | [optional] [default to null]
**FingerPrint** | **string** | Finger print for the private key | [optional] [default to null]
**PrivateKeyFileName** | **string** | File Name for the private key | [optional] [default to null]
**PrivateKeyPassword** | **string** | password for the private key.This is the passphrase that was used to encrypt the private key. | [optional] [default to null]
**RegionName** | **string** | home region for the tenant | [optional] [default to null]
**AccessNodes** | [**[]AccessNodeModel**](accessNodeModel.md) |  | [optional] [default to null]
**FbrUnixMediaAgent** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ActivityControl** | [***ActivityControlOptions**](ActivityControlOptions.md) |  | [optional] [default to null]
**Security** | [***VmHypervisorSecurityProp**](VMHypervisorSecurityProp.md) |  | [optional] [default to null]
**NewName** | **string** | The name of the hypervisor that has to be changed | [optional] [default to null]
**Settings** | [***HypervisorSettings**](hypervisorSettings.md) |  | [optional] [default to null]
**Credentials** | [***IdName**](IdName.md) |  | [optional] [default to null]
**SkipCredentialValidation** | **bool** | if credential validation has to be skipped. | [optional] [default to false]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

