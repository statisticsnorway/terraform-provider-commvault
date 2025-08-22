# KubernetesApplicationGroupFilterItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SkipStatelessApps** | **bool** | Specify whether to skip backup of stateless applications | [optional] [default to false]
**Applications** | [**[]KubernetesContentApplications**](KubernetesContentApplications.md) | List of applications to be added as content | [optional] [default to null]
**LabelSelectors** | [**[]KubernetesContentSelectors**](KubernetesContentSelectors.md) | List of label selectors to be added as content | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

