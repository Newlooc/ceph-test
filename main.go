package main

import (
	"fmt"
	cephutil "github.com/ceph/ceph-csi/pkg/util"
	rbd "github.com/ceph/go-ceph/rbd"
	"time"
)

func main() {
	cred, _ := cephutil.NewCredentials("kubernetes", "AQD+ZuJeZS+ZDBAA5Ox2st5Hg3OgQaYcEPzgfA==")
	defer cred.DeleteCredentials()

	connPool := cephutil.NewConnPool(2*time.Second, 10*time.Second)
	conn, _ := connPool.Get("kubernetes", "10.6.209.21:6789", cred.ID, cred.KeyFile)

	info, _ := conn.GetClusterStats()

	fmt.Printf("%+v\n", info)

	pools, _ := conn.ListPools()

	for _, pool := range pools {
		ctx, _ := conn.OpenIOContext(pool)
		ims, _ := rbd.GetImageNames(ctx)
		for _, im := range ims {
			fmt.Printf("=====================%s==================\n", im)
			imobj, _ := rbd.OpenImageReadOnly(ctx, im, rbd.NoSnapshot)
			info, err := imobj.Stat()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("%+v\n", info)
			fmt.Println("=======================================")
		}
	}

}
