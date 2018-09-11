package decision

import (
	"fmt"
	"math"
	"testing"

	"github.com/dms3-fs/go-bitswap/wantlist"
	cid "github.com/dms3-fs/go-cid"
	u "github.com/dms3-fs/go-fs-util"
	"github.com/dms3-p2p/go-p2p-peer"
	"github.com/dms3-p2p/go-testutil"
)

// FWIW: At the time of this commit, including a timestamp in task increases
// time cost of Push by 3%.
func BenchmarkTaskQueuePush(b *testing.B) {
	q := newPRQ()
	peers := []peer.ID{
		testutil.RandPeerIDFatal(b),
		testutil.RandPeerIDFatal(b),
		testutil.RandPeerIDFatal(b),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c := cid.NewCidV0(u.Hash([]byte(fmt.Sprint(i))))

		q.Push(&wantlist.Entry{Cid: c, Priority: math.MaxInt32}, peers[i%len(peers)])
	}
}
