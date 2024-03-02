package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(longestPalindrome("ararar"))
}

func longestPalindrome(s string) string {
    T := "^#" + strings.Join(strings.Split(s, ""), "#") + "#$"
    n := len(T)
    P := make([]int, n)
    C, R := 0, 0

	// C = 3  R = 6 // i = 5

	//       C     R
    // 0 1 2 3 4 5 6 7 8 9 10
	// # a # r # a # r # a #
//P // - 1 0 3 0
	t  := strings.Split(T,"")
	fmt.Println(t)

    for i := 1; i < n-1; i++ {
        if R > i {
            P[i] = min(R-i, P[2*C-i])
        }
		
		
        for T[i+1+P[i]] == T[i-1-P[i]] {
			r := i+1+P[i]
			l := i-1-P[i]

			fmt.Println(r,l)
            P[i]++
        }


        if i + P[i] > R {
            C, R = i, i + P[i]
        }
    }
    
    maxLen := 0
    centerIndex := 0
    for i, v := range P {
        if v > maxLen {
            maxLen = v
            centerIndex = i
        }
    }
    return s[(centerIndex-maxLen)/2 : (centerIndex+maxLen)/2]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}