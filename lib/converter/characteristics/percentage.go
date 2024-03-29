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

//TODO: replace with selected percentage id (old version uses multiple different percentages, the new version only this one)
const Percentage = "urn:infai:ses:characteristic:5caa707d-dc08-4f3b-bd9f-f08935c8dd3c"

func init() {
	register.Labels[Percentage] = "Percentage"
}
