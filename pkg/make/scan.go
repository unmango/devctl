package make

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type NodeType string

const (
	RuleNode        NodeType = "rule"
	VarNode         NodeType = "var"
	CommentNode     NodeType = "comment"
	UnsupportedNode NodeType = "unsupported"
)

type Node interface{}

type Scanner struct {
	s     *bufio.Scanner
	lines []string
	start int
	err   error
	node  NodeType
}

func (s *Scanner) Type() NodeType {
	return s.node
}

func (s *Scanner) Node() Node {
	switch s.node {
	case RuleNode:
		return Rule{
			Target: parseTarget(s.Lines()[0]),
		}
	case CommentNode:
		return fmt.Errorf("TODO")
	default:
		return nil
	}
}

func (s *Scanner) Lines() []string {
	return s.lines[s.start:]
}

func (s *Scanner) Err() error {
	return s.err
}

func (s *Scanner) Scan() bool {
	if len(s.lines) > 0 {
		s.start = len(s.lines) - 1
	}
	if s.start == 0 || s.node == UnsupportedNode {
		if !s.advance() {
			return false
		}
	}

	switch {
	case strings.Contains(s.line(), ":"):
		return s.rule()
	default:
		s.node = UnsupportedNode
		return true
	}
}

func (s *Scanner) advance() (cont bool) {
	if cont = s.s.Scan(); cont {
		s.lines = append(s.lines, s.s.Text())
	}

	return
}

func (s *Scanner) line() string {
	return s.lines[len(s.lines)-1]
}

func (s *Scanner) rule() bool {
	s.node = RuleNode
	for s.advance() {
		switch {
		// case strings.Contains(s.line(), "\t"):
		// 	continue
		default:
			return true
		}
	}

	return false
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{s: bufio.NewScanner(r)}
}

func parseTarget(l string) []string {
	if b, _, ok := strings.Cut(l, ":"); ok {
		return strings.Split(b, " ")
	} else {
		return []string{}
	}
}
