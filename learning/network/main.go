package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	DefaultTimeout     = 30 * time.Second
	DefaultConnections = 10000

	StoreFrontApi = "http://172.26.158.201:17050/api/v1/cart/basket/items"
)

var (
	// DefaultLocalAddr is the default local IP address an Attacker uses.
	DefaultLocalAddr = net.IPAddr{IP: net.IPv4zero}

	DefaultComputerAddr = net.IPAddr{IP: net.IPv4(172, 26, 158, 167)}
	// DefaultTLSConfig is the default tls.Config an Attacker uses.
	DefaultTLSConfig = &tls.Config{InsecureSkipVerify: true}
)

func main() {
	dialer := &net.Dialer{
		LocalAddr: &net.TCPAddr{IP: DefaultLocalAddr.IP, Zone: DefaultLocalAddr.Zone},
		KeepAlive: 30 * time.Second,
		Timeout:   DefaultTimeout,
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			Dial:  dialer.Dial,
			ResponseHeaderTimeout: DefaultTimeout,
			TLSClientConfig:       DefaultTLSConfig,
			TLSHandshakeTimeout:   10 * time.Second,
			MaxIdleConnsPerHost:   DefaultConnections,
		},
	}

	// local machine : http://172.26.158.167:17003/hue-ec-product/v1/products/categories/stores/1

	resp1, err := http.Get(StoreFrontApi)
	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", resp1)
	fmt.Println(resp1.StatusCode)
	bodyBytes, err2 := ioutil.ReadAll(resp1.Body)
	fmt.Println(err2)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	fmt.Println(resp1.Body)

	request, err := http.NewRequest("GET", "http://172.26.158.201:17050/api/v1/cart/basket/items", nil)

	fmt.Println(err)

	request.Header.Add("Accept-Charset", "UTF-8;q=1, ISO-8859-1;q=0")
	request.Header.Add("client-id", "tenant001")
	request.Header.Add("access-token", "eyJhbGciOiJSUzI1NiJ9.eyJjbGllbnRJZCI6InRlbmFudDAwMSIsImV4cGlyYXRpb25EYXRlIjoiMjAxNy0wOS0yMFQxMDoyMjowMy41MTdaW0V0Yy9VVENdIiwiZW5kcG9pbnRzIjpbIipAKiJdfQ.crG3ZMzv7RrmIyzmTmOC1QC3ZH3dX7NDqRs02_FjXSEhWHTacayLBtZWHX8Rk5WJPNpTKOQf6H2PaoA-GNEDboUvGqP_vOBD-14FOsC5vPdb7Td2eGJnhOcqAhz6JBB9lAkQB449piJUO4XNlootqU5nwqLTfCEZYUdfbMYAUqQ")
	request.Header.Add("storeId", "eccefec0-87d2-11e7-a3e1-4d3a169b023e")
	request.Header.Add("siteId", "31626270-87d3-11e7-a3e1-4d3a169b023e")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")

	checkError(err)

	response, err := client.Do(request)
	if response.Status != "200 OK" || err != nil {
		fmt.Println(response.Status)
		fmt.Println(err)
		//os.Exit(2)
	}

	chSet := getCharset(response)
	fmt.Printf("got charset %s\n", chSet)
	if chSet != "UTF-8" {
		fmt.Println("Cannot handle", chSet)
		//os.Exit(4)
	}

	fmt.Println("got body")
	fmt.Println(response)
	bodyBytes, err2 = ioutil.ReadAll(response.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	bodyString = string(bodyBytes)
	fmt.Println(bodyString)

}

func DNSLookUp(name string) {
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(addr.IP.String())
}

func getCharset(response *http.Response) string {
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		// guess
		return "UTF-8"
	}
	idx := strings.Index(contentType, "charset:")
	if idx == -1 {
		// guess
		return "UTF-8"
	}
	return strings.Trim(contentType[idx:], " ")
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
