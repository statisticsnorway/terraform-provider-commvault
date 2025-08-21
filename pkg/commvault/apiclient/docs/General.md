# General

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewAlias** | **string** | The company domain or NetBIOS name | [optional] [default to null]
**EmailSuffix** | **string** | Supported domains for the company | [optional] [default to null]
**AuthcodeForInstallation** | **bool** | Enable or disable authcode for installation. | [optional] [default to null]
**UpnInsteadOfEmail** | **bool** | Enable or disable the use of User Pricipal Name in place of an email address. | [optional] [default to null]
**TwoFactorAuth** | [***TwoFactorAuth**](TwoFactorAuth.md) |  | [optional] [default to null]
**ResellerMode** | **bool** | Enable reseller mode. A reseller is a user or user groups who can operate multiple tenant environments. The service provider can assign one or more resellers as the tenant operator for a company. The reseller can switch to any of their assigned tenant environment and operate on the company as a tenant user. Once enabled, the reseller mode cannot be disabled. | [optional] [default to null]
**EnableDataEncryption** | **bool** | Enable or disable data encryption | [optional] [default to null]
**AutoDiscoverApp** | **bool** | Enable or Disable Auto Discover Applications.When Auto discover applications is enabled, each member server of this company is searched once every 24 hours to discover any applications that need to be backed up. For any newly-discovered and unprotected applications, the agent software is automatically installed on the server. | [optional] [default to null]
**InfrastructureType** | [***InfrastructureTypes**](InfrastructureTypes.md) |  | [optional] [default to null]
**SupportedSolutions** | [**[]SupportedSolution**](SupportedSolution.md) |  | [optional] [default to null]
**AssignLaptopOwners** | [***OwnerAssignmentTypes**](OwnerAssignmentTypes.md) |  | [optional] [default to null]
**ServiceCommcells** | [**[]IdName**](IdName.md) | Service commcells associated with the master commcell. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

