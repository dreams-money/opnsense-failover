package opnsense

import (
	"github.com/dreams-money/opnsense-failover/config"
	"github.com/dreams-money/opnsense-failover/etcd"
)

func Failover(cfg config.Config) error {
	leader, err := etcd.GetLeaderName(cfg)
	if err != nil {
		return err
	}

	if leader == cfg.NodeName {
		return isPrimary(cfg)
	}

	return isReplica(leader, cfg)
}

func isPrimary(cfg config.Config) error {
	return removeVIPFromWireguardPeers(cfg)
}

func isReplica(leader string, cfg config.Config) error {
	err := removeVIPFromWireguardPeers(cfg)
	if err != nil {
		return err
	}

	return addVIPToWireguardPeer(leader, cfg)
}
