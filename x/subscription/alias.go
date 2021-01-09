// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/subscription/types
// ALIASGEN: github.com/sentinel-official/hub/x/subscription/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/subscription/querier
package subscription

import (
	"github.com/sentinel-official/hub/x/subscription/keeper"
	"github.com/sentinel-official/hub/x/subscription/querier"
	"github.com/sentinel-official/hub/x/subscription/types"
)

const (
	AttributeKeyOwner            = types.AttributeKeyOwner
	AttributeKeyAddress          = types.AttributeKeyAddress
	AttributeKeyID               = types.AttributeKeyID
	AttributeKeyStatus           = types.AttributeKeyStatus
	AttributeKeyNode             = types.AttributeKeyNode
	AttributeKeyCount            = types.AttributeKeyCount
	AttributeKeyPlan             = types.AttributeKeyPlan
	AttributeKeyConsumed         = types.AttributeKeyConsumed
	AttributeKeyAllocated        = types.AttributeKeyAllocated
	ModuleName                   = types.ModuleName
	ParamsSubspace               = types.ParamsSubspace
	QuerierRoute                 = types.QuerierRoute
	DefaultInactiveDuration      = types.DefaultInactiveDuration
	QuerySubscription            = types.QuerySubscription
	QuerySubscriptions           = types.QuerySubscriptions
	QuerySubscriptionsForNode    = types.QuerySubscriptionsForNode
	QuerySubscriptionsForPlan    = types.QuerySubscriptionsForPlan
	QuerySubscriptionsForAddress = types.QuerySubscriptionsForAddress
	QueryQuota                   = types.QueryQuota
	QueryQuotas                  = types.QueryQuotas
)

