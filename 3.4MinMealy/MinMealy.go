package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type automate struct {
	n, m, q0   int
	transition [][]int
	output     [][]string
}

func (a *automate) DFS(v int, used map[int]bool, order map[int]int, pos int) (map[int]bool, map[int]int, int) {
	used[v] = true
	order[pos] = v
	pos++
	for _, u := range a.transition[v] {
		if !used[u] {
			used, order, pos = a.DFS(u, used, order, pos)
		}
	}
	return used, order, pos
}

func (a *automate) reverse(list map[int]int) map[int]int {
	reversedList := make(map[int]int)
	for i := 0; i < len(list); i++ {
		reversedList[list[i]] = i
	}
	return reversedList
}

func (a *automate) canonic() {
	used := make(map[int]bool)
	order := make(map[int]int)
	pos := 0
	used, order, pos = a.DFS(a.q0, used, order, pos)
	orderReversed := a.reverse(order)
	a.n = len(order)
	newTr := make([][]int, a.n)
	newOut := make([][]string, a.n)
	a.q0 = 0
	for i := 0; i < len(order); i++ {
		newTr[i] = make([]int, a.m)
		newOut[i] = make([]string, a.m)
		for j := 0; j < a.m; j++ {
			newTr[i][j] = orderReversed[a.transition[order[i]][j]]
			newOut[i][j] = a.output[order[i]][j]
		}
	}
	a.transition = newTr
	a.output = newOut
}

func find(parent []int, v int) int {
	if parent[v] == v {
		return parent[v]
	} else {
		parent[v] = find(parent, parent[v])
		return parent[v]
	}
}

func union(parent []int, depth []int, u, v int) {
	root1 := find(parent, u)
	root2 := find(parent, v)
	if depth[root1] < depth[root2] {
		parent[root1] = parent[root2]
	} else {
		parent[root2] = parent[root1]
		if depth[root1] == depth[root2] && root1 != root2 {
			depth[root1]++
		}
	}
}

func minimize(a automate) automate {
	split1 := func() (int, []int) {
		parent := make([]int, a.n)
		depth := make([]int, a.n)
		m := a.n
		for i := 0; i < a.n; i++ {
			parent[i] = i
			depth[i] = 0
		}

		for i := 0; i < a.n; i++ {
			for j := i + 1; j < a.n; j++ {
				if find(parent, i) != find(parent, j) {
					checkEq := true
					for k := 0; k < a.m; k++ {
						if a.output[i][k] != a.output[j][k] {
							checkEq = false
							break
						}
					}
					if checkEq {
						union(parent, depth, i, j)
						m--
					}
				}
			}
		}

		partition := make([]int, a.n)
		for i := 0; i < a.n; i++ {
			partition[i] = find(parent, i)
		}
		return m, partition

	}

	split := func(partition *[]int) int {
		parent := make([]int, a.n)
		depth := make([]int, a.n)
		m := a.n
		for i := 0; i < a.n; i++ {
			parent[i] = i
			depth[i] = 0
		}

		for i := 0; i < a.n; i++ {
			for j := i + 1; j < a.n; j++ {
				if (*partition)[i] == (*partition)[j] && find(parent, i) != find(parent, j) {
					equalTransitions := true
					for k := 0; k < a.m; k++ {
						state1 := a.transition[i][k]
						state2 := a.transition[j][k]
						if (*partition)[state1] != (*partition)[state2] {
							equalTransitions = false
							break
						}
					}
					if equalTransitions {
						union(parent, depth, i, j)
						m--
					}
				}
			}
		}

		for i := 0; i < a.n; i++ {
			(*partition)[i] = find(parent, i)
		}

		return m
	}

	minimizedAutomaton := automate{
		n:          a.n,
		m:          a.m,
		q0:         a.q0,
		transition: make([][]int, a.n),
		output:     make([][]string, a.n),
	}
	for i := 0; i < a.n; i++ {
		minimizedAutomaton.transition[i] = make([]int, a.m)
		minimizedAutomaton.output[i] = make([]string, a.m)
		copy(minimizedAutomaton.transition[i], a.transition[i])
		copy(minimizedAutomaton.output[i], a.output[i])
	}

	m, partition := split1()
	for {
		mt := split(&partition)
		if m == mt {
			break
		}
		m = mt
	}

	minimizedTR := make([][]int, a.n)
	minimizedOut := make([][]string, a.n)
	added := make(map[int]bool)
	for i := 0; i < a.n; i++ {
		minimizedTR[i] = make([]int, a.m)
		minimizedOut[i] = make([]string, a.m)
	}

	for i := 0; i < a.n; i++ {
		v := partition[i]
		if !added[v] {
			added[v] = true
			for k := 0; k < a.m; k++ {
				minimizedTR[v][k] = partition[a.transition[v][k]]
				minimizedOut[v][k] = a.output[v][k]
			}
		}
	}
	minimizedAutomaton.transition = minimizedTR[:minimizedAutomaton.n]
	minimizedAutomaton.output = minimizedOut[:minimizedAutomaton.n]

	return minimizedAutomaton
}

func printMealy(a automate) {
	fmt.Println(a.n)
	fmt.Println(a.m)
	fmt.Println(a.q0)
	for i := 0; i < a.n; i++ {
		for j := 0; j < a.m; j++ {
			fmt.Printf("%d ", a.transition[i][j])
		}
		fmt.Println()
	}

	for i := 0; i < a.n; i++ {
		for j := 0; j < a.m; j++ {
			fmt.Printf("%s ", a.output[i][j])
		}
		fmt.Println()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n, m, q0 int
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	q0, _ = strconv.Atoi(scanner.Text())

	transition := make([][]int, n)
	output := make([][]string, n)
	for i := 0; i < n; i++ {
		transition[i] = make([]int, m)
		scanner.Scan()
		line := scanner.Text()
		numbers := strings.Fields(line)
		for j := 0; j < m; j++ {
			transition[i][j], _ = strconv.Atoi(numbers[j])
		}
	}

	for i := 0; i < n; i++ {
		output[i] = make([]string, m)
		scanner.Scan()
		line := scanner.Text()
		lineList := strings.Fields(line)
		for j := 0; j < m && j < len(lineList); j++ {
			if len(lineList[j]) != 0 {
				output[i][j] = lineList[j]
			}
		}
	}

	automaton := automate{n: n, m: m, q0: q0, transition: transition, output: output}
	automaton.canonic()
	automaton = minimize(automaton)
	automaton.canonic()
	fmt.Println("digraph {")
	fmt.Println("rankdir = LR")
	for i := 0; i < automaton.n; i++ {
		for j := 0; j < automaton.m; j++ {
			fmt.Printf("%d -> %d [label = \"%c(%s)\"]\n", i, automaton.transition[i][j], 'a'+j, automaton.output[i][j])
		}
	}
	fmt.Println("}")
}
