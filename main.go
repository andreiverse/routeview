package main

import (
	"log"
	"net"
	"strconv"

	common "andrei.vip/routeview/app"
	"andrei.vip/routeview/ip"
	"andrei.vip/routeview/peer"
)

func main() {
	app := common.NewApp()

	asn := ip.GetAsnForIp(app, net.ParseIP("38.122.35.199"))
	asn2 := ip.GetAsnForIp(app, net.ParseIP("131.125.129.85"))
	if asn != nil {
		log.Println("found asn: " + asn.As)
	}
	if asn2 != nil {
		log.Println("found asn2: " + asn2.As)
	}

	shared, err := peer.GetSharedFacilities(asn.As, asn2.As)

	if err != nil {
		panic(err)
	}

	for _, sfac := range shared {
		println("%s", strconv.Itoa(sfac))
	}
}
