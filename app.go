package app

import (
	"github.com/cosmos-sdk/x/auth"
	"github.com/tendermint/libs/log"

	bam "github.com/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/libs/tm-db"
)

const (
	appName = "nameservice"
)

type nameServiceApp struct {
	*bam.BaseApp
}

// NewNameServiceApp 添加nameservice的构造函数
func NewNameServiceApp(logger log.Logger, db dbm.DB) *nameServiceApp {

	// First define the top level codec that will be shared by the different modules. Note: Codec will be explained later
	cdc := MakeCodec()

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc))

	var app = &nameServiceApp{
		BaseApp: bApp,
		cdc:     cdc,
	}

	return app
}
