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
)

type SimpleRegister struct {
	casts map[CharacteristicId]map[CharacteristicId]CastFunction
}

func NewSimpleRegister(register []Entry) (this *SimpleRegister, err error) {
	this = &SimpleRegister{
		casts: map[CharacteristicId]map[CharacteristicId]CastFunction{},
	}
	for _, entry := range register {
		if _, ok := this.casts[entry.From]; !ok {
			this.casts[entry.From] = map[CharacteristicId]CastFunction{}
		}
		this.casts[entry.From][entry.To] = entry.Cast
	}
	return this, nil
}

func (this *SimpleRegister) GetCasts(from CharacteristicId, to CharacteristicId) (casts []CastFunction, err error) {
	if from == to {
		return []CastFunction{}, nil
	}
	castFrom, ok := this.casts[from]
	if !ok {
		return casts, errors.New("unknown cast from " + from)
	}
	cast, ok := castFrom[to]
	if !ok {
		return casts, errors.New("unknown cast from " + from + " to " + to)
	}
	casts = append(casts, cast)
	return casts, nil
}
