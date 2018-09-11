package bitswap

import (
	"context"

	bsnet "github.com/dms3-fs/go-bitswap/network"

	ds "github.com/dms3-fs/go-datastore"
	mockrouting "github.com/dms3-fs/go-fs-routing/mock"
	peer "github.com/dms3-p2p/go-p2p-peer"
	mockpeernet "github.com/dms3-p2p/go-p2p/p2p/net/mock"
	testutil "github.com/dms3-p2p/go-testutil"
)

type peernet struct {
	mockpeernet.Mocknet
	routingserver mockrouting.Server
}

func StreamNet(ctx context.Context, net mockpeernet.Mocknet, rs mockrouting.Server) (Network, error) {
	return &peernet{net, rs}, nil
}

func (pn *peernet) Adapter(p testutil.Identity) bsnet.BitSwapNetwork {
	client, err := pn.Mocknet.AddPeer(p.PrivateKey(), p.Address())
	if err != nil {
		panic(err.Error())
	}
	routing := pn.routingserver.ClientWithDatastore(context.TODO(), p, ds.NewMapDatastore())
	return bsnet.NewFromDms3FsHost(client, routing)
}

func (pn *peernet) HasPeer(p peer.ID) bool {
	for _, member := range pn.Mocknet.Peers() {
		if p == member {
			return true
		}
	}
	return false
}

var _ Network = (*peernet)(nil)
