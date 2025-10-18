package opnsense

import (
	"log"

	"github.com/dreams-money/opnsense-failover/config"
)

func SimpleCall(cfg config.Config) error {
	var peer config.Peer
	for _, peer = range cfg.Peers {
		break
	}
	_, err := getWireguardPeer(peer.OpnSenseWireguardPeerID, cfg)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
