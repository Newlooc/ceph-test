package main

import (
	cephutil "github.com/ceph/ceph-csi/pkg/util"
	//rbd "github.com/ceph/go-ceph/rbd"
	"time"
)

func main() {
	cred, _ := cephutil.NewCredentials("kubernetes", "AQD+ZuJeZS+ZDBAA5Ox2st5Hg3OgQaYcEPzgfA==")
	defer cred.DeleteCredentials()

	connPool := cephutil.NewConnPool(2*time.Second, 10*time.Second)
	conn, _ := connPool.Get("kubernetes", "10.6.209.21:6789", cred.ID, cred.KeyFile)

}
