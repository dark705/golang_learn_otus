package list

type Item struct {
	next  *Item
	prev  *Item
	list  *List
	value interface{}
}

func (i *Item) Next() *Item {
	if i.list != nil && i.next != &i.list.item {
		return i.next
	}
	return nil
}

func (i *Item) Prev() *Item {
	if i.list != nil && i.prev != &i.list.item {
		return i.prev
	}
	return nil
}

func (i *Item) Value() interface{} {
	if i.list != nil {
		return i.value
	}
	return nil
}

type List struct {
	item Item
	len  int
}

func (l *List) Len() int {
	return l.len
}

func (l *List) setZeroItemPoint() {
	l.item.next = &l.item
	l.item.prev = &l.item
}

func (l *List) First() *Item {
	if l.len == 0 {
		return nil
	}
	return l.item.next
}

func (l *List) Last() *Item {
	if l.len == 0 {
		return nil
	}
	return l.item.prev
}

func (l *List) add(v interface{}, i *Item) *Item {
	addedItem := Item{value: v}
	oldNext := i.next
	i.next = &addedItem
	addedItem.prev = i
	addedItem.next = oldNext
	oldNext.prev = &addedItem
	addedItem.list = l
	l.len++
	return &addedItem
}

func (l *List) Remove(i *Item) {
	if i.list == l {
		i.prev.next = i.next
		i.next.prev = i.prev
		i.prev = nil
		i.next = nil
		i.list = nil
		l.len--
	}
	if l.Len() == 0 {
		l.item.prev = nil
		l.item.next = nil
	}
}

func (l *List) PushFront(v interface{}) *Item {
	if l.item.next == nil {
		l.setZeroItemPoint()
	}
	return l.add(v, &l.item)
}

func (l *List) PushBack(v interface{}) *Item {
	if l.item.next == nil {
		l.setZeroItemPoint()
	}
	return l.add(v, l.item.prev)
}
