package main

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	res := make([]byte, len(s))

	for r:= 0; r < numRows; r++ {
		for i := r; i < len(s); i += 2 * (numRows - 1) {
			res = append(res, s[i])
			
			if r > 0 && r < numRows-1 &&
				i + 2 * (numRows - 1) - 2*r < len(s) {
				res = append(res, s[i + 2 * (numRows - 1) - 2*r])
			}
		}
	}

	return string(res[len(s):])
}
