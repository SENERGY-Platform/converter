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
	"encoding/json"
	"github.com/SENERGY-Platform/converter/lib/converter"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func init() {
	endpoints = append(endpoints, ConversionsEndpoint)
}

func ConversionsEndpoint(router *httprouter.Router) {
	resource := "/conversions"

	router.GET(resource+"/:from/:to", func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		inStr := request.URL.Query().Get("json")
		from := ps.ByName("from")
		to := ps.ByName("to")
		var in interface{}
		err := json.Unmarshal([]byte(inStr), &in)
		if err != nil {
			http.Error(writer, "expect valid json in 'json' query parameter", http.StatusBadRequest)
			return
		}
		out, err := converter.Cast(in, from, to)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		err = json.NewEncoder(writer).Encode(out)
		if err != nil {
			log.Println("ERROR: unable to encode response", err)
		}
	})

	router.POST(resource+"/:from/:to", func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		from := ps.ByName("from")
		to := ps.ByName("to")
		var in interface{}
		err := json.NewDecoder(request.Body).Decode(&in)
		if err != nil {
			http.Error(writer, "expect valid json in request body", http.StatusBadRequest)
			return
		}
		out, err := converter.Cast(in, from, to)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		writer.Header().Set("Content-Type", "application/json; charset=utf-8")
		err = json.NewEncoder(writer).Encode(out)
		if err != nil {
			log.Println("ERROR: unable to encode response", err)
		}
	})

}
