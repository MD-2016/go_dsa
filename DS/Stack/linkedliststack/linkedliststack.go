package linkedliststack

type Stack struct {
	list *singlylinkedlist[T]
}

func New() *Stack {
	return &Stack{list: &singlylinkedlist}
}
