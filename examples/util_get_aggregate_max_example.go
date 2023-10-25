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
	if aggregate, err := util.GetAggregateMax(c); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Aggregate: %s\n", aggregate)
	}
}
