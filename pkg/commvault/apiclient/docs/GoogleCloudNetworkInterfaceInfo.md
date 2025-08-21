# GoogleCloudNetworkInterfaceInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the network interface card. Ex: nic0. | [optional] [default to null]
**NetworkName** | **string** | Network Name. Ex: default. Defaults to source network name. | [optional] [default to Defaults to source network name.]
**SubnetId** | **string** | Subnet self-link. Format: googleapis.com/compute/v1/projects/&lt;Project ID&gt;/regions/&lt;Region&gt;/subnetworks/{Subnet Name}. Defaults to source subnet self-link. | [optional] [default to Defaults to source subnet]
**NetworkId** | **string** | Network self-link. Format: googleapis.com/compute/v1/projects/&lt;Project ID&gt;/global/networks/{Network Name}. Defaults to source network self-link. | [optional] [default to Defaults to source network self-link.]
**InternalIP** | **string** | Internal IP address. Defaults to the source internal IP after restore. | [optional] [default to Defaults to the source internal IP after restore.]
**ExternalIP** | **string** | External IP Address. Defaults to null and is not set after restore. | [optional] [default to Defaults to null and is not set after restore.]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

