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
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/SENERGY-Platform/converter/lib/converter"
	"github.com/SENERGY-Platform/converter/lib/converter/characteristics"
)

func TestConvertWithGetRequest(t *testing.T) {
	c, err := converter.New()
	if err != nil {
		t.Error(err)
		return
	}
	server := httptest.NewServer(GetRouter(c))
	defer server.Close()

	//hex to rgb
	hex := `"#ff00ff"`
	resp, err := http.Get(server.URL + "/conversions/" + url.PathEscape(characteristics.Hex) + "/" + url.PathEscape(characteristics.Rgb) + "?json=" + url.QueryEscape(hex))
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	rgbByte, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	if strings.TrimSpace(fmt.Sprint(string(rgbByte))) != `{"b":255,"g":0,"r":255}` {
		t.Error(string(rgbByte))
	}
}

func TestConvertWithPostRequest(t *testing.T) {
	c, err := converter.New()
	if err != nil {
		t.Error(err)
		return
	}
	server := httptest.NewServer(GetRouter(c))
	defer server.Close()

	//rgb to hex
	rgb := `{"b":255,"g":0,"r":255}`
	resp, err := http.Post(server.URL+"/conversions/"+url.PathEscape(characteristics.Rgb)+"/"+url.PathEscape(characteristics.Hex), "application/json", strings.NewReader(rgb))
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	rgbByte, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	if strings.TrimSpace(string(rgbByte)) != `"#ff00ff"` {
		t.Error(string(rgbByte))
	}
}

func TestPureExtension(t *testing.T) {
	c, err := converter.New()
	if err != nil {
		t.Error(err)
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
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	temp, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if strings.TrimSpace(string(temp)) != "42" {
		t.Error(string(temp))
	}
}

func TestMixedExtension(t *testing.T) {
	c, err := converter.New()
	if err != nil {
		t.Error(err)
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
		t.Error(err)
		return
	}
	defer resp.Body.Close()
	rgbByte, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	if strings.TrimSpace(string(rgbByte)) != "315.15" {
		t.Error(string(rgbByte))
	}
}

func TestExtensionWithLogic(t *testing.T) {
	c, err := converter.New()
	if err != nil {
		t.Error(err)
		return
	}
	server := httptest.NewServer(GetRouter(c))
	defer server.Close()

	actual, err := extensionExampleHelper(server.URL, true, []map[string]interface{}{
		{
			"from":             "foo",
			"to":               "bar",
			"distance":         -1,
			"formula":          "x ? 100 : 0",
			"placeholder_name": "x",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	if actual != "100" {
		t.Error(actual)
	}

	actual, err = extensionExampleHelper(server.URL, false, []map[string]interface{}{
		{
			"from":             "foo",
			"to":               "bar",
			"distance":         -1,
			"formula":          "x ? 100 : 0",
			"placeholder_name": "x",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	if actual != "0" {
		t.Error(actual)
	}

	actual, err = extensionExampleHelper(server.URL, 5, []map[string]interface{}{
		{
			"from":             "foo",
			"to":               "bar",
			"distance":         -1,
			"formula":          "x > 2 ? 100 : 0",
			"placeholder_name": "x",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	if actual != "100" {
		t.Error(actual)
	}

	actual, err = extensionExampleHelper(server.URL, 1, []map[string]interface{}{
		{
			"from":             "foo",
			"to":               "bar",
			"distance":         -1,
			"formula":          "x > 2 ? 100 : 0",
			"placeholder_name": "x",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	if actual != "0" {
		t.Error(actual)
	}

	actual, err = extensionExampleHelper(server.URL, 1, []map[string]interface{}{
		{
			"from":             "foo",
			"to":               "bar",
			"distance":         -1,
			"formula":          "x > 2",
			"placeholder_name": "x",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	if actual != "false" {
		t.Error(actual)
	}

	actual, err = extensionExampleHelper(server.URL, 5, []map[string]interface{}{
		{
			"from":             "foo",
			"to":               "bar",
			"distance":         -1,
			"formula":          "x > 2",
			"placeholder_name": "x",
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	if actual != "true" {
		t.Error(actual)
	}
}

func extensionExampleHelper(url string, input interface{}, extensions []map[string]interface{}) (output string, err error) {
	buf := bytes.NewBuffer(nil)
	json.NewEncoder(buf).Encode(map[string]interface{}{
		"input":      input,
		"extensions": extensions,
	})
	resp, err := http.Post(url+"/extended-conversions/foo/bar", "application/json", buf)
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()
	temp, err := io.ReadAll(resp.Body)
	if err != nil {
		return output, err
	}
	return strings.TrimSpace(string(temp)), nil
}
