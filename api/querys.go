package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/sphinx-plugin/message/npool"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/plugin/fil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	errInternal = status.Error(codes.Internal, "internal server error")
	debugFlag   = false
)

// 获取预签名信息
func (s *Server) GetSignInfo(ctx context.Context, in *npool.GetSignInfoRequest) (resp *npool.SignInfo, err error) {
	resp, err = fil.PluginServer.GetSignInfo(ctx, in)
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
func (s *Server) GetBalance(ctx context.Context, in *npool.GetBalanceRequest) (resp *npool.AccountBalance, err error) {
	resp, err = fil.PluginServer.GetBalance(ctx, in)
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
func (s *Server) BroadcastScript(ctx context.Context, in *npool.BroadcastScriptRequest) (resp *npool.BroadcastScriptResponse, err error) {
	resp, err = fil.PluginServer.BroadcastScript(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("broadcastscript error: %w", err)
		resp = &npool.BroadcastScriptResponse{}
		if debugFlag {
			err = errInternal
		}
	}
	return
}

// 交易状态查询
func (s *Server) GetTxStatus(ctx context.Context, in *npool.GetTxStatusRequest) (resp *npool.GetTxStatusResponse, err error) {
	resp, err = fil.PluginServer.GetTxStatus(ctx, in)
	if err != nil {
		logger.Sugar().Errorw("getinsitetxstatus error: %w", err)
		resp = &npool.GetTxStatusResponse{}
		if debugFlag {
			err = errInternal
		}
	}
	return
}

// TODO 账户交易查询
func (s *Server) GetTxJSON(ctx context.Context, in *npool.GetTxJSONRequest) (resp *npool.AccountTxJSON, err error) {
	resp = &npool.AccountTxJSON{}
	return
}
