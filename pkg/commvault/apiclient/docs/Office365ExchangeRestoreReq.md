# Office365ExchangeRestoreReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mailboxes** | [**[]Office365ExchMailboxRestoreEntity**](Office365ExchMailboxRestoreEntity.md) |  | [optional] [default to null]
**Groups** | [**[]NameGuid**](NameGUID.md) |  | [optional] [default to null]
**IfMessageExists** | **string** |  | [optional] [default to IF_MESSAGE_EXISTS.SKIP]
**DeletedItems** | **bool** | Include deleted items on Restore | [optional] [default to null]
**MatchEmailAddress** | **bool** | Match destination user based on the email address | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

