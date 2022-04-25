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

import "github.com/SENERGY-Platform/converter/lib/converter/register"

const WattPerSquareMeter = "urn:infai:ses:characteristic:a03ef4d9-cee5-4c4c-b797-e813db59dde5"
const WattPerSquareMeterName = "w/m²"

func init() {
	register.Labels[WattPerSquareMeter] = WattPerSquareMeterName
}
