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

package characteristics

import (
	"errors"
	"github.com/SENERGY-Platform/converter/lib/converter/register"
	"gopkg.in/go-playground/colors.v1"
	"log"
	"runtime/debug"
	"strings"
)

const Rgb = "urn:infai:ses:characteristic:5b4eea52-e8e5-4e80-9455-0382f81a1b43"
const RgbR = "urn:infai:ses:characteristic:dfe6be4a-650c-4411-8d87-062916b48951"
const RgbG = "urn:infai:ses:characteristic:5ef27837-4aca-43ad-b8f6-4d95cf9ed99e"
const RgbB = "urn:infai:ses:characteristic:590af9ef-3a5e-4edb-abab-177cb1320b17"

func init() {
	register.Labels[Rgb] = "Rgb"

	register.Add(Hex, Rgb, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		hexStr, ok := in.(string)
		if !ok {
			debug.PrintStack()
			return nil, errors.New("unable to interpret value as string")
		}
		if !strings.HasPrefix(hexStr, "#") {
			hexStr = "#" + hexStr
		}
		hex, err := colors.ParseHEX(hexStr)
		if err != nil {
			debug.PrintStack()
			return nil, err
		}
		rgb := hex.ToRGB()
		return map[string]int64{"r": int64(rgb.R), "g": int64(rgb.G), "b": int64(rgb.B)}, nil
	})

	register.Add(Rgb, Hex, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		rgbMap, ok := in.(map[string]interface{})
		if !ok {
			log.Println(in)
			debug.PrintStack()
			return nil, errors.New("unable to interpret value as map[string]interface{}")
		}
		r, ok := rgbMap["r"]
		if !ok {
			debug.PrintStack()
			return nil, errors.New("missing field r")
		}
		red, ok := r.(float64)
		if !ok {
			debug.PrintStack()
			return nil, errors.New("field r is not a number")
		}
		g, ok := rgbMap["g"]
		if !ok {
			debug.PrintStack()
			return nil, errors.New("missing field g")
		}
		green, ok := g.(float64)
		if !ok {
			debug.PrintStack()
			return nil, errors.New("field g is not a number")
		}
		b, ok := rgbMap["b"]
		if !ok {
			debug.PrintStack()
			return nil, errors.New("missing field b")
		}
		blue, ok := b.(float64)
		if !ok {
			debug.PrintStack()
			return nil, errors.New("field b is not a number")
		}
		rgb, err := colors.RGB(uint8(red), uint8(green), uint8(blue))
		if err != nil {
			debug.PrintStack()
			return nil, err
		}
		return rgb.ToHEX().String(), nil
	})
}
