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
	"github.com/SENERGY-Platform/converter/lib/converter/characteristics"
	"testing"
)

func TestRgbToHsb(t *testing.T) {
	t.Parallel()
	out, err := Cast(map[string]interface{}{"r": float64(54), "g": float64(94), "b": float64(169)}, characteristics.Rgb, characteristics.Hsb)
	if err != nil {
		t.Fatal(err)
	}
	hsb, ok := out.(map[string]int64)
	if !ok {
		t.Fatal(out)
	}

	if hsb["h"] != 219 {
		t.Fatal(hsb)
	}

	if hsb["s"] != 68 {
		t.Fatal(hsb)
	}

	if hsb["b"] != 66 {
		t.Fatal(hsb)
	}
}

func TestHexToHsb(t *testing.T) {
	t.Parallel()
	out, err := Cast("#365ea9", characteristics.Hex, characteristics.Hsb)
	if err != nil {
		t.Fatal(err)
	}
	hsb, ok := out.(map[string]int64)
	if !ok {
		t.Fatal(out)
	}

	if hsb["h"] != 219 {
		t.Fatal(hsb)
	}

	if hsb["s"] != 68 {
		t.Fatal(hsb)
	}

	if hsb["b"] != 66 {
		t.Fatal(hsb)
	}
}

func TestHsbToHex(t *testing.T) {
	t.Parallel()
	out, err := Cast(map[string]interface{}{"h": float64(219), "s": float64(68), "b": float64(66)}, characteristics.Hsb, characteristics.Hex)
	if err != nil {
		t.Fatal(err)
	}
	hex, ok := out.(string)
	if !ok {
		t.Fatal(out)
	}

	if hex != "#365ea8" {
		t.Fatal(hex)
	}
}

func TestHsbToRgb(t *testing.T) {
	t.Parallel()
	out, err := Cast(map[string]interface{}{"h": float64(219), "s": float64(68), "b": float64(66)}, characteristics.Hsb, characteristics.Rgb)
	if err != nil {
		t.Fatal(err)
	}
	rgb, ok := out.(map[string]int64)
	if !ok {
		t.Fatal(out)
	}

	if rgb["r"] != 54 {
		t.Fatal(rgb)
	}

	if rgb["g"] != 94 {
		t.Fatal(rgb)
	}

	if rgb["b"] != 168 {
		t.Fatal(rgb)
	}
}
