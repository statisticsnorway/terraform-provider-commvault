# ChangeServerDomainReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OldDomainName** | **string** | Current domain name of the servers that needs to be changed | [optional] [default to null]
**NewDomainName** | **string** | New domain name for the servers | [optional] [default to null]
**Servers** | [**[]IdName**](IdName.md) | List of servers that need domain name change | [optional] [default to null]
**ForceUpdateDatabase** | **bool** | Update database even if servers are unreachable | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

