package parser

// AST Node Interface
type (
	Node interface {
		TokenLiteral() string
		Evaluate() (float64, error)
		PartialEvaluate() (string, error)
	}

	Statement interface {
		Node
		String() string
	}

	Expression interface {
		Node
		String() string
	}
)
