package mapreader

import (
	"fmt"
	"regexp"
)

var nodeRegex, _ = regexp.Compile(`(\w+) = \((\w+), (\w+)\)`)

const (
	startNode = "AAA"
	endNode   = "ZZZ"
)

const (
	right = 'R'
	left  = 'L'
)

func GetStepsTillDestination(mapInput Map) int {
	return mapInput.ApplyDirectionsTillDestination()
}

func ParseMap(input []string) Map {
	directions := input[0]
	mapNodes := make(map[string]MapNode)
	for _, nodeData := range input[2:] {
		node := parseMapNode(nodeData)
		mapNodes[node.Id] = node
	}

	return Map{directions, mapNodes}
}

type MapNode struct {
	Id      string
	LeftId  string
	RightId string
}

func (n MapNode) getNextNode(direction rune) (string, error) {
	if direction == right {
		return n.RightId, nil
	} else if direction == left {
		return n.LeftId, nil
	} else {
		return "", fmt.Errorf("unknown direction id: %s", direction)
	}
}

type Map struct {
	Directions string
	MapNodes   map[string]MapNode
}

func (m Map) ApplyDirectionsTillDestination() int {
	currentNode := startNode
	directions := m.Directions
	moveCount := 0
	for ; currentNode != endNode; moveCount++ {
		direction := rune(directions[moveCount%len(directions)])
		currentNode, _ = m.MapNodes[currentNode].getNextNode(direction)
	}

	return moveCount
}

func parseMapNode(line string) MapNode {
	nodeData := nodeRegex.FindStringSubmatch(line)
	if len(nodeData) < 4 {
		panic("unexpected map node format")
	}

	return MapNode{nodeData[1], nodeData[2], nodeData[3]}
}
