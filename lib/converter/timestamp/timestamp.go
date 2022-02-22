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

package timestamp

import (
	"github.com/SENERGY-Platform/converter/lib/converter/base"
	"github.com/SENERGY-Platform/converter/lib/model"
)

var characteristicToConcept = &base.CastCharacteristicToConcept{}
var conceptToCharacteristic = &base.CastConceptToCharacteristic{}

const ConceptId = "urn:infai:ses:concept:9286d398-f49d-494f-8157-bea17947b1fa"
const ConceptName = "timestamp"

func init() {
	base.Concepts[ConceptId] = base.GetConceptCastFunction(characteristicToConcept, conceptToCharacteristic)
	base.ConceptRepo.Register(model.Concept{Id: ConceptId, Name: ConceptName, BaseCharacteristicId: IsoId}, []model.Characteristic{
		{
			Id:   UnixMilliSecondsId,
			Name: UnixMilliSecondsName,
			Type: model.Integer,
		},
		{
			Id:   UnixNanoSecondsId,
			Name: UnixNanoSecondsName,
			Type: model.Integer,
		},
		{
			Id:   UnixSecondsId,
			Name: UnixSecondsName,
			Type: model.Integer,
		},
		{
			Id:   IsoId,
			Name: IsoName,
			Type: model.String,
		},
	})
}