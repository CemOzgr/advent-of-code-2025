package day8

type UnionFind struct {
	parent []int
	sizes  []int
}

func MakeSet(size int) UnionFind {
	parent := make([]int, size)
	sizes := make([]int, size)

	for i := range parent {
		parent[i] = i
		sizes[i] = 1
	}

	return UnionFind{parent, sizes}
}

func (unionFind *UnionFind) Find(x int) int {
	for x != unionFind.parent[x] {
		x = unionFind.parent[x]
	}

	return x
}

func (unionFind *UnionFind) Union(e1 int, e2 int) {
	var root1, root2 = unionFind.Find(e1), unionFind.Find(e2)

	if root1 == root2 {
		return
	}

	if unionFind.sizes[root1] >= unionFind.sizes[root2] {
		unionFind.parent[root2] = root1
		unionFind.sizes[root1] += unionFind.sizes[root2]
		unionFind.sizes[root2] = 0

		return
	}

	unionFind.parent[root1] = root2
	unionFind.sizes[root2] += unionFind.sizes[root1]
	unionFind.sizes[root1] = 0
}
