import json
from ctypes import *

class Converter:
    def __init__(self, lib_location):
        self.lib = cdll.LoadLibrary(lib_location)
        self.lib.Cast.argtypes = [c_char_p, c_char_p, c_char_p]
        self.lib.Cast.restype = c_char_p
        self.lib.ListCharacteristics.argtypes = []
        self.lib.ListCharacteristics.restype = c_char_p

    def cast(self, value, from_characteristic, to_characteristic):
        json_value = json.dumps(value, ensure_ascii=False)
        json_value_c = c_char_p(json_value.encode('utf-8'))
        from_c = c_char_p(from_characteristic.encode('utf-8'))
        to_c = c_char_p(to_characteristic.encode('utf-8'))
        out_json = self.lib.Cast(json_value_c, from_c, to_c)
        return json.loads(out_json)

    def list_characteristics(self):
        result = self.lib.ListCharacteristics()
        return json.loads(result)


#c = Converter("/location/of/converter.so")
#print(c.cast(24, "urn:infai:ses:characteristic:5ba31623-0ccb-4488-bfb7-f73b50e03b5a", "urn:infai:ses:characteristic:75b2d113-1d03-4ef8-977a-8dbcbb31a683"))
#print(c.cast("24", "urn:infai:ses:characteristic:5ba31623-0ccb-4488-bfb7-f73b50e03b5a", "urn:infai:ses:characteristic:75b2d113-1d03-4ef8-977a-8dbcbb31a683"))
#print(c.list_characteristics())