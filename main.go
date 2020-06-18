package main

import (
	"fmt"
	cephutil "github.com/ceph/ceph-csi/pkg/util"
	rbd "github.com/ceph/go-ceph/rbd"
	"time"
)

func main() {
	cred, err := cephutil.NewCredentials("admin", "AQB48uBebNTYLhAA+PIZ8M0yqbyBK4cdSK6szA==")
	if err != nil {
		panic(err.Error())
	}
	defer cred.DeleteCredentials()

	connPool := cephutil.NewConnPool(2*time.Second, 10*time.Second)
	conn, err := connPool.Get("kubernetes", "10.6.209.21:6789", cred.ID, cred.KeyFile)
	if err != nil {
		panic(err.Error())
	}

	// Create ceph.conf for use with CLI commands
	if err = cephutil.WriteCephConfig(); err != nil {
		panic(err.Error())
	}

	info, err := conn.GetClusterStats()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", info)

	pools, err := conn.ListPools()
	if err != nil {
		panic(err.Error())
	}

	for _, pool := range pools {
		ctx, err := conn.OpenIOContext(pool)
		if err != nil {
			panic(err.Error())
		}
		ims, err := rbd.GetImageNames(ctx)
		if err != nil {
			panic(err.Error())
		}
		for _, im := range ims {
			fmt.Printf("=====================%s==================\n", im)
			imobj, err := rbd.OpenImageReadOnly(ctx, im, rbd.NoSnapshot)
			if err != nil {
				panic(err.Error())
			}
			info, err := imobj.Stat()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("%+v\n", info)
			fmt.Println("=======================================")
		}
	}
}
