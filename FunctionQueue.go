package FunctionQueue

type FunctionQueue struct {
	head *FunctionNode
	isRunning bool
	completion func()
}

func NewFunctionQueue() FunctionQueue {
	return FunctionQueue{nil, false, nil}
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
		} else if self.completion != nil {
			self.completion()
		}
	})
}

func (self *FunctionQueue) SetCompletion(completion func()) {
	self.completion = completion()
}
