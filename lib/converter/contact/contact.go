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

package contact

import (
	"github.com/SENERGY-Platform/converter/lib/converter/base"
	"github.com/SENERGY-Platform/converter/lib/model"
)

var characteristicToConcept = &base.CastCharacteristicToConcept{}
var conceptToCharacteristic = &base.CastConceptToCharacteristic{}

const Contact = "urn:infai:ses:concept:fb85f32b-1864-456d-b99f-1540892ffd02"

func init() {
	base.Concepts[Contact] = base.GetConceptCastFunction(characteristicToConcept, conceptToCharacteristic)
	base.ConceptRepo.Register(model.Concept{Id: Contact, Name: "contact", BaseCharacteristicId: ContactState}, []model.Characteristic{
		{
			Id:   ZwaveContactState,
			Name: "zwave contact state",
			Type: model.String,
		},
		{
			Id:   ContactState,
			Name: "contact state",
			Type: model.Boolean,
		},
	})
}
