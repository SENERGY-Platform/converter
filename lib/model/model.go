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

package model

var NullCharacteristic = Characteristic{
	Id:   "null",
	Name: "null",
}

var NullConcept = Concept{
	Id:   "null",
	Name: "null",
}

type Concept struct {
	Id                   string   `json:"id"`
	Name                 string   `json:"name"`
	CharacteristicIds    []string `json:"characteristic_ids"`
	BaseCharacteristicId string   `json:"base_characteristic_id"`
	RdfType              string   `json:"rdf_type"`
}

type Characteristic struct {
	Id                 string           `json:"id"`
	Name               string           `json:"name"`
	Type               Type             `json:"type"`
	MinValue           interface{}      `json:"min_value,omitempty"`
	MaxValue           interface{}      `json:"max_value,omitempty"`
	Value              interface{}      `json:"value,omitempty"`
	SubCharacteristics []Characteristic `json:"sub_characteristics"`
	RdfType            string           `json:"rdf_type"`
}

type Type string

const (
	String  Type = "https://schema.org/Text"
	Integer Type = "https://schema.org/Integer"
	Float   Type = "https://schema.org/Float"
	Boolean Type = "https://schema.org/Boolean"

	List      Type = "https://schema.org/ItemList"
	Structure Type = "https://schema.org/StructuredValue"
)
