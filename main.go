package main

import (
	common "andrei.vip/routeview/app"
	"andrei.vip/routeview/node"
)

func main() {
	app := common.NewApp()

	node0 := node.NodeFromIpString(app, "8.8.8.8")
	node1 := node.NodeFromIpString(app, "38.122.35.199")
	node2 := node.NodeFromIpString(app, "131.125.129.85")

	println(node0.String())
	println(node1.String())
	println(node2.String())
}
