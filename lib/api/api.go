/*
 * Copyright 2020 InfAI (CC SES)
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
	"github.com/SENERGY-Platform/converter/lib/api/util"
	"github.com/SENERGY-Platform/converter/lib/converter"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

var endpoints = []func(router *httprouter.Router, converter *converter.Converter){}

func Start(port string, converter *converter.Converter) (srv *http.Server, err error) {
	log.Println("start api")
	router := GetRouter(converter)
	log.Println("add logging and cors")
	corsHandler := util.NewCors(router)
	logger := util.NewLogger(corsHandler)
	log.Println("listen on port ", port)
	srv = &http.Server{Addr: ":" + port, Handler: logger}
	go func() { log.Println(srv.ListenAndServe()) }()
	return srv, nil
}

func GetRouter(converter *converter.Converter) (router *httprouter.Router) {
	router = httprouter.New()
	for _, e := range endpoints {
		log.Println("add endpoints: " + runtime.FuncForPC(reflect.ValueOf(e).Pointer()).Name())
		e(router, converter)
	}
	return
}
