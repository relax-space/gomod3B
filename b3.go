package gomod3B

import gomod3C_v6 "github.com/relax-space/gomod3C/v6"

func AddB(a, b int) int {
	return a + b
}

func MultiB(a, b int) int {
	return gomod3C_v6.MultiC(a, b)
}
