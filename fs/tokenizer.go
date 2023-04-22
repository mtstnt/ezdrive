package fs

func Tokenize(input string) ([]string, error) {
	var (
		tokens      = make([]string, 0)
		idx         = 0
		temp        = ""
		openerState = ""
	)

	for idx <= len(input) {
		if idx == len(input) || (idx < len(input) && input[idx] == ' ' && openerState == "") {
			// Handle adding to tokens list only if we meet a space (and quoted) or we meet the end of the input (quoted or unquoted)
			tokens = append(tokens, temp)
			temp = ""
		} else if temp == "" && (input[idx] == '"' || input[idx] == '\'') && openerState == "" {
			// Handle meeting quotations.
			openerState = string(input[idx])
		} else if openerState != "" && openerState == string(input[idx]) {
			// Handle meeting closing quotations.
			openerState = ""
		} else {
			temp += string(input[idx])
		}
		idx += 1
	}

	return tokens, nil
}
