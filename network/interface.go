package network

import (
	"context"

	bsmsg "github.com/dms3-fs/go-bitswap/message"

	cid "github.com/dms3-fs/go-cid"
	ifconnmgr "github.com/dms3-p2p/go-p2p-interface-connmgr"
	peer "github.com/dms3-p2p/go-p2p-peer"
	protocol "github.com/dms3-p2p/go-p2p-protocol"
)

var (
	// These two are equivalent, legacy
	ProtocolBitswapOne    protocol.ID = "/dms3fs/bitswap/1.0.0"
	ProtocolBitswapNoVers protocol.ID = "/dms3fs/bitswap"

	ProtocolBitswap protocol.ID = "/dms3fs/bitswap/1.1.0"
)

// BitSwapNetwork provides network connectivity for BitSwap sessions
type BitSwapNetwork interface {

	// SendMessage sends a BitSwap message to a peer.
	SendMessage(
		context.Context,
		peer.ID,
		bsmsg.BitSwapMessage) error

	// SetDelegate registers the Reciver to handle messages received from the
	// network.
	SetDelegate(Receiver)

	ConnectTo(context.Context, peer.ID) error

	NewMessageSender(context.Context, peer.ID) (MessageSender, error)

	ConnectionManager() ifconnmgr.ConnManager

	Routing
}

type MessageSender interface {
	SendMsg(context.Context, bsmsg.BitSwapMessage) error
	Close() error
	Reset() error
}

// Implement Receiver to receive messages from the BitSwapNetwork
type Receiver interface {
	ReceiveMessage(
		ctx context.Context,
		sender peer.ID,
		incoming bsmsg.BitSwapMessage)

	ReceiveError(error)

	// Connected/Disconnected warns bitswap about peer connections
	PeerConnected(peer.ID)
	PeerDisconnected(peer.ID)
}

type Routing interface {
	// FindProvidersAsync returns a channel of providers for the given key
	FindProvidersAsync(context.Context, *cid.Cid, int) <-chan peer.ID

	// Provide provides the key to the network
	Provide(context.Context, *cid.Cid) error
}
