package types

import (
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc"
)

type TxManager interface {
	TmQuery
	SendBatch(msgs Msgs, baseTx BaseTx) ([]ResultTx, Error)
	BuildAndSend(msg []Msg, baseTx BaseTx) (ResultTx, Error)
	BuildAndSign(msg []Msg, baseTx BaseTx) ([]byte, Error)
	BuildTxHash(msg []Msg, baseTx BaseTx) (string, Error)
	BuildAndSendWithAccount(addr string, accountNumber, sequence uint64, msg []Msg, baseTx BaseTx) (ResultTx, Error)
}

type Queries interface {
	StoreQuery
	AccountQuery
	TmQuery
}

type GRPCClient interface {
	GenConn() (*grpc.ClientConn, error)
}

type ParamQuery interface {
	QueryParams(module string, res Response) Error
}

type StoreQuery interface {
	QueryWithResponse(path string, data interface{}, result Response) error
	Query(path string, data interface{}) ([]byte, error)
	QueryStore(key HexBytes, storeName string, height int64, prove bool) (abci.ResponseQuery, error)
}

type AccountQuery interface {
	QueryAccount(address string) (BaseAccount, Error)
	QueryAddress(name, password string) (AccAddress, Error)
}

type TmQuery interface {
	QueryTx(hash string) (ResultQueryTx, error)
	QueryTxs(builder *EventQueryBuilder, page, size *int) (ResultSearchTxs, error)
	QueryBlock(height int64) (BlockDetail, error)
}

type TokenManager interface {
	QueryToken(denom string) (Token, error)
	SaveTokens(tokens ...Token)
	ToMinCoin(coin ...DecCoin) (Coins, Error)
	ToMainCoin(coin ...Coin) (DecCoins, Error)
}

type Logger interface {
	Logger() log.Logger
	SetLogger(log.Logger)
}

type BaseClient interface {
	TokenManager
	TxManager
	Queries
	TmClient
	Logger
	GRPCClient
	KeyManager
}
