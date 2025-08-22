# AlertDefinitionsTemplate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locale** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Email** | **string** | the message template for the email notification. Contains both email subject as well as body | [optional] [default to null]
**Console** | **string** | the message template for the console notification | [optional] [default to null]
**EventViewer** | **string** | the message template for the event viewer notification | [optional] [default to null]
**Webhook** | **string** | the message template for the webhook notification | [optional] [default to null]
**Workflow** | [***WorkflowEntity**](WorkflowEntity.md) |  | [optional] [default to null]
**WorkflowInputDetails** | **string** | the template for the workflow input entries values (empty value means the configured workflow does not need any inputs value) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

