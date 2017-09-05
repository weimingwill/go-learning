package loadbalancer

import "fmt"

type Pool []*Worker

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p Pool) Swap(i, j int) {
	fmt.Printf("Swap(%d, %d) and pool length is %d\n", i, j, len(p))
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
	fmt.Printf("Swap(%d, %d) and pool length is %d\n", i, j, len(p))
}

func (p *Pool) Push(x interface{}) {
	n := len(*p)
	worker := x.(*Worker)
	worker.index = n
	*p = append(*p, worker)
}

func (p *Pool) Pop() interface{} {
	old := *p
	fmt.Printf("Pop item %d\n", len(old)-1)
	n := len(*p)
	worker := old[n-1]
	worker.index = -1
	*p = old[0 : n-1]
	return worker
}
