package main

import (
	"github.com/FateBug403/FoFa/pkg/fofa"
	"log"
)

func main()  {

	FoFaClient,err := fofa.NewFoFa(&fofa.Options{
		Baseurl: "",
		Key:     "",
		Size:    10000,
	})
	if err != nil {
		log.Println(err)
		return
	}

	Result:=FoFaClient.SearchAllS([]string{"domain=\"enzyun.com\""})
	log.Println(Result.GetLinks())
}
