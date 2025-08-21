# TapeMediaOperationsReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaTypeId** | **int32** | MediaType list must be fetched from GET V4/Storage/Tape/{libraryId}/MediaType | [optional] [default to 0]
**SpareGroupId** | **int32** | SpareGroup list must be fetched from GET V4/Storage/Tape/{libraryId}/MediaType | [optional] [default to 0]
**DestSpareGroupId** | **int32** | In case of MOVE_MEDIA, this field should be set to specify the destination spare group | [optional] [default to 0]
**SlotList** | **[]int32** |  | [optional] [default to null]
**OperationType** | [***TapeMediaOperationType**](TapeMediaOperationType.md) |  | [optional] [default to null]
**MediaList** | **[]int32** | List of Media Ids for which operation needs to performed. NOTE : In case of QUICK_ERASE_SELECTED_MEDIA and FULL_ERASE_SELECTED_MEDIA single mediaId should be sent. | [optional] [default to null]
**UpdateBarcodeOpType** | [***UpdateBarcodeOpType**](UpdateBarcodeOpType.md) |  | [optional] [default to null]
**Barcode** | **string** | In case of REPLACE_BARCODE, media barcode is replaced with barcode sent. In other cases of UpdateBarcodeOpType, barcode sent is added or removed according to the type specified. For instance, in case of ADD_SUFFIX, the barcode sent will be appended to the end of barcodes of all the medias sent in request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

