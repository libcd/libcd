package parse

import "fmt"

type RunNode struct {
	NodeType `json:"type"`

	Name string `json:"name"`
}

func (n *RunNode) SetName(name string) *RunNode {
	n.Name = name
	return n
}

func NewRunNode() *RunNode {
	return &RunNode{NodeType: NodeRun}
}

func (n *RunNode) Validate() error {
	switch {
	case n.NodeType != NodeRun:
		return fmt.Errorf("Run Node uses an invalid type")
	case n.Name == "":
		return fmt.Errorf("Run Node has an invalid name")
	default:
		return nil
	}
}
