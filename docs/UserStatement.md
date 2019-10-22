# UserStatement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | User statement ID. | [default to null]
**UserId** | **int32** | User ID. | [default to null]
**Date** | [**time.Time**](time.Time.md) | User statement date. | [default to null]
**Balance** | **float64** |  | [default to null]
**Delta** | **float32** | Balance change amount. | [default to null]
**Type_** | **string** | Type of statement (what you have been charged for): *   **sms** for sending SMS *   **number** for renewing a dedicated number *   **schedule** for scheduling text messages *   **topup** for adding credits to your account  | [default to null]
**Value** | **string** | Value differs by **type**: *   for **sms**, it is the sent messages amount *   for **number**, it is a dedicated phone number *   for **schedule**, it is a scheduled messages amount *   for **top-up**, it is an invoice ID  | [default to null]
**Comment** | **string** | Optional comment. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


