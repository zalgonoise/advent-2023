package day20

import _ "embed"

//go:embed input/input.txt
var Input string

var TestInput1 = `
broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a
`

var TestInput2 = `
broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
`
