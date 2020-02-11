package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-resty/resty"
)

func main() {

	client := resty.New().
		//SetBasicAuth("bespindemo", "spanner").
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		//SetRedirectPolicy(resty.FlexibleRedirectPolicy(20)).
		SetMode("rest").
		SetHeaders(
			map[string]string{
				"Content-Type":   "application/json",
				"Accept":         "application/json",
				"X-Cmp-Customer": "00000000-0000-0000-0000-000000000000",
				"X-Cmp-User":     "00000000-0000-0000-0000-000000000000",
			})
	//COOLBOX_SERVICE_HOST ip
	//COOLBOX_SERVICE_PORT_HTTP port
	//ip := os.Getenv("FACILITIES_SERVICE_HOST")
	//port := os.Getenv("FACILITIES_SERVICE_PORT_HTTP")
	ip := "192.168.1.1"
	port := 8080
	
	baseURL := fmt.Sprintf("http://%s:%s", ip, port)
	reqURL := baseURL + "/hello"
	resp, err := client.R().Get(reqURL)
	fmt.Println("\nRequest URL:", reqURL)
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	fmt.Println()
	if resp.StatusCode() == 200 && resp.Body() != nil && err == nil {
		fmt.Println("You are authorised.")
	} else {
		fmt.Println("You are Unauthorised.")
	}

/* 	ip = os.Getenv("TRIDENT_SERVICE_HOST")
	port = os.Getenv("TRIDENT_SERVICE_PORT_HTTP")
	baseURL = fmt.Sprintf("http://%s:%s", ip, port)

	reqURL = baseURL + "/cmp/api/users/00000000-0000-0000-0000-000000000000"
	resp, err = client.R().Get(reqURL) */

	var rb resultBody

	err = json.Unmarshal(resp.Body(), &rb)
	if err != nil {
		fmt.Println("\nERROR:", err)
	}
	fmt.Println("\nRequest URL:", reqURL)
	if rb.ErrorCode != nil {
		fmt.Printf("\nErrorCode: %v", *rb.ErrorCode)
		fmt.Printf("\nErrorMessage: %s", *rb.Message)
	} /*else {
		fmt.Printf("\nCustomerName: %v", *rb.CustomerName)
		fmt.Printf("\nRequestorName: %v", *rb.RequestorName)
	}
	*/
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	fmt.Println()
}

type resultBody struct {
	CustomerName  *string `json:"customer_name"`
	RequestorName *string `json:"name"`
	ErrorCode     *int    `json:"error_code"`
	Message       *string `json:"message"`
}
