package grid

import (
	"container/heap"
)

type Numeric interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64
}

type Graph[T any] struct {
	Head Coord
	Tail Coord

	Map Map[T]
}

func (g Graph[T]) Root() Coord {
	return g.Head
}

func (g Graph[T]) Edges(c Coord) []Coord {
	edges := make([]Coord, 0, 4)

	north := Add(c, North)
	south := Add(c, South)
	east := Add(c, East)
	west := Add(c, West)

	if _, ok := g.Map.Items[north]; ok {
		edges = append(edges, north)
	}
	if _, ok := g.Map.Items[south]; ok {
		edges = append(edges, south)
	}
	if _, ok := g.Map.Items[east]; ok {
		edges = append(edges, east)
	}
	if _, ok := g.Map.Items[west]; ok {
		edges = append(edges, west)
	}

	return edges
}

func (g Graph[T]) IsLast(c Coord) bool {
	return c == g.Tail
}

type context struct {
	coord  Coord
	dir    Coord
	streak int
}

func AStar[T Numeric](g Graph[T], minSteps int, maxSteps int) int {
	var (
		startCtx = context{
			coord:  g.Head,
			dir:    Coord{},
			streak: 0,
		}
		pathQueue = &PriorityQueue[context]{}
		prevCtx   = make(map[context]context)
		heatLoss  = make(map[context]int)
	)

	heap.Init(pathQueue)
	heap.Push(pathQueue, QueueItem[context]{
		Item:     startCtx,
		Priority: 0,
	})

	prevCtx[startCtx] = startCtx
	heatLoss[startCtx] = 0

	for pathQueue.Len() > 0 {
		cur := heap.Pop(pathQueue).(QueueItem[context]).Item
		curLoss := heatLoss[cur]

		if g.IsLast(cur.coord) {
			return curLoss
		}

		edges := g.Edges(cur.coord)
		for i := range edges {
			dir := Sub(edges[i], cur.coord)
			streak := 1

			if dir == cur.dir {
				streak += cur.streak
			}

			nextCtx := context{
				coord:  edges[i],
				dir:    dir,
				streak: streak,
			}

			nextLoss := curLoss + int(g.Map.Items[edges[i]])
			if loss, ok := heatLoss[nextCtx]; ok && nextLoss >= loss {
				continue
			}

			if cur.streak < minSteps && dir != cur.dir && cur.coord != g.Head {
				continue
			}

			if streak > maxSteps {
				continue
			}

			if dir == Inverse(cur.dir) {
				continue
			}

			heatLoss[nextCtx] = nextLoss
			prevCtx[nextCtx] = cur

			priority := nextLoss + Manhattan(edges[i], g.Tail)
			queueItem := QueueItem[context]{Item: nextCtx, Priority: priority}
			heap.Push(pathQueue, queueItem)
		}
	}

	return -1
}
