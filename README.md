# API to convert between values
## Endpoints
### /conversions
### /extended-conversions
### /validate/extended-conversions

# expressions usable in extensions send to /extended-conversions

## govaluate documentation
https://github.com/Knetic/govaluate#what-operators-and-types-does-this-support
https://github.com/Knetic/govaluate/blob/master/MANUAL.md

## custom functions

### atoi

- parses string as integer
- signature: func(string)int64
- example: `atoi(x) == 42`

### atof

- parses string as float
- signature: func(string)float64
- example: `atof(x) == 42.13`

### ntoa

- formats number as string
- signature: func(number)string; number may be int, int16, int32, int64, float32, float64
- example: `ntoa(x) == "42"`

### strlen

- returns length of string
- signature: func(string)int
- example: `strlen("42") == 2`

### strIndex

- returns the index of the first instance of substr in s, or -1 if substr is not present in s
- signature: func(s string, substr string)int
- example: `strIndex("4.2", ".") == 1`

### substr

- returns a slice of the string s, from 'from' up to excluding 'to'
- signature: func(s string, from int, to int)string
- example: `substr("0123456789", 2, 4) == "23"`

### trimPrefix

- returns s without the provided leading prefix string. if s doesn't start with prefix, s is returned unchanged.
- signature: func(s string, prefix string)string
- example: `trimPrefix("foo:4.2", "foo:") == "4.2"`

### trimSuffix

- returns s without the provided trailing suffix string. if s doesn't end with suffix, s is returned unchanged.
- signature: func(s string, suffix string)string
- example: `trimSuffix("foo:4.2", ":4.2") == "foo"`

### replace

- returns a copy of the string s with all non-overlapping instances of old replaced by new. if old is empty, it matches at the beginning of the string
- signature: func(s string, old string, new string)string
- example: `replace("foo:4.2", ":", "/") == "foo/4.2"`

### toUpperCase

- returns s with all Unicode letters mapped to their upper case.
- signature: func(s string)string
- example: `toUpperCase("fooBAR") == "FOOBAR"`

### toLowerCase

- returns s with all Unicode letters mapped to their lower case.
- signature: func(s string)string
- example: `toLowerCase("fooBAR") == "foobar"`

# default conversions
to get a graph of possible conversions, call
```
go generate ./...
```
which prints a dot graph. you can use https://dreampuf.github.io/GraphvizOnline to generate images from the output.  

![conversion graph](./graphviz.svg)

# known dependent repositories

- https://github.com/SENERGY-Platform/marshaller
- https://github.com/SENERGY-Platform/external-task-worker
- https://github.com/SENERGY-Platform/mgw-external-task-worker
- https://github.com/SENERGY-Platform/device-command