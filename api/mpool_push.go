package api

import (
	"context"
	"net/http"

	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/env"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/plugin/fil"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
)

func (s Server) MpoolPush(ctx context.Context, in *sphinxplugin.MpoolPushRequest) (*sphinxplugin.MpoolPushResponse, error) {
	inMsg := in.GetMessage()
	inSign := in.GetSignature()

	to, err := address.NewFromString(inMsg.GetTo())
	if err != nil {

	}

	from, err := address.NewFromString(inMsg.GetFrom())
	if err != nil {

	}

	signType, err := fil.SignType(inSign.GetSignType())
	if err != nil {

	}
	signMsg := &types.SignedMessage{
		Message: types.Message{
			To:         to,
			From:       from,
			Method:     abi.MethodNum(inMsg.GetMethod()),
			Nonce:      inMsg.GetNonce(),
			Value:      abi.NewTokenAmount(1231243221000010),
			GasLimit:   655063,
			GasFeeCap:  abi.NewTokenAmount(2300),
			GasPremium: abi.NewTokenAmount(2250),
		},
		Signature: crypto.Signature{
			Type: signType,
			Data: inSign.GetData(),
		},
	}

	authToken, ok := env.LookupEnv(env.ENVCOINTOKEN)
	if !ok {

	}
	headers := http.Header{"Authorization": []string{"Bearer " + authToken}}

	addr, ok := env.LookupEnv(env.ENVCOINAPI)
	if !ok {

	}

	var api lotusapi.FullNodeStruct
	closer, err := jsonrpc.NewMergeClient(context.Background(), "ws://"+addr+"/rpc/v0", "Filecoin", []interface{}{&api.Internal, &api.CommonStruct.Internal}, headers)
	if err != nil {
	}
	defer closer()

	cid, err := api.MpoolPush(context.Background(), signMsg)
	if err != nil {
	}

	return &sphinxplugin.MpoolPushResponse{
		Info: &sphinxplugin.MpoolPushInfo{
			CID: cid.String(),
		},
	}, nil
}
