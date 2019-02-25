package main

import (
	uiza "github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/category"
	_ "github.com/uizaio/api-wrapper-go/testing"
	"log"
)

func main() {

	var typeCategory = uiza.CategoryFolderType
	params := &uiza.CategoryUpdateParams{
		ID:          uiza.String("Your category ID"),
		Name:        uiza.String(""),
		Type:        &typeCategory,
		Description: uiza.String(""),
		Icon:        uiza.String(""),
		OrderNumber: uiza.Int64(2)}
	response, _ := category.Update(params)
	log.Printf("%s", response)

}
