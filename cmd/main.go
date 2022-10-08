package main

import (
	"fmt"

	"github.com/zekhoi/learn-go-cli/pkg/services"
	"github.com/zekhoi/learn-go-cli/pkg/utils"
)

func main() {
	data := services.GetAllShorten()

	table := utils.CreateTable(data)
	fmt.Println(table.String())
}
