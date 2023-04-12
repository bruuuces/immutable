package immutable

type List[T any] struct {
	d    []T
	size int
}

func (l *List[T]) Get(index int) T {
	return l.d[index]
}
func (l *List[T]) Size() int {
	return l.size
}
func (l *List[T]) Iterator() ListIterator[T] {
	return ListIterator[T]{l.Size(), 0, l}
}

type ListIterator[T any] struct {
	len      int
	position int
	list     *List[T]
}

func (l *ListIterator[T]) hasNext() bool {
	return l.position < l.len
}

func (l *ListIterator[T]) Next() (int, T) {
	var empty T
	if !l.hasNext() {
		return -1, empty
	}
	position := l.position
	l.position++
	return position, l.list.Get(position)
}

type ListBuilder[T any] struct {
	d []T
}

func (l *ListBuilder[T]) Add(value T) {
	l.d = append(l.d, value)
}

func (l *ListBuilder[T]) Set(index int, value T) {
	l.d[index] = value
}
func (l *ListBuilder[T]) Build() *List[T] {
	data := make([]T, len(l.d))
	copy(data, l.d[:])
	return &List[T]{
		d:    data,
		size: len(data),
	}
}

func NewListBuilder[T any](values ...T) *ListBuilder[T] {
	return &ListBuilder[T]{values}
}
