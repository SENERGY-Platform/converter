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

package pressure

import (
	"github.com/SENERGY-Platform/converter/lib/converter/base"
	"github.com/SENERGY-Platform/converter/lib/model"
)

var characteristicToConcept = &base.CastCharacteristicToConcept{}
var conceptToCharacteristic = &base.CastConceptToCharacteristic{}

const Pressure = "urn:infai:ses:concept:9924ea1a-d96a-4882-88a2-165493956ec3"

func init() {
	base.Concepts[Pressure] = base.GetConceptCastFunction(characteristicToConcept, conceptToCharacteristic)
	base.ConceptRepo.Register(model.Concept{Id: Pressure, Name: "pressure", BaseCharacteristicId: HPa}, []model.Characteristic{
		{
			Id:   HPa,
			Name: "hPa",
			Type: model.Float,
		},
		{
			Id:   Bar,
			Name: "bar",
			Type: model.Float,
		},
	})
}
