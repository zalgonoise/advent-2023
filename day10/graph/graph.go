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

func (d *DFS[T]) Run(g graph[T]) T {
	root := g.Root()
	d.Queue = append(d.Queue, root)
	d.Len[root] = 0

	for len(d.Queue) > 0 {
		s := d.Queue[len(d.Queue)-1]
		d.Queue = d.Queue[:len(d.Queue)-1]
		n := d.Len[s]

		if _, ok := d.Cache[s]; ok {
			continue
		}

		d.Cache[s] = struct{}{}

		edges := g.Edges(s)
		for _, edge := range edges {
			if _, ok := d.Cache[edge]; ok {
				continue
			}

			d.Len[edge] = n + 1

			if g.IsLast(edge) {
				return edge
			}

			d.Queue = append(d.Queue, edge)
		}
	}

	return root
}
