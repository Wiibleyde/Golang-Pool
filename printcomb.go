package student

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for i := 0; i < 10; i++ {
		for j := i + 1; j < 10; j++ {
			for k := j + 1; k < 10; k++ {
				if i < j && j < k {
					z01.PrintRune(rune(i) + 48)
					z01.PrintRune(rune(j) + 48)
					z01.PrintRune(rune(k) + 48)
					if rune(i+48) != '7' || rune(j+48) != '8' || rune(k+48) != '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else {
						z01.PrintRune('\n')
					}
				}
			}
		}
	}
}
