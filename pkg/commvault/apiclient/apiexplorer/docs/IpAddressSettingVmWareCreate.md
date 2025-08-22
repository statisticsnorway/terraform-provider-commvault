# IpAddressSettingVmWareCreate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceIP** | **string** | The full IP address or an IP address pattern of the source VM. | [optional] [default to null]
**SourceSubnetMask** | **string** | The subnet mask of the source VM. | [optional] [default to null]
**SourceDefaultGateway** | **string** | The default gateway of the source VM. | [optional] [default to null]
**UseDhcp** | **bool** | Automatically assigns available IP addresses to be used with the specified destination network. | [optional] [default to false]
**DestinationIP** | **string** | The full IP address or an IP address pattern for the destination VM. Provide only if DHCP is not enabled. | [optional] [default to null]
**DestinationSubnetMask** | **string** | The subnet mask for the destination VM.  Provide only if DHCP is not enabled. | [optional] [default to null]
**DestinationDefaultGateway** | **string** | The default gateway for the destination VM.  Provide only if DHCP is not enabled. | [optional] [default to null]
**DestinationPreferredDNS** | **string** | The preferred DNS server for the destination VM.  Provide only if DHCP is not enabled | [optional] [default to null]
**DestinationAlternateDNS** | **string** | The alternate DNS server for the destination VM.  Provide only if DHCP is not enabled | [optional] [default to null]
**DestinationPreferredWINS** | **string** | The preferred WINS server for the destination VM.  Provide only if DHCP is not enabled | [optional] [default to null]
**DestinationAlternateWINS** | **string** | The alternate WINS server for the destination VM.  Provide only if DHCP is not enabled | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

