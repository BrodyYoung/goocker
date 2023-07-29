package network

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
	"goocker/common"
	"io/fs"
	"net"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	drivers  = map[string]NetworkDriver{}
	networks = map[string]*Network{}
)

type Network struct {
	Name    string
	IpRange *net.IPNet
	Driver  string
}

type Endpoint struct {
	Id          string           `json:"id"`
	Device      netlink.Veth     `json:"dev"`
	IpAddr      net.IPAddr       `json:"ip"`
	MacAddr     net.HardwareAddr `json:"mac"`
	network     *Network
	portMapping []string
}

type NetworkDriver interface {
	Name() string
	Create(subnet, name string) (Network, err)
	Delete(nw Network)
	Connect(nw Network, ep Endpoint)
	Disconnect(nw Network, ep Endpoint)
}

//转储文件
func (nw *Network) dump(dumpPath string) error {
	if _, err := os.Stat(dumpPath); err != nil && os.IsNotExist(err) {
		os.Mkdir(dumpPath, os.ModePerm)
	}
	pa := path.Join(dumpPath, nw.Name)
	nwFile, err := os.OpenFile(pa, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		logrus.Error(err)
		return err
	}
	defer nwFile.Close()
	mar, err := json.Marshal(nw)
	_, err = nwFile.Write(mar)
	return err
}

//加载文件
func (nw *Network) load(dumpPath string) error {
	_, err := os.ReadFile(dumpPath)
	return err

}

//删除文件
func (nw *Network) remove(dumpPath string) error {
	pa := path.Join(dumpPath, nw.Name)
	if _, err := os.Stat(pa); err != nil && os.IsNotExist(err) {
		logrus.Error(err)
		return err
	}
	return os.Remove(pa)
}

func Init() error {
	var bridgeDriver BridgeDriver
	drivers[bridgeDriver.Name()] = &BridgeDriver{}

	if _, err := os.Stat(common.DefaultNetworkPath); err != nil {
		os.Mkdir(common.DefaultNetworkPath, os.ModePerm)

	}

	err := filepath.Walk(common.DefaultNetworkPath, func(pa string, info fs.FileInfo, err error) error {

		if strings.HasSuffix(pa, "/") {
			return nil
		}
		_, nwName := path.Split(pa)

		nw := &Network{Name: nwName}

		if err := nw.load(pa); err != nil {
			return err
		}

		networks[nw.Name] = nw
	})
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Infof("network", networks)
	return nil

}

func CreateNetwork(driver, subnet, name string) error {
	//CIDR（Classless Inter-Domain Routing，无类域间路由选择）
	//是一个用于给用户分配IP地址以及在互联网上有效地路由IP数据包的对IP地址进行归类的方法
	_, ipNet, err := net.ParseCIDR(subnet)
	if err != nil {
		logrus.Error(err)
		return err
	}

	ip, err := ipAlloc.alloc()
	if err != nil {
		logrus.Error(err)
		return err
	}
	ipNet.IP = ip

	nw, err := drivers[driver].Create(subnet, name)
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = nw.dump(common.DefaultNetworkPath)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
