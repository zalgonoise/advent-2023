package graph

type graph[T comparable] interface {
	Root() T
	Edges(T) []T
	IsLast(T) bool
}

type BFS[T comparable] struct {
	Queue    []T
	Cache    map[T]struct{}
	Previous map[T]T
	Len      map[T]uint64
}

func NewBFS[T comparable]() *BFS[T] {
	return &BFS[T]{
		Queue:    []T{},
		Cache:    make(map[T]struct{}),
		Previous: make(map[T]T),
		Len:      make(map[T]uint64),
	}
}

func (b *BFS[T]) Run(g graph[T]) T {
	root := g.Root()
	b.Queue = append(b.Queue, root)
	b.Len[root] = 0

	for len(b.Queue) > 0 {
		cur := b.Queue[0]
		b.Queue = b.Queue[1:]
		count := b.Len[cur]

		if _, ok := b.Cache[cur]; ok {
			continue
		}

		b.Cache[cur] = struct{}{}

		edges := g.Edges(cur)
		for i := range edges {
			if _, ok := b.Cache[edges[i]]; ok {
				continue
			}

			b.Len[edges[i]] = count + 1
			b.Previous[edges[i]] = cur

			if g.IsLast(edges[i]) {
				return edges[i]
			}

			b.Queue = append(b.Queue, edges[i])
		}
	}

	return root
}
