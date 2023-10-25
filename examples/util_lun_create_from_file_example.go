package main

import (
	"fmt"
	"time"

	"github.com/overag3/go-ontap-rest/ontap"
	"github.com/overag3/go-ontap-rest/util"
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
	lunPath := "/vol/my_test_vol01/os_image_lun"
	filePath := "/vol/my_test_vol01/os_image"
	osType := "linux"
	if err := util.LunCreateFromFile(c, lunPath, filePath, osType); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
