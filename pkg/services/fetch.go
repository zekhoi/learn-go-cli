package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/zekhoi/learn-go-cli/pkg/entity"
)

func GetAllShorten() []entity.Shorten {

	response, err := http.Get("http://localhost:8080/api/v1/shorten")

	if err != nil {
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	// fmt.Println(string(data))

	var result entity.Response

	if err := json.Unmarshal(data, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	return result.Data

}
