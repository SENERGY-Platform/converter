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

type CastFunction func(in interface{}) (out interface{}, err error)
type CharacteristicId = string

type Entry struct {
	From     CharacteristicId
	To       CharacteristicId
	Distance int64
	Cast     CastFunction
}

var List = []Entry{}
var Labels = map[string]string{}

//NilCast makes no change to the value
func NilCast(in interface{}) (out interface{}, err error) {
	return in, nil
}

//Add registers cast functions. the functions may be nil to indicate that no value change on cast
func Add(from CharacteristicId, to CharacteristicId, distance int64, cast CastFunction) {
	List = append(List, Entry{
		From:     from,
		To:       to,
		Distance: distance,
		Cast:     cast,
	})
}

type Register interface {
	GetCasts(from CharacteristicId, to CharacteristicId) (casts []CastFunction, err error)
	Update(register []Entry) (err error)
}
