// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/provider/types
// ALIASGEN: github.com/sentinel-official/hub/x/provider/keeper
package provider

import (
	"github.com/sentinel-official/hub/x/provider/keeper"
	"github.com/sentinel-official/hub/x/provider/types"
)

const (
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	RegisterLegacyAminoCodec                = types.RegisterLegacyAminoCodec
	RegisterInterfaces                      = types.RegisterInterfaces
	NewGenesisState                         = types.NewGenesisState
	DefaultGenesisState                     = types.DefaultGenesisState
	ProviderKey                             = types.ProviderKey
	NewMsgRegisterRequest                   = types.NewMsgRegisterRequest
	NewMsgUpdateRequest                     = types.NewMsgUpdateRequest
	NewMsgServiceClient                     = types.NewMsgServiceClient
	RegisterMsgServiceServer                = types.RegisterMsgServiceServer
	NewParams                               = types.NewParams
	DefaultParams                           = types.DefaultParams
	ParamsKeyTable                          = types.ParamsKeyTable
	NewQueryProviderRequest                 = types.NewQueryProviderRequest
	NewQueryProvidersRequest                = types.NewQueryProvidersRequest
	NewQueryServiceClient                   = types.NewQueryServiceClient
	RegisterQueryServiceServer              = types.RegisterQueryServiceServer
	RegisterQueryServiceHandlerServer       = types.RegisterQueryServiceHandlerServer
	RegisterQueryServiceHandlerFromEndpoint = types.RegisterQueryServiceHandlerFromEndpoint
	RegisterQueryServiceHandler             = types.RegisterQueryServiceHandler
	RegisterQueryServiceHandlerClient       = types.RegisterQueryServiceHandlerClient
	NewKeeper                               = keeper.NewKeeper
	NewMsgServiceServer                     = keeper.NewMsgServiceServer
	NewQueryServiceServer                   = keeper.NewQueryServiceServer

	// variable aliases
	ModuleCdc                       = types.ModuleCdc
	ErrorMarshal                    = types.ErrorMarshal
	ErrorUnmarshal                  = types.ErrorUnmarshal
	ErrorUnknownMsgType             = types.ErrorUnknownMsgType
	ErrorUnknownQueryType           = types.ErrorUnknownQueryType
	ErrorInvalidField               = types.ErrorInvalidField
	ErrorDuplicateProvider          = types.ErrorDuplicateProvider
	ErrorProviderDoesNotExist       = types.ErrorProviderDoesNotExist
	ErrInvalidLengthEvents          = types.ErrInvalidLengthEvents
	ErrIntOverflowEvents            = types.ErrIntOverflowEvents
	ErrUnexpectedEndOfGroupEvents   = types.ErrUnexpectedEndOfGroupEvents
	ErrInvalidLengthGenesis         = types.ErrInvalidLengthGenesis
	ErrIntOverflowGenesis           = types.ErrIntOverflowGenesis
	ErrUnexpectedEndOfGroupGenesis  = types.ErrUnexpectedEndOfGroupGenesis
	ParamsSubspace                  = types.ParamsSubspace
	RouterKey                       = types.RouterKey
	StoreKey                        = types.StoreKey
	EventModuleName                 = types.EventModuleName
	ProviderKeyPrefix               = types.ProviderKeyPrefix
	ErrInvalidLengthMsg             = types.ErrInvalidLengthMsg
	ErrIntOverflowMsg               = types.ErrIntOverflowMsg
	ErrUnexpectedEndOfGroupMsg      = types.ErrUnexpectedEndOfGroupMsg
	DefaultDeposit                  = types.DefaultDeposit
	KeyDeposit                      = types.KeyDeposit
	ErrInvalidLengthParams          = types.ErrInvalidLengthParams
	ErrIntOverflowParams            = types.ErrIntOverflowParams
	ErrUnexpectedEndOfGroupParams   = types.ErrUnexpectedEndOfGroupParams
	ErrInvalidLengthProvider        = types.ErrInvalidLengthProvider
	ErrIntOverflowProvider          = types.ErrIntOverflowProvider
	ErrUnexpectedEndOfGroupProvider = types.ErrUnexpectedEndOfGroupProvider
	ErrInvalidLengthQuerier         = types.ErrInvalidLengthQuerier
	ErrIntOverflowQuerier           = types.ErrIntOverflowQuerier
	ErrUnexpectedEndOfGroupQuerier  = types.ErrUnexpectedEndOfGroupQuerier
)

type (
	EventModule                     = types.EventModule
	EventRegisterProvider           = types.EventRegisterProvider
	EventUpdateProvider             = types.EventUpdateProvider
	GenesisState                    = types.GenesisState
	MsgRegisterRequest              = types.MsgRegisterRequest
	MsgUpdateRequest                = types.MsgUpdateRequest
	MsgRegisterResponse             = types.MsgRegisterResponse
	MsgUpdateResponse               = types.MsgUpdateResponse
	MsgServiceClient                = types.MsgServiceClient
	MsgServiceServer                = types.MsgServiceServer
	UnimplementedMsgServiceServer   = types.UnimplementedMsgServiceServer
	Params                          = types.Params
	Providers                       = types.Providers
	Provider                        = types.Provider
	QueryProvidersRequest           = types.QueryProvidersRequest
	QueryProviderRequest            = types.QueryProviderRequest
	QueryProvidersResponse          = types.QueryProvidersResponse
	QueryProviderResponse           = types.QueryProviderResponse
	QueryServiceClient              = types.QueryServiceClient
	QueryServiceServer              = types.QueryServiceServer
	UnimplementedQueryServiceServer = types.UnimplementedQueryServiceServer
	Keeper                          = keeper.Keeper
)
