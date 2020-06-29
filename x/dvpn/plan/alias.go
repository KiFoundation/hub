// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/plan/types
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/plan/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/dvpn/plan/querier
package plan

import (
	"github.com/sentinel-official/hub/x/dvpn/plan/keeper"
	"github.com/sentinel-official/hub/x/dvpn/plan/querier"
	"github.com/sentinel-official/hub/x/dvpn/plan/types"
)

const (
	Codespace                  = types.Codespace
	EventTypeSetPlan           = types.EventTypeSetPlan
	EventTypeSetPlansCount     = types.EventTypeSetPlansCount
	EventTypeSetPlanStatus     = types.EventTypeSetPlanStatus
	EventTypeAddNodeForPlan    = types.EventTypeAddNodeForPlan
	EventTypeRemoveNodeForPlan = types.EventTypeRemoveNodeForPlan
	AttributeKeyAddress        = types.AttributeKeyAddress
	AttributeKeyID             = types.AttributeKeyID
	AttributeKeyCount          = types.AttributeKeyCount
	AttributeKeyStatus         = types.AttributeKeyStatus
	ModuleName                 = types.ModuleName
	QuerierRoute               = types.QuerierRoute
	QueryPlan                  = types.QueryPlan
	QueryPlans                 = types.QueryPlans
	QueryPlansForProvider      = types.QueryPlansForProvider
	QueryPlansForNode          = types.QueryPlansForNode
	QueryNodesForPlan          = types.QueryNodesForPlan
)

var (
	// functions aliases
	RegisterCodec                  = types.RegisterCodec
	ErrorMarshal                   = types.ErrorMarshal
	ErrorUnmarshal                 = types.ErrorUnmarshal
	ErrorUnknownMsgType            = types.ErrorUnknownMsgType
	ErrorUnknownQueryType          = types.ErrorUnknownQueryType
	ErrorInvalidField              = types.ErrorInvalidField
	ErrorProviderDoesNotExist      = types.ErrorProviderDoesNotExist
	ErrorPlanDoesNotExist          = types.ErrorPlanDoesNotExist
	ErrorNodeDoesNotExist          = types.ErrorNodeDoesNotExist
	ErrorUnauthorized              = types.ErrorUnauthorized
	ErrorDuplicateNode             = types.ErrorDuplicateNode
	ErrorNodeWasNotAdded           = types.ErrorNodeWasNotAdded
	NewGenesisState                = types.NewGenesisState
	DefaultGenesisState            = types.DefaultGenesisState
	PlanKey                        = types.PlanKey
	PlanForProviderKeyPrefix       = types.PlanForProviderKeyPrefix
	PlanForProviderKey             = types.PlanForProviderKey
	PlanForNodeKeyPrefix           = types.PlanForNodeKeyPrefix
	PlanForNodeKey                 = types.PlanForNodeKey
	NodeForPlanKeyPrefix           = types.NodeForPlanKeyPrefix
	NodeForPlanKey                 = types.NodeForPlanKey
	NewMsgAddPlan                  = types.NewMsgAddPlan
	NewMsgSetPlanStatus            = types.NewMsgSetPlanStatus
	NewMsgAddNodeForPlan           = types.NewMsgAddNodeForPlan
	NewMsgRemoveNodeForPlan        = types.NewMsgRemoveNodeForPlan
	NewQueryPlanParams             = types.NewQueryPlanParams
	NewQueryPlansForProviderParams = types.NewQueryPlansForProviderParams
	NewQueryPlansForNodeParams     = types.NewQueryPlansForNodeParams
	NewQueryNodesForPlanParams     = types.NewQueryNodesForPlanParams
	NewKeeper                      = keeper.NewKeeper
	Querier                        = querier.Querier

	// variable aliases
	ModuleCdc     = types.ModuleCdc
	RouterKey     = types.RouterKey
	StoreKey      = types.StoreKey
	PlansCountKey = types.PlansCountKey
	PlanKeyPrefix = types.PlanKeyPrefix
)

type (
	GenesisPlan                 = types.GenesisPlan
	GenesisPlans                = types.GenesisPlans
	GenesisState                = types.GenesisState
	MsgAddPlan                  = types.MsgAddPlan
	MsgSetPlanStatus            = types.MsgSetPlanStatus
	MsgAddNodeForPlan           = types.MsgAddNodeForPlan
	MsgRemoveNodeForPlan        = types.MsgRemoveNodeForPlan
	Plan                        = types.Plan
	Plans                       = types.Plans
	QueryPlanParams             = types.QueryPlanParams
	QueryPlansForProviderParams = types.QueryPlansForProviderParams
	QueryPlansForNodeParams     = types.QueryPlansForNodeParams
	QueryNodesForPlanParams     = types.QueryNodesForPlanParams
	Keeper                      = keeper.Keeper
)
