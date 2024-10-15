# Parser

## Introduction
This is a simple parser that reads a string of various asserts and validates if all the assert are valid or not.

```
Stmt := ‘assert’ Expr

Expr ::=
| [0-9]+[0-9]* ;; constants
| [a-zA-Z_]+[a-zA-Z0-9_]* ;; variables
| ‘(‘ Expr ‘)’
| Expr ‘+’ Expr ;; addition
| Expr ‘*’ Expr ;; multiplication
| Expr '/' Expr ;; division
| Expr ‘-’ Expr ;; less than
| Expr ‘==’ Expr ;; equality
| Expr ‘<=’ Expr ;; inequality
| `!` Expr ;; not
```

These are the various asserts that can be validated by the parser. 

* The parser can also simplify constant expression and try to evaluate the expression.
* For complete evaluation of expression with variables, a value map can be passed to the parser. 
* The parser will then try to evaluate the expression with the given values of the variables.

A simple example of the parser is as follows:
```
assert x * (2 * 3)
```

This assert will be evaluated to `assert (x * 6.00)`

If we further pass a value-map as 
```
{
    x: 0
}
```

The expression will be evaluated to 0.00
As 0 is considered as true value in the parser. The assert will be valid.


## Code Structure
```
.
├── constants/
│   └── tokens.go            [List of tokens]
├── lexer/
│   ├── lexer_test.go        [Test cases for the lexer]
│   └── lexer.go             [Lexer implementation]
├── parser/
│   ├── assert.go            [Assert Node for the parser]
│   ├── variable.go          [Variable Node for the parser]
│   ├── infix_expression.go  [Infix Node for the parser]
│   ├── number.go            [Number Node for the parser]
│   ├── parser_test.go       [Test cases for the parser]
│   ├── parser.go            [Parser implementation]
│   ├── prefix_expression.go [Prefix Node for the parser]
│   ├── program.go           [Entry point for evalutations of the asserts]  
│   ├── types.go             [Interfaces for the Node and various types]
│   └── utils.go             [Utility functions for the parser]
├── go.mod
├── go.sum          
├── main.go                  [Main file to run the parser]
└── Readme.md                [Readme file]
```
The code is divided into 2 main parts:
1. Lexer: The lexer reads the input string and converts it into tokens.
    * The lexer uses a simple state machine to read the input string.
    * Call NextToken() to get the next token from the lexer.
2. Parser: The parser reads the tokens and validates the expression.
    * The parser uses a recursive descent parser to validate the expression.
    * Parser will create a list of statements where each statement is an assert.
    * Each statement can be evaluated or partially evaluated based on the values of the variables. PS. Partial evaluation don't require values for the variables.

## Code Logic
### Lexer
The lexer interface consists of the following functions. A new custom lexer can be created by implementing the above interface. 

```go
type Lexerer interface {
	NewToken() constants.Token  // Create a new token
	readNumber() string         // Reads a number from the input string
	readVariable() string       // Reads a variabe like x, y, z from the input string 
	peakChar() byte             // Peeks the next character from the input string
	skipWhitespace()            // Skips the white spaces from the input string
}
```
The lexer implementation is fairly simple and self explanatory. Use the `NewLexer(input string)` function to create a new lexer. The object returned will implement the above interface. That object is then futher used by the parser to read the tokens and validate the expressions.

### Parser
The parser interface consists of the following functions. A new custom parser can be created by implementing the above interface. 

```go
type Parserer interface {
	Errors() []string  // Returns the list of errors
	ParseProgram() *Program  // Initiates the parsing of the program

	nextToken()  // Moves to the next token and updates the current and peek token
	printAST(any)  // Prints the AST of the given node
	parseStatement() Statement  // Parses a single statement
	parseAssertStatement() *AssertStatement	 // Parses an assert statement
	parseExpression(int) Expression  // Parses an expression with given precedence
	prefixParseFns(tokenType constants.TokenType) func() Expression  // Returns the prefix parse function for the given token
	infixParseFns(tokenType constants.TokenType) func(Expression) Expression  // Returns the infix parse function for the given token
	parseVariable() Expression  // Parses a variable and return that expression
	parseNumberLiteral() Expression  // Parses a number and return that expression
	parsePrefixExpression() Expression  // Parses a prefix expression
	parseInfixExpression(left Expression) Expression  // Parses an infix expression
	parseGroupedExpression() Expression  // Parses a grouped expression, having '(' and ')'
	noPrefixParseFnError(t constants.TokenType)  // Returns an error if no prefix parse function is found
	expectPeek(t constants.TokenType) bool  // Expects the next token to be of the given type
	peekTokenIs(t constants.TokenType) bool  // Checks if the next token is of the given type
	peekError(t constants.TokenType)  // Returns an error if the next token is not of the given type
	peekPrecedence() int  // Returns the precedence of the next token
	curPrecedence() int  // Returns the precedence of the current token
}
```
The parser object is created by calling the NewParser(lexer.Lexerer) function.
```go
type Parser struct {
	l         lexer.Lexerer	      // Lexer object
	curToken  constants.Token     // Current token
	peekToken constants.Token     // Next token
	errors    []string	      // List of errors
}
```

