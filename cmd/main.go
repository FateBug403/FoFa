package main

import (
	"github.com/FateBug403/FoFa/pkg/fofa"
	"log"
)

func main()  {

	FoFaClient,err := fofa.NewFoFa(&fofa.Options{
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
	log.Println(Result.GetHosts())
}
