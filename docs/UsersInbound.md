# UsersInbound

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Dedicated number ID. | [default to null]
**DisplayTimeFormat** | **string** | Format for representation of time | [optional] [default to null]
**Phone** | **string** | Dedicated phone number. | [optional] [default to null]
**User** | [***User**](User.md) |  | [default to null]
**PurchasedAt** | [**time.Time**](time.Time.md) | Time when the dedicated number was purchased. | [default to null]
**ExpireAt** | [**time.Time**](time.Time.md) | Dedicated number subscription expiration time. | [default to null]
**Status** | **string** | Number status: *   **U** for Unused. No messages have been sent from (or received to) this number. *   **A** for Active.  | [default to null]
**Country** | [***Country**](Country.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


