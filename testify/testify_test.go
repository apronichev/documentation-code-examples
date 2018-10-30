package testify

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type MySuite struct{ suite.Suite }

func TestSuite(t *testing.T) {
	suite.Run(t, new(MySuite))
}

func (s *MySuite) TestOne() {
	s.T().Run("subTest", func(t *testing.T) {
		t.Log("Hello")
	})
}

func (s *MySuite) TestTwo() {
	s.Equal("Hello\nWorld", "Hello\nWorld!")
}

func TestFoo(t *testing.T) {
}
