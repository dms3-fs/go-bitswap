package bitswap

import (
	bsnet "github.com/dms3-fs/go-bitswap/network"
	peer "github.com/dms3-p2p/go-p2p-peer"
	"github.com/dms3-p2p/go-testutil"
)

type Network interface {
	Adapter(testutil.Identity) bsnet.BitSwapNetwork

	HasPeer(peer.ID) bool
}
