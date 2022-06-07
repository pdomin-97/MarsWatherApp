package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Solar_day struct {
	ID       string `json:"SQL"`
	Sol_keys string `json:"sol_keys"`
	Validity string `json:"validity_checks"`
}

func main() {

	api_key := "MNSmJzA4KeJCIuSOeLsvHncRMNP9ByepOzdSce1C"
	url := "https://api.nasa.gov/insight_weather/?api_key=" + api_key + "&feedtype=json&ver=1.0"

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

}
