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
	"log"
	"log/slog"
	"os"
	"os/signal"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	"github.com/SENERGY-Platform/converter/lib/api"
	"github.com/SENERGY-Platform/converter/lib/converter"
	struct_logger "github.com/SENERGY-Platform/go-service-base/struct-logger"
)

func main() {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	setDefaultSlog(logLevel)
	c, err := converter.New()
	if err != nil {
		slog.Error("FATAL: unable to load converter", "error", err)
		log.Fatal("ERROR: unable to load converter", err)
	}
	srv, err := api.Start("8080", c)
	if err != nil {
		slog.Error("FATAL: unable to start api", "error", err)
		log.Fatal("ERROR: unable to start api", err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	sig := <-shutdown
	slog.Info("received shutdown signal", "signal", sig, "result", srv.Shutdown(context.Background()))
}

func setDefaultSlog(logLevel string) {
	info, ok := debug.ReadBuildInfo()
	project := ""
	org := ""
	if ok {
		if parts := strings.Split(info.Main.Path, "/"); len(parts) > 2 {
			project = strings.Join(parts[2:], "/")
			org = strings.Join(parts[:2], "/")
		}
	}
	logger := struct_logger.New(
		struct_logger.Config{
			Handler:    struct_logger.JsonHandlerSelector,
			Level:      logLevel,
			TimeFormat: time.RFC3339Nano,
			TimeUtc:    true,
			AddMeta:    true,
		},
		os.Stdout,
		org,
		project,
	)
	slog.SetDefault(logger)
	slog.SetLogLoggerLevel(slog.LevelInfo)
}
