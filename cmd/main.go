package main

import (
	"github.com/FateBug403/FoFa/pkg/config"
	"github.com/FateBug403/FoFa/pkg/fofa"
	"log"
)

func main()  {

	FoFaClient,err := fofa.NewFoFa(&config.FoFa{
		Baseurl: "https://fofa.info",
		Email:   "",
		Key:     "",
		Size:    10000,
	})
	if err != nil {
		log.Println(err)
		return
	}

	Result,err:=FoFaClient.SearchAll("icon_hash=\"-1830859634\"")
	if err != nil {
		log.Println(err)
		return
	}
	for _,value:=range Result.InFos{
		log.Println(value.Host)
	}
}
