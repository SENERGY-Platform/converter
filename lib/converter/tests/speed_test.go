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

package converter

import (
	"github.com/SENERGY-Platform/converter/lib/converter/characteristics"
	"math"
	"testing"
)

func FuzzSpeed(f *testing.F) {
	testcases := []float64{10, 0, 2, 30, -20, 2.3, -3.3}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig float64) {
		temp, err := Cast(orig, characteristics.Kmh, characteristics.MetersPerSecond)
		if err != nil {
			t.Error(err)
			return
		}
		temp, err = Cast(temp, characteristics.MetersPerSecond, characteristics.Kmh)
		if err != nil {
			t.Error(err)
			return
		}
		if v, ok := temp.(float64); !ok || math.Abs(v-orig) > 0.00001 {
			t.Errorf("Before: %v, after: %v", orig, v)
		}
	})
}
