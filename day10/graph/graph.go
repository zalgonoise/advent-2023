package graph

const minAlloc = 64

type graph[T comparable] interface {
	Root() T
	Edges(T) []T
	IsLast(T) bool
}

type DFS[T comparable] struct {
	Queue []T
	Cache map[T]struct{}
	Len   map[T]uint64
}

func NewDFS[T comparable]() *DFS[T] {
	return &DFS[T]{
		Queue: make([]T, 0, minAlloc),
		Cache: make(map[T]struct{}),
		Len:   make(map[T]uint64),
	}
}

func (b *DFS[T]) Run(g graph[T]) T {
	root := g.Root()
	b.Queue = append(b.Queue, root)
	b.Len[root] = 0

	for len(b.Queue) > 0 {
		s := b.Queue[len(b.Queue)-1]
		b.Queue = b.Queue[:len(b.Queue)-1]
		d := b.Len[s]

		if _, ok := b.Cache[s]; ok {
			continue
		}

		b.Cache[s] = struct{}{}

		edges := g.Edges(s)
		for _, edge := range edges {
			if _, ok := b.Cache[edge]; ok {
				continue
			}

			b.Len[edge] = d + 1

			if g.IsLast(edge) {
				return edge
			}

			b.Queue = append(b.Queue, edge)
		}
	}

	return root
}
