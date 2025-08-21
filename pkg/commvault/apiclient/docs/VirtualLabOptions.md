# VirtualLabOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalNetwork** | **string** | To use virtual lab VM to connect to the existing network | [optional] [default to null]
**GatewayTemplate** | **string** | Gateway template provision VMs to communicate to VMs outside the virtual lab | [optional] [default to null]
**GatewayNetwork** | **string** | Primary network associated on the Provisioning Gateway VM that has production network access | [optional] [default to null]
**ConfigureIsolatedNetwork** | **bool** | To use isolated network for the virtual lab VM | [optional] [default to null]
**IpSettings** | [**[]VirtualLabIpSettings**](VirtualLabIPSettings.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

