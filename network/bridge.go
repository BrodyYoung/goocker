package network

import (
	"github.com/sirupsen/logrus"
	"net"
)

//桥接网络模式
type BridgeNetworkDriver struct {
}

func (b *BridgeNetworkDriver) Name() string {
	return "bridge"
}

func (b *BridgeNetworkDriver) Create(subnet, name string) (*Network, error) {

	ip, ipRange, err := net.ParseCIDR(subnet)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	ipRange.IP = ip
	n := &Network{
		Name:    name,
		IpRange: ipRange,
		Driver:  b.Name(),
	}
	err = b.initBridge(n)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return n, nil
}

func (b *BridgeNetworkDriver) Delete(nw Network) {

}

func (b *BridgeNetworkDriver) Connect(nw Network, ep Endpoint) {

}

func (b *BridgeNetworkDriver) Disconnect(nw Network, ep Endpoint) {

}
