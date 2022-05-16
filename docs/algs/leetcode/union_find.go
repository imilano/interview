package leetcode


type UnionFind struct {

}

func (this *UnionFind) union(parent []int, p, q int) {
	xp := this.find(parent, p)
	xq := this.find(parent, q)
	if xp == xq {
		return
	}

	parent[p] = xp
}


func (this *UnionFind) find(parent []int, q int) int {
	if parent[q] != -1 {
		q = parent[q]
	}

	return q
}
