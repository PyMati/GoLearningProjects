package main

import "fmt"

const MaxNumber = 3

type Stack struct {
	items []int
}

type Tower struct {
	stack Stack
}

func (s *Stack) push(num int) {
	s.items = append(s.items, num)
}

func (s *Stack) top() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) pop() int {
	last_item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last_item
}

func (s *Stack) isEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

func (s *Stack) print() {
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Println(s.items[i])
	}
}

func printTowers(towers []Tower) {
	for i := 0; i < len(towers); i++ {
		fmt.Printf("Tower index %d\n", i)
		towers[i].stack.print()
	}
}

func checkWin(towers []Tower) bool {
	if len(towers[len(towers)-1].stack.items) == MaxNumber {
		return true
	}

	return false
}

func main() {
	first_stack := Stack{[]int{}}
	first_stack.push(3)
	first_stack.push(2)
	first_stack.push(1)
	first_tower := Tower{first_stack}

	second_stack := Stack{[]int{}}
	second_tower := Tower{second_stack}

	third_stack := Stack{[]int{}}
	third_tower := Tower{third_stack}

	towers := []Tower{first_tower, second_tower, third_tower}

	for {
		var fsTowerIndex int
		var scTowerIndex int

		printTowers(towers)

		fmt.Scan(&fsTowerIndex)
		fmt.Scan(&scTowerIndex)

		if fsTowerIndex > len(towers) || scTowerIndex > len(towers) {
			fmt.Printf("You provided bad index! Index must be greater than 0 and no longer than %d\n", len(towers)-1)
			continue
		}

		if !towers[scTowerIndex].stack.isEmpty() {
			if towers[fsTowerIndex].stack.top() < towers[scTowerIndex].stack.top() {
				towers[scTowerIndex].stack.push(towers[fsTowerIndex].stack.pop())
			} else {
				fmt.Println("You can't move bigger chunk on lower chunk!")
			}
		} else {
			towers[scTowerIndex].stack.push(towers[fsTowerIndex].stack.pop())
		}

		if checkWin(towers) {
			fmt.Println("You won!")
			break
		}

	}

}
