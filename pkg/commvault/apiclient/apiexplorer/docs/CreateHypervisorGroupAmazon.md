# CreateHypervisorGroupAmazon

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HypervisorType** | **string** |  | [default to null]
**AccessKey** | **string** | Access Key of Amazon login | [default to null]
**RoleARN** | **string** | Role ARN for STS assume role with IAM policy | [optional] [default to null]
**SecretKey** | **string** | secret Key of Amazon login | [default to null]
**Region** | **string** | AWS region if Iam role is used | [optional] [default to null]
**UseIamRole** | **bool** | if Iam Role is used | [default to null]
**EnableAWSAdminAccount** | **bool** |  | [optional] [default to null]
**UseServiceAccount** | **string** | Clientname to be used as Admin Account | [optional] [default to null]
**UseHostedInfrastructure** | **bool** | Use Metallic hosted infrastructure | [optional] [default to null]
**Name** | **string** | The name of the hypervisor group being created | [default to null]
**SkipCredentialValidation** | **bool** | if credential validation has to be skipped. | [optional] [default to false]
**AccessNodes** | [**[]AccessNodeModel**](accessNodeModel.md) |  | [optional] [default to null]
**Credentials** | [***IdName**](IdName.md) |  | [optional] [default to null]
**EtcdProtection** | [***EtcdProtectionItem**](EtcdProtectionItem.md) |  | [optional] [default to null]
**PlanEntity** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

