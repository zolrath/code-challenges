package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func perfectSquare(n *big.Int) bool {
	num := SqrtBig(n)
	return num.Mul(num, num).Cmp(n) == 0
}

func isFib(n *big.Int) string {
	five := big.NewInt(5)
	fourAdd := big.NewInt(4)
	fourSub := big.NewInt(4)
	square := n.Mul(n, n)
	five.Mul(five, square)
	if perfectSquare(fourSub.Add(five, fourSub)) || perfectSquare(fourAdd.Sub(five, fourAdd)) {
		return "IsFibo"
	}
	return "IsNotFibo"
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	t, _ := strconv.Atoi(s.Text())

	for i := 0; i < t; i++ {
		s.Scan()
		bi := new(big.Int)
		_, _ = fmt.Sscan(s.Text(), bi)
		fmt.Println(isFib(bi))
	}
}

// SqrtBig returns floor(sqrt(n)). It panics on n < 0.
func SqrtBig(n *big.Int) (x *big.Int) {
	switch n.Sign() {
	case -1:
		panic(-1)
	case 0:
		return big.NewInt(0)
	}

	var px, nx big.Int
	x = big.NewInt(0)
	x.SetBit(x, n.BitLen()/2+1, 1)
	for {
		nx.Rsh(nx.Add(x, nx.Div(n, x)), 1)
		if nx.Cmp(x) == 0 || nx.Cmp(&px) == 0 {
			break
		}
		px.Set(x)
		x.Set(&nx)
	}
	return
}
