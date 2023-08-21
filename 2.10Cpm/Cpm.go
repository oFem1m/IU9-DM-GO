package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Name         string
	Duration     int
	Dependencies []*Node
	Critical     bool
}

func (n *Node) AddDependency(dep *Node) {
	n.Dependencies = append(n.Dependencies, dep)
}

func (n *Node) SetCritical() {
	n.Critical = true
}

func (n *Node) IsDependentOn(node *Node) bool {
	for _, dep := range n.Dependencies {
		if dep == node {
			return true
		}
	}
	return false
}

func ParseWork(workStr string) (*Node, error) {
	workParts := strings.Split(workStr, "(")
	name := strings.TrimSpace(workParts[0])
	durationStr := strings.TrimRight(workParts[1], ")")
	duration := 0
	fmt.Sscanf(durationStr, "%d", &duration)

	node := &Node{
		Name:     name,
		Duration: duration,
	}

	return node, nil
}

func FindCriticalPath(node *Node) []*Node {
	visited := make(map[*Node]bool)
	path := []*Node{}
	dfs(node, visited, &path)
	return path
}

func dfs(node *Node, visited map[*Node]bool, path *[]*Node) {
	visited[node] = true

	for _, dep := range node.Dependencies {
		if !visited[dep] {
			dfs(dep, visited, path)
		}
	}

	*path = append(*path, node)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nodes := make(map[string]*Node)

	for scanner.Scan() {
		workStr := scanner.Text()
		workParts := strings.Split(workStr, "<")
		workParts[0] = strings.TrimSpace(workParts[0])
		workParts[1] = strings.TrimSpace(workParts[1])

		node, err := ParseWork(workParts[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		nodes[node.Name] = node

		if len(workParts[1]) > 0 {
			dependencies := strings.Split(workParts[1], "<")
			for _, depStr := range dependencies {
				depName := strings.TrimSpace(depStr)
				depNode, exists := nodes[depName]
				if !exists {
					depNode = &Node{Name: depName}
					nodes[depName] = depNode
				}
				node.AddDependency(depNode)
			}
		}
	}

	for _, node := range nodes {
		path := FindCriticalPath(node)
		for _, n := range path {
			n.SetCritical()
		}
	}

	fmt.Println("digraph {")
	for _, node := range nodes {
		color := "red"
		if !node.Critical {
			color = "blue"
		}
		fmt.Printf("  %s [label = \"%s(%d)\", color = %s]\n", node.Name, node.Name, node.Duration, color)
		for _, dep := range node.Dependencies {
			depColor := "blue"
			if dep.Critical && node.Critical {
				depColor = "red"
			}
			fmt.Printf("  %s -> %s [color = %s]\n", dep.Name, node.Name, depColor)
		}
	}
	fmt.Println("}")
}
