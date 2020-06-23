package ds

type Tree interface{
	GetLeft() Tree
	GetRight() Tree
	IsRootNode() bool
	GetValue() interface{}
	SetLeft(l Tree)
	SetRight(r Tree)
	Dfs() int
	Bfs() []Tree
	TreePrint() [][]interface{}
	GetLeaves() []Tree
}
