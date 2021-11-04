package core

import (
	"context"

	"github.com/NpoolPlatform/sphinx-service/message/npool"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/coininfo"
)

// 查询全部币种
func GetCoinInfos(ctx context.Context, req *npool.GetCoinInfosRequest) (cilist *npool.CoinInfoList, err error) {
	entResp, err := Client.CoinInfo.Query().All(ctx)
	tmp_cir := make([]*npool.CoinInfoRow, len(entResp))
	for i, row := range entResp {
		tmp_cir[i] = &npool.CoinInfoRow{
			Id:           row.ID,
			Name:         row.Name,
			Unit:         row.Unit,
			NeedSigninfo: row.NeedSigninfo,
		}
	}
	cilist = &npool.CoinInfoList{
		List: tmp_cir,
	}
	return
}

// 查询单个币种
func GetCoinInfo(ctx context.Context, req *npool.GetCoinInfoRequest) (coin_info *npool.CoinInfoRow, err error) {
	entResp, err := Client.CoinInfo.
		Query().
		Where(
			coininfo.ID(req.CoinId),
		).First(ctx)
	coin_info = &npool.CoinInfoRow{
		Id:           entResp.ID,
		NeedSigninfo: entResp.NeedSigninfo,
		Name:         entResp.Name,
		Unit:         entResp.Unit,
	}
	return
}
