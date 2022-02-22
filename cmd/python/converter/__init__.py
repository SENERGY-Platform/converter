#  Copyright 2021 InfAI (CC SES)
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

__version__ = '0.1.0'
__title__ = 'converter'
__description__ = 'python wrapper to use the converter shared lib.'
__url__ = 'https://github.com/SENERGY-Platform/converter/cmd/python'
__author__ = 'Ingo Rößner'
__license__ = 'Apache License 2.0'
__copyright__ = 'Copyright 2021 InfAI (CC SES)'


import json
from ctypes import *

class Converter:
    def __init__(self, lib_location):
        self.lib = cdll.LoadLibrary(lib_location)
        self.lib.Cast.argtypes = [c_char_p, c_char_p, c_char_p]
        self.lib.Cast.restype = c_char_p

    def cast(self, value, from_characteristic, to_characteristic):
        json_value = json.dumps(value, ensure_ascii=False)
        json_value_c = c_char_p(json_value.encode('utf-8'))
        from_c = c_char_p(from_characteristic.encode('utf-8'))
        to_c = c_char_p(to_characteristic.encode('utf-8'))
        out_json = self.lib.Cast(json_value_c, from_c, to_c)
        return json.loads(out_json)