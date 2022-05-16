/*
 * Copyright 2021 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/SENERGY-Platform/converter/lib/converter"
	"github.com/SENERGY-Platform/converter/lib/converter/characteristics"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

func ExampleConvertWithGetRequest() {
	c, err := converter.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	server := httptest.NewServer(GetRouter(c))
	defer server.Close()

	//hex to rgb
	hex := `"#ff00ff"`
	resp, err := http.Get(server.URL + "/conversions/" + url.PathEscape(characteristics.Hex) + "/" + url.PathEscape(characteristics.Rgb) + "?json=" + url.QueryEscape(hex))
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
	c, err := converter.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	server := httptest.NewServer(GetRouter(c))
	defer server.Close()

	//rgb to hex
	rgb := `{"b":255,"g":0,"r":255}`
	resp, err := http.Post(server.URL+"/conversions/"+url.PathEscape(characteristics.Rgb)+"/"+url.PathEscape(characteristics.Hex), "application/json", strings.NewReader(rgb))
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

func ExamplePureExtension() {
	c, err := converter.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	server := httptest.NewServer(GetRouter(c))
	defer server.Close()

	buf := bytes.NewBuffer(nil)
	json.NewEncoder(buf).Encode(map[string]interface{}{
		"input": 13,
		"extensions": []map[string]interface{}{
			{
				"from":             "temp",
				"to":               "bar",
				"distance":         -1,
				"formula":          "val - 10",
				"placeholder_name": "val",
			},
			{
				"from":             "foo",
				"to":               "temp",
				"distance":         -1,
				"formula":          "4*x",
				"placeholder_name": "x",
			},
		},
	})
	resp, err := http.Post(server.URL+"/extended-conversions/foo/bar", "application/json", buf)
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
	//42
}

func ExampleMixedExtension() {
	c, err := converter.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	server := httptest.NewServer(GetRouter(c))
	defer server.Close()

	buf := bytes.NewBuffer(nil)
	json.NewEncoder(buf).Encode(map[string]interface{}{
		"input": 13,
		"extensions": []map[string]interface{}{
			{
				"from":             characteristics.Kelvin,
				"to":               "bar",
				"distance":         -1,
				"formula":          "val - 10",
				"placeholder_name": "val",
			},
			{
				"from":             "foo",
				"to":               characteristics.Celsius,
				"distance":         -1,
				"formula":          "4*x",
				"placeholder_name": "x",
			},
		},
	})
	resp, err := http.Post(server.URL+"/extended-conversions/foo/bar", "application/json", buf)
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
	//315.15
}
