package fil

import (
	"context"
	"encoding/json"
	"time"

	"github.com/NpoolPlatform/sphinx-plugin/message/npool"
	"github.com/cyvadra/filecoin-client/types"
)

type Server npool.UnimplementedPluginServer

func (Server) GetSignInfo(ctx context.Context, in *npool.GetSignInfoRequest) (sio *npool.SignInfo, err error) {
	sioFIL, err := GetSignInfo(in.Address)
	if err != nil {
		return
	}
	js, err := json.Marshal(sioFIL)
	sio = &npool.SignInfo{
		Json: string(js),
	}
	return
}

func (Server) GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (acb *npool.AccountBalance, err error) {
	acbStr, err := GetBalance(in.Address)
	if err != nil {
		return
	}
	amountInt, amountDigits, amountString := DecomposeStringInt(acbStr)
	acb = &npool.AccountBalance{
		CoinId:       0,
		Address:      in.Address,
		TimestampUtc: time.Now().UnixNano(),
		AmountInt:    amountInt,
		AmountDigits: amountDigits,
		AmountString: amountString,
	}
	return
}

func (Server) BroadcastScript(ctx context.Context, in *npool.BroadcastScriptRequest) (resp *npool.BroadcastScriptResponse, err error) {
	msg := &types.SignedMessage{}
	err = json.Unmarshal([]byte(in.TransactionScript), msg)
	if err != nil {
		return
	}
	cid, err := BroadcastScript(msg)
	resp = &npool.BroadcastScriptResponse{
		TransactionIdChain: cid,
	}
	return
}

// default true
func (Server) GetTxStatus(ctx context.Context, in *npool.GetTxStatusRequest) (resp *npool.GetTxStatusResponse, err error) {
	msg, err := GetTxStatus(in.TransactionIdChain)
	if err != nil {
		return
	}
	amountInt, amountDigits, amountString := DecomposeStringInt(msg.Value.String())
	resp = &npool.GetTxStatusResponse{
		AmountInt:          amountInt,
		AmountDigits:       amountDigits,
		AmountString:       amountString,
		AddressFrom:        msg.From.String(),
		AddressTo:          msg.To.String(),
		TransactionIdChain: in.TransactionIdChain,
		CreatetimeUtc:      0,
		UpdatetimeUtc:      0,
		IsSuccess:          true,
		IsFailed:           false,
		NumBlocksConfirmed: -1,
	}
	return
}
