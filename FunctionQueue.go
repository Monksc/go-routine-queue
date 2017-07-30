package FunctionQueue

type FunctionQueue struct {
	head *FunctionNode
	isRunning bool
}

func NewFunctionQueue() FunctionQueue {
	return FunctionQueue{nil, false}
}

func (self *FunctionQueue) Add(function func ()) {
	node := newFunctionNode(function)

	if self.head != nil {
		self.head.add(node)
	} else {
		self.head = node
		self.run()
	}
}

func (self *FunctionQueue) run() {
	self.head.run(func() {
		self.head = self.head.next
		if self.head != nil {
			self.run()
		}
	})
}
