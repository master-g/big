package calc

// ASTNode represent a node of abstract syntax tree
type ASTNode struct {
	Token *Token
	Left  *ASTNode
	Right *ASTNode
}

func NewASTNode(t *Token, left, right *ASTNode) *ASTNode {
	return &ASTNode{
		Token: t,
		Left:  left,
		Right: right,
	}
}

type ASTNodeVisitor interface {
	Visit(node *ASTNode)
}

func (n *ASTNode) Accept(visitor ASTNodeVisitor) {
	visitor.Visit(n)
	if n.Right != nil {
		n.Right.Accept(visitor)
	}
	if n.Left != nil {
		n.Left.Accept(visitor)
	}
}

type ASTStack []*ASTNode

func NewASTStack() ASTStack {
	return ASTStack{}
}

func (s ASTStack) Empty() bool {
	return len(s) == 0
}

func (s ASTStack) Peek() *ASTNode {
	if len(s) == 0 {
		return nil
	}
	return s[len(s)-1]
}

func (s *ASTStack) Push(n *ASTNode) {
	*s = append(*s, n)
}

func (s *ASTStack) Pop() *ASTNode {
	if len(*s) == 0 {
		return nil
	}
	n := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return n
}

func (s *ASTStack) AddNode(t *Token) {
	if len(*s) < 2 {
		return
	}

	r := s.Pop()
	l := s.Pop()
	s.Push(NewASTNode(t, l, r))
}
