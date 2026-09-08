/*
Copyright © 2026 kira607 <kirill.lesckin@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package task

// Node is the base interface implemented by every AST node.
//
// AST nodes represent the syntactic structure of a query.
// They do not contain domain-specific knowledge about tasks.
type Node interface {
	node()
}

// Expression represents an AST node that can be used as an expression.
//
// Every expression is also a Node, but not every Node necessarily
// has to be an expression in a larger language.
type Expression interface {
	Node
	expression()
}

// Field represents a reference to a field.
//
// Examples:
//
//	priority
//	title
//	completed
//	due
//	scheduled
//
// The meaning of the field is resolved by the semantic layer.
type Field struct {
	Name string
}

func (Field) node()       {}
func (Field) expression() {}

// String represents a string literal.
//
// Example:
//
//	"brother"
type String struct {
	Value string
}

func (String) node()       {}
func (String) expression() {}

// Number represents a numeric literal.
//
// Example:
//
//	42
type Number struct {
	Value float64
}

func (Number) node()       {}
func (Number) expression() {}

// Boolean represents a boolean literal.
type Boolean struct {
	Value bool
}

func (Boolean) node()       {}
func (Boolean) expression() {}

// Identifier represents an unquoted identifier.
//
// Examples:
//
//	 medium
//	 high
//	 today
//	 2024-08-15
//
// Identifiers intentionally have no semantic meaning at the AST level.
// The semantic layer determines what a particular identifier means.
type Identifier struct {
	Value string
}

func (Identifier) node()       {}
func (Identifier) expression() {}

// UnaryOperator represents an operator that takes one operand.
type UnaryOperator int

const (
	// OperatorNot represents logical negation.
	//
	// Example:
	//
	//	not completed
	OperatorNot UnaryOperator = iota
)

// UnaryExpression represents an expression with one operand.
//
// Example:
//
//	not completed
//
// The semantic meaning of the operator is resolved later.
type UnaryExpression struct {
	Operator UnaryOperator
	Operand  Expression
}

func (UnaryExpression) node()       {}
func (UnaryExpression) expression() {}

// BinaryOperator represents an operator that takes two operands.
type BinaryOperator int

const (
	// OperatorAnd represents logical conjunction.
	//
	// Example:
	//
	//	completed and priority > medium
	OperatorAnd BinaryOperator = iota

	// OperatorOr represents logical disjunction.
	//
	// Example:
	//
	//	completed or priority > medium
	OperatorOr

	// OperatorEquals represents equality comparison.
	//
	// Example:
	//
	//	priority = medium
	OperatorEquals

	// OperatorNotEquals represents inequality comparison.
	//
	// Example:
	//
	//	priority != low
	OperatorNotEquals

	// OperatorGreaterThan represents a greater-than comparison.
	//
	// Example:
	//
	//	priority > medium
	OperatorGreaterThan

	// OperatorLessThan represents a less-than comparison.
	//
	// Example:
	//
	//	priority < high
	OperatorLessThan

	// OperatorGreaterOrEqual represents a greater-than-or-equal comparison.
	//
	// Example:
	//
	//	priority >= medium
	OperatorGreaterOrEqual

	// OperatorLessOrEqual represents a less-than-or-equal comparison.
	//
	// Example:
	//
	//	priority <= medium
	OperatorLessOrEqual

	// OperatorHas represents a containment or substring operation.
	//
	// Example:
	//
	//	title has "brother"
	OperatorHas

	// OperatorFuzzyMatch represents a fuzzy string match.
	//
	// Example:
	//
	//	title ~= "brother"
	OperatorFuzzyMatch
)

// BinaryExpression represents an expression with two operands.
//
// Examples:
//
//	priority > medium
//	title ~= "word"
//	title has "brother"
//
// BinaryExpression only describes the syntactic relationship between
// its operands. The semantic layer determines whether a particular
// operator is valid for the given operands.
type BinaryExpression struct {
	Operator BinaryOperator
	Left     Expression
	Right    Expression
}

func (BinaryExpression) node()       {}
func (BinaryExpression) expression() {}