var (
	// functions aliases
	RegisterCodec                              = types.RegisterCodec
	NewGenesisState                            = types.NewGenesisState
	DefaultGenesisState                        = types.DefaultGenesisState
	SubscriptionKey                            = types.SubscriptionKey
	GetSubscriptionForNodeKeyPrefix            = types.GetSubscriptionForNodeKeyPrefix
	SubscriptionForNodeKey                     = types.SubscriptionForNodeKey
	GetSubscriptionForPlanKeyPrefix            = types.GetSubscriptionForPlanKeyPrefix
	SubscriptionForPlanKey                     = types.SubscriptionForPlanKey
	GetActiveSubscriptionForAddressKeyPrefix   = types.GetActiveSubscriptionForAddressKeyPrefix
	ActiveSubscriptionForAddressKey            = types.ActiveSubscriptionForAddressKey
	GetInactiveSubscriptionForAddressKeyPrefix = types.GetInactiveSubscriptionForAddressKeyPrefix
	InactiveSubscriptionForAddressKey          = types.InactiveSubscriptionForAddressKey
	GetInactiveSubscriptionAtKeyPrefix         = types.GetInactiveSubscriptionAtKeyPrefix
	InactiveSubscriptionAtKey                  = types.InactiveSubscriptionAtKey
	GetQuotaKeyPrefix                          = types.GetQuotaKeyPrefix
	QuotaKey                                   = types.QuotaKey
	IDFromSubscriptionForNodeKey               = types.IDFromSubscriptionForNodeKey
	IDFromSubscriptionForPlanKey               = types.IDFromSubscriptionForPlanKey
	IDFromStatusSubscriptionForAddressKey      = types.IDFromStatusSubscriptionForAddressKey
	IDFromInactiveSubscriptionAtKey            = types.IDFromInactiveSubscriptionAtKey
	NewMsgSubscribeToNode                      = types.NewMsgSubscribeToNode
	NewMsgSubscribeToPlan                      = types.NewMsgSubscribeToPlan
	NewMsgCancel                               = types.NewMsgCancel
	NewMsgAddQuota                             = types.NewMsgAddQuota
	NewMsgUpdateQuota                          = types.NewMsgUpdateQuota
	NewParams                                  = types.NewParams
	DefaultParams                              = types.DefaultParams
	ParamsKeyTable                             = types.ParamsKeyTable
	NewQuerySubscriptionParams                 = types.NewQuerySubscriptionParams
	NewQuerySubscriptionsParams                = types.NewQuerySubscriptionsParams
	NewQuerySubscriptionsForNodeParams         = types.NewQuerySubscriptionsForNodeParams
	NewQuerySubscriptionsForPlanParams         = types.NewQuerySubscriptionsForPlanParams
	NewQuerySubscriptionsForAddressParams      = types.NewQuerySubscriptionsForAddressParams
	NewQueryQuotaParams                        = types.NewQueryQuotaParams
	NewQueryQuotasParams                       = types.NewQueryQuotasParams
	NewKeeper                                  = keeper.NewKeeper
	Querier                                    = querier.Querier

	// variable aliases
	ModuleCdc                               = types.ModuleCdc
	ErrorMarshal                            = types.ErrorMarshal
	ErrorUnmarshal                          = types.ErrorUnmarshal
	ErrorUnknownMsgType                     = types.ErrorUnknownMsgType
	ErrorUnknownQueryType                   = types.ErrorUnknownQueryType
	ErrorInvalidField                       = types.ErrorInvalidField
	ErrorPlanDoesNotExist                   = types.ErrorPlanDoesNotExist
	ErrorNodeDoesNotExist                   = types.ErrorNodeDoesNotExist
	ErrorUnauthorized                       = types.ErrorUnauthorized
	ErrorInvalidPlanStatus                  = types.ErrorInvalidPlanStatus
	ErrorPriceDoesNotExist                  = types.ErrorPriceDoesNotExist
	ErrorInvalidNodeStatus                  = types.ErrorInvalidNodeStatus
	ErrorSubscriptionDoesNotExist           = types.ErrorSubscriptionDoesNotExist
	ErrorInvalidSubscriptionStatus          = types.ErrorInvalidSubscriptionStatus
	ErrorCanNotSubscribe                    = types.ErrorCanNotSubscribe
	ErrorInvalidQuota                       = types.ErrorInvalidQuota
	ErrorDuplicateQuota                     = types.ErrorDuplicateQuota
	ErrorQuotaDoesNotExist                  = types.ErrorQuotaDoesNotExist
	ErrorCanNotAddQuota                     = types.ErrorCanNotAddQuota
	EventTypeSetCount                       = types.EventTypeSetCount
	EventTypeSet                            = types.EventTypeSet
	EventTypeCancel                         = types.EventTypeCancel
	EventTypeAddQuota                       = types.EventTypeAddQuota
	EventTypeUpdateQuota                    = types.EventTypeUpdateQuota
	RouterKey                               = types.RouterKey
	StoreKey                                = types.StoreKey
	EventModuleName                         = types.EventModuleName
	CountKey                                = types.CountKey
	SubscriptionKeyPrefix                   = types.SubscriptionKeyPrefix
	SubscriptionForNodeKeyPrefix            = types.SubscriptionForNodeKeyPrefix
	SubscriptionForPlanKeyPrefix            = types.SubscriptionForPlanKeyPrefix
	ActiveSubscriptionForAddressKeyPrefix   = types.ActiveSubscriptionForAddressKeyPrefix
	InactiveSubscriptionForAddressKeyPrefix = types.InactiveSubscriptionForAddressKeyPrefix
	InactiveSubscriptionAtKeyPrefix         = types.InactiveSubscriptionAtKeyPrefix
	QuotaKeyPrefix                          = types.QuotaKeyPrefix
	KeyInactiveDuration                     = types.KeyInactiveDuration
)

type (
	GenesisSubscription                = types.GenesisSubscription
	GenesisSubscriptions               = types.GenesisSubscriptions
	GenesisState                       = types.GenesisState
	MsgSubscribeToNode                 = types.MsgSubscribeToNode
	MsgSubscribeToPlan                 = types.MsgSubscribeToPlan
	MsgCancel                          = types.MsgCancel
	MsgAddQuota                        = types.MsgAddQuota
	MsgUpdateQuota                     = types.MsgUpdateQuota
	Params                             = types.Params
	QuerySubscriptionParams            = types.QuerySubscriptionParams
	QuerySubscriptionsParams           = types.QuerySubscriptionsParams
	QuerySubscriptionsForNodeParams    = types.QuerySubscriptionsForNodeParams
	QuerySubscriptionsForPlanParams    = types.QuerySubscriptionsForPlanParams
	QuerySubscriptionsForAddressParams = types.QuerySubscriptionsForAddressParams
	QueryQuotaParams                   = types.QueryQuotaParams
	QueryQuotasParams                  = types.QueryQuotasParams
	Quota                              = types.Quota
	Quotas                             = types.Quotas
	Subscription                       = types.Subscription
	Subscriptions                      = types.Subscriptions
	Keeper                             = keeper.Keeper
)
