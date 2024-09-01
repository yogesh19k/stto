package utils

/*
Single Line comments Map, Key: extension name ,
value:[Comment symbol]
*/
var comment_map map[string]string = map[string]string{
	"go":     "//",
	"c":      "//",
	"cpp":    "//",
	"js":     "//",
	"ts":     "//",
	"jsx":    "//",
	"tsx":    "//",
	"java":   "//",
	"rs":     "//",
	"swift":  "//",
	"kt":     "//",
	"php":    "//",
	"m":      "//",
	"groovy": "//",
	"cs":     "//",
	"scala":  "//",
	"zig":    "//",
	"gleam":  "//",
	"py":     "#",
	"r":      "#",
	"rb":     "#",
	"sh":     "#",
	"pl":     "#",
	"ex":     "#",
	"exs":    "#",
	"jl":     "#",
	"lua":    "--",
	"hs":     "--",
	"sql":    "--",
	"cbl":    "*",
	"erl":    "%",
	"clj":    ";;",
	"lisp":   ";",
}

/*
Multi Line comments Map, Key: extension name ,
value:[Opening symbol]:[Closing symbol]
*/
var multi_comment_map map[string]string = map[string]string{
	"go":     "/*:*/",
	"c":      "/*:*/",
	"cpp":    "/*:*/",
	"js":     "/*:*/",
	"ts":     "/*:*/",
	"jsx":    "/*:*/",
	"tsx":    "/*:*/",
	"java":   "/*:*/",
	"rs":     "/*:*/",
	"swift":  "/*:*/",
	"kt":     "/*:*/",
	"php":    "/*:*/",
	"m":      "/*:*/",
	"groovy": "/*:*/",
	"cs":     "/*:*/",
	"scala":  "/*:*/",
	"zig":    "/*:*/",
	"gleam":  "/*:*/",
	"py":     "\"\"\":\"\"\"",
	"sh":     ": ':'",
	"pl":     "/*:*/",
	"jl":     "#=:=#",
	"lua":    "--[[:]]--",
	"hs":     "/*:*/",
	"sql":    "/*:*/",
	"cbl":    "/*:*/",
	"erl":    "=begin:=cut",
}
