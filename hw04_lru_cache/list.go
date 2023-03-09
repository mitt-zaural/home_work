package hw04lrucache

type List interface {
	Len() int							// длина списка
	Front() *ListItem					// первый элемент списка
	Back() *ListItem					// последний элемент списка
	MoveToFront(i *ListItem)			// переместить элемент в начало
	PushFront(v interface{}) *ListItem	// добавить значение в начало
	PushBack(v interface{}) *ListItem	// добавить значение в конец
	Remove(i *ListItem)					// удалить элемент
}

type ListItem struct {
	Value interface{}	// значение
	Next  *ListItem		// следующий элемент
	Prev  *ListItem		// предыдущий элемент
}

type list struct {
	List // Remove me after realization.
	// Place your code here.
	root	ListItem
	//tail	*ListItem
	len	int
}

// создание нового пустого списка
func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

// первый элемент списка
func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.root.Next
}

// последний элемент списка
func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return	l.tail.Prev
}

// добавить значение в начало
func (l *list) PushFront(v interface{}) *ListItem {
	item := new(ListItem)		// новый элемент списка
	if item != nil {			// успешно создан?
		item.Value = v			// присваиваем элементу списка переданное значение
		if l.len == 0 {			// пустой список?
			l.root.Next = &item
			l.root.Prev = &item
		} else {				// нет
			item.Next = l.root.Prev
			l.root.Next = l.root.Prev
			l.root.Prev = &item
		}
		l.len++
	}
	return item
}

// добавить значение в конец
func (l *list) PushBack(v interface{}) *ListItem {
	item := new(ListItem)		// новый элемент списка
	if item != nil {			// успешно создан?
		item.Value = v			// присваиваем элементу списка переданное значение
		if l.len == 0 {			// пустой список?
			l.root.Next = &item
			l.root.Prev = &item
		} else {				// нет
			item.Prev = l.root.Next
			item.Next = &item
			l.root.Next = l.root.Prev
			l.root.Prev = 
		}
		l.len++
	}
	return item
}

// удалить элемент
func (l *list) Remove(i *ListItem) {
	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	i.Next = nil
	i.Prev = nil
	i.
	l.len--
}

// переместить элемент в начало
func (l *list) MoveToFront(i *ListItem) {
	if i.list != l || l.root.Next == i {
		return
	}
	i.Prev.Next = l.root.Next
	i.Next.Prev = i.root.
}
