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

package characteristics

import (
	"errors"
	"github.com/SENERGY-Platform/converter/lib/converter/register"
	"log"
	"reflect"
	"runtime/debug"
)

const MicroGramPerCubicMeterCO = "urn:infai:ses:characteristic:193ff994-60a4-4ed2-90a1-21b82622cb43"
const MilliGramPerCubicMeterCO = "urn:infai:ses:characteristic:e4e4fbb3-ef93-4702-b770-44e4186a00b4"
const PpmCO = "urn:infai:ses:characteristic:15af5b7b-8b17-4b18-b37e-2e91c839886b"

func init() {
	register.Labels[MicroGramPerCubicMeterCO] = "Microgram Per Cubic Meter (Carbon Monoxide)"
	register.Labels[MilliGramPerCubicMeterCO] = "Milligram Per Cubic Meter (Carbon Monoxide)"
	register.Labels[PpmCO] = "Parts per Million (Carbon Monoxide)"

	register.Add(MicroGramPerCubicMeterCO, MilliGramPerCubicMeterCO, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		switch microgram := in.(type) {
		case int:
			return float64(microgram) / 1000, nil
		case int32:
			return float64(microgram) / 1000, nil
		case int64:
			return float64(microgram) / 1000, nil
		case float32:
			return float64(microgram) / 1000, nil
		case float64:
			return microgram / 1000, nil
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
	})

	register.Add(MilliGramPerCubicMeterCO, MicroGramPerCubicMeterCO, register.NoLosses, func(in interface{}) (out interface{}, err error) {
		switch milligram := in.(type) {
		case int:
			return float64(milligram) * 1000, nil
		case int32:
			return float64(milligram) * 1000, nil
		case int64:
			return float64(milligram) * 1000, nil
		case float32:
			return float64(milligram) * 1000, nil
		case float64:
			return milligram * 1000, nil
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
	})

	//https://uk-air.defra.gov.uk/assets/documents/reports/cat06/0502160851_Conversion_Factors_Between_ppb_and.pdf
	//conversion presumes 20°C and 1013mb
	register.Add(PpmCO, MilliGramPerCubicMeterCO, register.MajorLoss, func(in interface{}) (out interface{}, err error) {
		switch ppm := in.(type) {
		case int:
			return float64(ppm) * 1.1642, nil
		case int32:
			return float64(ppm) * 1.1642, nil
		case int64:
			return float64(ppm) * 1.1642, nil
		case float32:
			return float64(ppm) * 1.1642, nil
		case float64:
			return ppm * 1.1642, nil
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
	})

	//https://uk-air.defra.gov.uk/assets/documents/reports/cat06/0502160851_Conversion_Factors_Between_ppb_and.pdf
	//conversion presumes 20°C and 1013mb
	register.Add(MilliGramPerCubicMeterCO, PpmCO, register.MajorLoss, func(in interface{}) (out interface{}, err error) {
		switch mgm3 := in.(type) {
		case int:
			return float64(mgm3) / 1.1642, nil
		case int32:
			return float64(mgm3) / 1.1642, nil
		case int64:
			return float64(mgm3) / 1.1642, nil
		case float32:
			return float64(mgm3) / 1.1642, nil
		case float64:
			return mgm3 / 1.1642, nil
		default:
			debug.PrintStack()
			log.Println("ERROR: ", reflect.TypeOf(in).String(), in)
			return nil, errors.New("unable to interpret value; input type is " + reflect.TypeOf(in).String())
		}
	})

}
