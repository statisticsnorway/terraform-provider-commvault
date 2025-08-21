# ArrayGenDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnapVendor** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Array** | [***IdName**](IdName.md) |  | [optional] [default to null]
**ControlHost** | **string** | Control Host name of the array | [optional] [default to null]
**Credential** | [***ArrayUserPassword**](ArrayUserPassword.md) |  | [optional] [default to null]
**SavedCredential** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Region** | [***IdName**](IdName.md) |  | [optional] [default to null]
**Flags** | **int32** | 0: default 1: Only selected arrays 2: Disable automatic cleanup 4: System created arrays 8: Auto created arrays 16: User created arrays | [optional] [default to FLAGS.0_]
**EnableNetAppCloudAccount** | **bool** | Set as true for NetApp cloud target arrays | [optional] [default to null]
**ArrayType** | **int32** | Used to differentiate between Primary, Secondary, and OCUM type of NetApp Array | [optional] [default to null]
**CloudVendorId** | **int32** | Id for cloud vendors associated to arrays | [optional] [default to null]
**UniqueIdentifier** | **string** | Unique identifier pertaining to each array | [optional] [default to null]
**Description** | **string** | User provided description of the array | [optional] [default to null]
**Company** | [***IdName**](IdName.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

