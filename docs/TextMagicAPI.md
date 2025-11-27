# \TextMagicAPI

All URIs are relative to *https://rest.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignContactsToList**](TextMagicAPI.md#AssignContactsToList) | **Put** /api/v2/lists/{id}/contacts | Assign contacts to a list
[**BlockContact**](TextMagicAPI.md#BlockContact) | **Post** /api/v2/contacts/block | Block a contact by phone number
[**BuyDedicatedNumber**](TextMagicAPI.md#BuyDedicatedNumber) | **Post** /api/v2/numbers | Buy a dedicated number
[**ClearAndAssignContactsToList**](TextMagicAPI.md#ClearAndAssignContactsToList) | **Post** /api/v2/lists/{id}/contacts | Reset list members to the specified contacts
[**CloseChatsBulk**](TextMagicAPI.md#CloseChatsBulk) | **Post** /api/v2/chats/close/bulk | Close chats (bulk)
[**CloseReadChats**](TextMagicAPI.md#CloseReadChats) | **Post** /api/v2/chats/close/read | Close read chats
[**CreateContact**](TextMagicAPI.md#CreateContact) | **Post** /api/v2/contacts/normalized | Add a new contact
[**CreateContactNote**](TextMagicAPI.md#CreateContactNote) | **Post** /api/v2/contacts/{id}/notes | Create a new contact note
[**CreateCustomField**](TextMagicAPI.md#CreateCustomField) | **Post** /api/v2/customfields | Add a new custom field
[**CreateEmailCampaign**](TextMagicAPI.md#CreateEmailCampaign) | **Post** /api/v2/email-campaigns | Send email campaign
[**CreateList**](TextMagicAPI.md#CreateList) | **Post** /api/v2/lists | Create a new list
[**CreateTag**](TextMagicAPI.md#CreateTag) | **Post** /api/v2/tags | Create tag
[**CreateTemplate**](TextMagicAPI.md#CreateTemplate) | **Post** /api/v2/templates | Create a template
[**DeleteAllContacts**](TextMagicAPI.md#DeleteAllContacts) | **Delete** /api/v2/contact/all | Delete contacts (bulk)
[**DeleteAllOutboundMessages**](TextMagicAPI.md#DeleteAllOutboundMessages) | **Delete** /api/v2/message/all | Delete all messages
[**DeleteAvatar**](TextMagicAPI.md#DeleteAvatar) | **Delete** /api/v2/user/avatar | Delete an avatar
[**DeleteChatMessages**](TextMagicAPI.md#DeleteChatMessages) | **Post** /api/v2/chats/{id}/messages/delete | Delete chat messages by ID(s)
[**DeleteChatsBulk**](TextMagicAPI.md#DeleteChatsBulk) | **Post** /api/v2/chats/delete | Delete chats (bulk)
[**DeleteContact**](TextMagicAPI.md#DeleteContact) | **Delete** /api/v2/contacts/{id} | Delete a contact
[**DeleteContactAvatar**](TextMagicAPI.md#DeleteContactAvatar) | **Delete** /api/v2/contacts/{id}/avatar | Delete an avatar
[**DeleteContactNote**](TextMagicAPI.md#DeleteContactNote) | **Delete** /api/v2/notes/{id} | Delete a contact note
[**DeleteContactNotesBulk**](TextMagicAPI.md#DeleteContactNotesBulk) | **Post** /api/v2/contacts/{id}/notes/delete | Delete contact notes (bulk)
[**DeleteContactsByIds**](TextMagicAPI.md#DeleteContactsByIds) | **Post** /api/v2/contacts/delete | Delete contacts by IDs (bulk)
[**DeleteContactsFromList**](TextMagicAPI.md#DeleteContactsFromList) | **Delete** /api/v2/lists/{id}/contacts | Unassign contacts from a list
[**DeleteCustomField**](TextMagicAPI.md#DeleteCustomField) | **Delete** /api/v2/customfields/{id} | Delete a custom field
[**DeleteDedicatedNumber**](TextMagicAPI.md#DeleteDedicatedNumber) | **Delete** /api/v2/numbers/{id} | Cancel a dedicated number subscription
[**DeleteInboundMessage**](TextMagicAPI.md#DeleteInboundMessage) | **Delete** /api/v2/replies/{id} | Delete a single inbound message
[**DeleteInboundMessagesBulk**](TextMagicAPI.md#DeleteInboundMessagesBulk) | **Post** /api/v2/replies/delete | Delete inbound messages (bulk)
[**DeleteList**](TextMagicAPI.md#DeleteList) | **Delete** /api/v2/lists/{id} | Delete a list
[**DeleteListAvatar**](TextMagicAPI.md#DeleteListAvatar) | **Delete** /api/v2/lists/{id}/avatar | Delete an avatar for a list
[**DeleteListContactsBulk**](TextMagicAPI.md#DeleteListContactsBulk) | **Post** /api/v2/lists/{id}/contacts/delete | Delete contacts from a list (bulk)
[**DeleteListsBulk**](TextMagicAPI.md#DeleteListsBulk) | **Post** /api/v2/lists/delete | Delete lists (bulk)
[**DeleteMessageSession**](TextMagicAPI.md#DeleteMessageSession) | **Delete** /api/v2/sessions/{id} | Delete a session
[**DeleteMessageSessionsBulk**](TextMagicAPI.md#DeleteMessageSessionsBulk) | **Post** /api/v2/sessions/delete | Delete sessions (bulk)
[**DeleteOutboundMessage**](TextMagicAPI.md#DeleteOutboundMessage) | **Delete** /api/v2/messages/{id} | Delete message
[**DeleteOutboundMessagesBulk**](TextMagicAPI.md#DeleteOutboundMessagesBulk) | **Post** /api/v2/messages/delete | Delete messages (bulk)
[**DeleteScheduledMessage**](TextMagicAPI.md#DeleteScheduledMessage) | **Delete** /api/v2/schedules/{id} | Delete a single scheduled message
[**DeleteScheduledMessagesBulk**](TextMagicAPI.md#DeleteScheduledMessagesBulk) | **Post** /api/v2/schedules/delete | Delete scheduled messages (bulk)
[**DeleteSenderId**](TextMagicAPI.md#DeleteSenderId) | **Delete** /api/v2/senderids/{id} | Delete a Sender ID
[**DeleteTemplate**](TextMagicAPI.md#DeleteTemplate) | **Delete** /api/v2/templates/{id} | Delete a template
[**DeleteTemplatesBulk**](TextMagicAPI.md#DeleteTemplatesBulk) | **Post** /api/v2/templates/delete | Delete templates (bulk)
[**DoCarrierLookup**](TextMagicAPI.md#DoCarrierLookup) | **Get** /api/v2/lookups/{phone} | Carrier Lookup
[**DoEmailLookup**](TextMagicAPI.md#DoEmailLookup) | **Get** /api/v2/email-lookups/{email} | Email Lookup
[**GetAllBulkSessions**](TextMagicAPI.md#GetAllBulkSessions) | **Get** /api/v2/bulks | Get all bulk sessions
[**GetAllChats**](TextMagicAPI.md#GetAllChats) | **Get** /api/v2/chats | Get all chats
[**GetAllInboundMessages**](TextMagicAPI.md#GetAllInboundMessages) | **Get** /api/v2/replies | Get all inbound messages
[**GetAllMessageSessions**](TextMagicAPI.md#GetAllMessageSessions) | **Get** /api/v2/sessions | Get all sessions
[**GetAllOutboundMessages**](TextMagicAPI.md#GetAllOutboundMessages) | **Get** /api/v2/messages | Get all messages
[**GetAllScheduledMessages**](TextMagicAPI.md#GetAllScheduledMessages) | **Get** /api/v2/schedules | Get all scheduled messages
[**GetAllTemplates**](TextMagicAPI.md#GetAllTemplates) | **Get** /api/v2/templates | Get all templates
[**GetAvailableDedicatedNumbers**](TextMagicAPI.md#GetAvailableDedicatedNumbers) | **Get** /api/v2/numbers/available | Find dedicated numbers available for purchase
[**GetAvailableSenderSettingOptions**](TextMagicAPI.md#GetAvailableSenderSettingOptions) | **Get** /api/v2/sources | Get available sender settings
[**GetBalanceNotificationOptions**](TextMagicAPI.md#GetBalanceNotificationOptions) | **Get** /api/v2/user/notification/balance/bundles | Returns the list of available balance options which can be used as a bound to determine when to send email to user with low balance notification. See https://my.textmagic.com/online/account/notifications/balance
[**GetBalanceNotificationSettings**](TextMagicAPI.md#GetBalanceNotificationSettings) | **Get** /api/v2/user/notification/balance | Get balance notification settings
[**GetBlockedContacts**](TextMagicAPI.md#GetBlockedContacts) | **Get** /api/v2/contacts/block/list | Get blocked contacts
[**GetBulkSession**](TextMagicAPI.md#GetBulkSession) | **Get** /api/v2/bulks/{id} | Get bulk session status
[**GetCallbackSettings**](TextMagicAPI.md#GetCallbackSettings) | **Get** /api/v2/callback/settings | Fetch callback URL settings
[**GetChat**](TextMagicAPI.md#GetChat) | **Get** /api/v2/chats/{id} | Get a single chat
[**GetChatByPhone**](TextMagicAPI.md#GetChatByPhone) | **Get** /api/v2/chats/{phone}/by/phone | Find chats by phone
[**GetChatMessages**](TextMagicAPI.md#GetChatMessages) | **Get** /api/v2/chats/{id}/message | Get chat messages
[**GetContact**](TextMagicAPI.md#GetContact) | **Get** /api/v2/contacts/{id} | Get the details of a specific contact
[**GetContactByPhone**](TextMagicAPI.md#GetContactByPhone) | **Get** /api/v2/contacts/phone/{phone} | Get the details of a specific contact by phone number
[**GetContactIfBlocked**](TextMagicAPI.md#GetContactIfBlocked) | **Get** /api/v2/contacts/block/phone | Check if a phone number is blocked
[**GetContactImportSessionProgress**](TextMagicAPI.md#GetContactImportSessionProgress) | **Get** /api/v2/contacts/import/progress/{id} | Check import progress
[**GetContactNote**](TextMagicAPI.md#GetContactNote) | **Get** /api/v2/notes/{id} | Get a contact note
[**GetContactNotes**](TextMagicAPI.md#GetContactNotes) | **Get** /api/v2/contacts/{id}/notes | Fetch notes assigned to a given contact
[**GetContacts**](TextMagicAPI.md#GetContacts) | **Get** /api/v2/contacts | Get all contacts
[**GetContactsAutocomplete**](TextMagicAPI.md#GetContactsAutocomplete) | **Get** /api/v2/contacts/autocomplete | Get contacts autocomplete suggestions
[**GetContactsByListId**](TextMagicAPI.md#GetContactsByListId) | **Get** /api/v2/lists/{id}/contacts | Get all contacts in a list
[**GetCountries**](TextMagicAPI.md#GetCountries) | **Get** /api/v2/countries | Get countries
[**GetCurrentUser**](TextMagicAPI.md#GetCurrentUser) | **Get** /api/v2/user | Get current account information
[**GetCustomField**](TextMagicAPI.md#GetCustomField) | **Get** /api/v2/customfields/{id} | Get the details of a specific custom field
[**GetCustomFields**](TextMagicAPI.md#GetCustomFields) | **Get** /api/v2/customfields | Get all custom fields
[**GetDedicatedNumber**](TextMagicAPI.md#GetDedicatedNumber) | **Get** /api/v2/numbers/{id} | Get the details of a specific dedicated number
[**GetEmailSenders**](TextMagicAPI.md#GetEmailSenders) | **Get** /api/v2/email-campaigns/email-senders | Get list of email senders
[**GetFavorites**](TextMagicAPI.md#GetFavorites) | **Get** /api/v2/contacts/favorite | Get favorite contacts and lists
[**GetInboundMessage**](TextMagicAPI.md#GetInboundMessage) | **Get** /api/v2/replies/{id} | Get a single inbound message
[**GetInboundMessagesNotificationSettings**](TextMagicAPI.md#GetInboundMessagesNotificationSettings) | **Get** /api/v2/user/notification/inbound | Get inbound messages notification settings
[**GetInvoices**](TextMagicAPI.md#GetInvoices) | **Get** /api/v2/invoices | Get all invoices
[**GetList**](TextMagicAPI.md#GetList) | **Get** /api/v2/lists/{id} | Get the details of a specific list
[**GetListContactsIds**](TextMagicAPI.md#GetListContactsIds) | **Get** /api/v2/lists/{id}/contacts/ids | Get all contact IDs in a list
[**GetLists**](TextMagicAPI.md#GetLists) | **Get** /api/v2/lists | Get all lists
[**GetListsOfContact**](TextMagicAPI.md#GetListsOfContact) | **Get** /api/v2/contacts/{id}/lists | Get a contact&#39;s lists
[**GetMessagePreview**](TextMagicAPI.md#GetMessagePreview) | **Get** /api/v2/messages/preview | Preview message
[**GetMessagePrice**](TextMagicAPI.md#GetMessagePrice) | **Get** /api/v2/messages/price/normalized | Check message price
[**GetMessageSession**](TextMagicAPI.md#GetMessageSession) | **Get** /api/v2/sessions/{id} | Get a session&#x60;s details
[**GetMessageSessionStat**](TextMagicAPI.md#GetMessageSessionStat) | **Get** /api/v2/sessions/{id}/stat | Get a session&#x60;s statistics
[**GetMessagesBySessionId**](TextMagicAPI.md#GetMessagesBySessionId) | **Get** /api/v2/sessions/{id}/messages | Get a session&#x60;s messages
[**GetMessagingCounters**](TextMagicAPI.md#GetMessagingCounters) | **Get** /api/v2/stats/messaging/data | Get sent/received messages counters values
[**GetMessagingStat**](TextMagicAPI.md#GetMessagingStat) | **Get** /api/v2/stats/messaging | Get messaging statistics
[**GetOutboundMessage**](TextMagicAPI.md#GetOutboundMessage) | **Get** /api/v2/messages/{id} | Get a single message
[**GetOutboundMessagesHistory**](TextMagicAPI.md#GetOutboundMessagesHistory) | **Get** /api/v2/history | Get history
[**GetScheduledMessage**](TextMagicAPI.md#GetScheduledMessage) | **Get** /api/v2/schedules/{id} | Get a single scheduled message
[**GetSenderId**](TextMagicAPI.md#GetSenderId) | **Get** /api/v2/senderids/{id} | Get the details of a specific Sender ID
[**GetSenderIds**](TextMagicAPI.md#GetSenderIds) | **Get** /api/v2/senderids | Get all your approved Sender IDs
[**GetSenderSettings**](TextMagicAPI.md#GetSenderSettings) | **Get** /api/v2/sender/settings/normalized | Get current sender settings
[**GetSpendingStat**](TextMagicAPI.md#GetSpendingStat) | **Get** /api/v2/stats/spending | Get spending statistics
[**GetTemplate**](TextMagicAPI.md#GetTemplate) | **Get** /api/v2/templates/{id} | Get a template&#x60;s details
[**GetTimezones**](TextMagicAPI.md#GetTimezones) | **Get** /api/v2/timezones | Get timezones
[**GetUnreadMessagesTotal**](TextMagicAPI.md#GetUnreadMessagesTotal) | **Get** /api/v2/chats/unread/count | Get unread messages number
[**GetUnsubscribedContact**](TextMagicAPI.md#GetUnsubscribedContact) | **Get** /api/v2/unsubscribers/{id} | Get the details of a specific unsubscribed contact
[**GetUnsubscribers**](TextMagicAPI.md#GetUnsubscribers) | **Get** /api/v2/unsubscribers | Get all unsubscribed contacts
[**GetUserDedicatedNumbers**](TextMagicAPI.md#GetUserDedicatedNumbers) | **Get** /api/v2/numbers | Get all your dedicated numbers
[**ImportContacts**](TextMagicAPI.md#ImportContacts) | **Post** /api/v2/contacts/import/normalized | Import contacts
[**MarkChatsReadBulk**](TextMagicAPI.md#MarkChatsReadBulk) | **Post** /api/v2/chats/read/bulk | Mark chats as read (bulk)
[**MarkChatsUnreadBulk**](TextMagicAPI.md#MarkChatsUnreadBulk) | **Post** /api/v2/chats/unread/bulk | Mark chats as unread (bulk)
[**MuteChat**](TextMagicAPI.md#MuteChat) | **Post** /api/v2/chats/mute | Mute chat sounds
[**MuteChatsBulk**](TextMagicAPI.md#MuteChatsBulk) | **Post** /api/v2/chats/mute/bulk | Mute chats (bulk)
[**Ping**](TextMagicAPI.md#Ping) | **Get** /api/v2/ping | Ping
[**ReopenChatsBulk**](TextMagicAPI.md#ReopenChatsBulk) | **Post** /api/v2/chats/reopen/bulk | Reopen chats (bulk)
[**RequestSenderId**](TextMagicAPI.md#RequestSenderId) | **Post** /api/v2/senderids | Apply for a new Sender ID
[**ScheduleEmailCampaign**](TextMagicAPI.md#ScheduleEmailCampaign) | **Post** /api/v2/email-campaigns/schedule | Schedule new email campaign
[**SearchChats**](TextMagicAPI.md#SearchChats) | **Get** /api/v2/chats/search | Find chats by message text
[**SearchChatsByIds**](TextMagicAPI.md#SearchChatsByIds) | **Get** /api/v2/chats/search/ids | Find chats (bulk)
[**SearchChatsByReceipent**](TextMagicAPI.md#SearchChatsByReceipent) | **Get** /api/v2/chats/search/recipients | Find chats by recipient
[**SearchContacts**](TextMagicAPI.md#SearchContacts) | **Get** /api/v2/contacts/search | Find contacts by given criteria
[**SearchInboundMessages**](TextMagicAPI.md#SearchInboundMessages) | **Get** /api/v2/replies/search | Find inbound messages
[**SearchLists**](TextMagicAPI.md#SearchLists) | **Get** /api/v2/lists/search | Find lists by given criteria
[**SearchOutboundMessages**](TextMagicAPI.md#SearchOutboundMessages) | **Get** /api/v2/messages/search | Find messages
[**SearchScheduledMessages**](TextMagicAPI.md#SearchScheduledMessages) | **Get** /api/v2/schedules/search | Find scheduled messages
[**SearchTemplates**](TextMagicAPI.md#SearchTemplates) | **Get** /api/v2/templates/search | Find templates by criteria
[**SendMessage**](TextMagicAPI.md#SendMessage) | **Post** /api/v2/messages | Send message
[**SetChatStatus**](TextMagicAPI.md#SetChatStatus) | **Post** /api/v2/chats/status | Change chat status
[**UnblockContact**](TextMagicAPI.md#UnblockContact) | **Post** /api/v2/contacts/unblock | Unblock a contact by phone number
[**UnblockContactsBulk**](TextMagicAPI.md#UnblockContactsBulk) | **Post** /api/v2/contacts/unblock/bulk | Unblock contacts (bulk)
[**UnmuteChatsBulk**](TextMagicAPI.md#UnmuteChatsBulk) | **Post** /api/v2/chats/unmute/bulk | Unmute chats (bulk)
[**UnsubscribeContact**](TextMagicAPI.md#UnsubscribeContact) | **Post** /api/v2/unsubscribers | Manually unsubscribe a contact
[**UpdateBalanceNotificationSettings**](TextMagicAPI.md#UpdateBalanceNotificationSettings) | **Put** /api/v2/user/notification/balance | Update balance notification settings
[**UpdateCallbackSettings**](TextMagicAPI.md#UpdateCallbackSettings) | **Put** /api/v2/callback/settings | Update callback URL settings
[**UpdateChatDesktopNotificationSettings**](TextMagicAPI.md#UpdateChatDesktopNotificationSettings) | **Put** /api/v2/user/desktop/notification | Update chat desktop notification settings
[**UpdateContact**](TextMagicAPI.md#UpdateContact) | **Put** /api/v2/contacts/{id}/normalized | Edit a contact
[**UpdateContactNote**](TextMagicAPI.md#UpdateContactNote) | **Put** /api/v2/notes/{id} | Update a contact note
[**UpdateCurrentUser**](TextMagicAPI.md#UpdateCurrentUser) | **Put** /api/v2/user | Edit current account info
[**UpdateCustomField**](TextMagicAPI.md#UpdateCustomField) | **Put** /api/v2/customfields/{id} | Edit a custom field
[**UpdateCustomFieldValue**](TextMagicAPI.md#UpdateCustomFieldValue) | **Put** /api/v2/customfields/{id}/update | Edit the custom field value of a specified contact
[**UpdateInboundMessagesNotificationSettings**](TextMagicAPI.md#UpdateInboundMessagesNotificationSettings) | **Put** /api/v2/user/notification/inbound | Update inbound messages notification settings
[**UpdateList**](TextMagicAPI.md#UpdateList) | **Put** /api/v2/lists/{id} | Edit a list
[**UpdateSenderSetting**](TextMagicAPI.md#UpdateSenderSetting) | **Put** /api/v2/sender/settings | Change sender settings
[**UpdateTemplate**](TextMagicAPI.md#UpdateTemplate) | **Put** /api/v2/templates/{id} | Update a template
[**UploadAvatar**](TextMagicAPI.md#UploadAvatar) | **Post** /api/v2/user/avatar | Upload an avatar
[**UploadContactAvatar**](TextMagicAPI.md#UploadContactAvatar) | **Post** /api/v2/contacts/{id}/avatar | Upload an avatar
[**UploadListAvatar**](TextMagicAPI.md#UploadListAvatar) | **Post** /api/v2/lists/{id}/avatar | Add an avatar for a list
[**UploadMessageAttachment**](TextMagicAPI.md#UploadMessageAttachment) | **Post** /api/v2/messages/attachment | Upload message attachment
[**UploadMessageMMSAttachment**](TextMagicAPI.md#UploadMessageMMSAttachment) | **Post** /api/v2/messages/mms/attachment | Upload message mms attachment



## AssignContactsToList

> ResourceLinkResponse AssignContactsToList(ctx, id).AssignContactsToListInputObject(assignContactsToListInputObject).Execute()

Assign contacts to a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	assignContactsToListInputObject := *openapiclient.NewAssignContactsToListRequest() // AssignContactsToListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.AssignContactsToList(context.Background(), id).AssignContactsToListInputObject(assignContactsToListInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.AssignContactsToList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignContactsToList`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.AssignContactsToList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignContactsToListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignContactsToListInputObject** | [**AssignContactsToListRequest**](AssignContactsToListRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BlockContact

> ResourceLinkResponse BlockContact(ctx).BlockContactInputObject(blockContactInputObject).Execute()

Block a contact by phone number



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	blockContactInputObject := *openapiclient.NewBlockContactRequest() // BlockContactRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.BlockContact(context.Background()).BlockContactInputObject(blockContactInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.BlockContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockContact`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.BlockContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBlockContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blockContactInputObject** | [**BlockContactRequest**](BlockContactRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuyDedicatedNumber

> BuyDedicatedNumber(ctx).BuyDedicatedNumberInputObject(buyDedicatedNumberInputObject).Execute()

Buy a dedicated number



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	buyDedicatedNumberInputObject := *openapiclient.NewBuyDedicatedNumberRequest() // BuyDedicatedNumberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.BuyDedicatedNumber(context.Background()).BuyDedicatedNumberInputObject(buyDedicatedNumberInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.BuyDedicatedNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBuyDedicatedNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **buyDedicatedNumberInputObject** | [**BuyDedicatedNumberRequest**](BuyDedicatedNumberRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ClearAndAssignContactsToList

> ResourceLinkResponse ClearAndAssignContactsToList(ctx, id).ClearAndAssignContactsToListInputObject(clearAndAssignContactsToListInputObject).Execute()

Reset list members to the specified contacts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	clearAndAssignContactsToListInputObject := *openapiclient.NewClearAndAssignContactsToListRequest() // ClearAndAssignContactsToListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.ClearAndAssignContactsToList(context.Background(), id).ClearAndAssignContactsToListInputObject(clearAndAssignContactsToListInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.ClearAndAssignContactsToList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClearAndAssignContactsToList`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.ClearAndAssignContactsToList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearAndAssignContactsToListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clearAndAssignContactsToListInputObject** | [**ClearAndAssignContactsToListRequest**](ClearAndAssignContactsToListRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloseChatsBulk

> CloseChatsBulk(ctx).CloseChatsBulkInputObject(closeChatsBulkInputObject).Execute()

Close chats (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	closeChatsBulkInputObject := *openapiclient.NewMarkChatsUnreadBulkRequest() // MarkChatsUnreadBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.CloseChatsBulk(context.Background()).CloseChatsBulkInputObject(closeChatsBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.CloseChatsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloseChatsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **closeChatsBulkInputObject** | [**MarkChatsUnreadBulkRequest**](MarkChatsUnreadBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloseReadChats

> CloseReadChats(ctx).Execute()

Close read chats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.CloseReadChats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.CloseReadChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloseReadChatsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContact

> ResourceLinkResponse CreateContact(ctx).CreateContactInputObject(createContactInputObject).Execute()

Add a new contact

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	createContactInputObject := *openapiclient.NewCreateContactRequest() // CreateContactRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.CreateContact(context.Background()).CreateContactInputObject(createContactInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.CreateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContact`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.CreateContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createContactInputObject** | [**CreateContactRequest**](CreateContactRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContactNote

> ResourceLinkResponse CreateContactNote(ctx, id).CreateContactNoteInputObject(createContactNoteInputObject).Execute()

Create a new contact note

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	createContactNoteInputObject := *openapiclient.NewCreateContactNoteRequest() // CreateContactNoteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.CreateContactNote(context.Background(), id).CreateContactNoteInputObject(createContactNoteInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.CreateContactNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateContactNote`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.CreateContactNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContactNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createContactNoteInputObject** | [**CreateContactNoteRequest**](CreateContactNoteRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCustomField

> ResourceLinkResponse CreateCustomField(ctx).CreateCustomFieldInputObject(createCustomFieldInputObject).Execute()

Add a new custom field

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	createCustomFieldInputObject := *openapiclient.NewCreateCustomFieldRequest() // CreateCustomFieldRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.CreateCustomField(context.Background()).CreateCustomFieldInputObject(createCustomFieldInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.CreateCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomField`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.CreateCustomField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createCustomFieldInputObject** | [**CreateCustomFieldRequest**](CreateCustomFieldRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmailCampaign

> CreateEmailCampaignResponse CreateEmailCampaign(ctx).CreateEmailCampaignInputObject(createEmailCampaignInputObject).Execute()

Send email campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	createEmailCampaignInputObject := *openapiclient.NewCreateEmailCampaignRequest() // CreateEmailCampaignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.CreateEmailCampaign(context.Background()).CreateEmailCampaignInputObject(createEmailCampaignInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.CreateEmailCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmailCampaign`: CreateEmailCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.CreateEmailCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEmailCampaignInputObject** | [**CreateEmailCampaignRequest**](CreateEmailCampaignRequest.md) |  | 

### Return type

[**CreateEmailCampaignResponse**](CreateEmailCampaignResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateList

> ResourceLinkResponse CreateList(ctx).CreateListInputObject(createListInputObject).Execute()

Create a new list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	createListInputObject := *openapiclient.NewCreateListRequest() // CreateListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.CreateList(context.Background()).CreateListInputObject(createListInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.CreateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateList`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.CreateList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createListInputObject** | [**CreateListRequest**](CreateListRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTag

> CreateTagResponse CreateTag(ctx).CreateTagInputObject(createTagInputObject).Execute()

Create tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	createTagInputObject := *openapiclient.NewCreateTagRequest() // CreateTagRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.CreateTag(context.Background()).CreateTagInputObject(createTagInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.CreateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTag`: CreateTagResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTagInputObject** | [**CreateTagRequest**](CreateTagRequest.md) |  | 

### Return type

[**CreateTagResponse**](CreateTagResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTemplate

> ResourceLinkResponse CreateTemplate(ctx).CreateTemplateInputObject(createTemplateInputObject).Execute()

Create a template



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	createTemplateInputObject := *openapiclient.NewCreateTemplateRequest() // CreateTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.CreateTemplate(context.Background()).CreateTemplateInputObject(createTemplateInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.CreateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTemplate`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.CreateTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTemplateInputObject** | [**CreateTemplateRequest**](CreateTemplateRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllContacts

> DeleteAllContacts(ctx).Execute()

Delete contacts (bulk)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteAllContacts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteAllContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllContactsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllOutboundMessages

> DeleteAllOutboundMessages(ctx).Execute()

Delete all messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteAllOutboundMessages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteAllOutboundMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllOutboundMessagesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAvatar

> DeleteAvatar(ctx).Execute()

Delete an avatar

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteAvatar(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAvatarRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChatMessages

> DeleteChatMessages(ctx, id).DeleteChatMessagesBulkInputObject(deleteChatMessagesBulkInputObject).Execute()

Delete chat messages by ID(s)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	deleteChatMessagesBulkInputObject := *openapiclient.NewDeleteChatMessagesRequest() // DeleteChatMessagesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteChatMessages(context.Background(), id).DeleteChatMessagesBulkInputObject(deleteChatMessagesBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteChatMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChatMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteChatMessagesBulkInputObject** | [**DeleteChatMessagesRequest**](DeleteChatMessagesRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteChatsBulk

> DeleteChatsBulk(ctx).DeleteChatsBulkInputObject(deleteChatsBulkInputObject).Execute()

Delete chats (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	deleteChatsBulkInputObject := *openapiclient.NewDeleteChatsBulkRequest() // DeleteChatsBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteChatsBulk(context.Background()).DeleteChatsBulkInputObject(deleteChatsBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteChatsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChatsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteChatsBulkInputObject** | [**DeleteChatsBulkRequest**](DeleteChatsBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContact

> DeleteContact(ctx, id).Execute()

Delete a contact



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteContact(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactAvatar

> DeleteContactAvatar(ctx, id).Execute()

Delete an avatar

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteContactAvatar(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteContactAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactNote

> DeleteContactNote(ctx, id).Execute()

Delete a contact note

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteContactNote(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteContactNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactNotesBulk

> DeleteContactNotesBulk(ctx, id).DeleteContactNotesBulkInputObject(deleteContactNotesBulkInputObject).Execute()

Delete contact notes (bulk)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	deleteContactNotesBulkInputObject := *openapiclient.NewDeleteContactNotesBulkRequest() // DeleteContactNotesBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteContactNotesBulk(context.Background(), id).DeleteContactNotesBulkInputObject(deleteContactNotesBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteContactNotesBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactNotesBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteContactNotesBulkInputObject** | [**DeleteContactNotesBulkRequest**](DeleteContactNotesBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactsByIds

> DeleteContactsByIds(ctx).DeleteContactsByIdsInputObject(deleteContactsByIdsInputObject).Execute()

Delete contacts by IDs (bulk)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	deleteContactsByIdsInputObject := *openapiclient.NewDeleteContactsByIdsRequest() // DeleteContactsByIdsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteContactsByIds(context.Background()).DeleteContactsByIdsInputObject(deleteContactsByIdsInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteContactsByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactsByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteContactsByIdsInputObject** | [**DeleteContactsByIdsRequest**](DeleteContactsByIdsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContactsFromList

> DeleteContactsFromList(ctx, id).DeleteContacsFromListObject(deleteContacsFromListObject).Execute()

Unassign contacts from a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	deleteContacsFromListObject := *openapiclient.NewDeleteContactsFromListRequest() // DeleteContactsFromListRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteContactsFromList(context.Background(), id).DeleteContacsFromListObject(deleteContacsFromListObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteContactsFromList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContactsFromListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteContacsFromListObject** | [**DeleteContactsFromListRequest**](DeleteContactsFromListRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomField

> DeleteCustomField(ctx, id).Execute()

Delete a custom field



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteCustomField(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDedicatedNumber

> DeleteDedicatedNumber(ctx, id).Execute()

Cancel a dedicated number subscription

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteDedicatedNumber(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteDedicatedNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDedicatedNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInboundMessage

> DeleteInboundMessage(ctx, id).Execute()

Delete a single inbound message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | The unique numeric ID for the inbound message.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteInboundMessage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteInboundMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The unique numeric ID for the inbound message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInboundMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInboundMessagesBulk

> DeleteInboundMessagesBulk(ctx).DeleteInboundMessagesBulkInputObject(deleteInboundMessagesBulkInputObject).Execute()

Delete inbound messages (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	deleteInboundMessagesBulkInputObject := *openapiclient.NewDeleteListsBulkRequest() // DeleteListsBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteInboundMessagesBulk(context.Background()).DeleteInboundMessagesBulkInputObject(deleteInboundMessagesBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteInboundMessagesBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInboundMessagesBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteInboundMessagesBulkInputObject** | [**DeleteListsBulkRequest**](DeleteListsBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteList

> DeleteList(ctx, id).Execute()

Delete a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteListAvatar

> DeleteListAvatar(ctx, id).Execute()

Delete an avatar for a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteListAvatar(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteListAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteListContactsBulk

> DeleteListContactsBulk(ctx, id).DeleteListContactsBulkInputObject(deleteListContactsBulkInputObject).Execute()

Delete contacts from a list (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	deleteListContactsBulkInputObject := *openapiclient.NewUnblockContactsBulkRequest() // UnblockContactsBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteListContactsBulk(context.Background(), id).DeleteListContactsBulkInputObject(deleteListContactsBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteListContactsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListContactsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteListContactsBulkInputObject** | [**UnblockContactsBulkRequest**](UnblockContactsBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteListsBulk

> DeleteListsBulk(ctx).DeleteListsBulkInputObject(deleteListsBulkInputObject).Execute()

Delete lists (bulk)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	deleteListsBulkInputObject := *openapiclient.NewDeleteListsBulkRequest() // DeleteListsBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteListsBulk(context.Background()).DeleteListsBulkInputObject(deleteListsBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteListsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteListsBulkInputObject** | [**DeleteListsBulkRequest**](DeleteListsBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessageSession

> DeleteMessageSession(ctx, id).Execute()

Delete a session



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteMessageSession(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteMessageSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMessageSessionsBulk

> DeleteMessageSessionsBulk(ctx).DeleteMessageSessionsBulkInputObject(deleteMessageSessionsBulkInputObject).Execute()

Delete sessions (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	deleteMessageSessionsBulkInputObject := *openapiclient.NewDeleteListsBulkRequest() // DeleteListsBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteMessageSessionsBulk(context.Background()).DeleteMessageSessionsBulkInputObject(deleteMessageSessionsBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteMessageSessionsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMessageSessionsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteMessageSessionsBulkInputObject** | [**DeleteListsBulkRequest**](DeleteListsBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOutboundMessage

> DeleteOutboundMessage(ctx, id).Execute()

Delete message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteOutboundMessage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteOutboundMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOutboundMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOutboundMessagesBulk

> DeleteOutboundMessagesBulk(ctx).DeleteOutboundMessagesBulkInputObject(deleteOutboundMessagesBulkInputObject).Execute()

Delete messages (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	deleteOutboundMessagesBulkInputObject := *openapiclient.NewDeleteOutboundMessagesBulkRequest() // DeleteOutboundMessagesBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteOutboundMessagesBulk(context.Background()).DeleteOutboundMessagesBulkInputObject(deleteOutboundMessagesBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteOutboundMessagesBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOutboundMessagesBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteOutboundMessagesBulkInputObject** | [**DeleteOutboundMessagesBulkRequest**](DeleteOutboundMessagesBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScheduledMessage

> DeleteScheduledMessage(ctx, id).Execute()

Delete a single scheduled message

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteScheduledMessage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteScheduledMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduledMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScheduledMessagesBulk

> DeleteScheduledMessagesBulk(ctx).DeleteScheduledMessagesBulkInputObject(deleteScheduledMessagesBulkInputObject).Execute()

Delete scheduled messages (bulk)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	deleteScheduledMessagesBulkInputObject := *openapiclient.NewDeleteScheduledMessagesBulkRequest() // DeleteScheduledMessagesBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteScheduledMessagesBulk(context.Background()).DeleteScheduledMessagesBulkInputObject(deleteScheduledMessagesBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteScheduledMessagesBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScheduledMessagesBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteScheduledMessagesBulkInputObject** | [**DeleteScheduledMessagesBulkRequest**](DeleteScheduledMessagesBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSenderId

> DeleteSenderId(ctx, id).Execute()

Delete a Sender ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteSenderId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteSenderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSenderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplate

> DeleteTemplate(ctx, id).Execute()

Delete a template

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTemplatesBulk

> DeleteTemplatesBulk(ctx).DeleteTemplatesBulkInputObject(deleteTemplatesBulkInputObject).Execute()

Delete templates (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	deleteTemplatesBulkInputObject := *openapiclient.NewDeleteContactNotesBulkRequest() // DeleteContactNotesBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.DeleteTemplatesBulk(context.Background()).DeleteTemplatesBulkInputObject(deleteTemplatesBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DeleteTemplatesBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTemplatesBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteTemplatesBulkInputObject** | [**DeleteContactNotesBulkRequest**](DeleteContactNotesBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoCarrierLookup

> DoCarrierLookupResponse DoCarrierLookup(ctx, phone).Country(country).Execute()

Carrier Lookup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	phone := "447860021130" // string | Phone number in [E.164 format](https://en.wikipedia.org/wiki/E.164) or in [National format](https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers). 
	country := "GB" // string | This option must be specified only if the phone number is in a **[National format](https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers)**.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.DoCarrierLookup(context.Background(), phone).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DoCarrierLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoCarrierLookup`: DoCarrierLookupResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.DoCarrierLookup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phone** | **string** | Phone number in [E.164 format](https://en.wikipedia.org/wiki/E.164) or in [National format](https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers).  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoCarrierLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | **string** | This option must be specified only if the phone number is in a **[National format](https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers)**.  | 

### Return type

[**DoCarrierLookupResponse**](DoCarrierLookupResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoEmailLookup

> DoEmailLookupResponse DoEmailLookup(ctx, email).Execute()

Email Lookup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	email := "john@sample.com" // string | Email address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.DoEmailLookup(context.Background(), email).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.DoEmailLookup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoEmailLookup`: DoEmailLookupResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.DoEmailLookup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | Email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoEmailLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DoEmailLookupResponse**](DoEmailLookupResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllBulkSessions

> GetAllBulkSessionsPaginatedResponse GetAllBulkSessions(ctx).Page(page).Limit(limit).Execute()

Get all bulk sessions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetAllBulkSessions(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetAllBulkSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllBulkSessions`: GetAllBulkSessionsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetAllBulkSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllBulkSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]

### Return type

[**GetAllBulkSessionsPaginatedResponse**](GetAllBulkSessionsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllChats

> GetAllChatsPaginatedResponse GetAllChats(ctx).Status(status).Page(page).Limit(limit).OrderBy(orderBy).Voice(voice).Flat(flat).Execute()

Get all chats

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	status := "a" // string | Fetch only (a)ctive, (c)losed or (d)eleted chats. (optional)
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")
	voice := int32(56) // int32 | Fetch results with voice calls. (optional) (default to 0)
	flat := int32(1) // int32 | Should additional contact info be included? (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetAllChats(context.Background()).Status(status).Page(page).Limit(limit).OrderBy(orderBy).Voice(voice).Flat(flat).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetAllChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllChats`: GetAllChatsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetAllChats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllChatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Fetch only (a)ctive, (c)losed or (d)eleted chats. | 
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]
 **voice** | **int32** | Fetch results with voice calls. | [default to 0]
 **flat** | **int32** | Should additional contact info be included? | [default to 0]

### Return type

[**GetAllChatsPaginatedResponse**](GetAllChatsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllInboundMessages

> GetAllInboundMessagesPaginatedResponse GetAllInboundMessages(ctx).Page(page).Limit(limit).OrderBy(orderBy).Direction(direction).Execute()

Get all inbound messages

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")
	direction := "direction_example" // string | Order direction. Default is desc. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetAllInboundMessages(context.Background()).Page(page).Limit(limit).OrderBy(orderBy).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetAllInboundMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllInboundMessages`: GetAllInboundMessagesPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetAllInboundMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllInboundMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]
 **direction** | **string** | Order direction. Default is desc. | [default to &quot;desc&quot;]

### Return type

[**GetAllInboundMessagesPaginatedResponse**](GetAllInboundMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllMessageSessions

> GetAllMessageSessionsPaginatedResponse GetAllMessageSessions(ctx).Page(page).Limit(limit).Execute()

Get all sessions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetAllMessageSessions(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetAllMessageSessions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllMessageSessions`: GetAllMessageSessionsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetAllMessageSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMessageSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]

### Return type

[**GetAllMessageSessionsPaginatedResponse**](GetAllMessageSessionsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllOutboundMessages

> GetAllOutboundMessagesPaginatedResponse GetAllOutboundMessages(ctx).Page(page).Limit(limit).LastId(lastId).Execute()

Get all messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	lastId := int32(56) // int32 | Filter results by ID, selecting all values lesser than the specified ID. Note that the \\'page\\' parameter is ignored when \\'lastId\\' is specified. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetAllOutboundMessages(context.Background()).Page(page).Limit(limit).LastId(lastId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetAllOutboundMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllOutboundMessages`: GetAllOutboundMessagesPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetAllOutboundMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllOutboundMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **lastId** | **int32** | Filter results by ID, selecting all values lesser than the specified ID. Note that the \\&#39;page\\&#39; parameter is ignored when \\&#39;lastId\\&#39; is specified. | 

### Return type

[**GetAllOutboundMessagesPaginatedResponse**](GetAllOutboundMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllScheduledMessages

> GetAllScheduledMessagesPaginatedResponse GetAllScheduledMessages(ctx).Page(page).Limit(limit).Status(status).OrderBy(orderBy).Direction(direction).Execute()

Get all scheduled messages

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	status := "status_example" // string | Fetch schedules with a specific status: a - actual, c - completed, x - all. (optional) (default to "x")
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")
	direction := "direction_example" // string | Order direction. Default is desc. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetAllScheduledMessages(context.Background()).Page(page).Limit(limit).Status(status).OrderBy(orderBy).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetAllScheduledMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllScheduledMessages`: GetAllScheduledMessagesPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetAllScheduledMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllScheduledMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **status** | **string** | Fetch schedules with a specific status: a - actual, c - completed, x - all. | [default to &quot;x&quot;]
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]
 **direction** | **string** | Order direction. Default is desc. | [default to &quot;desc&quot;]

### Return type

[**GetAllScheduledMessagesPaginatedResponse**](GetAllScheduledMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTemplates

> GetAllTemplatesPaginatedResponse GetAllTemplates(ctx).Page(page).Limit(limit).Execute()

Get all templates

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(1) // int32 | Fetch specified results page. (optional)
	limit := int32(10) // int32 | The number of results per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetAllTemplates(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetAllTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllTemplates`: GetAllTemplatesPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetAllTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | 
 **limit** | **int32** | The number of results per page. | 

### Return type

[**GetAllTemplatesPaginatedResponse**](GetAllTemplatesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableDedicatedNumbers

> GetAvailableDedicatedNumbersResponse GetAvailableDedicatedNumbers(ctx).Country(country).Prefix(prefix).Tollfree(tollfree).Execute()

Find dedicated numbers available for purchase

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	country := "GB" // string | The 2-letter dedicated number country ISO code.
	prefix := int32(447155) // int32 | Desired number prefix. Should include the country code (i.e. 447 for UK phone number format). Leave blank to get all the available numbers for the specified country. (optional)
	tollfree := int32(56) // int32 | Should we show only tollfree numbers (tollfree available only for US). (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetAvailableDedicatedNumbers(context.Background()).Country(country).Prefix(prefix).Tollfree(tollfree).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetAvailableDedicatedNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableDedicatedNumbers`: GetAvailableDedicatedNumbersResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetAvailableDedicatedNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableDedicatedNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | The 2-letter dedicated number country ISO code. | 
 **prefix** | **int32** | Desired number prefix. Should include the country code (i.e. 447 for UK phone number format). Leave blank to get all the available numbers for the specified country. | 
 **tollfree** | **int32** | Should we show only tollfree numbers (tollfree available only for US). | [default to 0]

### Return type

[**GetAvailableDedicatedNumbersResponse**](GetAvailableDedicatedNumbersResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvailableSenderSettingOptions

> GetAvailableSenderSettingOptionsResponse GetAvailableSenderSettingOptions(ctx).Country(country).Execute()

Get available sender settings



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	country := "US" // string | The 2-letter ISO country ID. If not specified, it returns all the available sender settings. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetAvailableSenderSettingOptions(context.Background()).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetAvailableSenderSettingOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableSenderSettingOptions`: GetAvailableSenderSettingOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetAvailableSenderSettingOptions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableSenderSettingOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | The 2-letter ISO country ID. If not specified, it returns all the available sender settings. | 

### Return type

[**GetAvailableSenderSettingOptionsResponse**](GetAvailableSenderSettingOptionsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalanceNotificationOptions

> GetBalanceNotificationOptionsResponse GetBalanceNotificationOptions(ctx).Execute()

Returns the list of available balance options which can be used as a bound to determine when to send email to user with low balance notification. See https://my.textmagic.com/online/account/notifications/balance

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetBalanceNotificationOptions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetBalanceNotificationOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalanceNotificationOptions`: GetBalanceNotificationOptionsResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetBalanceNotificationOptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceNotificationOptionsRequest struct via the builder pattern


### Return type

[**GetBalanceNotificationOptionsResponse**](GetBalanceNotificationOptionsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBalanceNotificationSettings

> GetBalanceNotificationSettingsResponse GetBalanceNotificationSettings(ctx).Execute()

Get balance notification settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetBalanceNotificationSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetBalanceNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBalanceNotificationSettings`: GetBalanceNotificationSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetBalanceNotificationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBalanceNotificationSettingsRequest struct via the builder pattern


### Return type

[**GetBalanceNotificationSettingsResponse**](GetBalanceNotificationSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockedContacts

> GetBlockedContactsPaginatedResponse GetBlockedContacts(ctx).Page(page).Limit(limit).Query(query).OrderBy(orderBy).Direction(direction).Execute()

Get blocked contacts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	query := "query_example" // string | Find blocked contacts by specified search query. (optional)
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")
	direction := "direction_example" // string | Order direction. Default is desc. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetBlockedContacts(context.Background()).Page(page).Limit(limit).Query(query).OrderBy(orderBy).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetBlockedContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockedContacts`: GetBlockedContactsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetBlockedContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockedContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **query** | **string** | Find blocked contacts by specified search query. | 
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]
 **direction** | **string** | Order direction. Default is desc. | [default to &quot;desc&quot;]

### Return type

[**GetBlockedContactsPaginatedResponse**](GetBlockedContactsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkSession

> BulkSession GetBulkSession(ctx, id).Execute()

Get bulk session status

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetBulkSession(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetBulkSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkSession`: BulkSession
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetBulkSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BulkSession**](BulkSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCallbackSettings

> GetCallbackSettingsResponse GetCallbackSettings(ctx).Execute()

Fetch callback URL settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetCallbackSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetCallbackSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallbackSettings`: GetCallbackSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetCallbackSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallbackSettingsRequest struct via the builder pattern


### Return type

[**GetCallbackSettingsResponse**](GetCallbackSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChat

> Chat GetChat(ctx, id).Execute()

Get a single chat

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetChat(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChat`: Chat
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Chat**](Chat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatByPhone

> Chat GetChatByPhone(ctx, phone).Upsert(upsert).Reopen(reopen).Execute()

Find chats by phone

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	phone := "447860021130" // string | 
	upsert := int32(56) // int32 | Create a new chat if not found. (optional) (default to 0)
	reopen := int32(56) // int32 | Reopen chat if found or do not change status. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetChatByPhone(context.Background(), phone).Upsert(upsert).Reopen(reopen).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetChatByPhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatByPhone`: Chat
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetChatByPhone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatByPhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsert** | **int32** | Create a new chat if not found. | [default to 0]
 **reopen** | **int32** | Reopen chat if found or do not change status. | [default to 0]

### Return type

[**Chat**](Chat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChatMessages

> GetChatMessagesPaginatedResponse GetChatMessages(ctx, id).Page(page).Limit(limit).Query(query).Start(start).End(end).Direction(direction).Voice(voice).IncludeNotes(includeNotes).Execute()

Get chat messages

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	query := "query_example" // string | Find messages by specified search query. (optional)
	start := "start_example" // string | Return messages since specified timestamp only. Required when `end` parameter specified. (optional)
	end := "end_example" // string | Return messages up to specified timestamp only. Required when `start` parameter specified. (optional)
	direction := "direction_example" // string | Order direction. Default is desc. (optional) (default to "desc")
	voice := int32(56) // int32 | Fetch results with voice calls. (optional) (default to 0)
	includeNotes := int32(56) // int32 | Fetch results with messenger notes. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetChatMessages(context.Background(), id).Page(page).Limit(limit).Query(query).Start(start).End(end).Direction(direction).Voice(voice).IncludeNotes(includeNotes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetChatMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChatMessages`: GetChatMessagesPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetChatMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChatMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **query** | **string** | Find messages by specified search query. | 
 **start** | **string** | Return messages since specified timestamp only. Required when &#x60;end&#x60; parameter specified. | 
 **end** | **string** | Return messages up to specified timestamp only. Required when &#x60;start&#x60; parameter specified. | 
 **direction** | **string** | Order direction. Default is desc. | [default to &quot;desc&quot;]
 **voice** | **int32** | Fetch results with voice calls. | [default to 0]
 **includeNotes** | **int32** | Fetch results with messenger notes. | [default to 0]

### Return type

[**GetChatMessagesPaginatedResponse**](GetChatMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContact

> Contact GetContact(ctx, id).Execute()

Get the details of a specific contact

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | Contact ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetContact(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContact`: Contact
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Contact ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Contact**](Contact.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactByPhone

> Contact GetContactByPhone(ctx, phone).Execute()

Get the details of a specific contact by phone number

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	phone := "447860021130" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetContactByPhone(context.Background(), phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetContactByPhone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactByPhone`: Contact
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetContactByPhone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phone** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactByPhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Contact**](Contact.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactIfBlocked

> Contact GetContactIfBlocked(ctx).Phone(phone).Execute()

Check if a phone number is blocked

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	phone := "447860021130" // string | Phone number to check.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetContactIfBlocked(context.Background()).Phone(phone).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetContactIfBlocked``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactIfBlocked`: Contact
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetContactIfBlocked`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContactIfBlockedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phone** | **string** | Phone number to check. | 

### Return type

[**Contact**](Contact.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactImportSessionProgress

> GetContactImportSessionProgressResponse GetContactImportSessionProgress(ctx, id).Execute()

Check import progress



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetContactImportSessionProgress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetContactImportSessionProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactImportSessionProgress`: GetContactImportSessionProgressResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetContactImportSessionProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactImportSessionProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContactImportSessionProgressResponse**](GetContactImportSessionProgressResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactNote

> ContactNote GetContactNote(ctx, id).Execute()

Get a contact note

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetContactNote(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetContactNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactNote`: ContactNote
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetContactNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContactNote**](ContactNote.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactNotes

> GetContactNotesPaginatedResponse GetContactNotes(ctx, id).Page(page).Limit(limit).Execute()

Fetch notes assigned to a given contact

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetContactNotes(context.Background(), id).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetContactNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactNotes`: GetContactNotesPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetContactNotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]

### Return type

[**GetContactNotesPaginatedResponse**](GetContactNotesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContacts

> GetContactsPaginatedResponse GetContacts(ctx).Page(page).Limit(limit).Shared(shared).OrderBy(orderBy).Direction(direction).Execute()

Get all contacts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	shared := int32(56) // int32 | Should shared contacts be included? (optional) (default to 0)
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")
	direction := "direction_example" // string | Order direction. Default is desc. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetContacts(context.Background()).Page(page).Limit(limit).Shared(shared).OrderBy(orderBy).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContacts`: GetContactsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **shared** | **int32** | Should shared contacts be included? | [default to 0]
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]
 **direction** | **string** | Order direction. Default is desc. | [default to &quot;desc&quot;]

### Return type

[**GetContactsPaginatedResponse**](GetContactsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactsAutocomplete

> []GetContactsAutocompleteResponseItem GetContactsAutocomplete(ctx).Query(query).Limit(limit).Lists(lists).Execute()

Get contacts autocomplete suggestions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	query := "A" // string | Find recipients by specified search query.
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	lists := int32(56) // int32 | Should lists be returned or not? (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetContactsAutocomplete(context.Background()).Query(query).Limit(limit).Lists(lists).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetContactsAutocomplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactsAutocomplete`: []GetContactsAutocompleteResponseItem
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetContactsAutocomplete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContactsAutocompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Find recipients by specified search query. | 
 **limit** | **int32** | The number of results per page. | [default to 10]
 **lists** | **int32** | Should lists be returned or not? | [default to 0]

### Return type

[**[]GetContactsAutocompleteResponseItem**](GetContactsAutocompleteResponseItem.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContactsByListId

> GetContactsByListIdPaginatedResponse GetContactsByListId(ctx, id).Page(page).Limit(limit).OrderBy(orderBy).Direction(direction).Execute()

Get all contacts in a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | Given group ID.
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")
	direction := "direction_example" // string | Order direction. Default is desc. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetContactsByListId(context.Background(), id).Page(page).Limit(limit).OrderBy(orderBy).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetContactsByListId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContactsByListId`: GetContactsByListIdPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetContactsByListId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Given group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContactsByListIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]
 **direction** | **string** | Order direction. Default is desc. | [default to &quot;desc&quot;]

### Return type

[**GetContactsByListIdPaginatedResponse**](GetContactsByListIdPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountries

> []Country GetCountries(ctx).Execute()

Get countries

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetCountries(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetCountries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCountries`: []Country
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


### Return type

[**[]Country**](Country.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentUser

> User GetCurrentUser(ctx).Execute()

Get current account information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetCurrentUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUser`: User
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern


### Return type

[**User**](User.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomField

> UserCustomField GetCustomField(ctx, id).Execute()

Get the details of a specific custom field

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetCustomField(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomField`: UserCustomField
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserCustomField**](UserCustomField.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomFields

> GetCustomFieldsPaginatedResponse GetCustomFields(ctx).Page(page).Limit(limit).Execute()

Get all custom fields

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetCustomFields(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetCustomFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomFields`: GetCustomFieldsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetCustomFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]

### Return type

[**GetCustomFieldsPaginatedResponse**](GetCustomFieldsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDedicatedNumber

> UsersInbound GetDedicatedNumber(ctx, id).Execute()

Get the details of a specific dedicated number

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetDedicatedNumber(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetDedicatedNumber``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDedicatedNumber`: UsersInbound
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetDedicatedNumber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDedicatedNumberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UsersInbound**](UsersInbound.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEmailSenders

> GetEmailSendersResponse GetEmailSenders(ctx).DomainId(domainId).Execute()

Get list of email senders



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	domainId := int32(56) // int32 | Filter email senders by specific domain ID. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetEmailSenders(context.Background()).DomainId(domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetEmailSenders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEmailSenders`: GetEmailSendersResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetEmailSenders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailSendersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **int32** | Filter email senders by specific domain ID. | 

### Return type

[**GetEmailSendersResponse**](GetEmailSendersResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFavorites

> GetFavoritesPaginatedResponse GetFavorites(ctx).Page(page).Limit(limit).Query(query).Execute()

Get favorite contacts and lists

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	query := "A" // string | Find contacts or lists by specified search query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetFavorites(context.Background()).Page(page).Limit(limit).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetFavorites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFavorites`: GetFavoritesPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetFavorites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFavoritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **query** | **string** | Find contacts or lists by specified search query. | 

### Return type

[**GetFavoritesPaginatedResponse**](GetFavoritesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboundMessage

> MessageIn GetInboundMessage(ctx, id).Execute()

Get a single inbound message

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1782832) // int32 | The unique numeric ID for the inbound message.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetInboundMessage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetInboundMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboundMessage`: MessageIn
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetInboundMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The unique numeric ID for the inbound message. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboundMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageIn**](MessageIn.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInboundMessagesNotificationSettings

> GetInboundMessagesNotificationSettingsResponse GetInboundMessagesNotificationSettings(ctx).Execute()

Get inbound messages notification settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetInboundMessagesNotificationSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetInboundMessagesNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInboundMessagesNotificationSettings`: GetInboundMessagesNotificationSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetInboundMessagesNotificationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboundMessagesNotificationSettingsRequest struct via the builder pattern


### Return type

[**GetInboundMessagesNotificationSettingsResponse**](GetInboundMessagesNotificationSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInvoices

> GetInvoicesPaginatedResponse GetInvoices(ctx).Page(page).Limit(limit).Execute()

Get all invoices



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetInvoices(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetInvoices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInvoices`: GetInvoicesPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetInvoices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]

### Return type

[**GetInvoicesPaginatedResponse**](GetInvoicesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetList

> List GetList(ctx, id).Execute()

Get the details of a specific list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetList(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetList`: List
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**List**](List.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListContactsIds

> []int32 GetListContactsIds(ctx, id).Execute()

Get all contact IDs in a list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetListContactsIds(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetListContactsIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListContactsIds`: []int32
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetListContactsIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListContactsIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]int32**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLists

> GetListsPaginatedResponse GetLists(ctx).Page(page).Limit(limit).OrderBy(orderBy).Direction(direction).FavoriteOnly(favoriteOnly).OnlyMine(onlyMine).Execute()

Get all lists

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | The current fetched page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")
	direction := "direction_example" // string | Order direction. Default is desc. (optional) (default to "desc")
	favoriteOnly := int32(56) // int32 | Return only favorited lists. (optional) (default to 0)
	onlyMine := int32(56) // int32 | Return only current user lists. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetLists(context.Background()).Page(page).Limit(limit).OrderBy(orderBy).Direction(direction).FavoriteOnly(favoriteOnly).OnlyMine(onlyMine).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLists`: GetListsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The current fetched page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]
 **direction** | **string** | Order direction. Default is desc. | [default to &quot;desc&quot;]
 **favoriteOnly** | **int32** | Return only favorited lists. | [default to 0]
 **onlyMine** | **int32** | Return only current user lists. | [default to 0]

### Return type

[**GetListsPaginatedResponse**](GetListsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListsOfContact

> GetListsOfContactPaginatedResponse GetListsOfContact(ctx, id).Page(page).Limit(limit).Execute()

Get a contact's lists



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetListsOfContact(context.Background(), id).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetListsOfContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetListsOfContact`: GetListsOfContactPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetListsOfContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListsOfContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]

### Return type

[**GetListsOfContactPaginatedResponse**](GetListsOfContactPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessagePreview

> GetMessagePreviewResponse GetMessagePreview(ctx).Text(text).TemplateId(templateId).SendingTime(sendingTime).SendingDateTime(sendingDateTime).SendingTimezone(sendingTimezone).Contacts(contacts).Lists(lists).Phones(phones).CutExtra(cutExtra).PartsCount(partsCount).ReferenceId(referenceId).From(from).Rule(rule).CreateChat(createChat).Tts(tts).Local(local).LocalCountry(localCountry).Execute()

Preview message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	text := "Test message test" // string | Message text. Required if **template_id** is not set. (optional)
	templateId := int32(1) // int32 | Template used instead of message text. Required if **text** is not set. (optional)
	sendingTime := int32(1565606455) // int32 | DEPRECATED, consider using sendingDateTime and sendingTimezone parameters instead: Optional (required with rrule set). Message sending time is in unix timestamp format. Default is now. (optional)
	sendingDateTime := "2020-05-27 13:02:33" // string | Sending time is in Y-m-d H:i:s format (e.g. 2016-05-27 13:02:33). This time is relative to the sendingTimezone. (optional)
	sendingTimezone := "America/Buenos_Aires" // string | The ID or ISO-name of the timezone used for sending when the sendingDateTime parameter is set, e.g. if you specify sendingDateTime = \\\"2016-05-27 13:02:33\\\" and sendingTimezone = \\\"America/Buenos_Aires\\\", your message will be sent on May 27, 2016 13:02:33 Buenos Aires time, or 16:02:33 UTC. Default is the account timezone. (optional)
	contacts := "1,2,3,4" // string | Comma-separated array of contact resources id message will be sent to. (optional)
	lists := "1,2,3,4" // string | Comma-separated array of list resources id message will be sent to. (optional)
	phones := "447860021130,447860021131" // string | Comma-separated array of E.164 phone numbers message will be sent to. (optional)
	cutExtra := int32(56) // int32 | Should sending method cut extra characters which not fit supplied partsCount or return 400 Bad request response instead. (optional) (default to 0)
	partsCount := int32(56) // int32 | Maximum message parts count (Textmagic allows sending of 1 to 6 message parts). (optional) (default to 6)
	referenceId := int32(1) // int32 | Custom message reference id which can be used in your application infrastructure. (optional)
	from := "Test Sender ID" // string | One of the allowed Sender ID (phone number or alphanumeric sender ID). If the specified Sender ID is not allowed for some destinations, a fallback default Sender ID will be used to ensure delivery. See [Get timezones](https://docs.textmagic.com/#tag/Sender-IDs). (optional)
	rule := "FREQ=YEARLY;BYMONTH=1;BYMONTHDAY=1;COUNT=1" // string | An iCal RRULE parameter to create recurrent scheduled messages. When used, sendingTime is mandatory as the start point of sending. See https://www.textmagic.com/free-tools/rrule-generator for format details. (optional)
	createChat := int32(56) // int32 | Should the sending method try to create new Chat(if not exist) with specified recipients? (optional) (default to 0)
	tts := int32(56) // int32 | Send Text-to-Speech message. (optional) (default to 0)
	local := int32(56) // int32 | Treat phone numbers passed in the \\'phones\\' field as local. (optional) (default to 0)
	localCountry := "US" // string | The 2-letter ISO country code for local phone numbers, used when \\'local\\' is set to true. Default is the account country. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetMessagePreview(context.Background()).Text(text).TemplateId(templateId).SendingTime(sendingTime).SendingDateTime(sendingDateTime).SendingTimezone(sendingTimezone).Contacts(contacts).Lists(lists).Phones(phones).CutExtra(cutExtra).PartsCount(partsCount).ReferenceId(referenceId).From(from).Rule(rule).CreateChat(createChat).Tts(tts).Local(local).LocalCountry(localCountry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetMessagePreview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessagePreview`: GetMessagePreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetMessagePreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagePreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **string** | Message text. Required if **template_id** is not set. | 
 **templateId** | **int32** | Template used instead of message text. Required if **text** is not set. | 
 **sendingTime** | **int32** | DEPRECATED, consider using sendingDateTime and sendingTimezone parameters instead: Optional (required with rrule set). Message sending time is in unix timestamp format. Default is now. | 
 **sendingDateTime** | **string** | Sending time is in Y-m-d H:i:s format (e.g. 2016-05-27 13:02:33). This time is relative to the sendingTimezone. | 
 **sendingTimezone** | **string** | The ID or ISO-name of the timezone used for sending when the sendingDateTime parameter is set, e.g. if you specify sendingDateTime &#x3D; \\\&quot;2016-05-27 13:02:33\\\&quot; and sendingTimezone &#x3D; \\\&quot;America/Buenos_Aires\\\&quot;, your message will be sent on May 27, 2016 13:02:33 Buenos Aires time, or 16:02:33 UTC. Default is the account timezone. | 
 **contacts** | **string** | Comma-separated array of contact resources id message will be sent to. | 
 **lists** | **string** | Comma-separated array of list resources id message will be sent to. | 
 **phones** | **string** | Comma-separated array of E.164 phone numbers message will be sent to. | 
 **cutExtra** | **int32** | Should sending method cut extra characters which not fit supplied partsCount or return 400 Bad request response instead. | [default to 0]
 **partsCount** | **int32** | Maximum message parts count (Textmagic allows sending of 1 to 6 message parts). | [default to 6]
 **referenceId** | **int32** | Custom message reference id which can be used in your application infrastructure. | 
 **from** | **string** | One of the allowed Sender ID (phone number or alphanumeric sender ID). If the specified Sender ID is not allowed for some destinations, a fallback default Sender ID will be used to ensure delivery. See [Get timezones](https://docs.textmagic.com/#tag/Sender-IDs). | 
 **rule** | **string** | An iCal RRULE parameter to create recurrent scheduled messages. When used, sendingTime is mandatory as the start point of sending. See https://www.textmagic.com/free-tools/rrule-generator for format details. | 
 **createChat** | **int32** | Should the sending method try to create new Chat(if not exist) with specified recipients? | [default to 0]
 **tts** | **int32** | Send Text-to-Speech message. | [default to 0]
 **local** | **int32** | Treat phone numbers passed in the \\&#39;phones\\&#39; field as local. | [default to 0]
 **localCountry** | **string** | The 2-letter ISO country code for local phone numbers, used when \\&#39;local\\&#39; is set to true. Default is the account country. | 

### Return type

[**GetMessagePreviewResponse**](GetMessagePreviewResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessagePrice

> GetMessagePriceResponse GetMessagePrice(ctx).IncludeBlocked(includeBlocked).Text(text).TemplateId(templateId).SendingTime(sendingTime).SendingDateTime(sendingDateTime).SendingTimezone(sendingTimezone).Contacts(contacts).Lists(lists).Phones(phones).CutExtra(cutExtra).PartsCount(partsCount).ReferenceId(referenceId).From(from).Rule(rule).CreateChat(createChat).Tts(tts).Local(local).LocalCountry(localCountry).Execute()

Check message price



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	includeBlocked := int32(56) // int32 | Should we show the pricing for blocked contacts? (optional) (default to 0)
	text := "Test message test" // string | Message text. Required if the **template_id** is not set. (optional)
	templateId := int32(1) // int32 | Template used instead of message text. Required if the **text** is not set. (optional)
	sendingTime := int32(1565606455) // int32 | DEPRECATED, consider using the sendingDateTime and sendingTimezone parameters instead: optional (required with rrule set). Message sending time is in unix timestamp format. Default is now. (optional)
	sendingDateTime := "2020-05-27 13:02:33" // string | Sending time is in Y-m-d H:i:s format (e.g. 2016-05-27 13:02:33). This time is relative to the sendingTimezone. (optional)
	sendingTimezone := "America/Buenos_Aires" // string | The ID or ISO-name of the timezone used for sending when sendingDateTime parameter is set, e.g. if you specify sendingDateTime = \\\"2016-05-27 13:02:33\\\" and sendingTimezone = \\\"America/Buenos_Aires\\\", your message will be sent on May 27, 2016 13:02:33 Buenos Aires time, or 16:02:33 UTC. Default is the account timezone. (optional)
	contacts := "1,2,3,4" // string | Comma-separated array of contact resources id message will be sent to. (optional)
	lists := "1,2,3,4" // string | Comma-separated array of list resources id message will be sent to. (optional)
	phones := "447860021130,447860021131" // string | Comma-separated array of E.164 phone numbers message will be sent to. (optional)
	cutExtra := int32(56) // int32 | Should sending method cut extra characters which not fit supplied partsCount or return 400 Bad request response instead. (optional) (default to 0)
	partsCount := int32(56) // int32 | Maximum message parts count (Textmagic allows sending 1 to 6 message parts). (optional) (default to 6)
	referenceId := int32(1) // int32 | Custom message reference id which can be used in your application infrastructure. (optional)
	from := "Test Sender ID" // string | One of the allowed Sender ID (phone number or alphanumeric sender ID). If the specified Sender ID is not allowed for some destinations, a fallback default Sender ID will be used to ensure delivery. See [Get timezones](https://docs.textmagic.com/#tag/Sender-IDs). (optional)
	rule := "FREQ=YEARLY;BYMONTH=1;BYMONTHDAY=1;COUNT=1" // string | An iCal RRULE parameter to create recurrent scheduled messages. When used, sendingTime is mandatory as the start point of sending. See https://www.textmagic.com/free-tools/rrule-generator for format details. (optional)
	createChat := int32(56) // int32 | Should the sending method try to create new Chat (if not exist) with specified recipients? (optional) (default to 0)
	tts := int32(56) // int32 | Send a Text-to-Speech message. (optional) (default to 0)
	local := int32(56) // int32 | Treat phone numbers passed in the \\'phones\\' field as local. (optional) (default to 0)
	localCountry := "US" // string | The 2-letter ISO country code for local phone numbers, used when \\'local\\' is set to true. Default is the account country. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetMessagePrice(context.Background()).IncludeBlocked(includeBlocked).Text(text).TemplateId(templateId).SendingTime(sendingTime).SendingDateTime(sendingDateTime).SendingTimezone(sendingTimezone).Contacts(contacts).Lists(lists).Phones(phones).CutExtra(cutExtra).PartsCount(partsCount).ReferenceId(referenceId).From(from).Rule(rule).CreateChat(createChat).Tts(tts).Local(local).LocalCountry(localCountry).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetMessagePrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessagePrice`: GetMessagePriceResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetMessagePrice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagePriceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeBlocked** | **int32** | Should we show the pricing for blocked contacts? | [default to 0]
 **text** | **string** | Message text. Required if the **template_id** is not set. | 
 **templateId** | **int32** | Template used instead of message text. Required if the **text** is not set. | 
 **sendingTime** | **int32** | DEPRECATED, consider using the sendingDateTime and sendingTimezone parameters instead: optional (required with rrule set). Message sending time is in unix timestamp format. Default is now. | 
 **sendingDateTime** | **string** | Sending time is in Y-m-d H:i:s format (e.g. 2016-05-27 13:02:33). This time is relative to the sendingTimezone. | 
 **sendingTimezone** | **string** | The ID or ISO-name of the timezone used for sending when sendingDateTime parameter is set, e.g. if you specify sendingDateTime &#x3D; \\\&quot;2016-05-27 13:02:33\\\&quot; and sendingTimezone &#x3D; \\\&quot;America/Buenos_Aires\\\&quot;, your message will be sent on May 27, 2016 13:02:33 Buenos Aires time, or 16:02:33 UTC. Default is the account timezone. | 
 **contacts** | **string** | Comma-separated array of contact resources id message will be sent to. | 
 **lists** | **string** | Comma-separated array of list resources id message will be sent to. | 
 **phones** | **string** | Comma-separated array of E.164 phone numbers message will be sent to. | 
 **cutExtra** | **int32** | Should sending method cut extra characters which not fit supplied partsCount or return 400 Bad request response instead. | [default to 0]
 **partsCount** | **int32** | Maximum message parts count (Textmagic allows sending 1 to 6 message parts). | [default to 6]
 **referenceId** | **int32** | Custom message reference id which can be used in your application infrastructure. | 
 **from** | **string** | One of the allowed Sender ID (phone number or alphanumeric sender ID). If the specified Sender ID is not allowed for some destinations, a fallback default Sender ID will be used to ensure delivery. See [Get timezones](https://docs.textmagic.com/#tag/Sender-IDs). | 
 **rule** | **string** | An iCal RRULE parameter to create recurrent scheduled messages. When used, sendingTime is mandatory as the start point of sending. See https://www.textmagic.com/free-tools/rrule-generator for format details. | 
 **createChat** | **int32** | Should the sending method try to create new Chat (if not exist) with specified recipients? | [default to 0]
 **tts** | **int32** | Send a Text-to-Speech message. | [default to 0]
 **local** | **int32** | Treat phone numbers passed in the \\&#39;phones\\&#39; field as local. | [default to 0]
 **localCountry** | **string** | The 2-letter ISO country code for local phone numbers, used when \\&#39;local\\&#39; is set to true. Default is the account country. | 

### Return type

[**GetMessagePriceResponse**](GetMessagePriceResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageSession

> MessageSession GetMessageSession(ctx, id).Execute()

Get a session`s details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | Session ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetMessageSession(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetMessageSession``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageSession`: MessageSession
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetMessageSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | Session ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageSession**](MessageSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageSessionStat

> GetMessageSessionStatResponse GetMessageSessionStat(ctx, id).IncludeDeleted(includeDeleted).Execute()

Get a session`s statistics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	includeDeleted := int32(56) // int32 | Search also in deleted messages. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetMessageSessionStat(context.Background(), id).IncludeDeleted(includeDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetMessageSessionStat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessageSessionStat`: GetMessageSessionStatResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetMessageSessionStat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageSessionStatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDeleted** | **int32** | Search also in deleted messages. | [default to 0]

### Return type

[**GetMessageSessionStatResponse**](GetMessageSessionStatResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessagesBySessionId

> GetMessagesBySessionIdPaginatedResponse GetMessagesBySessionId(ctx, id).Page(page).Limit(limit).Statuses(statuses).IncludeDeleted(includeDeleted).Execute()

Get a session`s messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	statuses := "statuses_example" // string | Find messages by status. (optional)
	includeDeleted := int32(56) // int32 | Search also in deleted messages. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetMessagesBySessionId(context.Background(), id).Page(page).Limit(limit).Statuses(statuses).IncludeDeleted(includeDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetMessagesBySessionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessagesBySessionId`: GetMessagesBySessionIdPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetMessagesBySessionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagesBySessionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **statuses** | **string** | Find messages by status. | 
 **includeDeleted** | **int32** | Search also in deleted messages. | [default to 0]

### Return type

[**GetMessagesBySessionIdPaginatedResponse**](GetMessagesBySessionIdPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessagingCounters

> GetMessagingCountersResponse GetMessagingCounters(ctx).Execute()

Get sent/received messages counters values



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetMessagingCounters(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetMessagingCounters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessagingCounters`: GetMessagingCountersResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetMessagingCounters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagingCountersRequest struct via the builder pattern


### Return type

[**GetMessagingCountersResponse**](GetMessagingCountersResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessagingStat

> []MessagingStatItem GetMessagingStat(ctx).By(by).Start(start).End(end).Execute()

Get messaging statistics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	by := "off" // string | *   **off** - to get total values per specified time interval; *   **day** - to show values grouped by day; *   **month** - to show values grouped by month; *   **year** - to show values grouped by year.  (optional) (default to "off")
	start := int32(1430438400) // int32 | Time period start in [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time) format. The default is 7 days prior.  (optional)
	end := int32(1431648000) // int32 | Time period start in [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time) format. The default is today.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetMessagingStat(context.Background()).By(by).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetMessagingStat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMessagingStat`: []MessagingStatItem
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetMessagingStat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMessagingStatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **by** | **string** | *   **off** - to get total values per specified time interval; *   **day** - to show values grouped by day; *   **month** - to show values grouped by month; *   **year** - to show values grouped by year.  | [default to &quot;off&quot;]
 **start** | **int32** | Time period start in [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time) format. The default is 7 days prior.  | 
 **end** | **int32** | Time period start in [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time) format. The default is today.  | 

### Return type

[**[]MessagingStatItem**](MessagingStatItem.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutboundMessage

> MessageOut GetOutboundMessage(ctx, id).Execute()

Get a single message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetOutboundMessage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetOutboundMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutboundMessage`: MessageOut
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetOutboundMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutboundMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageOut**](MessageOut.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutboundMessagesHistory

> GetOutboundMessagesHistoryPaginatedResponse GetOutboundMessagesHistory(ctx).Limit(limit).LastId(lastId).Query(query).OrderBy(orderBy).Direction(direction).Execute()

Get history



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	lastId := int32(56) // int32 | Filter results by ID, selecting all values lesser than the specified ID. (optional)
	query := "query_example" // string | Find message by specified search query. (optional)
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")
	direction := "direction_example" // string | Order direction. Default is desc. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetOutboundMessagesHistory(context.Background()).Limit(limit).LastId(lastId).Query(query).OrderBy(orderBy).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetOutboundMessagesHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutboundMessagesHistory`: GetOutboundMessagesHistoryPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetOutboundMessagesHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOutboundMessagesHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The number of results per page. | [default to 10]
 **lastId** | **int32** | Filter results by ID, selecting all values lesser than the specified ID. | 
 **query** | **string** | Find message by specified search query. | 
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]
 **direction** | **string** | Order direction. Default is desc. | [default to &quot;desc&quot;]

### Return type

[**GetOutboundMessagesHistoryPaginatedResponse**](GetOutboundMessagesHistoryPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScheduledMessage

> MessagesIcs GetScheduledMessage(ctx, id).Execute()

Get a single scheduled message

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetScheduledMessage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetScheduledMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScheduledMessage`: MessagesIcs
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetScheduledMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduledMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessagesIcs**](MessagesIcs.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSenderId

> SenderId GetSenderId(ctx, id).Execute()

Get the details of a specific Sender ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetSenderId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetSenderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSenderId`: SenderId
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetSenderId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSenderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SenderId**](SenderId.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSenderIds

> GetSenderIdsPaginatedResponse GetSenderIds(ctx).Page(page).Limit(limit).Execute()

Get all your approved Sender IDs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetSenderIds(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetSenderIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSenderIds`: GetSenderIdsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetSenderIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSenderIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]

### Return type

[**GetSenderIdsPaginatedResponse**](GetSenderIdsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSenderSettings

> GetSenderSettingsResponse GetSenderSettings(ctx).Country(country).Execute()

Get current sender settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	country := "US" // string | Return sender settings enabled for sending to a specified country. Should be 2 upper-case characters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetSenderSettings(context.Background()).Country(country).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetSenderSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSenderSettings`: GetSenderSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetSenderSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSenderSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** | Return sender settings enabled for sending to a specified country. Should be 2 upper-case characters. | 

### Return type

[**GetSenderSettingsResponse**](GetSenderSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSpendingStat

> GetSpendingStatPaginatedResponse GetSpendingStat(ctx).Page(page).Limit(limit).Start(start).End(end).Execute()

Get spending statistics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	start := "2018-11-11 11:11" // string | Time period start in [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time) format. The default is 7 days prior.  (optional)
	end := "2019-11-11 11:11" // string | Time period start in [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time) format. The default is today.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetSpendingStat(context.Background()).Page(page).Limit(limit).Start(start).End(end).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetSpendingStat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSpendingStat`: GetSpendingStatPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetSpendingStat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSpendingStatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **start** | **string** | Time period start in [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time) format. The default is 7 days prior.  | 
 **end** | **string** | Time period start in [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time) format. The default is today.  | 

### Return type

[**GetSpendingStatPaginatedResponse**](GetSpendingStatPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTemplate

> MessageTemplate GetTemplate(ctx, id).Execute()

Get a template`s details



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTemplate`: MessageTemplate
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTimezones

> map[string]interface{} GetTimezones(ctx).Full(full).Execute()

Get timezones



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	full := int32(56) // int32 | Return full info about timezones in array (0 or 1). Default is 0. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetTimezones(context.Background()).Full(full).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetTimezones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTimezones`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetTimezones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTimezonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **full** | **int32** | Return full info about timezones in array (0 or 1). Default is 0. | [default to 0]

### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnreadMessagesTotal

> GetUnreadMessagesTotalResponse GetUnreadMessagesTotal(ctx).Execute()

Get unread messages number



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetUnreadMessagesTotal(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetUnreadMessagesTotal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnreadMessagesTotal`: GetUnreadMessagesTotalResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetUnreadMessagesTotal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnreadMessagesTotalRequest struct via the builder pattern


### Return type

[**GetUnreadMessagesTotalResponse**](GetUnreadMessagesTotalResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnsubscribedContact

> UnsubscribedContact GetUnsubscribedContact(ctx, id).Execute()

Get the details of a specific unsubscribed contact

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetUnsubscribedContact(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetUnsubscribedContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnsubscribedContact`: UnsubscribedContact
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetUnsubscribedContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnsubscribedContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UnsubscribedContact**](UnsubscribedContact.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnsubscribers

> GetUnsubscribersPaginatedResponse GetUnsubscribers(ctx).Page(page).Limit(limit).Execute()

Get all unsubscribed contacts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetUnsubscribers(context.Background()).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetUnsubscribers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnsubscribers`: GetUnsubscribersPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetUnsubscribers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUnsubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]

### Return type

[**GetUnsubscribersPaginatedResponse**](GetUnsubscribersPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserDedicatedNumbers

> GetUserDedicatedNumbersPaginatedResponse GetUserDedicatedNumbers(ctx).Page(page).Limit(limit).SurveyId(surveyId).Execute()

Get all your dedicated numbers

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	surveyId := int32(56) // int32 | Fetch only those numbers that are ready for the survey. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.GetUserDedicatedNumbers(context.Background()).Page(page).Limit(limit).SurveyId(surveyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.GetUserDedicatedNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserDedicatedNumbers`: GetUserDedicatedNumbersPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.GetUserDedicatedNumbers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserDedicatedNumbersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **surveyId** | **int32** | Fetch only those numbers that are ready for the survey. | 

### Return type

[**GetUserDedicatedNumbersPaginatedResponse**](GetUserDedicatedNumbersPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportContacts

> ResourceLinkResponse ImportContacts(ctx).Column(column).File(file).ListId(listId).ListName(listName).Execute()

Import contacts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	column := "0:firstName;1:lastName;3:phone;4:email" // string | Import file column mapping. The string must contain sub-strings of mapping in format `columnNumber:field` glued by `;`. For example: `0:firstName;1:lastName;3:phone;4:email` where the value before `:` is a number of the column in the file, and the value after `:` is a field of the newly created contact or the ID of a custom field. Numbers of columns begin from zero. Allowed built-in contact fields are: `firstName`, `lastName`, `phone`, `email`. Existing of `phone` mapping is required. 
	file := os.NewFile(1234, "some_file") // *os.File | File containing contacts in csv or xls(x) formats.
	listId := int32(443) // int32 | List that ID contacts will be imported to. Ignored if `listName` is specified.  (optional)
	listName := "A new list" // string | List name. This list will be created during import. If such name is already taken, an ordinal (1, 2, ...) will be added to the end. Ignored if `listId` is specified.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.ImportContacts(context.Background()).Column(column).File(file).ListId(listId).ListName(listName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.ImportContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ImportContacts`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.ImportContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **column** | **string** | Import file column mapping. The string must contain sub-strings of mapping in format &#x60;columnNumber:field&#x60; glued by &#x60;;&#x60;. For example: &#x60;0:firstName;1:lastName;3:phone;4:email&#x60; where the value before &#x60;:&#x60; is a number of the column in the file, and the value after &#x60;:&#x60; is a field of the newly created contact or the ID of a custom field. Numbers of columns begin from zero. Allowed built-in contact fields are: &#x60;firstName&#x60;, &#x60;lastName&#x60;, &#x60;phone&#x60;, &#x60;email&#x60;. Existing of &#x60;phone&#x60; mapping is required.  | 
 **file** | ***os.File** | File containing contacts in csv or xls(x) formats. | 
 **listId** | **int32** | List that ID contacts will be imported to. Ignored if &#x60;listName&#x60; is specified.  | 
 **listName** | **string** | List name. This list will be created during import. If such name is already taken, an ordinal (1, 2, ...) will be added to the end. Ignored if &#x60;listId&#x60; is specified.  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkChatsReadBulk

> MarkChatsReadBulk(ctx).MarkChatsReadBulkInputObject(markChatsReadBulkInputObject).Execute()

Mark chats as read (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	markChatsReadBulkInputObject := *openapiclient.NewMarkChatsUnreadBulkRequest() // MarkChatsUnreadBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.MarkChatsReadBulk(context.Background()).MarkChatsReadBulkInputObject(markChatsReadBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.MarkChatsReadBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkChatsReadBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markChatsReadBulkInputObject** | [**MarkChatsUnreadBulkRequest**](MarkChatsUnreadBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkChatsUnreadBulk

> MarkChatsUnreadBulk(ctx).MarkChatsUnreadBulkInputObject(markChatsUnreadBulkInputObject).Execute()

Mark chats as unread (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	markChatsUnreadBulkInputObject := *openapiclient.NewMarkChatsUnreadBulkRequest() // MarkChatsUnreadBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.MarkChatsUnreadBulk(context.Background()).MarkChatsUnreadBulkInputObject(markChatsUnreadBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.MarkChatsUnreadBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkChatsUnreadBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markChatsUnreadBulkInputObject** | [**MarkChatsUnreadBulkRequest**](MarkChatsUnreadBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteChat

> ResourceLinkResponse MuteChat(ctx).MuteChatInputObject(muteChatInputObject).Execute()

Mute chat sounds

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	muteChatInputObject := *openapiclient.NewMuteChatRequest() // MuteChatRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.MuteChat(context.Background()).MuteChatInputObject(muteChatInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.MuteChat``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MuteChat`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.MuteChat`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMuteChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **muteChatInputObject** | [**MuteChatRequest**](MuteChatRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteChatsBulk

> MuteChatsBulk(ctx).MuteChatsBulkInputObject(muteChatsBulkInputObject).Execute()

Mute chats (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	muteChatsBulkInputObject := *openapiclient.NewMuteChatsBulkRequest() // MuteChatsBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.MuteChatsBulk(context.Background()).MuteChatsBulkInputObject(muteChatsBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.MuteChatsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMuteChatsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **muteChatsBulkInputObject** | [**MuteChatsBulkRequest**](MuteChatsBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ping

> PingResponse Ping(ctx).Execute()

Ping



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.Ping(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.Ping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Ping`: PingResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPingRequest struct via the builder pattern


### Return type

[**PingResponse**](PingResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReopenChatsBulk

> ReopenChatsBulk(ctx).ReopenChatsBulkInputObject(reopenChatsBulkInputObject).Execute()

Reopen chats (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	reopenChatsBulkInputObject := *openapiclient.NewMarkChatsUnreadBulkRequest() // MarkChatsUnreadBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.ReopenChatsBulk(context.Background()).ReopenChatsBulkInputObject(reopenChatsBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.ReopenChatsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReopenChatsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reopenChatsBulkInputObject** | [**MarkChatsUnreadBulkRequest**](MarkChatsUnreadBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestSenderId

> ResourceLinkResponse RequestSenderId(ctx).RequestSenderIdInputObject(requestSenderIdInputObject).Execute()

Apply for a new Sender ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	requestSenderIdInputObject := *openapiclient.NewRequestSenderIdRequest() // RequestSenderIdRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.RequestSenderId(context.Background()).RequestSenderIdInputObject(requestSenderIdInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.RequestSenderId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestSenderId`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.RequestSenderId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestSenderIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestSenderIdInputObject** | [**RequestSenderIdRequest**](RequestSenderIdRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleEmailCampaign

> ScheduleEmailCampaignResponse ScheduleEmailCampaign(ctx).ScheduleEmailCampaignInputObject(scheduleEmailCampaignInputObject).Execute()

Schedule new email campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	scheduleEmailCampaignInputObject := *openapiclient.NewScheduleEmailCampaignRequest() // ScheduleEmailCampaignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.ScheduleEmailCampaign(context.Background()).ScheduleEmailCampaignInputObject(scheduleEmailCampaignInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.ScheduleEmailCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleEmailCampaign`: ScheduleEmailCampaignResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.ScheduleEmailCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleEmailCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scheduleEmailCampaignInputObject** | [**ScheduleEmailCampaignRequest**](ScheduleEmailCampaignRequest.md) |  | 

### Return type

[**ScheduleEmailCampaignResponse**](ScheduleEmailCampaignResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchChats

> SearchChatsPaginatedResponse SearchChats(ctx).Page(page).Limit(limit).Query(query).Execute()

Find chats by message text

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	query := "query_example" // string | Find chats by specified search query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.SearchChats(context.Background()).Page(page).Limit(limit).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.SearchChats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchChats`: SearchChatsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.SearchChats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchChatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **query** | **string** | Find chats by specified search query. | 

### Return type

[**SearchChatsPaginatedResponse**](SearchChatsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchChatsByIds

> SearchChatsByIdsPaginatedResponse SearchChatsByIds(ctx).Page(page).Limit(limit).Ids(ids).Execute()

Find chats (bulk)

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	ids := "ids_example" // string | Find chats by ID(s). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.SearchChatsByIds(context.Background()).Page(page).Limit(limit).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.SearchChatsByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchChatsByIds`: SearchChatsByIdsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.SearchChatsByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchChatsByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **ids** | **string** | Find chats by ID(s). | 

### Return type

[**SearchChatsByIdsPaginatedResponse**](SearchChatsByIdsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchChatsByReceipent

> SearchChatsByReceipentPaginatedResponse SearchChatsByReceipent(ctx).Page(page).Limit(limit).Query(query).OrderBy(orderBy).Execute()

Find chats by recipient



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	query := "query_example" // string | Find chats by specified search query. (optional)
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.SearchChatsByReceipent(context.Background()).Page(page).Limit(limit).Query(query).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.SearchChatsByReceipent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchChatsByReceipent`: SearchChatsByReceipentPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.SearchChatsByReceipent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchChatsByReceipentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **query** | **string** | Find chats by specified search query. | 
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]

### Return type

[**SearchChatsByReceipentPaginatedResponse**](SearchChatsByReceipentPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchContacts

> SearchContactsPaginatedResponse SearchContacts(ctx).Page(page).Limit(limit).Shared(shared).Ids(ids).ListId(listId).IncludeBlocked(includeBlocked).Query(query).Local(local).ExactMatch(exactMatch).Country(country).OrderBy(orderBy).Direction(direction).TagIds(tagIds).Execute()

Find contacts by given criteria

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	shared := int32(56) // int32 | Should shared contacts be included? (optional) (default to 0)
	ids := "ids_example" // string | Find contacts by IDs. (optional)
	listId := int32(56) // int32 | Find contacts by List ID. (optional)
	includeBlocked := int32(56) // int32 | Should blocked contacts be included? (optional)
	query := "query_example" // string | Find contacts by specified search query. (optional)
	local := int32(56) // int32 | Treat phone number passed in the \"query\" field as local. Default is 0. (optional) (default to 0)
	exactMatch := int32(56) // int32 | Return only exactly matching contacts. Default is 0. (optional) (default to 0)
	country := "country_example" // string | The 2-letter ISO country code for local phone numbers, used when \"local\" is set to true. Default is the account country. (optional)
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")
	direction := "direction_example" // string | Order direction. Default is desc. (optional) (default to "desc")
	tagIds := "tagIds_example" // string | Find contacts by tag ID(s). Multiple IDs can be separated by comma. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.SearchContacts(context.Background()).Page(page).Limit(limit).Shared(shared).Ids(ids).ListId(listId).IncludeBlocked(includeBlocked).Query(query).Local(local).ExactMatch(exactMatch).Country(country).OrderBy(orderBy).Direction(direction).TagIds(tagIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.SearchContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchContacts`: SearchContactsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.SearchContacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **shared** | **int32** | Should shared contacts be included? | [default to 0]
 **ids** | **string** | Find contacts by IDs. | 
 **listId** | **int32** | Find contacts by List ID. | 
 **includeBlocked** | **int32** | Should blocked contacts be included? | 
 **query** | **string** | Find contacts by specified search query. | 
 **local** | **int32** | Treat phone number passed in the \&quot;query\&quot; field as local. Default is 0. | [default to 0]
 **exactMatch** | **int32** | Return only exactly matching contacts. Default is 0. | [default to 0]
 **country** | **string** | The 2-letter ISO country code for local phone numbers, used when \&quot;local\&quot; is set to true. Default is the account country. | 
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]
 **direction** | **string** | Order direction. Default is desc. | [default to &quot;desc&quot;]
 **tagIds** | **string** | Find contacts by tag ID(s). Multiple IDs can be separated by comma. | 

### Return type

[**SearchContactsPaginatedResponse**](SearchContactsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchInboundMessages

> SearchInboundMessagesPaginatedResponse SearchInboundMessages(ctx).Page(page).Limit(limit).Ids(ids).Query(query).OrderBy(orderBy).Direction(direction).Expand(expand).Execute()

Find inbound messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	ids := "ids_example" // string | Find message by ID(s). (optional)
	query := "query_example" // string | Find recipients by specified search query. (optional)
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")
	direction := "direction_example" // string | Order direction. Default is desc. (optional) (default to "desc")
	expand := int32(56) // int32 | Expand by adding firstName, lastName and contactId. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.SearchInboundMessages(context.Background()).Page(page).Limit(limit).Ids(ids).Query(query).OrderBy(orderBy).Direction(direction).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.SearchInboundMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchInboundMessages`: SearchInboundMessagesPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.SearchInboundMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchInboundMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **ids** | **string** | Find message by ID(s). | 
 **query** | **string** | Find recipients by specified search query. | 
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]
 **direction** | **string** | Order direction. Default is desc. | [default to &quot;desc&quot;]
 **expand** | **int32** | Expand by adding firstName, lastName and contactId. | [default to 0]

### Return type

[**SearchInboundMessagesPaginatedResponse**](SearchInboundMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchLists

> SearchListsPaginatedResponse SearchLists(ctx).Page(page).Limit(limit).Ids(ids).Query(query).OnlyMine(onlyMine).OnlyDefault(onlyDefault).OrderBy(orderBy).Direction(direction).Execute()

Find lists by given criteria

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	ids := "1,2,3,4" // string | Find lists by IDs. (optional)
	query := "A" // string | Find lists by specified search query. (optional)
	onlyMine := int32(56) // int32 | Return only current user lists. (optional) (default to 0)
	onlyDefault := int32(56) // int32 | Return only default lists. (optional) (default to 0)
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")
	direction := "direction_example" // string | Order direction. Default is desc. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.SearchLists(context.Background()).Page(page).Limit(limit).Ids(ids).Query(query).OnlyMine(onlyMine).OnlyDefault(onlyDefault).OrderBy(orderBy).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.SearchLists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchLists`: SearchListsPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.SearchLists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchListsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **ids** | **string** | Find lists by IDs. | 
 **query** | **string** | Find lists by specified search query. | 
 **onlyMine** | **int32** | Return only current user lists. | [default to 0]
 **onlyDefault** | **int32** | Return only default lists. | [default to 0]
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]
 **direction** | **string** | Order direction. Default is desc. | [default to &quot;desc&quot;]

### Return type

[**SearchListsPaginatedResponse**](SearchListsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchOutboundMessages

> SearchOutboundMessagesPaginatedResponse SearchOutboundMessages(ctx).Page(page).Limit(limit).LastId(lastId).Ids(ids).SessionId(sessionId).Statuses(statuses).IncludeDeleted(includeDeleted).Query(query).Execute()

Find messages



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	lastId := int32(56) // int32 | Filter results by ID, selecting all values lesser than the specified ID. Note that the \\'page\\' parameter is ignored when \\'lastId\\' is specified. (optional)
	ids := "ids_example" // string | Find message by ID(s). (optional)
	sessionId := int32(56) // int32 | Find messages by session ID. (optional)
	statuses := "q" // string | Find messages by status. (optional)
	includeDeleted := int32(56) // int32 | Search also in deleted messages. (optional) (default to 0)
	query := "query_example" // string | Find messages by specified search query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.SearchOutboundMessages(context.Background()).Page(page).Limit(limit).LastId(lastId).Ids(ids).SessionId(sessionId).Statuses(statuses).IncludeDeleted(includeDeleted).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.SearchOutboundMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchOutboundMessages`: SearchOutboundMessagesPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.SearchOutboundMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchOutboundMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **lastId** | **int32** | Filter results by ID, selecting all values lesser than the specified ID. Note that the \\&#39;page\\&#39; parameter is ignored when \\&#39;lastId\\&#39; is specified. | 
 **ids** | **string** | Find message by ID(s). | 
 **sessionId** | **int32** | Find messages by session ID. | 
 **statuses** | **string** | Find messages by status. | 
 **includeDeleted** | **int32** | Search also in deleted messages. | [default to 0]
 **query** | **string** | Find messages by specified search query. | 

### Return type

[**SearchOutboundMessagesPaginatedResponse**](SearchOutboundMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchScheduledMessages

> SearchScheduledMessagesPaginatedResponse SearchScheduledMessages(ctx).Page(page).Limit(limit).Query(query).Ids(ids).Status(status).OrderBy(orderBy).Direction(direction).Execute()

Find scheduled messages

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	query := "query_example" // string | Find messages by specified search query. (optional)
	ids := "ids_example" // string | Find schedules by ID(s). (optional)
	status := "status_example" // string | Fetch schedules with a specific status: a - actual, c - completed, x - all. (optional) (default to "x")
	orderBy := "orderBy_example" // string | Order results by some field. Default is id. (optional) (default to "id")
	direction := "direction_example" // string | Order direction. Default is desc. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.SearchScheduledMessages(context.Background()).Page(page).Limit(limit).Query(query).Ids(ids).Status(status).OrderBy(orderBy).Direction(direction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.SearchScheduledMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchScheduledMessages`: SearchScheduledMessagesPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.SearchScheduledMessages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchScheduledMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **query** | **string** | Find messages by specified search query. | 
 **ids** | **string** | Find schedules by ID(s). | 
 **status** | **string** | Fetch schedules with a specific status: a - actual, c - completed, x - all. | [default to &quot;x&quot;]
 **orderBy** | **string** | Order results by some field. Default is id. | [default to &quot;id&quot;]
 **direction** | **string** | Order direction. Default is desc. | [default to &quot;desc&quot;]

### Return type

[**SearchScheduledMessagesPaginatedResponse**](SearchScheduledMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTemplates

> SearchTemplatesPaginatedResponse SearchTemplates(ctx).Page(page).Limit(limit).Ids(ids).Name(name).Content(content).Execute()

Find templates by criteria



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	page := int32(56) // int32 | Fetch specified results page. (optional) (default to 1)
	limit := int32(56) // int32 | The number of results per page. (optional) (default to 10)
	ids := "ids_example" // string | Find template by ID(s). (optional)
	name := "name_example" // string | Find template by name. (optional)
	content := "content_example" // string | Find template by content. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.SearchTemplates(context.Background()).Page(page).Limit(limit).Ids(ids).Name(name).Content(content).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.SearchTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTemplates`: SearchTemplatesPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.SearchTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Fetch specified results page. | [default to 1]
 **limit** | **int32** | The number of results per page. | [default to 10]
 **ids** | **string** | Find template by ID(s). | 
 **name** | **string** | Find template by name. | 
 **content** | **string** | Find template by content. | 

### Return type

[**SearchTemplatesPaginatedResponse**](SearchTemplatesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessage

> SendMessageResponse SendMessage(ctx).SendMessageInputObject(sendMessageInputObject).Execute()

Send message



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	sendMessageInputObject := *openapiclient.NewSendMessageRequest() // SendMessageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.SendMessage(context.Background()).SendMessageInputObject(sendMessageInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.SendMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendMessage`: SendMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.SendMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendMessageInputObject** | [**SendMessageRequest**](SendMessageRequest.md) |  | 

### Return type

[**SendMessageResponse**](SendMessageResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetChatStatus

> ResourceLinkResponse SetChatStatus(ctx).SetChatStatusInputObject(setChatStatusInputObject).Execute()

Change chat status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	setChatStatusInputObject := *openapiclient.NewSetChatStatusRequest() // SetChatStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.SetChatStatus(context.Background()).SetChatStatusInputObject(setChatStatusInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.SetChatStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetChatStatus`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.SetChatStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetChatStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setChatStatusInputObject** | [**SetChatStatusRequest**](SetChatStatusRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnblockContact

> UnblockContact(ctx).UnblockContactInputObject(unblockContactInputObject).Execute()

Unblock a contact by phone number



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	unblockContactInputObject := *openapiclient.NewBlockContactRequest() // BlockContactRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.UnblockContact(context.Background()).UnblockContactInputObject(unblockContactInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UnblockContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnblockContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unblockContactInputObject** | [**BlockContactRequest**](BlockContactRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnblockContactsBulk

> UnblockContactsBulk(ctx).UnblockContactsBulkInputObject(unblockContactsBulkInputObject).Execute()

Unblock contacts (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	unblockContactsBulkInputObject := *openapiclient.NewUnblockContactsBulkRequest() // UnblockContactsBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.UnblockContactsBulk(context.Background()).UnblockContactsBulkInputObject(unblockContactsBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UnblockContactsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnblockContactsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unblockContactsBulkInputObject** | [**UnblockContactsBulkRequest**](UnblockContactsBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnmuteChatsBulk

> UnmuteChatsBulk(ctx).UnmuteChatsBulkInputObject(unmuteChatsBulkInputObject).Execute()

Unmute chats (bulk)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	unmuteChatsBulkInputObject := *openapiclient.NewUnmuteChatsBulkRequest() // UnmuteChatsBulkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.UnmuteChatsBulk(context.Background()).UnmuteChatsBulkInputObject(unmuteChatsBulkInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UnmuteChatsBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnmuteChatsBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unmuteChatsBulkInputObject** | [**UnmuteChatsBulkRequest**](UnmuteChatsBulkRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsubscribeContact

> ResourceLinkResponse UnsubscribeContact(ctx).UnsubscribeContactInputObject(unsubscribeContactInputObject).Execute()

Manually unsubscribe a contact



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	unsubscribeContactInputObject := *openapiclient.NewUnsubscribeContactRequest() // UnsubscribeContactRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.UnsubscribeContact(context.Background()).UnsubscribeContactInputObject(unsubscribeContactInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UnsubscribeContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsubscribeContact`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.UnsubscribeContact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unsubscribeContactInputObject** | [**UnsubscribeContactRequest**](UnsubscribeContactRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBalanceNotificationSettings

> UpdateBalanceNotificationSettings(ctx).UpdateBalanceNotificationSettingsInputObject(updateBalanceNotificationSettingsInputObject).Execute()

Update balance notification settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	updateBalanceNotificationSettingsInputObject := *openapiclient.NewUpdateBalanceNotificationSettingsRequest() // UpdateBalanceNotificationSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.UpdateBalanceNotificationSettings(context.Background()).UpdateBalanceNotificationSettingsInputObject(updateBalanceNotificationSettingsInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UpdateBalanceNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBalanceNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateBalanceNotificationSettingsInputObject** | [**UpdateBalanceNotificationSettingsRequest**](UpdateBalanceNotificationSettingsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCallbackSettings

> UpdateCallbackSettings(ctx).UpdateCallbackSettingsInputObject(updateCallbackSettingsInputObject).Execute()

Update callback URL settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	updateCallbackSettingsInputObject := *openapiclient.NewUpdateCallbackSettingsRequest() // UpdateCallbackSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.UpdateCallbackSettings(context.Background()).UpdateCallbackSettingsInputObject(updateCallbackSettingsInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UpdateCallbackSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallbackSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCallbackSettingsInputObject** | [**UpdateCallbackSettingsRequest**](UpdateCallbackSettingsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateChatDesktopNotificationSettings

> UpdateChatDesktopNotificationSettings(ctx).UpdateChatDesktopNotificationSettingsInputObject(updateChatDesktopNotificationSettingsInputObject).Execute()

Update chat desktop notification settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	updateChatDesktopNotificationSettingsInputObject := *openapiclient.NewUpdateChatDesktopNotificationSettingsRequest() // UpdateChatDesktopNotificationSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.UpdateChatDesktopNotificationSettings(context.Background()).UpdateChatDesktopNotificationSettingsInputObject(updateChatDesktopNotificationSettingsInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UpdateChatDesktopNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChatDesktopNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateChatDesktopNotificationSettingsInputObject** | [**UpdateChatDesktopNotificationSettingsRequest**](UpdateChatDesktopNotificationSettingsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContact

> ResourceLinkResponse UpdateContact(ctx, id).UpdateContactInputObject(updateContactInputObject).Execute()

Edit a contact

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	updateContactInputObject := *openapiclient.NewUpdateContactRequest() // UpdateContactRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.UpdateContact(context.Background(), id).UpdateContactInputObject(updateContactInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UpdateContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContact`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.UpdateContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateContactInputObject** | [**UpdateContactRequest**](UpdateContactRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContactNote

> ResourceLinkResponse UpdateContactNote(ctx, id).UpdateContactNoteInputObject(updateContactNoteInputObject).Execute()

Update a contact note

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	updateContactNoteInputObject := *openapiclient.NewUpdateContactNoteRequest() // UpdateContactNoteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.UpdateContactNote(context.Background(), id).UpdateContactNoteInputObject(updateContactNoteInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UpdateContactNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateContactNote`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.UpdateContactNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContactNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateContactNoteInputObject** | [**UpdateContactNoteRequest**](UpdateContactNoteRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCurrentUser

> UpdateCurrentUserResponse UpdateCurrentUser(ctx).UpdateCurrentUserInputObject(updateCurrentUserInputObject).Execute()

Edit current account info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	updateCurrentUserInputObject := *openapiclient.NewUpdateCurrentUserRequest() // UpdateCurrentUserRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.UpdateCurrentUser(context.Background()).UpdateCurrentUserInputObject(updateCurrentUserInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UpdateCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCurrentUser`: UpdateCurrentUserResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.UpdateCurrentUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateCurrentUserInputObject** | [**UpdateCurrentUserRequest**](UpdateCurrentUserRequest.md) |  | 

### Return type

[**UpdateCurrentUserResponse**](UpdateCurrentUserResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomField

> ResourceLinkResponse UpdateCustomField(ctx, id).UpdateCustomFieldInputObject(updateCustomFieldInputObject).Execute()

Edit a custom field

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	updateCustomFieldInputObject := *openapiclient.NewCreateCustomFieldRequest() // CreateCustomFieldRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.UpdateCustomField(context.Background(), id).UpdateCustomFieldInputObject(updateCustomFieldInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UpdateCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomField`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.UpdateCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomFieldInputObject** | [**CreateCustomFieldRequest**](CreateCustomFieldRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomFieldValue

> ResourceLinkResponse UpdateCustomFieldValue(ctx, id).UpdateCustomFieldValueInputObject(updateCustomFieldValueInputObject).Execute()

Edit the custom field value of a specified contact

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(554) // int32 | 
	updateCustomFieldValueInputObject := *openapiclient.NewUpdateCustomFieldValueRequest() // UpdateCustomFieldValueRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.UpdateCustomFieldValue(context.Background(), id).UpdateCustomFieldValueInputObject(updateCustomFieldValueInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UpdateCustomFieldValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomFieldValue`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.UpdateCustomFieldValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFieldValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomFieldValueInputObject** | [**UpdateCustomFieldValueRequest**](UpdateCustomFieldValueRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInboundMessagesNotificationSettings

> UpdateInboundMessagesNotificationSettings(ctx).UpdateInboundMessagesNotificationSettingsInputObject(updateInboundMessagesNotificationSettingsInputObject).Execute()

Update inbound messages notification settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	updateInboundMessagesNotificationSettingsInputObject := *openapiclient.NewUpdateInboundMessagesNotificationSettingsRequest() // UpdateInboundMessagesNotificationSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.UpdateInboundMessagesNotificationSettings(context.Background()).UpdateInboundMessagesNotificationSettingsInputObject(updateInboundMessagesNotificationSettingsInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UpdateInboundMessagesNotificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInboundMessagesNotificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateInboundMessagesNotificationSettingsInputObject** | [**UpdateInboundMessagesNotificationSettingsRequest**](UpdateInboundMessagesNotificationSettingsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateList

> ResourceLinkResponse UpdateList(ctx, id).UpdateListObject(updateListObject).Execute()

Edit a list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	updateListObject := *openapiclient.NewUpdateListRequest() // UpdateListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.UpdateList(context.Background(), id).UpdateListObject(updateListObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UpdateList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateList`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.UpdateList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateListObject** | [**UpdateListRequest**](UpdateListRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSenderSetting

> UpdateSenderSetting(ctx).UpdateSenderSettingInputObject(updateSenderSettingInputObject).Execute()

Change sender settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	updateSenderSettingInputObject := *openapiclient.NewUpdateSenderSettingRequest() // UpdateSenderSettingRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.UpdateSenderSetting(context.Background()).UpdateSenderSettingInputObject(updateSenderSettingInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UpdateSenderSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSenderSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateSenderSettingInputObject** | [**UpdateSenderSettingRequest**](UpdateSenderSettingRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTemplate

> ResourceLinkResponse UpdateTemplate(ctx, id).UpdateTemplateInputObject(updateTemplateInputObject).Execute()

Update a template

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	updateTemplateInputObject := *openapiclient.NewCreateTemplateRequest() // CreateTemplateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.UpdateTemplate(context.Background(), id).UpdateTemplateInputObject(updateTemplateInputObject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UpdateTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTemplate`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.UpdateTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTemplateInputObject** | [**CreateTemplateRequest**](CreateTemplateRequest.md) |  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadAvatar

> UploadAvatar(ctx).Image(image).Execute()

Upload an avatar

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	image := os.NewFile(1234, "some_file") // *os.File | User avatar. Should be a PNG or JPG file not more than 10 MB.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TextMagicAPI.UploadAvatar(context.Background()).Image(image).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UploadAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **image** | ***os.File** | User avatar. Should be a PNG or JPG file not more than 10 MB. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadContactAvatar

> ResourceLinkResponse UploadContactAvatar(ctx, id).Image(image).Execute()

Upload an avatar

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	image := os.NewFile(1234, "some_file") // *os.File | Contact avatar. Should be a PNG or JPG file not more than 10 MB.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.UploadContactAvatar(context.Background(), id).Image(image).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UploadContactAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadContactAvatar`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.UploadContactAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadContactAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **image** | ***os.File** | Contact avatar. Should be a PNG or JPG file not more than 10 MB. | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadListAvatar

> ResourceLinkResponse UploadListAvatar(ctx, id).Image(image).Execute()

Add an avatar for a list



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	id := int32(1) // int32 | 
	image := os.NewFile(1234, "some_file") // *os.File | List avatar. Should be a PNG or JPG file not more than 10 MB.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.UploadListAvatar(context.Background(), id).Image(image).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UploadListAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadListAvatar`: ResourceLinkResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.UploadListAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadListAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **image** | ***os.File** | List avatar. Should be a PNG or JPG file not more than 10 MB. | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadMessageAttachment

> UploadMessageAttachmentResponse UploadMessageAttachment(ctx).File(file).Execute()

Upload message attachment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | Attachment. Supports .jpg, .gif, .png, .pdf, .txt, .csv, .doc, .docx, .xls, .xlsx, .ppt, .pptx & .vcf file formats.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.UploadMessageAttachment(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UploadMessageAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadMessageAttachment`: UploadMessageAttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.UploadMessageAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadMessageAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | Attachment. Supports .jpg, .gif, .png, .pdf, .txt, .csv, .doc, .docx, .xls, .xlsx, .ppt, .pptx &amp; .vcf file formats. | 

### Return type

[**UploadMessageAttachmentResponse**](UploadMessageAttachmentResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadMessageMMSAttachment

> UploadMessageAttachmentResponse UploadMessageMMSAttachment(ctx).File(file).Execute()

Upload message mms attachment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/textmagic/textmagic-rest-go-v2"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File | Attachment. Supports .jpg, .gif, .png, .pdf, .txt, .csv, .doc, .docx, .xls, .xlsx, .ppt, .pptx & .vcf file formats.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TextMagicAPI.UploadMessageMMSAttachment(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TextMagicAPI.UploadMessageMMSAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadMessageMMSAttachment`: UploadMessageAttachmentResponse
	fmt.Fprintf(os.Stdout, "Response from `TextMagicAPI.UploadMessageMMSAttachment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadMessageMMSAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | Attachment. Supports .jpg, .gif, .png, .pdf, .txt, .csv, .doc, .docx, .xls, .xlsx, .ppt, .pptx &amp; .vcf file formats. | 

### Return type

[**UploadMessageAttachmentResponse**](UploadMessageAttachmentResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

