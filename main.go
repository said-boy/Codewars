package main

import (
	"fmt"

	"github.com/said-boy/codewars/code6Kyu"
)

func main(){
	a := []rune{'n','s','n','s','n','s','n','s','n','s'}
	fmt.Println(code6Kyu.IsValidWalk(a))
}