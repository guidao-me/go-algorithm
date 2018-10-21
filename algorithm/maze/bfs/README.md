**走迷宫问题**

迷宫是许多小方格构成的矩形，在每个小方格中有的是墙（用1表示），有的是路（用0表示）。走迷宫就是从一个小方格沿上、下、左、右四个方向到邻近的方格，当然不能穿墙。

**问题描述**
实现生成迷宫（从文本中读取数据生成），输出迷宫（图案方式），探索迷宫路径（最短路径），输出迷宫路径（图案方式)。

输入的文件内容格式为：
```
6 5
0 1 0 1 1
0 0 1 1 1
1 0 0 1 1
0 1 0 0 1
1 1 1 0 0
1 1 1 1 0

开头一行表示地图为6行5列，地图中0表示路，1表示障碍物。现在规定从左上角进入迷宫从右下角出迷宫。
```

**算法分析**
在任意位置，理论上有四个位置可以走，上下左右。现在位于位置0，走一步可以走的位置为：
![](https://i.imgur.com/FQB6hDj.png)

也就是当前位置为0，可以走的位置为1，相同道理，走2步可以走的位置为：
![](https://i.imgur.com/BMWOnz6.png)

当位于0时，可以探索的位置为周围的四个1，广度优先的算法是根据状态可分为两种状态，当前位置和下一步可以为的位置，当前位置为0，下一步可以为的位置为1，通过队列将下一步位置先保存起来，遍历下一步可以为的4个1位置，再分别根据位于4个1位置时探索下一步的位置。如果当位于0位置找到一个下一步位置1之后再直接进行下一步探索则成了深度优先算法。

以起始点为出发，将这个点按照某个方向规则走，周围所有可以走的位置依次写入到队列，这个点走完之后，队首出队列，走这个点四周所有可以走的点，按照之前的规则，将周围可走的点依次放入到队列中，重复之前过程，直到队列为空或者走到终点。在这个过程中在借助一个辅助矩阵，存放起点到位置的最短路径。

**源码分析**
```go
//坐标结构
type point struct {
	i int
	j int
}

//分别代表上左下右
var dirs = [4]point{
	{-1, 0},
	{ 0, -1},
	{ 1, 0},
	{ 0, 1},
}

//读取地图文件
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
		
		//读取文件过程中遇到换行读取为0
		fmt.Fscanf(file, "%d", &changeLine)

		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

func walkMaze(maze [][]int, start, end point)([][]int,bool) {
	steps := make([][]int, len(maze))

	for i := range steps {
		steps[i] = make([]int, len(maze[0]))
	}
	
	//保存四周下一步可以出现位置
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		for _, dir := range dirs {
			next := cur.add(dir)
			//1代表障碍物
			nextValue, ok := next.at(maze)
			if !ok || nextValue == 1 {
				continue
			}
			//下一步必须是还没走过的位置
			nextValue, ok = next.at(steps)
			if !ok || nextValue != 0 {
				continue
			}
			//不能走到其实位置
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
	//当走到终点时，则终点位置会被赋值为走过的步数
	if steps[len(steps)-1][len(steps[0])-1] == 0 {
		return steps,false
	}

	return steps,true
}
```

<a href="https://github.com/guidao-me/go-algorithm/blob/master/algorithm/maze/bfs/main.go">完整源码</a>