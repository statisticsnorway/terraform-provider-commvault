# AzureVmInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceVMGuid** | **string** | Azure VM ID | [optional] [default to null]
**Name** | **string** | Azure VM name to be set after restore. Defaults to source VM name. | [optional] [default to Defaults to source VM name]
**ResourceGroup** | **string** | Azure Resource Group Name. Defaults to source VM resource group name. | [optional] [default to Defaults to source VM resource group name]
**Region** | **string** | Azure region name. Ex: For azure region (West US 3) the region value will be westus3. For reference: learn.microsoft.com/en-us/rest/api/resources/subscriptions/list-locations | [optional] [default to Defaults to source VM region]
**AvailabilityZone** | **string** | Azure availablity zones. Defaults to Auto. Values can be as follows: &#x27;Auto&#x27; (Select from the source VM config), &#x27;None&#x27; (Will not restore to any zone), &#x27;{Zone Number}&#x27; (Availablity zone number the restored VM should be in. Ex: 2). | [optional] [default to Defaults to Auto]
**StorageAccount** | **string** | Azure staging storage account. Defaults to source storage account. | [default to Defaults to source storage account]
**VmSize** | **string** | VM size to be after restore. Defaults to --Auto Select--. Ex: Standard_A1. | [optional] [default to --Auto Select--]
**DiskType** | **string** | Disk type to set after restore. Ex: Standard_LRS. | [optional] [default to null]
**CreatePublicIP** | **bool** | Boolean to specify if VM should have publice IP. Defaults to false. | [optional] [default to false]
**RestoreAsManagedVM** | **bool** | Boolean to specify if VM need to be restored as managed VM. Defaults to true. | [optional] [default to true]
**SecurityGroupId** | **string** | Azure security group to be set for the VM. Defaults to source configuration. | [optional] [default to Defaults to source configuration]
**DiskEncryptionSetTypeId** | **string** | Azure Disk Encryption Type. Values can be: &#x27;EncryptionAtRestWithCustomerKey&#x27; or &#x27;EncryptionAtRestWithPlatformAndCustomerKeys&#x27; or &#x27;EncryptionAtRestWithPlatformAndCustomerKeys&#x27;. For reference: learn.microsoft.com/en-us/dotnet/api/microsoft.azure.management.compute.models.encryptiontype?view&#x3D;azure-dotnet&amp;viewFallbackFrom&#x3D;azure-dotnet | [optional] [default to null]
**DiskEncryptionSetId** | **string** | Azure Disk Encryption Set ID. Format: /subscriptions/{Subscription ID}/resourceGroups/{Resource Group Name}/providers/Microsoft.Compute/diskEncryptionSets/{Disk Encryption Set}. | [optional] [default to null]
**Nics** | [**[]AzureDestinationInfoNic**](AzureDestinationInfoNic.md) | Azure VM network interface list | [optional] [default to null]
**VmTags** | [**[]NameValue**](NameValue.md) | Azure VM tag list | [optional] [default to null]
**RestoreSourceVMTags** | **bool** | Restore source VM tags. | [optional] [default to true]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

