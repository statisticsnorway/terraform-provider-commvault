# FirewallTopologyDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topology** | [***IdName**](IdName.md) |  | [optional] [default to null]
**TopologyType** | [***FirewallTopologyTypes**](firewallTopologyTypes.md) |  | [optional] [default to null]
**ClientType** | [***FirewallClientType**](firewallClientType.md) |  | [optional] [default to null]
**UseWildCardProxy** | **bool** | Flag determining wether network gateways are used to connect all infrastructure machines | [optional] [default to null]
**FirewallGroups** | [**[]FirewallTopologyGroups**](FirewallTopologyGroups.md) |  | [optional] [default to null]
**EncryptTraffic** | **bool** | Flag determining if we want the data from tunnel to use HTTPS protocol | [optional] [default to null]
**TunnelProtocol** | [***FirewallTopologyTunnelProtocol**](FirewallTopologyTunnelProtocol.md) |  | [optional] [default to null]
**TunnelsPerRoute** | **int32** | The number of tunnel connections per route | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

