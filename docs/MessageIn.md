# MessageIn

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the inbound message. | [default to null]
**Sender** | **string** | The sender’s phone number. | [default to null]
**Receiver** | **string** | The receiver’s phone number (i.e. your dedicated or shared reply number). | [default to null]
**MessageTime** | [**time.Time**](time.Time.md) | The time when the message reached the TextMagic API endpoint. | [default to null]
**Text** | **string** | The text from the received message. | [default to null]
**ContactId** | **int32** | Sender contact ID. | [optional] [default to null]
**FirstName** | **string** | Sender contact first name. | [optional] [default to null]
**LastName** | **string** | Sender contact last name. | [optional] [default to null]
**Avatar** | **string** |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


