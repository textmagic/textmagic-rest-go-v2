# BulkSession

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Bulk Session ID. | [default to null]
**Status** | **string** | * **n** - bulk session is just created * **w** - work in progress * **f** - failed * **c** - completed with success * **s** - suspended  | [default to null]
**ItemsProcessed** | **int32** | Amount of messages which is already processed. | [default to null]
**ItemsTotal** | **int32** | Total amount of messages to be processed. | [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | Creation date and time of a Bulk Session. | [default to null]
**Session** | [***MessageSession**](MessageSession.md) |  | [default to null]
**Text** | **string** | Message text of a Bulk Session. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


