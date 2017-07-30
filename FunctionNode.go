package FunctionQueue

type FunctionNode struct {
	next *FunctionNode
	function func ()
}

func newFunctionNode(function func ()) *FunctionNode {
	return &FunctionNode{nil, function}
}

func (self *FunctionNode) add(node *FunctionNode) {
	if self.next == nil {
		self.next = node
	} else {
		self.next.add(node)
	}
}

func (self *FunctionNode) run(completion func()) {
	go func() {
		self.function()
		completion()
	}()
}

func (self *FunctionNode) size() int {
	if self.next != nil {
		return 1 + self.next.size()
	}
	return 0
}