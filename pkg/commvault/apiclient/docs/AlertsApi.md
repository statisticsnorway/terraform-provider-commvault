# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertDefinitions**](AlertsApi.md#CreateAlertDefinitions) | **Post** /V4/AlertDefinitions | Create Alert Definitions
[**DeleteAlertCustomRule**](AlertsApi.md#DeleteAlertCustomRule) | **Put** /Alert/CustomRule/{alertRuleId}/Delete | Delete Alert Custom Rule
[**DeleteAlertDefinitions**](AlertsApi.md#DeleteAlertDefinitions) | **Delete** /V4/AlertDefinitions/{id} | Delete AlertDefinitions
[**DeleteMultipleTriggeredalerts**](AlertsApi.md#DeleteMultipleTriggeredalerts) | **Put** /V4/TriggeredAlerts/Action/Delete | Delete Multiple Triggered alerts
[**DeleteTriggeredAlerts**](AlertsApi.md#DeleteTriggeredAlerts) | **Delete** /V4/TriggeredAlerts/{id} | Delete Triggered Alerts
[**DisableAlertCustomRule**](AlertsApi.md#DisableAlertCustomRule) | **Put** /Alert/CustomRule/{alertRuleId}/Disable | Disable Alert Custom Rule
[**DisableAlertDefinitions**](AlertsApi.md#DisableAlertDefinitions) | **Put** /V4/AlertDefinitions/{id}/Disable | Disable Alert Definitions
[**EditAlertDefinitions**](AlertsApi.md#EditAlertDefinitions) | **Put** /V4/AlertDefinitions/{id} | Edit Alert Definitions
[**EnableAlertCustomRule**](AlertsApi.md#EnableAlertCustomRule) | **Put** /Alert/CustomRule/{alertRuleId}/Enable | Enable Alert Custom Rule
[**EnableAlertDefinitions**](AlertsApi.md#EnableAlertDefinitions) | **Put** /V4/AlertDefinitions/{id}/Enable | Enable Alert Definitions
[**GetAlertCustomRuleDetails**](AlertsApi.md#GetAlertCustomRuleDetails) | **Get** /Alert/CustomRule/{alertRuleId} | Alert custom rule details
[**GetAlertCustomRulesList**](AlertsApi.md#GetAlertCustomRulesList) | **Get** /Alert/CustomRule | Get Alert Custom Rules List
[**GetAlertDefinitionsDetails**](AlertsApi.md#GetAlertDefinitionsDetails) | **Get** /V4/AlertDefinitions/{id} | Get Alert DefinitionsDetails
[**GetAlertDefinitionsList**](AlertsApi.md#GetAlertDefinitionsList) | **Get** /V4/AlertDefinitions | Get the list of Alert Definitions
[**GetAlertType**](AlertsApi.md#GetAlertType) | **Get** /V4/AlertType | Get Alert Type and Alert Custom Rules list
[**GetAlertsTriggered**](AlertsApi.md#GetAlertsTriggered) | **Get** /V4/TriggeredAlerts | Get Alerts Triggered
[**GetTriggeredAlertDetails**](AlertsApi.md#GetTriggeredAlertDetails) | **Get** /V4/TriggeredAlerts/{id} | Get details of triggered alert
[**ModifyTriggeredAlertsNotes**](AlertsApi.md#ModifyTriggeredAlertsNotes) | **Put** /V4/TriggeredAlerts/{id}/Notes | Modify Triggered AlertsNotes
[**PinTriggeredAlerts**](AlertsApi.md#PinTriggeredAlerts) | **Put** /V4/TriggeredAlerts/{id}/Pin | Pin TriggeredAlerts
[**ReadTriggeredAlerts**](AlertsApi.md#ReadTriggeredAlerts) | **Put** /V4/TriggeredAlerts/{id}/Read | Read Triggered Alerts
[**UnpinTriggeredAlerts**](AlertsApi.md#UnpinTriggeredAlerts) | **Put** /V4/TriggeredAlerts/{id}/Unpin | Unpin Triggered Alerts
[**UnreadTriggeredAlerts**](AlertsApi.md#UnreadTriggeredAlerts) | **Put** /V4/TriggeredAlerts/{id}/Unread | Unread TriggeredAlerts

# **CreateAlertDefinitions**
> IdNameGuid CreateAlertDefinitions(ctx, optional)
Create Alert Definitions

Create Alert Definitions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertsApiCreateAlertDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertsApiCreateAlertDefinitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AlertDefinitonsCreate**](AlertDefinitonsCreate.md)|  | 

### Return type

[**IdNameGuid**](IdNameGUID.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAlertCustomRule**
> AlertRuleGenericResp DeleteAlertCustomRule(ctx, alertRuleId)
Delete Alert Custom Rule

Delete the alert custom rule by using the given alert rule id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertRuleId** | **int32**| Unique Id of the alert custom rule to delete it | 

### Return type

[**AlertRuleGenericResp**](AlertRuleGenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAlertDefinitions**
> GenericResp DeleteAlertDefinitions(ctx, id)
Delete AlertDefinitions

Delete alert definitions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMultipleTriggeredalerts**
> GenericResp DeleteMultipleTriggeredalerts(ctx, optional)
Delete Multiple Triggered alerts

Delete multiple triggered alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertsApiDeleteMultipleTriggeredalertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertsApiDeleteMultipleTriggeredalertsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AlertIds**](AlertIds.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTriggeredAlerts**
> GenericResp DeleteTriggeredAlerts(ctx, id)
Delete Triggered Alerts

Delete Triggered Alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableAlertCustomRule**
> AlertRuleGenericResp DisableAlertCustomRule(ctx, alertRuleId)
Disable Alert Custom Rule

Disable the alert custom rule by using the given alert rule id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertRuleId** | **int32**| Unique Id of the alert custom rule to disable it | 

### Return type

[**AlertRuleGenericResp**](AlertRuleGenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableAlertDefinitions**
> GenericResp DisableAlertDefinitions(ctx, id)
Disable Alert Definitions

Disable Alert Definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditAlertDefinitions**
> GenericResp EditAlertDefinitions(ctx, id, optional)
Edit Alert Definitions

Edit Alert Definition details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***AlertsApiEditAlertDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertsApiEditAlertDefinitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AlertDefinitionsEdit**](AlertDefinitionsEdit.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableAlertCustomRule**
> AlertRuleGenericResp EnableAlertCustomRule(ctx, alertRuleId)
Enable Alert Custom Rule

Enable the alert custom rule by using the given alert rule id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertRuleId** | **int32**| Unique Id of the alert custom rule to enable it | 

### Return type

[**AlertRuleGenericResp**](AlertRuleGenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableAlertDefinitions**
> GenericResp EnableAlertDefinitions(ctx, id)
Enable Alert Definitions

Enable Alert Definitions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertCustomRuleDetails**
> AlertRulesDetailsList GetAlertCustomRuleDetails(ctx, alertRuleId)
Alert custom rule details

Get the details of the custom alert rule by Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **alertRuleId** | **int32**| Unique Id of the alert custom rule to get details | 

### Return type

[**AlertRulesDetailsList**](AlertRulesDetailsList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertCustomRulesList**
> AlertRulesDetailsList GetAlertCustomRulesList(ctx, )
Get Alert Custom Rules List

Get the list of alert custom rules

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AlertRulesDetailsList**](AlertRulesDetailsList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertDefinitionsDetails**
> AlertDefinitionsDetails GetAlertDefinitionsDetails(ctx, id, optional)
Get Alert DefinitionsDetails

Get details of alert definition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***AlertsApiGetAlertDefinitionsDetailsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertsApiGetAlertDefinitionsDetailsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **additionalProperties** | **optional.Bool**|  | 

### Return type

[**AlertDefinitionsDetails**](AlertDefinitionsDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertDefinitionsList**
> AlertDefinitionsList GetAlertDefinitionsList(ctx, optional)
Get the list of Alert Definitions

Get the list of Alert Definitions

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertsApiGetAlertDefinitionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertsApiGetAlertDefinitionsListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **additionalProperties** | **optional.Bool**| To get extra meta data details for the api | 

### Return type

[**AlertDefinitionsList**](AlertDefinitionsList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertType**
> AlertTypeResp GetAlertType(ctx, )
Get Alert Type and Alert Custom Rules list

Get the list of alert types and alert custom rules

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AlertTypeResp**](AlertTypeResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAlertsTriggered**
> AlertsTriggeredListResp GetAlertsTriggered(ctx, )
Get Alerts Triggered

Get List Of Alerts Triggered

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AlertsTriggeredListResp**](AlertsTriggeredListResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTriggeredAlertDetails**
> TriggeredAlertsDetails GetTriggeredAlertDetails(ctx, id)
Get details of triggered alert

Get details of triggered alert

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**TriggeredAlertsDetails**](TriggeredAlertsDetails.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyTriggeredAlertsNotes**
> GenericResp ModifyTriggeredAlertsNotes(ctx, id, optional)
Modify Triggered AlertsNotes

Add, Modify and Delete notes for triggered alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***AlertsApiModifyTriggeredAlertsNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertsApiModifyTriggeredAlertsNotesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of UpdateTriggeredAlertsNotes**](UpdateTriggeredAlertsNotes.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PinTriggeredAlerts**
> GenericResp PinTriggeredAlerts(ctx, id)
Pin TriggeredAlerts

Pin triggered alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReadTriggeredAlerts**
> GenericResp ReadTriggeredAlerts(ctx, id)
Read Triggered Alerts

Mark a triggered alert as read

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnpinTriggeredAlerts**
> GenericResp UnpinTriggeredAlerts(ctx, id)
Unpin Triggered Alerts

Unpin triggered alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnreadTriggeredAlerts**
> GenericResp UnreadTriggeredAlerts(ctx, id)
Unread TriggeredAlerts

Mark a triggered alert as unread

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

