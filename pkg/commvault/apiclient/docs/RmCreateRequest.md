# RmCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the request | [default to null]
**Type_** | [***RmRequestType**](RMRequestType.md) |  | [default to null]
**DeleteFromBackup** | **bool** |  | [optional] [default to false]
**EnableRedaction** | **bool** | This option redacts sensitive information from the files in the request | [optional] [default to false]
**EnableDocumentChaining** | **bool** | If additional entities are found in a document, include documents that contain those additional entities in search results | [optional] [default to false]
**Requestor** | **string** |  | [default to null]
**Entities** | [**[]NameValues**](NameValues.md) | List of personal entities with their values to form the criteria for the request | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

