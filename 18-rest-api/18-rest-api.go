package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

func main() {

	response, err := http.Get("http://httpbin.org/get")
	if err != nil {
		fmt.Printf("The Http request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Printf(string(data))
	}

	fmt.Println("********** End Of Get Api Response **********")

	//jsonData := map[string]string{"firstname" : "nic","lastname" :"http://httpbin.org/post"}
	//jsonValue, _ :=json.Marshal(jsonData)
	//response, err = http.Post("http://httpbin.org/post","application/json",bytes.NewBuffer(jsonValue))
	//if err != nil {
	//	fmt.Printf("The Http request failed with error %s\n",err)
	//}else{
	//	data, _ := ioutil.ReadAll(response.Body)
	//	fmt.Printf(string(data))
	//}
	//
	//fmt.Println("********** End Of Post Api Response **********")

	/*
	Another way to do POST request with header
	 */

	jsonData := map[string]string{"firstname": "nic", "lastname": "http://httpbin.org/post"}
	jsonValue, _ := json.Marshal(jsonData)

	/*
	starting of request object
	*/

	request, _ := http.NewRequest("POST", "http://httpbin.org/post", bytes.NewBuffer(jsonValue))
	request.Header.Set("Context-Type", "application/json")
	client := &http.Client{}

	/*
	ending of request object then we send it with the client
	 */

	response, err = client.Do(request)

	if err != nil {
		fmt.Printf("The Http request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Printf(string(data))
	}

	fmt.Println("********** End Of Post Api With Header Response **********")

}
