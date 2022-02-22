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

package register

import (
	"errors"
	"github.com/RyanCarrier/dijkstra"
)

type GraphRegister struct {
	graph *dijkstra.Graph
	casts map[CharacteristicId]map[CharacteristicId]CastFunction
}

func NewGraphRegister(register []Entry) (this *GraphRegister, err error) {
	this = &GraphRegister{
		graph: dijkstra.NewGraph(),
		casts: map[CharacteristicId]map[CharacteristicId]CastFunction{},
	}
	for _, entry := range register {
		if _, ok := this.casts[entry.From]; !ok {
			this.casts[entry.From] = map[CharacteristicId]CastFunction{}
		}
		this.casts[entry.From][entry.To] = entry.Cast
		if _, err := this.graph.GetMapping(entry.From); err != nil {
			this.graph.AddMappedVertex(entry.From)
		}
		if _, err := this.graph.GetMapping(entry.To); err != nil {
			this.graph.AddMappedVertex(entry.To)
		}
		err = this.graph.AddMappedArc(entry.From, entry.To, entry.Distance)
		if err != nil {
			return this, err
		}
	}
	return this, nil
}

func (this *GraphRegister) GetCasts(from CharacteristicId, to CharacteristicId) (casts []CastFunction, err error) {
	if from == to {
		return []CastFunction{}, nil
	}
	path, err := this.path(from, to)
	if err != nil {
		return casts, err
	}
	for _, element := range path {
		castFrom, ok := this.casts[element.From]
		if !ok {
			return casts, errors.New("unknown cast from " + element.From)
		}
		cast, ok := castFrom[element.To]
		if !ok {
			return casts, errors.New("unknown cast from " + element.From + " to " + element.To)
		}
		casts = append(casts, cast)
	}
	return casts, nil
}

type PathElement struct {
	From CharacteristicId
	To   CharacteristicId
}

func (this *GraphRegister) path(from CharacteristicId, to CharacteristicId) (path []PathElement, err error) {
	fromId, err := this.graph.GetMapping(from)
	if err != nil {
		return path, err
	}
	toId, err := this.graph.GetMapping(to)
	if err != nil {
		return path, err
	}
	idPath, err := this.graph.Shortest(fromId, toId)
	if err != nil {
		return path, err
	}
	prev := ""
	for _, idElement := range idPath.Path {
		if prev != "" {
			element := PathElement{From: prev}
			element.To, err = this.graph.GetMapped(idElement)
			if err != nil {
				return path, err
			}
			prev = element.To
		} else {
			prev, err = this.graph.GetMapped(idElement)
			if err != nil {
				return path, err
			}
		}
	}
	return path, nil
}
