package node

import (
	"fmt"
	"net"
	"strings"

	"andrei.vip/routeview/app"
	"andrei.vip/routeview/ip"
)

type Node struct {
	AS      ip.AS
	IP      net.IP
	IPStr   string
	Domains []string
}

// String implements the fmt.Stringer interface for pretty printing
func (node *Node) String() string {
	res := fmt.Sprintf("[%s] %s", node.AS.String(), node.IPStr)
	if len(node.Domains) > 0 {
		res += fmt.Sprintf(" (%s)", strings.Join(node.Domains, ", "))
	}
	return res
}

// NodeFromIpString fetches ASN data and performs reverse DNS lookup
func NodeFromIpString(app *app.App, ipStr string) *Node {
	parsedIp := net.ParseIP(ipStr)

	// Handle invalid IP cases gracefully
	if parsedIp == nil {
		return nil
	}

	// Reverse DNS lookup
	domains, err := net.LookupAddr(ipStr)
	if err != nil {
		if !strings.Contains(err.Error(), "no such host") {
			println("[error] ", err.Error())
		}

		domains = make([]string, 0)
	}

	as := ip.GetAsnForIp(app, parsedIp)

	return &Node{
		AS:      as,
		IP:      parsedIp,
		Domains: domains,
		IPStr:   ipStr,
	}
}
