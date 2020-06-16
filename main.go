package main

import (
	"fmt"
	rados "github.com/ceph/go-ceph/rados"
	//rbd "github.com/ceph/go-ceph/rbd"
)

func main() {
	conn, _ := rados.NewConn()
	conn.ReadDefaultConfigFile()
	conn.Connect()

	info, _ := conn.GetClusterStats()

	fmt.Printf("%+v", info)
}
