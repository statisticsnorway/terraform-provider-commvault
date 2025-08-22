# InstallMediaAgent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostNames** | **[]string** | Host names where MediaAgent package is required to be installed | [default to null]
**Username** | **string** | Username to access hostnames | [default to null]
**Password** | **string** | password should be in base64 encoded string. password is optional only if SSHKeyFilePassphrase is specified | [optional] [default to null]
**OSType** | [***MediaAgentInstallOsType**](MediaAgentInstallOSType.md) |  | [optional] [default to null]
**InstallLocation** | **string** | Holds install path according to the chosen OSType | [optional] [default to null]
**RebootIfRequired** | **bool** |  | [optional] [default to false]
**SSHKeyPath** | **string** | Applicable only for Unix OSType. Path provided should be valid on the CommServ Client | [optional] [default to null]
**SSHKeyFilePassphrase** | **string** | Applicable only for Unix OSType and should be in base64 encoded string | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

