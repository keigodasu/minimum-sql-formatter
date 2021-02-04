package minimum_sql_formatter

import "testing"

func TestNextToken(t *testing.T) {
	t.Run("symbol tokens", func(t *testing.T) {
		input := `,;`

		tests := []struct{
			expectedType    TokenType
			expectedLiteral string
		} {
			{COMMA, ","},
			{SEMICOLON, ";"},
		}

		lexer := NewLexer(input)

		for _, tt := range tests {
			tok := lexer.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("got %v, want %v", tok.Type, tt.expectedType)
			}

			if tok.Literal!= tt.expectedLiteral {
				t.Fatalf("got %v, want %v", tok.Literal, tt.expectedLiteral)
			}
		}
	})
	
	t.Run("keyword tokens", func(t *testing.T) {
		input := ` SELECT col1, col2 
from test_table; 
`
		tests := []struct{
			expectedType    TokenType
			expectedLiteral string
		} {
			{SELECT, "SELECT"},
			{IDENT, "COL1"},
			{COMMA, ","},
			{IDENT, "COL2"},
			{FROM, "FROM"},
			{IDENT, "TEST_TABLE"},
			{SEMICOLON, ";"},
		}

		lexer := NewLexer(input)

		for _, tt := range tests {
			tok := lexer.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("got %v, want %v", tok.Type, tt.expectedType)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("got %v, want %v", tok.Literal, tt.expectedLiteral)
			}
		}
	})
}
