package main

import (
	"fmt"
	rados "github.com/ceph/go-ceph/rados"
	rbd "github.com/ceph/go-ceph/rbd"
)

func main() {
	conn, _ := rados.NewConn()
	conn.ReadDefaultConfigFile()
	conn.Connect()

	info, _ := conn.GetClusterStats()

	fmt.Printf("%+v", info)

	pools, _ := conn.ListPools()

	for _, pool := range pools {
		ctx, _ := conn.OpenIOContext(pool)
		ims, _ := rbd.GetImageNames(ctx)
		for _, im := range ims {
			fmt.Printf("=====================%s==================\n", im)
			imobj := rbd.GetImage(ctx, im)
			info, _ := imobj.Stat()
			fmt.Printf("%+s\n", info)
			fmt.Println("=======================================")
		}
	}
}
