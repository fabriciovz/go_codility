package solution1

/*
type Nails struct {
	nails []*Nail
}

func (n *Nails) Len() int {
	return len(n.nails)
}

func (n *Nails) Swap(i, j int) {
	n.nails[i], n.nails[j] = n.nails[j], n.nails[i]
}

//From the above discussion, if we check the planks sorted in the ascending order by their end positions,
//it is clear that we can forget any nail that appears before the beginning position of the plank that we are currently checking.
func (n *Nails) Less(i, j int) bool {
	return n.nails[i].value < n.nails[j].value
}

func (n *Nails) Add(nail *Nail) {
	n.nails = append(n.nails, nail)
}

func (n *Nails) Get(i int) *Nail {
	return n.nails[i]
}
*/
