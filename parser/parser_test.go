package parser

import (
	"pscript/lexer"
	"testing"
)

func TestNew(t *testing.T) {
	input := "token1 token2"

	lexer := lexer.New(input)

	tests := struct {
		expectedCurTokenValue  string
		expectedPeekTokenValue string
	}{
		expectedCurTokenValue:  "token1",
		expectedPeekTokenValue: "token2",
	}

	actual := New(lexer)

	if actual.CurToken.Literal != tests.expectedCurTokenValue {
		t.Errorf("FAIL on CurToken value: expected %s, got %s", tests.expectedCurTokenValue, actual.CurToken.Literal)
	}

	if actual.PeekToken.Literal != tests.expectedPeekTokenValue {
		t.Errorf("FAIL on PeekToken value: expected %s, got %s", tests.expectedPeekTokenValue, actual.CurToken.Literal)
	}
}
