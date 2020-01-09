package api

import (
	"fmt"
	"github.com/SENERGY-Platform/converter/lib/converter/color"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

func ExampleConvertWithGetRequest() {
	server := httptest.NewServer(GetRouter())
	defer server.Close()

	//hex to rgb
	hex := `"#ff00ff"`
	resp, err := http.Get(server.URL + "/conversions/" + url.PathEscape(color.Hex) + "/" + url.PathEscape(color.Rgb) + "?json=" + url.QueryEscape(hex))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	rgbByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(rgbByte))

	//output:
	//{"b":255,"g":0,"r":255}
}

func ExampleConvertWithPostRequest() {
	server := httptest.NewServer(GetRouter())
	defer server.Close()

	//rgb to hex
	rgb := `{"b":255,"g":0,"r":255}`
	resp, err := http.Post(server.URL+"/conversions/"+url.PathEscape(color.Rgb)+"/"+url.PathEscape(color.Hex), "application/json", strings.NewReader(rgb))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	rgbByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(rgbByte))

	//output:
	//"#ff00ff"
}
