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

package main

import (
	"context"
	"github.com/SENERGY-Platform/converter/lib/api"
	"github.com/SENERGY-Platform/converter/lib/converter"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c, err := converter.New()
	if err != nil {
		log.Fatal("ERROR: unable to load converter", err)
	}
	srv, err := api.Start("8080", c)
	if err != nil {
		log.Fatal("ERROR: unable to start api", err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	sig := <-shutdown
	log.Println("received shutdown signal", sig)
	log.Println(srv.Shutdown(context.Background()))
}
