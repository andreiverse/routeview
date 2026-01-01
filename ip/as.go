package ip

import (
	"log"
	"net"
	"strings"

	"andrei.vip/routeview/app"
)

type Asn struct {
	Net  net.IPNet
	As   string
	Cidr int
}

func GetAsnForIp(app *app.App, ip net.IP) *Asn {

	var asn *Asn = nil

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
			asn = &Asn{
				Net:  *net,
				As:   e[5],
				Cidr: cidr,
			}
		}
	}

	return asn
}
