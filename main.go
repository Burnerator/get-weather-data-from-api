package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//Type of the first API response
type firstResponse struct {
	Datos string `json:"datos,omitempty"`
}

//Type of the second API response
//We request all fields but only filter to optain some of them
type secondResponse struct {
	Fecha   string `json:"fecha,omitempty"`
	Nombre  string `json:"nombre,omitempty"`
	PresMin string `json:"presMin,omitempty"`
	PresMax string `json:"presMax,omitempty"`
	TMin    string `json:"tMin,omitempty"`
	TMax    string `json:"tMax,omitempty"`
	TMed    string `json:"tMed,omitempty"`
	Prec    string `json:"prec,omitempty"`
}

//Make http request to web server
func makeRequest(url string, verb string, headers map[string]string) *http.Response {
	req, _ := http.NewRequest(verb, url, nil)

	//Parsing http headers
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	res, _ := http.DefaultClient.Do(req)

	//Printing if http status code is not OK
	if res.StatusCode != 200 {
		fmt.Println(res)
	}

	//Returning http response. Any case.
	return res
}

//Writing a file with a content. We can configure permissions file
func writeFile(content string, pathFile string, permissions os.FileMode) {
	data := []byte(content)
	err := ioutil.WriteFile(pathFile, data, permissions)

	//In case of error
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	//Fetching input params
	if len(os.Args) != 6 {
		fmt.Println("Number of params must be 6")
		os.Exit(1)
	}

	startDate := os.Args[1]
	endDate := os.Args[2]
	station := os.Args[3]
	resFile := os.Args[4]
	apiKey := os.Args[5]

	//Creating url to make GET operation over API (valores climatológicos diarios)
	url := "https://opendata.aemet.es/opendata/api/valores/climatologicos/diarios/datos"
	url = url + "/fechaini/" + startDate + "T00:00:00UTC/fechafin/" + endDate + "T23:59:59UTC/estacion/" + station + "/?api_key=" + apiKey

	//Creating Http headers. They will be used in several request
	headers := map[string]string{
		"cache-control": "no-cache",
	}

	//First request
	var first firstResponse
	_ = json.NewDecoder(makeRequest(url, "GET", headers).Body).Decode(&first)

	//Second request
	var list []secondResponse
	_ = json.NewDecoder(makeRequest(first.Datos, "GET", headers).Body).Decode(&list)

	//Generating results. CSV format.
	results := "Fecha;Nombre;TMin;TMax;PresMin;PresMax;Prec\n"
	for _, item := range list {
		results = results + item.Fecha + ";" + item.Nombre + ";" + item.TMin + ";" + item.TMax + ";" + item.PresMin + ";" + item.PresMax + ";" + item.Prec + "\n"
	}

	//Printing results.
	//Screen
	fmt.Print(results)

	//File
	writeFile(results, resFile, 0644)
}
