package graph

import (
	"bytes"
	"fmt"
	"io"

	"github.com/libcd/libcd"
	"github.com/libcd/libcd/parse"
)

// Create creates a new graphviz dot document from the spec file and returns
// the contents of the file as a []byte.
func Create(spec *libcd.Spec) []byte {
	var buf bytes.Buffer
	WriteTo(spec, &buf)
	return buf.Bytes()
}

// WriteTo creates a new graphviz dot document from the spec file and writes the
// document to the io.Writer.
func WriteTo(spec *libcd.Spec, w io.Writer) {
	g := grapher{writer: w, parent: "start"}
	g.writeHeader()
	g.writeNode(spec.Nodes.ListNode)
	g.writeTrailer()
}

type grapher struct {
	writer  io.Writer
	counter int

	parent string
	child  string
}

func (g *grapher) writeNode(node parse.Node) {
	g.counter++

	switch v := node.(type) {
	case *parse.ListNode:
		g.writeList(v)
	case *parse.DeferNode:
		g.writeDefer(v)
	case *parse.RecoverNode:
		g.writeRecover(v)
	case *parse.ParallelNode:
		g.writeParallel(v)
	case *parse.RunNode:
		g.writeRun(v)
	}
}

func (g *grapher) writeList(node *parse.ListNode) {
	for _, node := range node.Body {
		g.writeNode(node)
	}
}

func (g *grapher) writeDefer(node *parse.DeferNode) {
	g.writeNode(node.Body)
	g.writeNode(node.Defer)
}

func (g *grapher) writeRecover(node *parse.RecoverNode) {
	g.writeNode(node.Body)
}

func (g *grapher) writeParallel(node *parse.ParallelNode) {
	for _, node := range node.Body {
		g.writeNode(node)
	}
}

func (g *grapher) writeRun(node *parse.RunNode) {
	g.child = fmt.Sprintf("%s_%d", node.Name, g.counter)
	g.writeLink(g.parent, g.child)
	g.parent = g.child
}

func (g *grapher) writeHeader() {
	io.WriteString(g.writer, "digraph G { compound=true; ")
}

func (g *grapher) writeTrailer() {
	io.WriteString(g.writer, "start [shape=Mdiamond]; end [shape=Mdiamond]; }")
}

func (g *grapher) writeSubHeader(name string) {
	fmt.Fprintf(g.writer, "subgraph %s {", name)
}

func (g *grapher) writeSubTrailer() {
	io.WriteString(g.writer, "};")
}

func (g *grapher) writeLink(parent, child string) {
	if parent == "" {
		fmt.Fprintf(g.writer, "%s;", child)
		return
	}
	fmt.Fprintf(g.writer, "%s -> %s;", parent, child)
}
