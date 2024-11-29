package core

import (
	"fmt"
	"testing"
)

func printVal(val int) {
	switch val {
	case 0:
		fmt.Println("[ ] [*] [ ] [ ] [ ]")
		//fmt.Println()
	case 1:
		fmt.Println("[ ] [ ] [ ] [*] [ ]")
		//fmt.Println()
	case 2:
		fmt.Println("[ ] [ ] [*] [ ] [ ]")
		//fmt.Println()
	case 3:
		fmt.Println("[*] [ ] [ ] [ ] [ ]")
		//fmt.Println()
	case 4:
		fmt.Println("[ ] [ ] [ ] [ ] [*]")
		//fmt.Println()
	default:
		fmt.Println()
	}
}

func TestRandomByTwo(t *testing.T) {
	list := RandomBy(TwoNotesMap, 100, ResetRules(RuleSameFoot, RuleReverse))
	arr := list.ToArray()
	for _, val := range arr {
		printVal(val)
	}
}

func TestRandomByFour(t *testing.T) {
	list := RandomBy(FourNotesRunMap, 100, ResetRules(RuleSameFoot, RuleReverse, RuleDiagonal, RuleNoRepeat))
	arr := list.ToArray()
	for _, val := range arr {
		printVal(val)
	}
}
