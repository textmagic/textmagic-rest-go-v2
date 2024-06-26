/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type GetStateResponse struct {
	SystemCacheClear int32 `json:"systemCacheClear"`
	SystemExit int32 `json:"systemExit"`
	SystemAlert int32 `json:"systemAlert"`
	SystemAccountStateChanged int32 `json:"systemAccountStateChanged"`
	SystemAccountClosed int32 `json:"systemAccountClosed"`
	SystemAccountAdditionalFields int32 `json:"systemAccountAdditionalFields"`
	SystemAccountPermissionsChanged int32 `json:"systemAccountPermissionsChanged"`
	UserProfileChanged int32 `json:"userProfileChanged"`
	UserBalanceChanged int32 `json:"userBalanceChanged"`
	UserImpersonationEnd int32 `json:"userImpersonationEnd"`
	MessageDeleted int32 `json:"messageDeleted"`
	MessageIncoming int32 `json:"messageIncoming"`
	MessageIncomingDeleted int32 `json:"messageIncomingDeleted"`
	MessageStateChanged int32 `json:"messageStateChanged"`
	MessageBulkEnd int32 `json:"messageBulkEnd"`
	MessageWipeEnd int32 `json:"messageWipeEnd"`
	MessageSent int32 `json:"messageSent"`
	MessageSessionDeleted int32 `json:"messageSessionDeleted"`
	MessageCacheClear int32 `json:"messageCacheClear"`
	MessageIncomingCacheClear int32 `json:"messageIncomingCacheClear"`
	MessageScheduleAdded int32 `json:"messageScheduleAdded"`
	MessageScheduleStateChanged int32 `json:"messageScheduleStateChanged"`
	MessageScheduleDeleted int32 `json:"messageScheduleDeleted"`
	MessageScheduleNotSentStateChanged int32 `json:"messageScheduleNotSentStateChanged"`
	MessageScheduleCacheClear int32 `json:"messageScheduleCacheClear"`
	MessageTemplateCacheClear int32 `json:"messageTemplateCacheClear"`
	CallFinished int32 `json:"callFinished"`
	ChatCreated int32 `json:"chatCreated"`
	ChatMarkedAsRead int32 `json:"chatMarkedAsRead"`
	ChatMuted int32 `json:"chatMuted"`
	ChatUnmuted int32 `json:"chatUnmuted"`
	ChatPinned int32 `json:"chatPinned"`
	ChatUnpinned int32 `json:"chatUnpinned"`
	ChatDeleted int32 `json:"chatDeleted"`
	ChatClosed int32 `json:"chatClosed"`
	ChatReopened int32 `json:"chatReopened"`
	ChatCacheClear int32 `json:"chatCacheClear"`
	ChatRead int32 `json:"chatRead"`
	ChatUnread int32 `json:"chatUnread"`
	ContactAdded int32 `json:"contactAdded"`
	ContactDeleted int32 `json:"contactDeleted"`
	ContactStateChanged int32 `json:"contactStateChanged"`
	ListAdded int32 `json:"listAdded"`
	ListDeleted int32 `json:"listDeleted"`
	ListStateChanged int32 `json:"listStateChanged"`
	ContactWipeEnd int32 `json:"contactWipeEnd"`
	ContactImportEnd int32 `json:"contactImportEnd"`
	ContactCacheClear int32 `json:"contactCacheClear"`
	ListCacheClear int32 `json:"listCacheClear"`
	CustomFieldsCacheClear int32 `json:"customFieldsCacheClear"`
	ProgressCarrierBulkLookup int32 `json:"progressCarrierBulkLookup"`
	ProgressEmailBulkLookup int32 `json:"progressEmailBulkLookup"`
	ProgressSubAccountBulkImport int32 `json:"progressSubAccountBulkImport"`
	ProgressContactBulkImport int32 `json:"progressContactBulkImport"`
	ForceRefreshWebApp int32 `json:"forceRefreshWebApp"`
	ChatSenderSettingsChanged int32 `json:"chatSenderSettingsChanged"`
	CountrySenderSettingsChanged int32 `json:"countrySenderSettingsChanged"`
	ChatSummaryChunk int32 `json:"chatSummaryChunk"`
	ChatWaysToReplyChunk int32 `json:"chatWaysToReplyChunk"`
	ChatSuggestedReplyChunk int32 `json:"chatSuggestedReplyChunk"`
	UserSubscriptionChanged int32 `json:"userSubscriptionChanged"`
	UserSubscriptionDeleted int32 `json:"userSubscriptionDeleted"`
}
