/*
 * Copyright 2022 InfAI (CC SES)
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
	"fmt"
	_ "github.com/SENERGY-Platform/converter/lib/converterV2/characteristics"
	"github.com/SENERGY-Platform/converter/lib/converterV2/register"
	"github.com/emicklei/dot"
	"strconv"
)

//go:generate go run dot.go ExportDefaultDotGraph

func main() {
	ExportDotGraph(register.List, register.Labels)
}

func ExportDotGraph(registry []register.Entry, labels map[string]string) {
	g := dot.NewGraph(dot.Directed)
	for _, e := range registry {
		fromNode := g.Node(e.From)
		if l, ok := labels[e.From]; ok {
			fromNode.Label(l)
		}
		toNode := g.Node(e.To)
		if l, ok := labels[e.To]; ok {
			toNode.Label(l)
		}
		g.Edge(fromNode, toNode, strconv.Itoa(int(e.Distance)))
	}
	fmt.Println(g.String())
}
