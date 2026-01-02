package ip

import (
	"log"
	"net"
	"strconv"
	"strings"

	"andrei.vip/routeview/app"
)

type AS struct {
	Net  net.IPNet
	As   string
	Cidr int
}

func (a AS) String() string {
	return a.Net.IP.String() + "/" + strconv.Itoa(a.Cidr) + "[" + a.As + "]"
}

func GetAsnForIp(app *app.App, ip net.IP) AS {
	var asn *AS = nil

	for _, e := range app.Asns {
		network := e[0]
		if !strings.Contains(network, "/") {
			network += "/32"
		}

		_, net, err := net.ParseCIDR(network)

		if err != nil {
			continue
		}

		if !net.Contains(ip) {
			continue
		}

		cidr, _ := net.Mask.Size()
		log.Println("found prefix for ip " + ip.String() + ": " + net.String() + " " + strings.Join(e, ", "))

		if asn == nil || asn.Cidr < cidr {
			asn = &AS{
				Net:  *net,
				As:   e[5],
				Cidr: cidr,
			}
		}
	}

	return *asn
}
