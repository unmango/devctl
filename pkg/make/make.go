package make

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Rule struct {
	Target  []string
	PreReqs []string
	Recipe  []string
}

type Makefile struct {
	Rules []Rule
}

func Parse(r io.Reader) (*Makefile, error) {
	m := &Makefile{}
	s := NewScanner(r)
	for s.Scan() {
		switch n := s.Node().(type) {
		case Rule:
			m.Rules = append(m.Rules, n)
		}
	}

	if s.Err() != nil {
		return nil, s.Err()
	}

	return nil, nil
}

func parseRule(s *bufio.Scanner) (*Rule, error) {
	line := s.Text()
	elem := strings.Split(line, ":")

	if len(elem) != 2 {
		return nil, fmt.Errorf("invalid rule: %s", s.Text())
	}

	return nil, nil
}
