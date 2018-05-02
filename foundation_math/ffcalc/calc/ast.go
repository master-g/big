package calc

// ASTNode represent a node of abstract syntax tree
type ASTNode struct {
	Token *Token
	Left  *ASTNode
	Right *ASTNode
}

// NewASTNode returns a new ASTNode
func NewASTNode(token *Token, left, right *ASTNode) *ASTNode {
	return &ASTNode{
		Token: token,
		Left:  left,
		Right: right,
	}
}

// ASTNodeVisitor defines visitor pattern interface for AST
type ASTNodeVisitor interface {
	// Visit interface
	Visit(node *ASTNode)
}

// Accept visitor interface
func (n *ASTNode) Accept(visitor ASTNodeVisitor) {
	visitor.Visit(n)
	if n.Right != nil {
		n.Right.Accept(visitor)
	}
	if n.Left != nil {
		n.Left.Accept(visitor)
	}
}

// ASTStack stack structure for ASTNodes
type ASTStack []*ASTNode

// NewASTStack returns new ASTStack
func NewASTStack() ASTStack {
	return ASTStack{}
}

// Empty returns true if stack is empty
func (s ASTStack) Empty() bool {
	return len(s) == 0
}

// Peek returns the top element of the stack without popping it
func (s ASTStack) Peek() *ASTNode {
	if len(s) == 0 {
		return nil
	}
	return s[len(s)-1]
}

// Push a element to the top of the stack
func (s *ASTStack) Push(n *ASTNode) {
	*s = append(*s, n)
}

// Pop a element from the top of the stack
func (s *ASTStack) Pop() *ASTNode {
	if len(*s) == 0 {
		return nil
	}
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}

// AddNode create and push a new ASTNode by popping left and right node from the stack
func (s *ASTStack) AddNode(t *Token) {
	if len(*s) < 2 {
		return
	}

	r := s.Pop()
	l := s.Pop()
	s.Push(NewASTNode(t, l, r))
}
