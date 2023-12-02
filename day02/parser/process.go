package parser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/zalgonoise/parse"
)

const minAlloc = 16

func Parse(input string) ([]Game, error) {
	split := strings.Split(input, "\n")

	games := make([]Game, 0, len(split))
	errs := make([]error, 0, len(split))

	for i := range split {
		if split[i] == "" {
			continue
		}

		game, err := parse.Run([]byte(split[i]), StateFunc, ParseFunc, ProcessFunc)
		if err != nil {
			errs = append(errs, err)

			continue
		}

		games = append(games, game)
	}

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return games, nil
}

func ProcessFunc(t *parse.Tree[Token, byte]) (Game, error) {
	if err := Validate(t); err != nil {
		return Game{}, nil
	}

	nodes := t.List()
	game := Game{
		Cubes: make([]Cubes, 0, minAlloc),
	}

	id, err := strconv.Atoi(string(nodes[0].Edges[0].Value))
	if err != nil {
		return game, err
	}

	game.ID = id

	// iterate through each set of Cubes
	for i := range nodes[0].Edges[1].Edges {
		cubes := &Cubes{}
		node := nodes[0].Edges[1].Edges[i]

		switch node.Type {
		case TokenNum:
			// OK as-is
		case TokenSemicolon:
			node = node.Edges[0]
		}

		if err = processCube(cubes, node); err != nil {
			return game, err
		}

		game.Cubes = append(game.Cubes, *cubes)
	}

	return game, nil
}

func processCube(cubes *Cubes, n *parse.Node[Token, byte]) error {
	var num int
	var err error

	if num, err = strconv.Atoi(string(n.Value)); err != nil {
		return err
	}

	if len(n.Edges) > 0 && n.Edges[0].Type == TokenAlpha {
		switch strings.ToLower(string(n.Edges[0].Value)) {
		case "red":
			cubes.Red = num
		case "green":
			cubes.Green = num
		case "blue":
			cubes.Blue = num
		}
	}

	if len(n.Edges) > 1 {
		errs := make([]error, 0, len(n.Edges)-1)

		for i := 1; i < len(n.Edges); i++ {
			if n.Edges[i].Type == TokenComma && len(n.Edges[i].Edges) == 1 {
				if innerErr := processCube(cubes, n.Edges[i].Edges[0]); innerErr != nil {
					errs = append(errs, err)
				}
			}
		}

		err = errors.Join(errs...)
	}

	return err
}

var (
	errEmptyNodes         = errors.New("empty nodes list")
	errUnexpectedNumEdges = errors.New("unexpected number of edges")
	errUnexpectedType     = errors.New("unexpected token type")
)

func Validate(t *parse.Tree[Token, byte]) error {
	nodes := t.List()

	if len(nodes) == 0 {
		return errEmptyNodes
	}

	if len(nodes[0].Edges) != 2 {
		return fmt.Errorf("%w: nodes[0].Edges -> len %d", errUnexpectedNumEdges, len(nodes[0].Edges))
	}

	if nodes[0].Edges[0].Type != TokenNum {
		return fmt.Errorf("%w: nodes[0].Edges[0].Type: %v", errUnexpectedType, nodes[0].Edges[0].Type)
	}

	if nodes[0].Edges[1].Type != TokenColon {
		return fmt.Errorf("%w: nodes[0].Edges[1].Type: %v", errUnexpectedType, nodes[0].Edges[1].Type)
	}

	if len(nodes[0].Edges[1].Edges) < 1 {
		return fmt.Errorf("%w: nodes[0].Edges[1].Edges -> len %d",
			errUnexpectedNumEdges, len(nodes[0].Edges[1].Edges),
		)
	}

	return nil
}
