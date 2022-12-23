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
	"github.com/SENERGY-Platform/models/go/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func init() {
	endpoints = append(endpoints, ValidationEndpoint)
}

func ValidationEndpoint(router *httprouter.Router, converter *converter.Converter) {
	resource := "/validate/extended-conversions"

	type ValidationRequest struct {
		Nodes      []string                    `json:"nodes"`
		Extensions []models.ConverterExtension `json:"extensions"`
	}

	router.POST(resource, func(writer http.ResponseWriter, request *http.Request, ps httprouter.Params) {
		r := ValidationRequest{}
		err := json.NewDecoder(request.Body).Decode(&r)
		if err != nil {
			http.Error(writer, "expect valid json in request body", http.StatusBadRequest)
			return
		}
		err = converter.ValidateExtensions(r.Nodes, r.Extensions)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		writer.WriteHeader(http.StatusOK)
	})

}
