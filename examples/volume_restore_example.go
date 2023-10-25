package main

import (
	"fmt"
	"time"

	"github.com/overag3/go-ontap-rest/ontap"
)

func main() {
	c := ontap.NewClient(
		"https://mytestsvm.example.com",
		&ontap.ClientOptions{
			BasicAuthUser:     "vsadmin",
			BasicAuthPassword: "secret",
			SSLVerify:         false,
			Debug:             true,
			Timeout:           60 * time.Second,
		},
	)
	var parameters []string
	parameters = []string{"name=my_test_vol01"}
	volumes, _, err := c.VolumeGetIter(parameters)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(volumes) > 0 {
		parameters = []string{"name=my_snapshot_test01"}
		if snapshots, _, err := c.SnapshotGetIter(volumes[0].Uuid, parameters); err != nil {
			fmt.Println(err)
		} else {
			if len(snapshots) > 0 {
				parameters = []string{"restore_to.snapshot.uuid=" + snapshots[0].Uuid}
				volume := ontap.Volume{}
				_, err := c.VolumeModify(volumes[0].GetRef(), &volume, parameters)
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("success")
				}
			} else {
				fmt.Println("no snapshot found")
			}
		}
	} else {
		fmt.Println("no volume found")
	}
}
