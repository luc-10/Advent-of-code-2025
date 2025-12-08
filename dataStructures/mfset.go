package dataStructures

type Mfset struct {
	parent []int
	rank   []int
	sets   int
}

func NewMfset(n int) *Mfset {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &Mfset{parent, rank, n}
}

func (mfs *Mfset) Find(x int) int {
	if mfs.parent[x] != x {
		mfs.parent[x] = mfs.Find(mfs.parent[x])
	}
	return mfs.parent[x]
}

func (mfs *Mfset) Merge(x, y int) {
	rx := mfs.Find(x)
	ry := mfs.Find(y)
	if rx != ry {
		if mfs.rank[rx] > mfs.rank[ry] {
			mfs.parent[ry] = rx
		} else if mfs.rank[ry] > mfs.rank[rx] {
			mfs.parent[rx] = ry
		} else {
			mfs.parent[rx] = ry
			mfs.rank[ry] = mfs.rank[ry] + 1
		}
		mfs.sets--
	}
}

func (mfs *Mfset) CountSets() int {
	return mfs.sets
}
