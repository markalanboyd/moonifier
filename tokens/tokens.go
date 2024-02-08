package tokens

// GetTokenSet returns a map of Lua language tokens where each token is a key
// with a boolean value of true. This set includes keywords, operators, and
// symbols used in the Lua language syntax.
func GetTokenSet() map[string]bool {
	tokens := make(map[string]bool)

	tokens["function"] = true
	tokens["do"] = true
	tokens["end"] = true
	tokens["while"] = true
	tokens["repeat"] = true
	tokens["until"] = true
	tokens["if"] = true
	tokens["then"] = true
	tokens["else"] = true
	tokens["elseif"] = true
	tokens["for"] = true
	tokens["in"] = true
	tokens["local"] = true
	tokens["return"] = true
	tokens["break"] = true
	tokens["="] = true
	tokens[","] = true
	tokens["."] = true
	tokens[":"] = true
	tokens["["] = true
	tokens["]"] = true
	tokens["("] = true
	tokens[")"] = true
	tokens["{"] = true
	tokens["}"] = true
	tokens["+"] = true
	tokens["-"] = true
	tokens["*"] = true
	tokens["/"] = true
	tokens["^"] = true
	tokens["%"] = true
	tokens[".."] = true
	tokens["..."] = true
	tokens["<"] = true
	tokens["<="] = true
	tokens[">"] = true
	tokens[">="] = true
	tokens["=="] = true
	tokens["~="] = true
	tokens["and"] = true
	tokens["or"] = true
	tokens["not"] = true
	tokens["#"] = true

	return tokens
}
