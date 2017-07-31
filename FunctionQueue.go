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
	self.completion = completion
}

func (self *FunctionQueue) Size() int {
	if self.head != nil {
		return self.head.size()
	}

	return 0
}

func (self *FunctionQueue) Remove(index int) *FunctionNode {

	if index == 0 {
		panic("Cannot remove Queue at index 0")
	}

	traverse := self.head
	for i:=1; i < index; i++ {
		traverse = traverse.next
	}

	node := traverse.next
	traverse.next = traverse.next.next

	return node
}

func (self *FunctionQueue) RemoveAllButCurrent() {
	if self.head != nil {
		self.head.next = nil
	}
}
