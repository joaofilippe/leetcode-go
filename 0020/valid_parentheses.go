package validparentheses

import "strings"

/*
	Vamos utilizar a técnica dos dois ponteiros
	O = n*log(n)
*/

func isValid(s string) bool {
	s = strings.TrimSpace(s)
	// se o tamanho da string for ímpar, então não é válido, porque algum parêntese não tem par
	if len(s)%2 != 0 || len(s) == 0 {
		return false
	}

	// vamos criar um dicionário para o
	pMap := make(map[rune]rune)
	pMap['('], pMap['['], pMap['{'] = ')', ']', '}'

	/* 
	() 40 41
	[] 91 93
	{} 123 125
	*/


	rStack := make([]rune, 0)

	for _, r := range s {
		_, isOpening := pMap[r]

		if isOpening {
			rStack = append(rStack, r)
		} else if len(rStack) == 0 || pMap[rStack[len(rStack)-1]] != r {
			return false
		} else {
			rStack = rStack[:len(rStack)-1]
		}
	}

	return len(rStack) == 0
}
