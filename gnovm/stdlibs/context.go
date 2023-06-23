package stdlibs

import (
	"github.com/gnolang/gno/tm2/pkg/crypto"
	"github.com/gnolang/gno/tm2/pkg/sdk"
	"github.com/gnolang/gno/tm2/pkg/std"
)

type ExecContext struct {
	ChainID       string
	Height        int64
	Timestamp     int64 // seconds
	TimestampNano int64 // nanoseconds, only used for testing.
	Msg           sdk.Msg
	OrigCaller    crypto.Bech32Address
	OrigPkgAddr   crypto.Bech32Address
	OrigSend      std.Coins
	OrigSendSpent *std.Coins // mutable
	Banker        Banker
}

func (e ExecContext) GetTimestamp() int64     { return e.Timestamp }
func (e ExecContext) GetTimestampNano() int64 { return e.TimestampNano }
func (e ExecContext) GetChainID() string      { return e.ChainID }
func (e ExecContext) GetHeight() int64        { return e.Height }