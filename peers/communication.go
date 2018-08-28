package peers

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Subdomains targets for messaging
var Routes = map[Message]string{
	Announce: "/announce",
	Route:    "/route",
	Relay:    "/relay",
	Exit:     "/exit",
}

// Check if a peer is still online by sending
// a request to it.
func PingPeer(peer Peer) bool {
	client := http.Client{
		Timeout: time.Second * 5,
	}
	resp, err := client.Get(peer.Address + ":" + peer.Port)
	return err == nil && resp.StatusCode != 200
}

// Announce our availability to a peer. Also
// asks for their list of known peers.
func AnnounceTo(peer Peer) ([]Peer, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	uri := Routes[Announce]
	_, err := client.Get(peer.Address + ":" + uri)
	if err != nil {
		return nil, err
	}
	return []Peer{}, nil
}

// Parse the response from an announce request
func parseAnnounce(req *http.Request) ([]Peer, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read announce response body")
	}
	var addrs []NodeAddr
	if err := json.Unmarshal(body, &addrs); err != nil {
		return nil, errors.Wrap(err, "Failed to parse announce response body")
	}
	var peers []Peer
	for _, addr := range addrs {
		peers = append(peers, Peer{
			Address:   addr.Address,
			Port:      addr.Port,
			LastAlive: time.Now(),
		})
	}
	return peers, nil
}
