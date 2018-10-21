package main

import (
	"fmt"
	"os"
)

type point struct {
	i int
	j int
}

var dirs = [4]point{
	{-1, 0},
	{ 0, -1},
	{ 1, 0},
	{ 0, 1},
}

func readMaze(fileName string) [][]int {
	file, err := os.Open(fileName)
	if err != nil {
		panic("open file fail")
	}

	var row, col int
	var changeLine int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)

		fmt.Fscanf(file, "%d", &changeLine)

		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

func (p point) add(dir point) point {
	return point{p.i + dir.i, p.j + dir.j}
}

func (p point) at(m [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(m) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(m[0]) {
		return 0,false
	}

	return m[p.i][p.j],true
}

func walkMaze(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		for _, dir := range dirs {
			next := cur.add(dir)

			nextValue, ok := next.at(maze)
			if !ok || nextValue == 1 {
				continue
			}

			nextValue, ok = next.at(steps)
			if !ok || nextValue != 0 {
				continue
			}

			if next == start {
				continue
			}

			Q = append(Q,next)

			mazeValue, ok := cur.at(steps)
			if ok {
				steps[next.i][next.j] = mazeValue + 1
			}

		}
	}

	return steps
}

func main() {
	maze := readMaze("input.in")

	for _, value := range maze {
		fmt.Printf("%d\n", value)
	}

	steps := walkMaze(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _,value := range steps {
		fmt.Printf("%3d\n",value)
	}
}
