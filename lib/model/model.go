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
