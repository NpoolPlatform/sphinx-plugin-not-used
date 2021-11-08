package fil

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/cyvadra/filecoin-client"
	"github.com/cyvadra/filecoin-client/local"
	"github.com/cyvadra/filecoin-client/types"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/shopspring/decimal"
)

// todo
func GetSignInfoRaw(addr string) (json string, err error) { return }
func GetTxStatus(CID string) (json string, err error)     { return }
func GetTxJSON(addr string) (json string, err error)      { return }

func GetBalanceRaw(addr string) (str string, err error) {
	addrStd, err := address.NewFromString(addr)
	if err != nil {
		return
	}
	bal, err := Client.WalletBalance(context.Background(), addrStd)
	str = bal.String()
	return
}

func BroadcastScriptRaw(s *types.SignedMessage) (err error) {
	// 消息广播
	mid, err := Client.MpoolPush(context.Background(), s)
	if err != nil {
		fmt.Println("消息广播失败")
		fmt.Println(err)
	} else {
		fmt.Println("消息发送成功，message id:")
		fmt.Println(mid.String())
	}
	return
}