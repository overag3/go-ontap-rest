package main

import (
	"fmt"
	"os"
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
	var file *os.File
	volumeName := "my_test_vol01"
	filePath := "examples/util_upload_file_example.go"
	localFile := "./util_upload_file_example.go"
	file, err := os.Open(localFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	if bytesUploaded, err := util.UploadFileAPI(c, volumeName, filePath, file); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Uploaded %d bytes\n", bytesUploaded)
	}
}