### Program
Program is the entry point for the evaluation of the asserts. 
```go
type ProgramEvaluator interface {
	SetValueMap(map[string]float64)			// Set the value map for the variables	
	Evaluate() ([]float64, []error, bool)		// Evaluate the asserts
	PartialEvaluate() ([]string, []error, bool)     // Partially evaluate and simplify the asserts without the values of the variables
}
```

### PrintVisitor
PrintVisitor is a simple visitor that prints the AST of the parser. It uses the Visitor pattern to traverse the AST and print the nodes. The indentation is used to show the depth of the nodes in the AST and is increased recursively as we go deeper into the AST. It's invoked after all the statements are parsed.

```go
type PrintVisitor interface {
	VisitProgram(*Program, int) // Goes inside a program and tries to print the statements inside it
	VisitExpression(Expression, int) // Goes inside an expression and tries to print the nodes inside as per the type of the node, i.e. prefix, infix, etc.
}

```

## Features
1. The parser is mostly used for validating the asserts. A simple example of an assert is `assert x * (2 * 3)`. 
2. If the value of x is 0, the assert will be valid as the expression will be evaluated to 0.00.
3. Multiple asserts can be passed to the parser. The parser will validate all the asserts. If all the asserts are valid, the parser will return true.
4. Each assert is considered as a statement. The parser will create a list of statements and validate each statement.
5. The parser comes with 2 main functions `Evaluate()` and `PartialEvaluate()`. 
	* `Evaluate()` will evaluate the asserts with the given values of the variables. If the asserts are valid, it will return true.
	* `PartialEvaluate()` will simplify the asserts without the values of the variables. It will return the simplified asserts. Partial evaluation don't require the values of the variables.
6. The parser can also evaluate the expression with the given values of the variables. The values of the variables can be passed to the parser using the `SetValueMap(map[string]float64)` function.
7. A partially evaluated assert can be further evaluated with the values of the variables. The parser will try to evaluate the expression with the given values of the variables. Do check out the test cases in `TestDualParser()` in parser_test.go for more details.
8. The parser can also evaluate basic expressions like `1 + 2 * 3`. The parser will evaluate the expression to `7.00`. It will be marked as failed as the result is not 0.00.

## How to run the parser
Add your asserts and valueMap if any in the main.go 
```go
assertValues := `
	assert (2 + 6) * (x - 3) 
	assert (y + 6) 
	assert z * 2
`
valueMap := map[string]float64{
	"x": 3, 
	"y": -6, 
	"z": 0,
}
```

This part of the code will create a lexer and parser object. It will partially evaluate the asserts and print the simplified results.
```go
l := lexer.NewLexer(assertValues)
p := parser.NewParser(l)
program := p.ParseProgram()

simplifiedResult, errors, isSuccess := program.PartialEvaluate()
if !isSuccess {
	for _, err := range errors {
		fmt.Println(err)
	}
	os.Exit(1)
}

fmt.Println("Simplified results :-")
for _, result := range simplifiedResult {
	fmt.Println(result)
}

```

It also generates and prints out the AST using the PrintVisitor which uses the Visitor pattern to traverse the AST and print the nodes. Here is the partial evalution result of the AST for the above asserts.
```
AST :-
Program:
    -> Statement:
        AssertStatement:
            InfixExpression:
                Left:
                    InfixExpression:
                        Left:
                            NumberLiteral: 2.00
                        Operator: +
                        Right:
                            NumberLiteral: 6.00
                Operator: *
                Right:
                    InfixExpression:
                        Left:
                            Variable: x
                        Operator: -
                        Right:
                            NumberLiteral: 3.00
    -> Statement:
        AssertStatement:
            InfixExpression:
                Left:
                    Variable: y
                Operator: +
                Right:
                    NumberLiteral: 6.00
    -> Statement:
        AssertStatement:
            InfixExpression:
                Left:
                    Variable: z
                Operator: *
                Right:
                    NumberLiteral: 2.00
```


This part of the code will add the valueMap to the parser and evaluate the asserts with the given values of the variables. It can use the original asserts or even the `partiallyEvaluated` resposes as well. The parser will return `isSuccess` as true if all the asserts are valid or false if any of the assert is invalid.
```go
fmt.Println("\nAdding value map for the asserts")

program.SetValueMap(valueMap)
fmt.Println()

l = lexer.NewLexer(combinedPartialResults)
p = parser.NewParser(l)
program = p.ParseProgram()

_, errors, isSuccess = program.Evaluate()
if !isSuccess {
	for _, err := range errors {
		fmt.Println(err)
	}
	fmt.Println("Asserts Failed [X]")
	os.Exit(1)
}

fmt.Println("Asserts Passed [✓]")
```

For the above asserts and valueMap, the parser will return the following output. All the asserts will be satisfied and will get result as 0, thus the parser will return `Asserts Passed [✓]`.
```
Evaluating statement 1
Assert value: 0
Result: 0.000000


Evaluating statement 2
Assert value: 0
Result: 0.000000


Evaluating statement 3
Assert value: 0
Result: 0.000000

Asserts Passed [✓]
```

Also check out the test cases in parser_test.go for more examples and use cases.

## Future Improvements
1. The parser can be further improved to support more arithmetic operations like powers, factorials, log() etc.
2. Addition of various other operators like OR, AND, XOR, etc.
3. Better error handling and reporting.
4. More robust test cases and also better testing for edge cases.








