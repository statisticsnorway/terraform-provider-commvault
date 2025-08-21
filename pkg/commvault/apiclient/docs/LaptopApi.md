# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLaptopList**](LaptopApi.md#GetLaptopList) | **Get** /V4/Laptop | Get Laptop Devices (New)
[**PostLaptop**](LaptopApi.md#PostLaptop) | **Post** /V4/Laptop | create a laptop

# **GetLaptopList**
> LaptopsList GetLaptopList(ctx, optional)
Get Laptop Devices (New)

Get list of laptops

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaptopApiGetLaptopListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaptopApiGetLaptopListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **edgemode** | **optional.Bool**|  | 
 **additionalProperties** | **optional.Bool**|  | 

### Return type

[**LaptopsList**](LaptopsList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostLaptop**
> AddLaptopResp PostLaptop(ctx, optional)
create a laptop

Simplified API to create a laptop

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaptopApiPostLaptopOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaptopApiPostLaptopOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AddLaptop**](AddLaptop.md)|  | 

### Return type

[**AddLaptopResp**](AddLaptopResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

