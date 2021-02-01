/*
 * Copyright 2019 InfAI (CC SES)
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

package converter

import (
	"errors"
	"github.com/SENERGY-Platform/converter/lib/converter/base"
	_ "github.com/SENERGY-Platform/converter/lib/converter/battery"
	_ "github.com/SENERGY-Platform/converter/lib/converter/binary"
	_ "github.com/SENERGY-Platform/converter/lib/converter/brightness"
	_ "github.com/SENERGY-Platform/converter/lib/converter/color"
	_ "github.com/SENERGY-Platform/converter/lib/converter/contact"
	_ "github.com/SENERGY-Platform/converter/lib/converter/energy"
	_ "github.com/SENERGY-Platform/converter/lib/converter/example"
	_ "github.com/SENERGY-Platform/converter/lib/converter/humidity"
	_ "github.com/SENERGY-Platform/converter/lib/converter/luminiscence"
	_ "github.com/SENERGY-Platform/converter/lib/converter/particleamount"
	_ "github.com/SENERGY-Platform/converter/lib/converter/power"
	_ "github.com/SENERGY-Platform/converter/lib/converter/pressure"
	_ "github.com/SENERGY-Platform/converter/lib/converter/speedlevel"
	_ "github.com/SENERGY-Platform/converter/lib/converter/temperature"
	_ "github.com/SENERGY-Platform/converter/lib/converter/time"
	_ "github.com/SENERGY-Platform/converter/lib/converter/ultraviolet"
	"github.com/SENERGY-Platform/converter/lib/model"
	"runtime/debug"
)

var ConceptRepo = base.ConceptRepo

func Cast(in interface{}, from string, to string) (out interface{}, err error) {
	fromConceptId, err := base.ConceptRepo.GetConceptOfCharacteristic(from)
	if err != nil {
		return nil, err
	}
	toConceptId, err := base.ConceptRepo.GetConceptOfCharacteristic(to)
	if err != nil {
		return nil, err
	}
	if fromConceptId != toConceptId {
		return nil, errors.New("expect " + fromConceptId + " and " + toConceptId + " to have the same concept")
	}
	if from == model.NullCharacteristic.Id || to == model.NullCharacteristic.Id || fromConceptId == model.NullConcept.Id {
		return in, nil
	}
	return Concepts(fromConceptId)(from)(in)(to)
}

func Concepts(conceptId string) base.FindCastFromCharacteristicToConceptFunction {
	result, ok := base.Concepts[conceptId]
	if !ok {
		debug.PrintStack()
		return base.GetErrorFindCastFromCharacteristicToConceptFunction(errors.New("concept not found"))
	}
	return result
}
