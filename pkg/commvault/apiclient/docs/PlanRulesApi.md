# {{classname}}

All URIs are relative to *https://{ServerURL}/commandcenter/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePlanRule**](PlanRulesApi.md#CreatePlanRule) | **Post** /V4/Plan/Rule | Create a Plan Rule
[**DeletePlanRule**](PlanRulesApi.md#DeletePlanRule) | **Delete** /V4/Plan/Rule/{ruleId} | Delete selected plan rule
[**GetPlanRuleAssociatedEntities**](PlanRulesApi.md#GetPlanRuleAssociatedEntities) | **Get** /V4/Plan/Rule/Entities | Get entities applicable for plan association
[**GetPlanRuleDetails**](PlanRulesApi.md#GetPlanRuleDetails) | **Get** /V4/Plan/Rule/{ruleId} | Api to fetch plan rules details.
[**GetPlanRuleEntitiesExclude**](PlanRulesApi.md#GetPlanRuleEntitiesExclude) | **Get** /V4/Plan/Rule/Entities/Exclude | Get excluded workloads from plan rule evaluation framework
[**GetPlanRuleSettings**](PlanRulesApi.md#GetPlanRuleSettings) | **Get** /V4/Plan/Rule/Settings | Fetch plan rule execution settings
[**GetPlanRules**](PlanRulesApi.md#GetPlanRules) | **Get** /V4/Plan/Rule | Fetch plan rules list.
[**PutPlanRuleAssociatedEntities**](PlanRulesApi.md#PutPlanRuleAssociatedEntities) | **Put** /V4/Plan/Rule/Entities | Process list of entities against plan rules.
[**PutPlanRuleEntitiesAssociate**](PlanRulesApi.md#PutPlanRuleEntitiesAssociate) | **Put** /V4/Plan/Rule/Entities/Associate | Associate workloads to selected plan and plan rule
[**PutPlanRuleEntitiesExclude**](PlanRulesApi.md#PutPlanRuleEntitiesExclude) | **Put** /V4/Plan/Rule/Entities/Exclude | Exclude workloads from plan rule evaluation framework
[**PutPlanRuleEntitiesInclude**](PlanRulesApi.md#PutPlanRuleEntitiesInclude) | **Put** /V4/Plan/Rule/Entities/Include | Include workloads to be applicable for plan rule evaluation.
[**PutPlanRuleSettings**](PlanRulesApi.md#PutPlanRuleSettings) | **Put** /V4/Plan/Rule/Settings | Put PlanRuleSettings
[**UpdatePlanRule**](PlanRulesApi.md#UpdatePlanRule) | **Put** /V4/Plan/Rule | Update a Plan Rule
[**UpdatePlanRuleRank**](PlanRulesApi.md#UpdatePlanRuleRank) | **Put** /V4/Plan/Rule/Rank | Update rank for multiple rules at a time.

# **CreatePlanRule**
> IdName CreatePlanRule(ctx, optional)
Create a Plan Rule

Api to create plan rules. Request body will suggest plan against which rule need to be created and type of entities(workloads type, server groups, regions and all) that need to be referenced for applying that rule against a workload.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlanRulesApiCreatePlanRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanRulesApiCreatePlanRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of CreatePlanEntityRule**](CreatePlanEntityRule.md)|  | 

### Return type

[**IdName**](IdName.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePlanRule**
> GenericResp DeletePlanRule(ctx, ruleId)
Delete selected plan rule

Api to delete selected plan rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleId** | **int32**| Id of the rule to update in Plan | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlanRuleAssociatedEntities**
> PlanRuleApplicableEntitiesList GetPlanRuleAssociatedEntities(ctx, )
Get entities applicable for plan association

API to fetch list of entities that are applicable for plan association via plan assignment rules.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PlanRuleApplicableEntitiesList**](PlanRuleApplicableEntitiesList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlanRuleDetails**
> PlanEntityRuleInfo GetPlanRuleDetails(ctx, ruleId)
Api to fetch plan rules details.

Api to fetch plan rules details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ruleId** | **int32**| Id of the rule to update in Plan | 

### Return type

[**PlanEntityRuleInfo**](PlanEntityRuleInfo.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlanRuleEntitiesExclude**
> PlanRuleExcludedEntitiesList GetPlanRuleEntitiesExclude(ctx, )
Get excluded workloads from plan rule evaluation framework

API to get excluded workloads from plan rule evaluation framework

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PlanRuleExcludedEntitiesList**](PlanRuleExcludedEntitiesList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlanRuleSettings**
> PlanRuleExecutionSettings GetPlanRuleSettings(ctx, )
Fetch plan rule execution settings

Fetch plan rule execution settings

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PlanRuleExecutionSettings**](PlanRuleExecutionSettings.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPlanRules**
> PlanEntityRuleList GetPlanRules(ctx, )
Fetch plan rules list.

Api to fetch plan rules list.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PlanEntityRuleList**](PlanEntityRuleList.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPlanRuleAssociatedEntities**
> GenericResp PutPlanRuleAssociatedEntities(ctx, optional)
Process list of entities against plan rules.

API to send request to process plan rules against specific set of entities.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlanRulesApiPutPlanRuleAssociatedEntitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanRulesApiPutPlanRuleAssociatedEntitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ExecutePlanRules**](ExecutePlanRules.md)| Request body will consist of entities that are need to be evaluated against plan rules. | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPlanRuleEntitiesAssociate**
> PutPlanRuleEntitiesAssociate(ctx, optional)
Associate workloads to selected plan and plan rule

API to associate workloads to selected plan and plan rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlanRulesApiPutPlanRuleEntitiesAssociateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanRulesApiPutPlanRuleEntitiesAssociateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AssociatePlanRuleEntityList**](AssociatePlanRuleEntityList.md)|  | 

### Return type

 (empty response body)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPlanRuleEntitiesExclude**
> GenericResp PutPlanRuleEntitiesExclude(ctx, optional)
Exclude workloads from plan rule evaluation framework

API to exclude workloads from plan rule evaluation framework

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlanRulesApiPutPlanRuleEntitiesExcludeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanRulesApiPutPlanRuleEntitiesExcludeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ExcludeEntitiesFromPlanRuleEvaluation**](ExcludeEntitiesFromPlanRuleEvaluation.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPlanRuleEntitiesInclude**
> GenericResp PutPlanRuleEntitiesInclude(ctx, optional)
Include workloads to be applicable for plan rule evaluation.

API to include workloads to be applicable for plan rule evaluation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlanRulesApiPutPlanRuleEntitiesIncludeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanRulesApiPutPlanRuleEntitiesIncludeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of IncludeEntitiesForPlanRuleEvaluation**](IncludeEntitiesForPlanRuleEvaluation.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutPlanRuleSettings**
> GenericResp PutPlanRuleSettings(ctx, optional)
Put PlanRuleSettings

API to set plan rule execution settings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlanRulesApiPutPlanRuleSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanRulesApiPutPlanRuleSettingsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PlanRuleExecutionSettings**](PlanRuleExecutionSettings.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePlanRule**
> GenericResp UpdatePlanRule(ctx, optional)
Update a Plan Rule

Api to update existing plan rules. Request body will suggest which rule need to be updated and type of entities(workloads type, server groups, regions and all) that need to be referenced for applying that rule against a workload.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlanRulesApiUpdatePlanRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanRulesApiUpdatePlanRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpdatePlanEntityRule**](UpdatePlanEntityRule.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePlanRuleRank**
> GenericResp UpdatePlanRuleRank(ctx, optional)
Update rank for multiple rules at a time.

API to update rank for multiple rules at a time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlanRulesApiUpdatePlanRuleRankOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlanRulesApiUpdatePlanRuleRankOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpdatePlanEntityRuleRanks**](UpdatePlanEntityRuleRanks.md)|  | 

### Return type

[**GenericResp**](GenericResp.md)

### Authorization

[bearerToken](../README.md#bearerToken), [token](../README.md#token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

