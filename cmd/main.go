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

	Result:=FoFaClient.SearchAllS([]string{"domain=\"enzyun.com\"","domain=\"wesvr.cn\"","domain=\"wesvr.com\""})
	log.Println(Result.GetHosts())
}
