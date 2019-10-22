# TextMagic\TextMagicApi

All URIs are relative to *http://rest.textmagic.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignContactsToList**](TextMagicApi.md#AssignContactsToList) | **Put** /api/v2/lists/{id}/contacts | Assign contacts to a list
[**BlockContact**](TextMagicApi.md#BlockContact) | **Post** /api/v2/contacts/block | Block contact by phone number
[**BuyDedicatedNumber**](TextMagicApi.md#BuyDedicatedNumber) | **Post** /api/v2/numbers | Buy a dedicated number
[**CancelVerification**](TextMagicApi.md#CancelVerification) | **Delete** /api/v2/verify/{verifyId} | Cancel verification process
[**CheckPhoneVerificationCodeTFA**](TextMagicApi.md#CheckPhoneVerificationCodeTFA) | **Put** /api/v2/verify | Step 2: Check the verification code 
[**ClearAndAssignContactsToList**](TextMagicApi.md#ClearAndAssignContactsToList) | **Post** /api/v2/lists/{id}/contacts | Reset list members to the specified contacts
[**CloseChatsBulk**](TextMagicApi.md#CloseChatsBulk) | **Post** /api/v2/chats/close/bulk | Close chats (bulk)
[**CloseReadChats**](TextMagicApi.md#CloseReadChats) | **Post** /api/v2/chats/close/read | Close read chats
[**CloseSubaccount**](TextMagicApi.md#CloseSubaccount) | **Delete** /api/v2/subaccounts/{id} | Close sub-account
[**CreateContact**](TextMagicApi.md#CreateContact) | **Post** /api/v2/contacts/normalized | Add a new contact
[**CreateContactNote**](TextMagicApi.md#CreateContactNote) | **Post** /api/v2/contacts/{id}/notes | Create a new contact note
[**CreateCustomField**](TextMagicApi.md#CreateCustomField) | **Post** /api/v2/customfields | Add a new custom field
[**CreateList**](TextMagicApi.md#CreateList) | **Post** /api/v2/lists | Create a new list
[**CreateTemplate**](TextMagicApi.md#CreateTemplate) | **Post** /api/v2/templates | Create a template
[**DeleteAllContacts**](TextMagicApi.md#DeleteAllContacts) | **Delete** /api/v2/contact/all | Delete contacts (bulk)
[**DeleteAllOutboundMessages**](TextMagicApi.md#DeleteAllOutboundMessages) | **Delete** /api/v2/message/all | Delete all messages
[**DeleteAvatar**](TextMagicApi.md#DeleteAvatar) | **Delete** /api/v2/user/avatar | Delete an avatar
[**DeleteChatMessages**](TextMagicApi.md#DeleteChatMessages) | **Post** /api/v2/chats/{id}/messages/delete | Delete chat messages by ID(s)
[**DeleteChatsBulk**](TextMagicApi.md#DeleteChatsBulk) | **Post** /api/v2/chats/delete | Delete chats (bulk)
[**DeleteContact**](TextMagicApi.md#DeleteContact) | **Delete** /api/v2/contacts/{id} | Delete a contact
[**DeleteContactAvatar**](TextMagicApi.md#DeleteContactAvatar) | **Delete** /api/v2/contacts/{id}/avatar | Delete an avatar
[**DeleteContactNote**](TextMagicApi.md#DeleteContactNote) | **Delete** /api/v2/notes/{id} | Delete a contact note
[**DeleteContactNotesBulk**](TextMagicApi.md#DeleteContactNotesBulk) | **Post** /api/v2/contacts/{id}/notes/delete | Delete contact notes (bulk)
[**DeleteContactsByIds**](TextMagicApi.md#DeleteContactsByIds) | **Post** /api/v2/contacts/delete | Delete contacts by IDs (bulk)
[**DeleteContactsFromList**](TextMagicApi.md#DeleteContactsFromList) | **Delete** /api/v2/lists/{id}/contacts | Unassign contacts from a list
[**DeleteCustomField**](TextMagicApi.md#DeleteCustomField) | **Delete** /api/v2/customfields/{id} | Delete a custom field
[**DeleteDedicatedNumber**](TextMagicApi.md#DeleteDedicatedNumber) | **Delete** /api/v2/numbers/{id} | Cancel dedicated number subscription
[**DeleteInboundMessage**](TextMagicApi.md#DeleteInboundMessage) | **Delete** /api/v2/replies/{id} | Delete a single inbound message
[**DeleteInboundMessagesBulk**](TextMagicApi.md#DeleteInboundMessagesBulk) | **Post** /api/v2/replies/delete | Delete inbound messages (bulk)
[**DeleteList**](TextMagicApi.md#DeleteList) | **Delete** /api/v2/lists/{id} | Delete a list
[**DeleteListAvatar**](TextMagicApi.md#DeleteListAvatar) | **Delete** /api/v2/lists/{id}/avatar | Delete an avatar for the list
[**DeleteListContactsBulk**](TextMagicApi.md#DeleteListContactsBulk) | **Post** /api/v2/lists/{id}/contacts/delete | Delete contacts from list (bulk)
[**DeleteListsBulk**](TextMagicApi.md#DeleteListsBulk) | **Post** /api/v2/lists/delete | Delete lists (bulk)
[**DeleteMessageSession**](TextMagicApi.md#DeleteMessageSession) | **Delete** /api/v2/sessions/{id} | Delete a session
[**DeleteMessageSessionsBulk**](TextMagicApi.md#DeleteMessageSessionsBulk) | **Post** /api/v2/sessions/delete | Delete sessions (bulk)
[**DeleteOutboundMessage**](TextMagicApi.md#DeleteOutboundMessage) | **Delete** /api/v2/messages/{id} | Delete message
[**DeleteOutboundMessagesBulk**](TextMagicApi.md#DeleteOutboundMessagesBulk) | **Post** /api/v2/messages/delete | Delete messages (bulk)
[**DeleteScheduledMessage**](TextMagicApi.md#DeleteScheduledMessage) | **Delete** /api/v2/schedules/{id} | Delete a single scheduled message
[**DeleteScheduledMessagesBulk**](TextMagicApi.md#DeleteScheduledMessagesBulk) | **Post** /api/v2/schedules/delete | Delete scheduled messages (bulk)
[**DeleteSenderId**](TextMagicApi.md#DeleteSenderId) | **Delete** /api/v2/senderids/{id} | Delete a Sender ID
[**DeleteTemplate**](TextMagicApi.md#DeleteTemplate) | **Delete** /api/v2/templates/{id} | Delete a template
[**DeleteTemplatesBulk**](TextMagicApi.md#DeleteTemplatesBulk) | **Post** /api/v2/templates/delete | Delete templates (bulk)
[**DoCarrierLookup**](TextMagicApi.md#DoCarrierLookup) | **Get** /api/v2/lookups/{phone} | Carrier Lookup
[**DoEmailLookup**](TextMagicApi.md#DoEmailLookup) | **Get** /api/v2/email-lookups/{email} | Email Lookup
[**GetAllBulkSessions**](TextMagicApi.md#GetAllBulkSessions) | **Get** /api/v2/bulks | Get all bulk sessions
[**GetAllChats**](TextMagicApi.md#GetAllChats) | **Get** /api/v2/chats | Get all chats
[**GetAllInboundMessages**](TextMagicApi.md#GetAllInboundMessages) | **Get** /api/v2/replies | Get all inbound messages
[**GetAllMessageSessions**](TextMagicApi.md#GetAllMessageSessions) | **Get** /api/v2/sessions | Get all sessions
[**GetAllOutboundMessages**](TextMagicApi.md#GetAllOutboundMessages) | **Get** /api/v2/messages | Get all messages
[**GetAllScheduledMessages**](TextMagicApi.md#GetAllScheduledMessages) | **Get** /api/v2/schedules | Get all scheduled messages
[**GetAllTemplates**](TextMagicApi.md#GetAllTemplates) | **Get** /api/v2/templates | Get all templates
[**GetAvailableDedicatedNumbers**](TextMagicApi.md#GetAvailableDedicatedNumbers) | **Get** /api/v2/numbers/available | Find dedicated numbers available for purchase
[**GetAvailableSenderSettingOptions**](TextMagicApi.md#GetAvailableSenderSettingOptions) | **Get** /api/v2/sources | Get available sender settings
[**GetBalanceNotificationOptions**](TextMagicApi.md#GetBalanceNotificationOptions) | **Get** /api/v2/user/notification/balance/bundles | Returns the list of available balance options which can be used as a bound to determine when to send email to user with low balance notification. See https://my.textmagic.com/online/account/notifications/balance
[**GetBalanceNotificationSettings**](TextMagicApi.md#GetBalanceNotificationSettings) | **Get** /api/v2/user/notification/balance | Get balance notification settings
[**GetBlockedContacts**](TextMagicApi.md#GetBlockedContacts) | **Get** /api/v2/contacts/block/list | Get blocked contacts
[**GetBulkSession**](TextMagicApi.md#GetBulkSession) | **Get** /api/v2/bulks/{id} | Get bulk session status
[**GetCallbackSettings**](TextMagicApi.md#GetCallbackSettings) | **Get** /api/v2/callback/settings | Fetch callback URL settings
[**GetChat**](TextMagicApi.md#GetChat) | **Get** /api/v2/chats/{id} | Get a single chat
[**GetChatByPhone**](TextMagicApi.md#GetChatByPhone) | **Get** /api/v2/chats/{phone}/by/phone | Find chats by phone
[**GetChatMessages**](TextMagicApi.md#GetChatMessages) | **Get** /api/v2/chats/{id}/message | Get chat messages
[**GetContact**](TextMagicApi.md#GetContact) | **Get** /api/v2/contacts/{id} | Get the details of a specific contact
[**GetContactByPhone**](TextMagicApi.md#GetContactByPhone) | **Get** /api/v2/contacts/phone/{phone} | Get the details of a specific contact by phone number
[**GetContactIfBlocked**](TextMagicApi.md#GetContactIfBlocked) | **Get** /api/v2/contacts/block/phone | Check is that phone number blocked
[**GetContactImportSessionProgress**](TextMagicApi.md#GetContactImportSessionProgress) | **Get** /api/v2/contacts/import/progress/{id} | Check import progress
[**GetContactNote**](TextMagicApi.md#GetContactNote) | **Get** /api/v2/notes/{id} | Get a contact note
[**GetContactNotes**](TextMagicApi.md#GetContactNotes) | **Get** /api/v2/contacts/{id}/notes | Fetch notes assigned to the given contact.
[**GetContacts**](TextMagicApi.md#GetContacts) | **Get** /api/v2/contacts | Get all contacts
[**GetContactsAutocomplete**](TextMagicApi.md#GetContactsAutocomplete) | **Get** /api/v2/contacts/autocomplete | Get contacts autocomplete suggestions
[**GetContactsByListId**](TextMagicApi.md#GetContactsByListId) | **Get** /api/v2/lists/{id}/contacts | Get all contacts in a list
[**GetCountries**](TextMagicApi.md#GetCountries) | **Get** /api/v2/countries | Get countries
[**GetCurrentUser**](TextMagicApi.md#GetCurrentUser) | **Get** /api/v2/user | Get current account information
[**GetCustomField**](TextMagicApi.md#GetCustomField) | **Get** /api/v2/customfields/{id} | Get the details of a specific custom field
[**GetCustomFields**](TextMagicApi.md#GetCustomFields) | **Get** /api/v2/customfields | Get all custom fields
[**GetDedicatedNumber**](TextMagicApi.md#GetDedicatedNumber) | **Get** /api/v2/numbers/{id} | Get the details of a specific dedicated number
[**GetFavourites**](TextMagicApi.md#GetFavourites) | **Get** /api/v2/contacts/favorite | Get favorite contacts and lists
[**GetInboundMessage**](TextMagicApi.md#GetInboundMessage) | **Get** /api/v2/replies/{id} | Get a single inbound message
[**GetInboundMessagesNotificationSettings**](TextMagicApi.md#GetInboundMessagesNotificationSettings) | **Get** /api/v2/user/notification/inbound | Get inbound messages notification settings
[**GetInvoices**](TextMagicApi.md#GetInvoices) | **Get** /api/v2/invoices | Get all invoices
[**GetList**](TextMagicApi.md#GetList) | **Get** /api/v2/lists/{id} | Get the details of a specific list
[**GetListContactsIds**](TextMagicApi.md#GetListContactsIds) | **Get** /api/v2/lists/{id}/contacts/ids | Get all contacts IDs in a list
[**GetLists**](TextMagicApi.md#GetLists) | **Get** /api/v2/lists | Get all lists
[**GetListsOfContact**](TextMagicApi.md#GetListsOfContact) | **Get** /api/v2/contacts/{id}/lists | Get contact&#39;s lists
[**GetMessagePreview**](TextMagicApi.md#GetMessagePreview) | **Get** /api/v2/messages/preview | Preview message
[**GetMessagePrice**](TextMagicApi.md#GetMessagePrice) | **Get** /api/v2/messages/price/normalized | Check message price
[**GetMessageSession**](TextMagicApi.md#GetMessageSession) | **Get** /api/v2/sessions/{id} | Get a session details
[**GetMessageSessionStat**](TextMagicApi.md#GetMessageSessionStat) | **Get** /api/v2/sessions/{id}/stat | Get a session statistics
[**GetMessagesBySessionId**](TextMagicApi.md#GetMessagesBySessionId) | **Get** /api/v2/sessions/{id}/messages | Get a session messages
[**GetMessagingCounters**](TextMagicApi.md#GetMessagingCounters) | **Get** /api/v2/stats/messaging/data | Get sent/received messages counters values
[**GetMessagingStat**](TextMagicApi.md#GetMessagingStat) | **Get** /api/v2/stats/messaging | Get messaging statistics
[**GetOutboundMessage**](TextMagicApi.md#GetOutboundMessage) | **Get** /api/v2/messages/{id} | Get a single message
[**GetOutboundMessagesHistory**](TextMagicApi.md#GetOutboundMessagesHistory) | **Get** /api/v2/history | Get history
[**GetScheduledMessage**](TextMagicApi.md#GetScheduledMessage) | **Get** /api/v2/schedules/{id} | Get a single scheduled message
[**GetSenderId**](TextMagicApi.md#GetSenderId) | **Get** /api/v2/senderids/{id} | Get the details of a specific Sender ID
[**GetSenderIds**](TextMagicApi.md#GetSenderIds) | **Get** /api/v2/senderids | Get all your approved Sender IDs
[**GetSenderSettings**](TextMagicApi.md#GetSenderSettings) | **Get** /api/v2/sender/settings/normalized | Get current sender settings
[**GetSpendingStat**](TextMagicApi.md#GetSpendingStat) | **Get** /api/v2/stats/spending | Get spending statistics
[**GetSubaccount**](TextMagicApi.md#GetSubaccount) | **Get** /api/v2/subaccounts/{id} | Get sub-account information
[**GetSubaccounts**](TextMagicApi.md#GetSubaccounts) | **Get** /api/v2/subaccounts | Get sub-accounts list
[**GetSubaccountsWithTokens**](TextMagicApi.md#GetSubaccountsWithTokens) | **Post** /api/v2/subaccounts/tokens/list | Get all sub-accounts with their REST API tokens associated with app name
[**GetTemplate**](TextMagicApi.md#GetTemplate) | **Get** /api/v2/templates/{id} | Get a template details
[**GetTimezones**](TextMagicApi.md#GetTimezones) | **Get** /api/v2/timezones | Get timezones
[**GetUnreadMessagesTotal**](TextMagicApi.md#GetUnreadMessagesTotal) | **Get** /api/v2/chats/unread/count | Get unread messages number
[**GetUnsubscribedContact**](TextMagicApi.md#GetUnsubscribedContact) | **Get** /api/v2/unsubscribers/{id} | Get the details of a specific unsubscribed contact
[**GetUnsubscribers**](TextMagicApi.md#GetUnsubscribers) | **Get** /api/v2/unsubscribers | Get all unsubscribed contacts
[**GetUserDedicatedNumbers**](TextMagicApi.md#GetUserDedicatedNumbers) | **Get** /api/v2/numbers | Get all your dedicated numbers
[**InviteSubaccount**](TextMagicApi.md#InviteSubaccount) | **Post** /api/v2/subaccounts | Invite a new sub-account
[**MarkChatsReadBulk**](TextMagicApi.md#MarkChatsReadBulk) | **Post** /api/v2/chats/read/bulk | Mark chats as read (bulk)
[**MarkChatsUnreadBulk**](TextMagicApi.md#MarkChatsUnreadBulk) | **Post** /api/v2/chats/unread/bulk | Mark chats as unread (bulk)
[**MuteChat**](TextMagicApi.md#MuteChat) | **Post** /api/v2/chats/mute | Mute chat sounds
[**MuteChatsBulk**](TextMagicApi.md#MuteChatsBulk) | **Post** /api/v2/chats/mute/bulk | Mute chats (bulk)
[**Ping**](TextMagicApi.md#Ping) | **Get** /api/v2/ping | Ping
[**ReopenChatsBulk**](TextMagicApi.md#ReopenChatsBulk) | **Post** /api/v2/chats/reopen/bulk | Reopen chats (bulk)
[**RequestNewSubaccountToken**](TextMagicApi.md#RequestNewSubaccountToken) | **Post** /api/v2/subaccounts/tokens | Request a new REST API token for sub-account
[**RequestSenderId**](TextMagicApi.md#RequestSenderId) | **Post** /api/v2/senderids | Apply for a new Sender ID
[**SearchChats**](TextMagicApi.md#SearchChats) | **Get** /api/v2/chats/search | Find chats by message text
[**SearchChatsByIds**](TextMagicApi.md#SearchChatsByIds) | **Get** /api/v2/chats/search/ids | Find chats (bulk)
[**SearchChatsByReceipent**](TextMagicApi.md#SearchChatsByReceipent) | **Get** /api/v2/chats/search/recipients | Find chats by recipient
[**SearchContacts**](TextMagicApi.md#SearchContacts) | **Get** /api/v2/contacts/search | Find contacts by given criteria
[**SearchInboundMessages**](TextMagicApi.md#SearchInboundMessages) | **Get** /api/v2/replies/search | Find inbound messages
[**SearchLists**](TextMagicApi.md#SearchLists) | **Get** /api/v2/lists/search | Find lists by given criteria
[**SearchOutboundMessages**](TextMagicApi.md#SearchOutboundMessages) | **Get** /api/v2/messages/search | Find messages
[**SearchScheduledMessages**](TextMagicApi.md#SearchScheduledMessages) | **Get** /api/v2/schedules/search | Find scheduled messages
[**SearchTemplates**](TextMagicApi.md#SearchTemplates) | **Get** /api/v2/templates/search | Find templates by criteria
[**SendMessage**](TextMagicApi.md#SendMessage) | **Post** /api/v2/messages | Send message
[**SendPhoneVerificationCodeTFA**](TextMagicApi.md#SendPhoneVerificationCodeTFA) | **Post** /api/v2/verify | Step 1: Send a verification code 
[**SetChatStatus**](TextMagicApi.md#SetChatStatus) | **Post** /api/v2/chats/status | Change chat status
[**UnblockContact**](TextMagicApi.md#UnblockContact) | **Post** /api/v2/contacts/unblock | Unblock contact by phone number.
[**UnblockContactsBulk**](TextMagicApi.md#UnblockContactsBulk) | **Post** /api/v2/contacts/unblock/bulk | Unblock contacts (bulk)
[**UnmuteChatsBulk**](TextMagicApi.md#UnmuteChatsBulk) | **Post** /api/v2/chats/unmute/bulk | Unmute chats (bulk)
[**UnsubscribeContact**](TextMagicApi.md#UnsubscribeContact) | **Post** /api/v2/unsubscribers | Manually unsubscribe a contact
[**UpdateBalanceNotificationSettings**](TextMagicApi.md#UpdateBalanceNotificationSettings) | **Put** /api/v2/user/notification/balance | Update balance notification settings
[**UpdateCallbackSettings**](TextMagicApi.md#UpdateCallbackSettings) | **Put** /api/v2/callback/settings | Update callback URL settings
[**UpdateChatDesktopNotificationSettings**](TextMagicApi.md#UpdateChatDesktopNotificationSettings) | **Put** /api/v2/user/desktop/notification | Update chat desktop notification settings
[**UpdateContact**](TextMagicApi.md#UpdateContact) | **Put** /api/v2/contacts/{id}/normalized | Edit a contact
[**UpdateContactNote**](TextMagicApi.md#UpdateContactNote) | **Put** /api/v2/notes/{id} | Update a contact note
[**UpdateCurrentUser**](TextMagicApi.md#UpdateCurrentUser) | **Put** /api/v2/user | Edit current account info
[**UpdateCustomField**](TextMagicApi.md#UpdateCustomField) | **Put** /api/v2/customfields/{id} | Edit a custom field
[**UpdateCustomFieldValue**](TextMagicApi.md#UpdateCustomFieldValue) | **Put** /api/v2/customfields/{id}/update | Edit the custom field value of a specified contact
[**UpdateInboundMessagesNotificationSettings**](TextMagicApi.md#UpdateInboundMessagesNotificationSettings) | **Put** /api/v2/user/notification/inbound | Update inbound messages notification settings
[**UpdateList**](TextMagicApi.md#UpdateList) | **Put** /api/v2/lists/{id} | Edit a list
[**UpdateSenderSetting**](TextMagicApi.md#UpdateSenderSetting) | **Put** /api/v2/sender/settings | Change sender settings
[**UpdateTemplate**](TextMagicApi.md#UpdateTemplate) | **Put** /api/v2/templates/{id} | Update a template


# **AssignContactsToList**
> ResourceLinkResponse AssignContactsToList(ctx, assignContactsToListInputObject, id)
Assign contacts to a list

> Unlike all other PUT requests, this command does not need old contact IDs to be submitted. For example, if you have a list with contacts 150, 151 and 152 and you want to add contact ID 153, you only need to submit 153 as a parameter of PUT /api/v2/lists/{id}/contacts. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **assignContactsToListInputObject** | [**AssignContactsToListInputObject**](AssignContactsToListInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BlockContact**
> ResourceLinkResponse BlockContact(ctx, blockContactInputObject)
Block contact by phone number

Block contact from inbound and outbound communication by phone number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **blockContactInputObject** | [**BlockContactInputObject**](BlockContactInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BuyDedicatedNumber**
> BuyDedicatedNumber(ctx, buyDedicatedNumberInputObject)
Buy a dedicated number

To buy a dedicated number, you first need to find an available number matching your criteria using the `/api/v2/numbers/available` command described above.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **buyDedicatedNumberInputObject** | [**BuyDedicatedNumberInputObject**](BuyDedicatedNumberInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelVerification**
> CancelVerification(ctx, verifyId)
Cancel verification process

You can cancel the verification not earlier than 30 seconds after the initial request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verifyId** | **string**| the verifyId that you received in Step 1. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckPhoneVerificationCodeTFA**
> CheckPhoneVerificationCodeTFA(ctx, checkPhoneVerificationCodeInputObject)
Step 2: Check the verification code 

Check received code from user with the code which was actually sent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **checkPhoneVerificationCodeInputObject** | [**CheckPhoneVerificationCodeInputObject**](CheckPhoneVerificationCodeInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearAndAssignContactsToList**
> ResourceLinkResponse ClearAndAssignContactsToList(ctx, clearAndAssignContactsToListInputObject, id)
Reset list members to the specified contacts



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clearAndAssignContactsToListInputObject** | [**ClearAndAssignContactsToListInputObject**](ClearAndAssignContactsToListInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloseChatsBulk**
> CloseChatsBulk(ctx, closeChatsBulkInputObject)
Close chats (bulk)

Close chats by chat ids or close all chats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **closeChatsBulkInputObject** | [**CloseChatsBulkInputObject**](CloseChatsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloseReadChats**
> CloseReadChats(ctx, )
Close read chats

Close all chats that have no unread messages.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloseSubaccount**
> CloseSubaccount(ctx, id)
Close sub-account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContact**
> ResourceLinkResponse CreateContact(ctx, createContactInputObject)
Add a new contact



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createContactInputObject** | [**CreateContactInputObject**](CreateContactInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContactNote**
> ResourceLinkResponse CreateContactNote(ctx, createContactNoteInputObject, id)
Create a new contact note



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createContactNoteInputObject** | [**CreateContactNoteInputObject**](CreateContactNoteInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomField**
> ResourceLinkResponse CreateCustomField(ctx, createCustomFieldInputObject)
Add a new custom field



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createCustomFieldInputObject** | [**CreateCustomFieldInputObject**](CreateCustomFieldInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateList**
> ResourceLinkResponse CreateList(ctx, createListInputObject)
Create a new list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createListInputObject** | [**CreateListInputObject**](CreateListInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTemplate**
> ResourceLinkResponse CreateTemplate(ctx, createTemplateInputObject)
Create a template

There are times when creating a new template makes sense (such as when targeting specific clients or improving your business strategies).  You can create new SMS templates for marketing purposes or SMS templates for business campaigns. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **createTemplateInputObject** | [**CreateTemplateInputObject**](CreateTemplateInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllContacts**
> DeleteAllContacts(ctx, )
Delete contacts (bulk)



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllOutboundMessages**
> DeleteAllOutboundMessages(ctx, )
Delete all messages

Delete all messages.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAvatar**
> DeleteAvatar(ctx, )
Delete an avatar



### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChatMessages**
> DeleteChatMessages(ctx, deleteChatMessagesBulkInputObject, id)
Delete chat messages by ID(s)

Delete messages from chat by given messages ID(s).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteChatMessagesBulkInputObject** | [**DeleteChatMessagesBulkInputObject**](DeleteChatMessagesBulkInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteChatsBulk**
> DeleteChatsBulk(ctx, deleteChatsBulkInputObject)
Delete chats (bulk)

Delete chats by given ID(s) or delete all chats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteChatsBulkInputObject** | [**DeleteChatsBulkInputObject**](DeleteChatsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContact**
> DeleteContact(ctx, id)
Delete a contact

> This command removes your contact completely. If it was assigned or saved to a shared list, it will disappear from there too. If you only need to remove a contact from selected lists, instead use the Contact assignment command in the Lists section rather than deleting the contact. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContactAvatar**
> DeleteContactAvatar(ctx, id)
Delete an avatar



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContactNote**
> DeleteContactNote(ctx, id)
Delete a contact note



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContactNotesBulk**
> DeleteContactNotesBulk(ctx, id, deleteContactNotesBulkInputObject)
Delete contact notes (bulk)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
  **deleteContactNotesBulkInputObject** | [**DeleteContactNotesBulkInputObject**](DeleteContactNotesBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContactsByIds**
> DeleteContactsByIds(ctx, deleteContactsByIdsInputObject)
Delete contacts by IDs (bulk)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteContactsByIdsInputObject** | [**DeleteContactsByIdsInputObject**](DeleteContactsByIdsInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContactsFromList**
> DeleteContactsFromList(ctx, deleteContacsFromListObject, id)
Unassign contacts from a list

> When you remove contacts from a specific list, they will be deleted permanently, unless they are first saved in another list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteContacsFromListObject** | [**DeleteContacsFromListObject**](DeleteContacsFromListObject.md)|  | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomField**
> DeleteCustomField(ctx, id)
Delete a custom field

> When a custom field is deleted, all the information that was added to contacts under this custom field will also be lost. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDedicatedNumber**
> DeleteDedicatedNumber(ctx, id)
Cancel dedicated number subscription



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInboundMessage**
> DeleteInboundMessage(ctx, id)
Delete a single inbound message

> Note, deleted inbound message will disappear from TextMagic Online, chats, and any other place they are referenced.  So, be careful! 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The unique numeric ID for the inbound message. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteInboundMessagesBulk**
> DeleteInboundMessagesBulk(ctx, deleteInboundMessagesBulkInputObject)
Delete inbound messages (bulk)

> Note, deleted inbound message will disappear from TextMagic Online, chats, and any other place they are referenced.  So, be careful! 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteInboundMessagesBulkInputObject** | [**DeleteInboundMessagesBulkInputObject**](DeleteInboundMessagesBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteList**
> DeleteList(ctx, id)
Delete a list

This command has no parameters. If successful, this command will return the standard delete response (204 No Content), otherwise a standard error response will be returned.  When you delete a list, the contacts in it are deleted as well unless they were saved in other list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListAvatar**
> DeleteListAvatar(ctx, id)
Delete an avatar for the list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListContactsBulk**
> DeleteListContactsBulk(ctx, deleteListContactsBulkInputObject, id)
Delete contacts from list (bulk)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteListContactsBulkInputObject** | [**DeleteListContactsBulkInputObject**](DeleteListContactsBulkInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteListsBulk**
> DeleteListsBulk(ctx, deleteListsBulkInputObject)
Delete lists (bulk)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteListsBulkInputObject** | [**DeleteListsBulkInputObject**](DeleteListsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMessageSession**
> DeleteMessageSession(ctx, id)
Delete a session

Delete a message session, together with all nested messages. > You will not be refunded for any deleted sent sessions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMessageSessionsBulk**
> DeleteMessageSessionsBulk(ctx, deleteMessageSessionsBulkInputObject)
Delete sessions (bulk)

Delete messages sessions, together with all nested messages, by given ID(s) or delete all messages sessions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteMessageSessionsBulkInputObject** | [**DeleteMessageSessionsBulkInputObject**](DeleteMessageSessionsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOutboundMessage**
> DeleteOutboundMessage(ctx, id)
Delete message

Delete a single message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOutboundMessagesBulk**
> DeleteOutboundMessagesBulk(ctx, deleteOutboundMessagesBulkInputObject)
Delete messages (bulk)

Delete outbound messages by given ID(s) or delete all outbound messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteOutboundMessagesBulkInputObject** | [**DeleteOutboundMessagesBulkInputObject**](DeleteOutboundMessagesBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScheduledMessage**
> DeleteScheduledMessage(ctx, id)
Delete a single scheduled message



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScheduledMessagesBulk**
> DeleteScheduledMessagesBulk(ctx, deleteScheduledMessagesBulkInputObject)
Delete scheduled messages (bulk)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteScheduledMessagesBulkInputObject** | [**DeleteScheduledMessagesBulkInputObject**](DeleteScheduledMessagesBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteSenderId**
> DeleteSenderId(ctx, id)
Delete a Sender ID



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTemplate**
> DeleteTemplate(ctx, id)
Delete a template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTemplatesBulk**
> DeleteTemplatesBulk(ctx, deleteTemplatesBulkInputObject)
Delete templates (bulk)

Delete template by given ID(s) or delete all templates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deleteTemplatesBulkInputObject** | [**DeleteTemplatesBulkInputObject**](DeleteTemplatesBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoCarrierLookup**
> DoCarrierLookupResponse DoCarrierLookup(ctx, phone, optional)
Carrier Lookup

This API call allows you to retrieve additional information about a phone number: region-specific phone number formatting, carrier, phone type (landline/mobile) and country information.  > Numbers can be checked one by one. You cannot check multiple numbers in one request.   

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phone** | **string**| Phone number in [E.164 format](https://en.wikipedia.org/wiki/E.164) or in [National format](https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers).  | 
 **optional** | ***DoCarrierLookupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DoCarrierLookupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **country** | **optional.String**| This option must be specified only if the phone number in a **[National format](https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers)**.  | 

### Return type

[**DoCarrierLookupResponse**](DoCarrierLookupResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DoEmailLookup**
> DoEmailLookupResponse DoEmailLookup(ctx, email)
Email Lookup

To get more details about an email address or to check if it is a valid email, you can use the Email Lookup command. To upload and check emails in bulk, please use our [Web app](https://my.textmagic.com/online/email-lookup/).  This API call allows you to retrieve additional information about an email address, such as mailbox detection, syntax checks, DNS validation, deliverability status, and many more helpful values (see the table below).  > Emails must be checked one by one. You cannot check multiple emails in one request. To upload and check emails in bulk, please use our [Web app](https://my.textmagic.com/online/email-lookup/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**| Email address. | 

### Return type

[**DoEmailLookupResponse**](DoEmailLookupResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllBulkSessions**
> GetAllBulkSessionsPaginatedResponse GetAllBulkSessions(ctx, optional)
Get all bulk sessions



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllBulkSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllBulkSessionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetAllBulkSessionsPaginatedResponse**](GetAllBulkSessionsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllChats**
> GetAllChatsPaginatedResponse GetAllChats(ctx, optional)
Get all chats



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllChatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllChatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **optional.String**| Fetch only (a)ctive, (c)losed or (d)eleted chats | 
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **voice** | **optional.Int32**| Fetch results with voice calls | [default to 0]
 **flat** | **optional.Int32**| Should additional contact info be included | [default to 0]

### Return type

[**GetAllChatsPaginatedResponse**](GetAllChatsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllInboundMessages**
> GetAllInboundMessagesPaginatedResponse GetAllInboundMessages(ctx, optional)
Get all inbound messages



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllInboundMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllInboundMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetAllInboundMessagesPaginatedResponse**](GetAllInboundMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllMessageSessions**
> GetAllMessageSessionsPaginatedResponse GetAllMessageSessions(ctx, optional)
Get all sessions

Get all message sending sessions. > This list contains all of your sessions, including those which were sent but not via API 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllMessageSessionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllMessageSessionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetAllMessageSessionsPaginatedResponse**](GetAllMessageSessionsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllOutboundMessages**
> GetAllOutboundMessagesPaginatedResponse GetAllOutboundMessages(ctx, optional)
Get all messages

Get all user oubound messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllOutboundMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllOutboundMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **lastId** | **optional.Int32**| Filter results by ID, selecting all values lesser than the specified ID. Note that \\&#39;page\\&#39; parameter is ignored when \\&#39;lastId\\&#39; is specified | 

### Return type

[**GetAllOutboundMessagesPaginatedResponse**](GetAllOutboundMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllScheduledMessages**
> GetAllScheduledMessagesPaginatedResponse GetAllScheduledMessages(ctx, optional)
Get all scheduled messages



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllScheduledMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllScheduledMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **status** | **optional.String**| Fetch schedules with the specific status: a - actual, c - completed, x - all | [default to x]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetAllScheduledMessagesPaginatedResponse**](GetAllScheduledMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTemplates**
> GetAllTemplatesPaginatedResponse GetAllTemplates(ctx, optional)
Get all templates



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | 
 **limit** | **optional.Int32**| The number of results per page. | 

### Return type

[**GetAllTemplatesPaginatedResponse**](GetAllTemplatesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableDedicatedNumbers**
> GetAvailableDedicatedNumbersResponse GetAvailableDedicatedNumbers(ctx, country, optional)
Find dedicated numbers available for purchase



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **country** | **string**| Two-letter dedicated number country ISO code. | 
 **optional** | ***GetAvailableDedicatedNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAvailableDedicatedNumbersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prefix** | **optional.Int32**| Desired number prefix. Should include country code (i.e. 447 for UK phone number format). Leave blank to get all the available numbers for the specified country. | 
 **tollfree** | **optional.Int32**| Should we show only tollfree numbers (tollfree available only for US). | [default to 0]

### Return type

[**GetAvailableDedicatedNumbersResponse**](GetAvailableDedicatedNumbersResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvailableSenderSettingOptions**
> GetAvailableSenderSettingOptionsResponse GetAvailableSenderSettingOptions(ctx, optional)
Get available sender settings

Get all available sender setting options which could be used in \"from\" parameter of POST messages method.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAvailableSenderSettingOptionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAvailableSenderSettingOptionsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **optional.String**| Two-letter ISO country ID. If not specified, it returns all the available sender settings. | 

### Return type

[**GetAvailableSenderSettingOptionsResponse**](GetAvailableSenderSettingOptionsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBalanceNotificationOptions**
> GetBalanceNotificationOptionsResponse GetBalanceNotificationOptions(ctx, )
Returns the list of available balance options which can be used as a bound to determine when to send email to user with low balance notification. See https://my.textmagic.com/online/account/notifications/balance



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetBalanceNotificationOptionsResponse**](GetBalanceNotificationOptionsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBalanceNotificationSettings**
> GetBalanceNotificationSettingsResponse GetBalanceNotificationSettings(ctx, )
Get balance notification settings



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetBalanceNotificationSettingsResponse**](GetBalanceNotificationSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBlockedContacts**
> GetBlockedContactsPaginatedResponse GetBlockedContacts(ctx, optional)
Get blocked contacts



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetBlockedContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetBlockedContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **query** | **optional.String**| Find blocked contacts by specified search query | 
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetBlockedContactsPaginatedResponse**](GetBlockedContactsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBulkSession**
> BulkSession GetBulkSession(ctx, id)
Get bulk session status



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**BulkSession**](BulkSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCallbackSettings**
> GetCallbackSettingsResponse GetCallbackSettings(ctx, )
Fetch callback URL settings



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetCallbackSettingsResponse**](GetCallbackSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChat**
> Chat GetChat(ctx, id)
Get a single chat



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**Chat**](Chat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChatByPhone**
> Chat GetChatByPhone(ctx, phone, optional)
Find chats by phone



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phone** | **string**|  | 
 **optional** | ***GetChatByPhoneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChatByPhoneOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsert** | **optional.Int32**| Create a new chat if not found | [default to 0]
 **reopen** | **optional.Int32**| Reopen chat if found or do not change status | [default to 0]

### Return type

[**Chat**](Chat.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChatMessages**
> GetChatMessagesPaginatedResponse GetChatMessages(ctx, id, optional)
Get chat messages



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetChatMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetChatMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **query** | **optional.String**| Find messages by specified search query | 
 **start** | **optional.Int32**| Return messages since specified timestamp only | 
 **end** | **optional.Int32**| Return messages up to specified timestamp only | 
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]
 **voice** | **optional.Int32**| Fetch results with voice calls | [default to 0]

### Return type

[**GetChatMessagesPaginatedResponse**](GetChatMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContact**
> Contact GetContact(ctx, id)
Get the details of a specific contact



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The contact id | 

### Return type

[**Contact**](Contact.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactByPhone**
> Contact GetContactByPhone(ctx, phone)
Get the details of a specific contact by phone number



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phone** | **string**|  | 

### Return type

[**Contact**](Contact.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactIfBlocked**
> Contact GetContactIfBlocked(ctx, phone)
Check is that phone number blocked



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **phone** | **string**| Phone number to check | 

### Return type

[**Contact**](Contact.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactImportSessionProgress**
> GetContactImportSessionProgressResponse GetContactImportSessionProgress(ctx, id)
Check import progress

Get contact import session progress.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GetContactImportSessionProgressResponse**](GetContactImportSessionProgressResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactNote**
> ContactNote GetContactNote(ctx, id)
Get a contact note



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**ContactNote**](ContactNote.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactNotes**
> GetContactNotesPaginatedResponse GetContactNotes(ctx, id, optional)
Fetch notes assigned to the given contact.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetContactNotesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContactNotesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetContactNotesPaginatedResponse**](GetContactNotesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContacts**
> GetContactsPaginatedResponse GetContacts(ctx, optional)
Get all contacts



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **shared** | **optional.Int32**| Should shared contacts to be included | [default to 0]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetContactsPaginatedResponse**](GetContactsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactsAutocomplete**
> GetContactsAutocompleteResponse GetContactsAutocomplete(ctx, query, optional)
Get contacts autocomplete suggestions

Get contacts autocomplete suggestions by given search term

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **query** | **string**| Find recipients by specified search query | 
 **optional** | ***GetContactsAutocompleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContactsAutocompleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **lists** | **optional.Int32**| Should lists be returned or not | [default to 0]

### Return type

[**GetContactsAutocompleteResponse**](GetContactsAutocompleteResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContactsByListId**
> GetContactsByListIdPaginatedResponse GetContactsByListId(ctx, id, optional)
Get all contacts in a list

A useful synonym for \"contacts/search\" command with provided \"listId\" parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Given group Id. | 
 **optional** | ***GetContactsByListIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetContactsByListIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetContactsByListIdPaginatedResponse**](GetContactsByListIdPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCountries**
> GetCountriesResponse GetCountries(ctx, )
Get countries



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetCountriesResponse**](GetCountriesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCurrentUser**
> User GetCurrentUser(ctx, )
Get current account information



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomField**
> UserCustomField GetCustomField(ctx, id)
Get the details of a specific custom field



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**UserCustomField**](UserCustomField.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCustomFields**
> GetCustomFieldsPaginatedResponse GetCustomFields(ctx, optional)
Get all custom fields



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetCustomFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetCustomFieldsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetCustomFieldsPaginatedResponse**](GetCustomFieldsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDedicatedNumber**
> UsersInbound GetDedicatedNumber(ctx, id)
Get the details of a specific dedicated number



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**UsersInbound**](UsersInbound.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFavourites**
> GetFavouritesPaginatedResponse GetFavourites(ctx, optional)
Get favorite contacts and lists



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetFavouritesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFavouritesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **query** | **optional.String**| Find contacts or lists by specified search query | 

### Return type

[**GetFavouritesPaginatedResponse**](GetFavouritesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInboundMessage**
> MessageIn GetInboundMessage(ctx, id)
Get a single inbound message



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The unique numeric ID for the inbound message. | 

### Return type

[**MessageIn**](MessageIn.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInboundMessagesNotificationSettings**
> GetInboundMessagesNotificationSettingsResponse GetInboundMessagesNotificationSettings(ctx, )
Get inbound messages notification settings



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetInboundMessagesNotificationSettingsResponse**](GetInboundMessagesNotificationSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInvoices**
> GetInvoicesPaginatedResponse GetInvoices(ctx, optional)
Get all invoices

With the TextMagic API, you can check the invoices and transactions for your account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetInvoicesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetInvoicesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetInvoicesPaginatedResponse**](GetInvoicesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetList**
> List GetList(ctx, id)
Get the details of a specific list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**List**](List.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListContactsIds**
> GetListContactsIdsResponse GetListContactsIds(ctx, id)
Get all contacts IDs in a list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**GetListContactsIdsResponse**](GetListContactsIdsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLists**
> GetListsPaginatedResponse GetLists(ctx, optional)
Get all lists



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetListsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| The current fetched page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]
 **favoriteOnly** | **optional.Int32**| Return only favorite lists | [default to 0]
 **onlyMine** | **optional.Int32**| Return only current user lists | [default to 0]

### Return type

[**GetListsPaginatedResponse**](GetListsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetListsOfContact**
> GetListsOfContactPaginatedResponse GetListsOfContact(ctx, id, optional)
Get contact's lists

Get all the lists in which the contact is included

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetListsOfContactOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetListsOfContactOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetListsOfContactPaginatedResponse**](GetListsOfContactPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagePreview**
> GetMessagePreviewResponse GetMessagePreview(ctx, optional)
Preview message

Get messages preview (with tags merged) up to 100 messages per session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetMessagePreviewOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMessagePreviewOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **optional.String**| Message text. Required if **template_id** is not set. | 
 **templateId** | **optional.Int32**| Template used instead of message text. Required if **text** is not set. | 
 **sendingTime** | **optional.Int32**| DEPRECATED, consider using sendingDateTime and sendingTimezone parameters instead: Optional (required with rrule set). Message sending time in unix timestamp format. Default is now. | 
 **sendingDateTime** | **optional.String**| Sending time in Y-m-d H:i:s format (e.g. 2016-05-27 13:02:33). This time is relative to sendingTimezone. | 
 **sendingTimezone** | **optional.String**| ID or ISO-name of timezone used for sending when sendingDateTime parameter is set. E.g. if you specify sendingDateTime &#x3D; \\\&quot;2016-05-27 13:02:33\\\&quot; and sendingTimezone &#x3D; \\\&quot;America/Buenos_Aires\\\&quot;, your message will be sent at May 27, 2016 13:02:33 Buenos Aires time, or 16:02:33 UTC. Default is account timezone. | 
 **contacts** | **optional.String**| Comma separated array of contact resources id message will be sent to. | 
 **lists** | **optional.String**| Comma separated array of list resources id message will be sent to. | 
 **phones** | **optional.String**| Comma separated array of E.164 phone numbers message will be sent to. | 
 **cutExtra** | **optional.Int32**| Should sending method cut extra characters which not fit supplied partsCount or return 400 Bad request response instead. | [default to 0]
 **partsCount** | **optional.Int32**| Maximum message parts count (TextMagic allows sending 1 to 6 message parts). | [default to 6]
 **referenceId** | **optional.Int32**| Custom message reference id which can be used in your application infrastructure. | 
 **from** | **optional.String**| One of allowed Sender ID (phone number or alphanumeric sender ID). If specified Sender ID is not allowed for some destinations, a fallback default Sender ID will be used to ensure delivery. See [Get timezones](http://docs.textmagictesting.com/#tag/Sender-IDs). | 
 **rule** | **optional.String**| iCal RRULE parameter to create recurrent scheduled messages. When used, sendingTime is mandatory as start point of sending. See https://www.textmagic.com/free-tools/rrule-generator for format details. | 
 **createChat** | **optional.Int32**| Should sending method try to create new Chat(if not exist) with specified recipients. | [default to 0]
 **tts** | **optional.Int32**| Send Text to Speech message. | [default to 0]
 **local** | **optional.Int32**| Treat phone numbers passed in \\&#39;phones\\&#39; field as local. | [default to 0]
 **localCountry** | **optional.String**| 2-letter ISO country code for local phone numbers, used when \\&#39;local\\&#39; is set to true. Default is account country. | 

### Return type

[**GetMessagePreviewResponse**](GetMessagePreviewResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagePrice**
> GetMessagePriceResponse GetMessagePrice(ctx, optional)
Check message price

Check pricing for a new outbound message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetMessagePriceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMessagePriceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeBlocked** | **optional.Int32**| Should we show pricing for the blocked contacts. | [default to 0]
 **text** | **optional.String**| Message text. Required if **template_id** is not set. | 
 **templateId** | **optional.Int32**| Template used instead of message text. Required if **text** is not set. | 
 **sendingTime** | **optional.Int32**| DEPRECATED, consider using sendingDateTime and sendingTimezone parameters instead: Optional (required with rrule set). Message sending time in unix timestamp format. Default is now. | 
 **sendingDateTime** | **optional.String**| Sending time in Y-m-d H:i:s format (e.g. 2016-05-27 13:02:33). This time is relative to sendingTimezone. | 
 **sendingTimezone** | **optional.String**| ID or ISO-name of timezone used for sending when sendingDateTime parameter is set. E.g. if you specify sendingDateTime &#x3D; \\\&quot;2016-05-27 13:02:33\\\&quot; and sendingTimezone &#x3D; \\\&quot;America/Buenos_Aires\\\&quot;, your message will be sent at May 27, 2016 13:02:33 Buenos Aires time, or 16:02:33 UTC. Default is account timezone. | 
 **contacts** | **optional.String**| Comma separated array of contact resources id message will be sent to. | 
 **lists** | **optional.String**| Comma separated array of list resources id message will be sent to. | 
 **phones** | **optional.String**| Comma separated array of E.164 phone numbers message will be sent to. | 
 **cutExtra** | **optional.Int32**| Should sending method cut extra characters which not fit supplied partsCount or return 400 Bad request response instead. | [default to 0]
 **partsCount** | **optional.Int32**| Maximum message parts count (TextMagic allows sending 1 to 6 message parts). | [default to 6]
 **referenceId** | **optional.Int32**| Custom message reference id which can be used in your application infrastructure. | 
 **from** | **optional.String**| One of allowed Sender ID (phone number or alphanumeric sender ID). If specified Sender ID is not allowed for some destinations, a fallback default Sender ID will be used to ensure delivery. See [Get timezones](http://docs.textmagictesting.com/#tag/Sender-IDs). | 
 **rule** | **optional.String**| iCal RRULE parameter to create recurrent scheduled messages. When used, sendingTime is mandatory as start point of sending. See https://www.textmagic.com/free-tools/rrule-generator for format details. | 
 **createChat** | **optional.Int32**| Should sending method try to create new Chat(if not exist) with specified recipients. | [default to 0]
 **tts** | **optional.Int32**| Send Text to Speech message. | [default to 0]
 **local** | **optional.Int32**| Treat phone numbers passed in \\&#39;phones\\&#39; field as local. | [default to 0]
 **localCountry** | **optional.String**| 2-letter ISO country code for local phone numbers, used when \\&#39;local\\&#39; is set to true. Default is account country. | 

### Return type

[**GetMessagePriceResponse**](GetMessagePriceResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessageSession**
> MessageSession GetMessageSession(ctx, id)
Get a session details

Get a specific session’s details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| a session ID | 

### Return type

[**MessageSession**](MessageSession.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessageSessionStat**
> GetMessageSessionStatResponse GetMessageSessionStat(ctx, id, optional)
Get a session statistics



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetMessageSessionStatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMessageSessionStatOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeDeleted** | **optional.Int32**| Search also in deleted messages | [default to 0]

### Return type

[**GetMessageSessionStatResponse**](GetMessageSessionStatResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagesBySessionId**
> GetMessagesBySessionIdPaginatedResponse GetMessagesBySessionId(ctx, id, optional)
Get a session messages

A useful synonym for \"messages/search\" command with provided \"sessionId\" parameter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***GetMessagesBySessionIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMessagesBySessionIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **statuses** | **optional.String**| Find messages by status | 
 **includeDeleted** | **optional.Int32**| Search also in deleted messages | [default to 0]

### Return type

[**GetMessagesBySessionIdPaginatedResponse**](GetMessagesBySessionIdPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagingCounters**
> GetMessagingCountersResponse GetMessagingCounters(ctx, )
Get sent/received messages counters values

Get total contacts, sent messages and received messages counters values.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetMessagingCountersResponse**](GetMessagingCountersResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagingStat**
> GetMessagingStatResponse GetMessagingStat(ctx, optional)
Get messaging statistics



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetMessagingStatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetMessagingStatOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **by** | **optional.String**| *   **off** to get total values per specified time interval *   **day** to show values grouped by day *   **month** to show values grouped by month *   **year** to show values grouped by year  | [default to off]
 **start** | **optional.Int32**| Time period start in [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time) format. The default is 7 days prior.  | 
 **end** | **optional.Int32**| Time period start in [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time) format. The default is today.  | 

### Return type

[**GetMessagingStatResponse**](GetMessagingStatResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutboundMessage**
> MessageOut GetOutboundMessage(ctx, id)
Get a single message

Get a single outgoing message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**MessageOut**](MessageOut.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOutboundMessagesHistory**
> GetOutboundMessagesHistoryPaginatedResponse GetOutboundMessagesHistory(ctx, optional)
Get history

Get outbound messages history.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetOutboundMessagesHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetOutboundMessagesHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **lastId** | **optional.Int32**| Filter results by ID, selecting all values lesser than the specified ID. | 
 **query** | **optional.String**| Find message by specified search query | 
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**GetOutboundMessagesHistoryPaginatedResponse**](GetOutboundMessagesHistoryPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetScheduledMessage**
> MessagesIcs GetScheduledMessage(ctx, id)
Get a single scheduled message



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**MessagesIcs**](MessagesIcs.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSenderId**
> SenderId GetSenderId(ctx, id)
Get the details of a specific Sender ID



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**SenderId**](SenderId.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSenderIds**
> GetSenderIdsPaginatedResponse GetSenderIds(ctx, optional)
Get all your approved Sender IDs



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSenderIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSenderIdsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetSenderIdsPaginatedResponse**](GetSenderIdsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSenderSettings**
> GetSenderSettingsResponse GetSenderSettings(ctx, optional)
Get current sender settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSenderSettingsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSenderSettingsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **optional.String**| Return sender settings enabled for sending to specified country. Two upper case characters | 

### Return type

[**GetSenderSettingsResponse**](GetSenderSettingsResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSpendingStat**
> GetSpendingStatPaginatedResponse GetSpendingStat(ctx, optional)
Get spending statistics



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSpendingStatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSpendingStatOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **start** | **optional.String**| Time period start in [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time) format. The default is 7 days prior.  | 
 **end** | **optional.String**| Time period start in [UNIX timestamp](https://en.wikipedia.org/wiki/Unix_time) format. The default is today.  | 

### Return type

[**GetSpendingStatPaginatedResponse**](GetSpendingStatPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubaccount**
> User GetSubaccount(ctx, id)
Get sub-account information



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**User**](User.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubaccounts**
> User GetSubaccounts(ctx, optional)
Get sub-accounts list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetSubaccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSubaccountsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**User**](User.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSubaccountsWithTokens**
> GetSubaccountsWithTokensResponse GetSubaccountsWithTokens(ctx, getSubaccountsWithTokensInputObject, optional)
Get all sub-accounts with their REST API tokens associated with app name

Get all sub-accounts with their REST API tokens associated with specified app name. When more than one token related to app name, last key will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **getSubaccountsWithTokensInputObject** | [**GetSubaccountsWithTokensInputObject**](GetSubaccountsWithTokensInputObject.md)|  | 
 **optional** | ***GetSubaccountsWithTokensOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetSubaccountsWithTokensOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Float32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetSubaccountsWithTokensResponse**](GetSubaccountsWithTokensResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTemplate**
> MessageTemplate GetTemplate(ctx, id)
Get a template details

Get a single template.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**MessageTemplate**](MessageTemplate.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTimezones**
> GetTimezonesResponse GetTimezones(ctx, optional)
Get timezones

Return all available timezone IDs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetTimezonesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetTimezonesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **full** | **optional.Int32**| Return full info about timezones in array (0 or 1). Default is 0 | [default to 0]

### Return type

[**GetTimezonesResponse**](GetTimezonesResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnreadMessagesTotal**
> GetUnreadMessagesTotalResponse GetUnreadMessagesTotal(ctx, )
Get unread messages number

Get total amount of unread messages in the current user chats.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetUnreadMessagesTotalResponse**](GetUnreadMessagesTotalResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnsubscribedContact**
> UnsubscribedContact GetUnsubscribedContact(ctx, id)
Get the details of a specific unsubscribed contact



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 

### Return type

[**UnsubscribedContact**](UnsubscribedContact.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUnsubscribers**
> GetUnsubscribersPaginatedResponse GetUnsubscribers(ctx, optional)
Get all unsubscribed contacts

When one of your message recipients sends a request with one of the [STOP-words](https://www.textmagic.com/sms-stop-command/), they will be immediately opted-out of your send lists and their contact status will change to an unsubscribed contact. To retrieve information on all contacts who have unsubscribed, use: 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUnsubscribersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUnsubscribersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]

### Return type

[**GetUnsubscribersPaginatedResponse**](GetUnsubscribersPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserDedicatedNumbers**
> GetUserDedicatedNumbersPaginatedResponse GetUserDedicatedNumbers(ctx, optional)
Get all your dedicated numbers



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetUserDedicatedNumbersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetUserDedicatedNumbersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **surveyId** | **optional.Int32**| Fetch only that numbers which are ready for the survey | 

### Return type

[**GetUserDedicatedNumbersPaginatedResponse**](GetUserDedicatedNumbersPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InviteSubaccount**
> InviteSubaccount(ctx, inviteSubaccountInputObject)
Invite a new sub-account



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **inviteSubaccountInputObject** | [**InviteSubaccountInputObject**](InviteSubaccountInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkChatsReadBulk**
> MarkChatsReadBulk(ctx, markChatsReadBulkInputObject)
Mark chats as read (bulk)

Mark several chats as read by chat ids or mark all chats as read

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **markChatsReadBulkInputObject** | [**MarkChatsReadBulkInputObject**](MarkChatsReadBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarkChatsUnreadBulk**
> MarkChatsUnreadBulk(ctx, markChatsUnreadBulkInputObject)
Mark chats as unread (bulk)

Mark several chats as UNread by chat ids or mark all chats as UNread

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **markChatsUnreadBulkInputObject** | [**MarkChatsUnreadBulkInputObject**](MarkChatsUnreadBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MuteChat**
> ResourceLinkResponse MuteChat(ctx, muteChatInputObject)
Mute chat sounds



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **muteChatInputObject** | [**MuteChatInputObject**](MuteChatInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MuteChatsBulk**
> MuteChatsBulk(ctx, muteChatsBulkInputObject)
Mute chats (bulk)

Mute several chats by chat ids or mute all chats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **muteChatsBulkInputObject** | [**MuteChatsBulkInputObject**](MuteChatsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ping**
> PingResponse Ping(ctx, )
Ping

Make a simple ping request

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PingResponse**](PingResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReopenChatsBulk**
> ReopenChatsBulk(ctx, reopenChatsBulkInputObject)
Reopen chats (bulk)

Reopen chats by chat ids or reopen all chats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **reopenChatsBulkInputObject** | [**ReopenChatsBulkInputObject**](ReopenChatsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestNewSubaccountToken**
> User RequestNewSubaccountToken(ctx, requestNewSubaccountTokenInputObject)
Request a new REST API token for sub-account

Returning user object, key and app name.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestNewSubaccountTokenInputObject** | [**RequestNewSubaccountTokenInputObject**](RequestNewSubaccountTokenInputObject.md)|  | 

### Return type

[**User**](User.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSenderId**
> ResourceLinkResponse RequestSenderId(ctx, requestSenderIdInputObject)
Apply for a new Sender ID

> Sender IDs are shared between all of your sub-accounts.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestSenderIdInputObject** | [**RequestSenderIdInputObject**](RequestSenderIdInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchChats**
> SearchChatsPaginatedResponse SearchChats(ctx, optional)
Find chats by message text



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchChatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchChatsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **query** | **optional.String**| Find chats by specified search query | 

### Return type

[**SearchChatsPaginatedResponse**](SearchChatsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchChatsByIds**
> SearchChatsByIdsPaginatedResponse SearchChatsByIds(ctx, optional)
Find chats (bulk)



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchChatsByIdsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchChatsByIdsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **ids** | **optional.String**| Find chats by ID(s) | 

### Return type

[**SearchChatsByIdsPaginatedResponse**](SearchChatsByIdsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchChatsByReceipent**
> SearchChatsByReceipentPaginatedResponse SearchChatsByReceipent(ctx, optional)
Find chats by recipient

Find chats by recipient (contact, list name or phone number).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchChatsByReceipentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchChatsByReceipentOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **query** | **optional.String**| Find chats by specified search query | 
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]

### Return type

[**SearchChatsByReceipentPaginatedResponse**](SearchChatsByReceipentPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchContacts**
> SearchContactsPaginatedResponse SearchContacts(ctx, optional)
Find contacts by given criteria



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchContactsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchContactsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **shared** | **optional.Int32**| Should shared contacts to be included | [default to 0]
 **ids** | **optional.String**| Find contact by ID(s) | 
 **listId** | **optional.Int32**| Find contact by List ID | 
 **includeBlocked** | **optional.Int32**| Should blocked contacts to be included | 
 **query** | **optional.String**| Find contacts by specified search query | 
 **local** | **optional.Int32**| Treat phone number passed in &#39;query&#39; field as local. Default is 0 | [default to 0]
 **country** | **optional.String**| 2-letter ISO country code for local phone numbers, used when &#39;local&#39; is set to true. Default is account country | 
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**SearchContactsPaginatedResponse**](SearchContactsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchInboundMessages**
> SearchInboundMessagesPaginatedResponse SearchInboundMessages(ctx, optional)
Find inbound messages

Find inbound messages by given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchInboundMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchInboundMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **ids** | **optional.String**| Find message by ID(s) | 
 **query** | **optional.String**| Find recipients by specified search query | 
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]
 **expand** | **optional.Int32**| Expand by adding firstName, lastName and contactId | [default to 0]

### Return type

[**SearchInboundMessagesPaginatedResponse**](SearchInboundMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLists**
> SearchListsPaginatedResponse SearchLists(ctx, optional)
Find lists by given criteria



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchListsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchListsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **ids** | **optional.String**| Find lists by ID(s) | 
 **query** | **optional.String**| Find lists by specified search query | 
 **onlyMine** | **optional.Int32**| Return only current user lists | [default to 0]
 **onlyDefault** | **optional.Int32**| Return only default lists | [default to 0]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**SearchListsPaginatedResponse**](SearchListsPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOutboundMessages**
> SearchOutboundMessagesPaginatedResponse SearchOutboundMessages(ctx, optional)
Find messages

Find outbound messages by given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchOutboundMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchOutboundMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **lastId** | **optional.Int32**| Filter results by ID, selecting all values lesser than the specified ID. Note that \\&#39;page\\&#39; parameter is ignored when \\&#39;lastId\\&#39; is specified | 
 **ids** | **optional.String**| Find message by ID(s) | 
 **sessionId** | **optional.Int32**| Find messages by session ID | 
 **statuses** | **optional.String**| Find messages by status | 
 **includeDeleted** | **optional.Int32**| Search also in deleted messages | [default to 0]
 **query** | **optional.String**| Find messages by specified search query | 

### Return type

[**SearchOutboundMessagesPaginatedResponse**](SearchOutboundMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchScheduledMessages**
> SearchScheduledMessagesPaginatedResponse SearchScheduledMessages(ctx, optional)
Find scheduled messages



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchScheduledMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchScheduledMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **query** | **optional.String**| Find messages by specified search query | 
 **ids** | **optional.String**| Find schedules by ID(s) | 
 **status** | **optional.String**| Fetch schedules with the specific status: a - actual, c - completed, x - all | [default to x]
 **orderBy** | **optional.String**| Order results by some field. Default is id | [default to id]
 **direction** | **optional.String**| Order direction. Default is desc | [default to desc]

### Return type

[**SearchScheduledMessagesPaginatedResponse**](SearchScheduledMessagesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTemplates**
> SearchTemplatesPaginatedResponse SearchTemplates(ctx, optional)
Find templates by criteria

Find user templates by given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchTemplatesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| Fetch specified results page. | [default to 1]
 **limit** | **optional.Int32**| The number of results per page. | [default to 10]
 **ids** | **optional.String**| Find template by ID(s) | 
 **name** | **optional.String**| Find template by name | 
 **content** | **optional.String**| Find template by content | 

### Return type

[**SearchTemplatesPaginatedResponse**](SearchTemplatesPaginatedResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendMessage**
> SendMessageResponse SendMessage(ctx, sendMessageInputObject)
Send message

The main entrypoint to send messages. See examples above for the reference.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sendMessageInputObject** | [**SendMessageInputObject**](SendMessageInputObject.md)|  | 

### Return type

[**SendMessageResponse**](SendMessageResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendPhoneVerificationCodeTFA**
> SendPhoneVerificationCodeResponse SendPhoneVerificationCodeTFA(ctx, sendPhoneVerificationCodeInputObject)
Step 1: Send a verification code 

Sends verification code to specified phone number.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sendPhoneVerificationCodeInputObject** | [**SendPhoneVerificationCodeInputObject**](SendPhoneVerificationCodeInputObject.md)|  | 

### Return type

[**SendPhoneVerificationCodeResponse**](SendPhoneVerificationCodeResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetChatStatus**
> ResourceLinkResponse SetChatStatus(ctx, setChatStatusInputObject)
Change chat status

Set status of the chat given by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **setChatStatusInputObject** | [**SetChatStatusInputObject**](SetChatStatusInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnblockContact**
> UnblockContact(ctx, unblockContactInputObject)
Unblock contact by phone number.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unblockContactInputObject** | [**UnblockContactInputObject**](UnblockContactInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnblockContactsBulk**
> UnblockContactsBulk(ctx, unblockContactsBulkInputObject)
Unblock contacts (bulk)

Unblock several contacts by blocked contact ids or unblock all contacts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unblockContactsBulkInputObject** | [**UnblockContactsBulkInputObject**](UnblockContactsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnmuteChatsBulk**
> UnmuteChatsBulk(ctx, unmuteChatsBulkInputObject)
Unmute chats (bulk)

Unmute several chats by chat ids or unmute all chats

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unmuteChatsBulkInputObject** | [**UnmuteChatsBulkInputObject**](UnmuteChatsBulkInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnsubscribeContact**
> ResourceLinkResponse UnsubscribeContact(ctx, unsubscribeContactInputObject)
Manually unsubscribe a contact

> Please note, if you unsubscribe a contact, this action cannot be reversed. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **unsubscribeContactInputObject** | [**UnsubscribeContactInputObject**](UnsubscribeContactInputObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBalanceNotificationSettings**
> UpdateBalanceNotificationSettings(ctx, updateBalanceNotificationSettingsInputObject)
Update balance notification settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateBalanceNotificationSettingsInputObject** | [**UpdateBalanceNotificationSettingsInputObject**](UpdateBalanceNotificationSettingsInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCallbackSettings**
> UpdateCallbackSettings(ctx, updateCallbackSettingsInputObject)
Update callback URL settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateCallbackSettingsInputObject** | [**UpdateCallbackSettingsInputObject**](UpdateCallbackSettingsInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateChatDesktopNotificationSettings**
> UpdateChatDesktopNotificationSettings(ctx, updateChatDesktopNotificationSettingsInputObject)
Update chat desktop notification settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateChatDesktopNotificationSettingsInputObject** | [**UpdateChatDesktopNotificationSettingsInputObject**](UpdateChatDesktopNotificationSettingsInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContact**
> ResourceLinkResponse UpdateContact(ctx, updateContactInputObject, id)
Edit a contact



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateContactInputObject** | [**UpdateContactInputObject**](UpdateContactInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContactNote**
> ResourceLinkResponse UpdateContactNote(ctx, updateContactNoteInputObject, id)
Update a contact note



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateContactNoteInputObject** | [**UpdateContactNoteInputObject**](UpdateContactNoteInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCurrentUser**
> UpdateCurrentUserResponse UpdateCurrentUser(ctx, updateCurrentUserInputObject)
Edit current account info



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateCurrentUserInputObject** | [**UpdateCurrentUserInputObject**](UpdateCurrentUserInputObject.md)|  | 

### Return type

[**UpdateCurrentUserResponse**](UpdateCurrentUserResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomField**
> ResourceLinkResponse UpdateCustomField(ctx, updateCustomFieldInputObject, id)
Edit a custom field



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateCustomFieldInputObject** | [**UpdateCustomFieldInputObject**](UpdateCustomFieldInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomFieldValue**
> ResourceLinkResponse UpdateCustomFieldValue(ctx, updateCustomFieldValueInputObject, id)
Edit the custom field value of a specified contact



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateCustomFieldValueInputObject** | [**UpdateCustomFieldValueInputObject**](UpdateCustomFieldValueInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateInboundMessagesNotificationSettings**
> UpdateInboundMessagesNotificationSettings(ctx, updateInboundMessagesNotificationSettingsInputObject)
Update inbound messages notification settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateInboundMessagesNotificationSettingsInputObject** | [**UpdateInboundMessagesNotificationSettingsInputObject**](UpdateInboundMessagesNotificationSettingsInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateList**
> ResourceLinkResponse UpdateList(ctx, id, optional)
Edit a list



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**|  | 
 **optional** | ***UpdateListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateListObject** | [**optional.Interface of UpdateListObject**](UpdateListObject.md)|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSenderSetting**
> UpdateSenderSetting(ctx, updateSenderSettingInputObject)
Change sender settings



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateSenderSettingInputObject** | [**UpdateSenderSettingInputObject**](UpdateSenderSettingInputObject.md)|  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTemplate**
> ResourceLinkResponse UpdateTemplate(ctx, updateTemplateInputObject, id)
Update a template



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **updateTemplateInputObject** | [**UpdateTemplateInputObject**](UpdateTemplateInputObject.md)|  | 
  **id** | **int32**|  | 

### Return type

[**ResourceLinkResponse**](ResourceLinkResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

