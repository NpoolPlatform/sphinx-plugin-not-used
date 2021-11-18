package api

import (
	"context"
	"log"
	"net/http"

	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/env"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	lotusapi "github.com/filecoin-project/lotus/api"
)

func (s *Server) MpoolGetNonce(ctx context.Context, in *sphinxplugin.MpoolGetNonceRequest) (*sphinxplugin.MpoolGetNonceResponse, error) {
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
		log.Fatalf("connecting with lotus failed: %s", err)
	}
	defer closer()

	from, err := address.NewFromString(in.GetAddress())
	if err != nil {

	}
	nonce, err := api.MpoolGetNonce(context.Background(), from)
	if err != nil {

	}
	return &sphinxplugin.MpoolGetNonceResponse{
		Info: &sphinxplugin.MpoolGetNonceInfo{
			Nonce: nonce,
		},
	}, nil
}
