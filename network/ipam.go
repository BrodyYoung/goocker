package network

import "net"

//ip地址管理
type IPAM struct {
}

var ipAlloc = &IPAM{}

func (i IPAM) alloc() (net.IP, error) {

	return nil,nil
}
