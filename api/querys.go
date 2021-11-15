package api

import (
	"context"

	"github.com/NpoolPlatform/sphinx-plugin/message/npool"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/plugin/fil"
)

var (
	errInternal = status.Error(codes.Internal, "internal server error")
	debugFlag   = false
)

// 获取预签名信息
func (s *Server) GetSignInfo(ctx context.Context, in *npool.GetSignInfoRequest) (ret *npool.SignInfo, err error) {
	resp, err = fil.Server.GetSignInfo(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("getsigninfo error: %w", err)
		resp = &npool.SignInfo{}
		if debugFlag {
			err = errInternal
		}
	}
	return
}

// 余额查询
func (s *Server) GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (ret *npool.AccountBalance, err error) {
	resp, err = fil.Server.GetBalance(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("getbalance error: %w", err)
		resp = &npool.AccountBalance{}
		if debugFlag {
			err = errInternal
		}
	}
	return
}

// 广播交易
func (s *Server) BroadcastScript(ctx context.Context, in *npool.BroadcastScriptRequest) (ret *npool.broadcastScriptResponse, err error) {
	resp, err = fil.Server.BroadcastScript(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("broadcastscript error: %w", err)
		resp = &npool.broadcastScriptResponse{}
		if debugFlag {
			err = errInternal
		}
	}
	return
}

// 交易状态查询
func (s *Server) GetInsiteTxStatus(ctx context.Context, in *npool.GetInsiteTxStatusRequest) (ret *npool.GetInsiteTxStatusResponse, err error) {
	resp, err = fil.Server.GetInsiteTxStatus(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("getinsitetxstatus error: %w", err)
		resp = &npool.GetInsiteTxStatus{}
		if debugFlag {
			err = errInternal
		}
	}
	return
}

// TODO 账户交易查询
func (s *Server) GetTxJSON(ctx context.Context, in *npool.GetTxJSONRequest) (ret *npool.AccountTxJSON, err error) {
	ret = &npool.AccountTxJSON{}
	return
}
