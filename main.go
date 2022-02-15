package main

import (
	"encoding/json"
	"log"

	"github.com/yusufpapurcu/wmi"
)

type SoftwareLicensingProduct struct {
	Name                   string `json:"Name"`
	RemainingAppReArmCount int    `json:"RemainingAppReArmCount"`
	GracePeriodRemaining   int    `json:"GracePeriodRemaining"`
}

func main() {
	var dst []SoftwareLicensingProduct
	q := wmi.CreateQuery(&dst, "")
	err := wmi.Query(q, &dst)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range dst {
		data, _ := json.Marshal(v)
		println(string(data))
	}
}
