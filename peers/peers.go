package peers

import (
	"time"
)

// A Tor node that we have connected to
// previously
type Peer struct {
	Address   string
	Port      string
	LastAlive time.Time
}

// Session details for a certain route
type Session struct {
	PublicKey []byte
	Start     time.Time
}

// Different ways a peer can interact
// with other peers.
type Message int

const (
	Announce Message = iota // Announce availability & request peer list
	Route                   // Route a request from the beginning
	Exit                    // Make a request to an external site
	Relay                   // Relay a request to another node
)

// Assume peer goes offline after 30 minutes
// of no contact. If they don't respond immediately,
// also mark offline.
const liveTimeout = time.Minute * 30

// Peers identified by their address
var uidToPeer map[string]Peer

// All encountered peers
var Peers []Peer = []Peer{
	Peer{
		Address:   "",
		Port:      "8080",
		LastAlive: time.Now(),
	},
}

// Returns true if the last contact was
// made less than "liveTimeout" duration ago
func IsLive(peer Peer) bool {
	return time.Since(peer.LastAlive) < liveTimeout
}
