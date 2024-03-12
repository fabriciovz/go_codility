package solution1

import (
	"fmt"
	"sort"
)

// 62% https://app.codility.com/demo/results/trainingS93D2V-C5S/
func Solution(A []int, B []int, C []int) int {
	if !IsValidInput(A, B, C) {
		return -1
	}
	var nails Nails
	var nailsSeen = make(map[int]int, 0)
	for i := 0; i < len(C); i++ {
		_, ok := nailsSeen[C[i]]
		if !ok {
			nail := NewNail(i+1, C[i])
			nails.Add(nail)
			nailsSeen[C[i]] = i
		}
	}
	sort.Sort(&nails)
	var planks Planks
	for i := 0; i < len(A); i++ {
		plank := NewPlank(A[i], B[i])
		planks.Add(plank)
	}
	sort.Sort(&planks)
	for i := 0; i < planks.Len()-2; i++ {
		for i > 0 && planks.Get(i).end > planks.Get(i+1).end {
			i--
			planks.Remove(i)
		}
	}
	maxMin := 0
	for i := range planks.GetAll() {
		minIndex := GetMinNailIndex(nails, planks.GetAll()[i])
		if minIndex == -1 {
			return -1
		}
		maxMin = Max(maxMin, minIndex)
	}

	return maxMin
}

func GetMinNailIndex(nails Nails, plank *Plank) int {
	start := 0
	end := nails.Len() - 1
	result := -1
	for start <= end {
		center := (start + end) / 2
		if nails.Get(center).value < plank.start {
			start = center + 1
		} else if nails.Get(center).value > plank.end {
			end = center - 1
		} else {
			result = center
			end = center - 1
		}
	}
	if result == -1 {
		return -1
	}
	minIndex := nails.Get(result).index
	for i := result + 1; i < nails.Len(); i++ {
		if nails.Get(i).value <= plank.end {
			minIndex = Min(minIndex, nails.Get(i).index)
		} else {
			return minIndex
		}
	}
	return minIndex
}

func Min(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func Max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

type Nails struct {
	nails []*Nail
}

func (n *Nails) Len() int {
	return len(n.nails)
}

func (n *Nails) Swap(i, j int) {
	n.nails[i], n.nails[j] = n.nails[j], n.nails[i]
}

func (n *Nails) Less(i, j int) bool {
	return n.nails[i].value < n.nails[j].value
}

func (n *Nails) Add(nail *Nail) {
	n.nails = append(n.nails, nail)
}

func (n *Nails) Get(i int) *Nail {
	return n.nails[i]
}

type Nail struct {
	index int
	value int
}

func NewNail(index, value int) *Nail {
	return &Nail{index: index, value: value}
}

func (n *Nail) ToString() {
	fmt.Printf("Nail [index=%d, value=%d]", n.index, n.value)
}

type Planks struct {
	planks []*Plank
}

func (n *Planks) Len() int {
	return len(n.planks)
}

func (n *Planks) Swap(i, j int) {
	n.planks[i], n.planks[j] = n.planks[j], n.planks[i]
}

func (n *Planks) Less(i, j int) bool {
	return n.planks[i].start < n.planks[j].start
}

func (n *Planks) Remove(i int) {
	n.planks = append(n.planks[:i], n.planks[i+1:]...)
}

func (n *Planks) Add(plank *Plank) {
	n.planks = append(n.planks, plank)
}

func (n *Planks) Get(i int) *Plank {
	return n.planks[i]
}

func (n *Planks) GetAll() []*Plank {
	return n.planks
}

type Plank struct {
	start int
	end   int
}

func NewPlank(start int, end int) *Plank {
	return &Plank{start: start, end: end}
}

func (p *Plank) ToString() {
	fmt.Printf("Plank [start=%d, end=%d]", p.start, p.end)
}

func IsValidInput(A []int, B []int, C []int) bool {
	lenA := len(A)
	lenB := len(B)
	lenC := len(C)

	if !(lenA == lenB) {
		return false
	}
	if lenA < 1 || lenA > 30000 {
		return false
	}
	if lenC < 1 || lenC > 30000 {
		return false
	}
	for i := 0; i < lenA; i++ {
		if !(A[i] <= B[i]) {
			return false
		}
		if A[i] < 1 || A[i] > 2*lenC {
			return false
		}
		if B[i] < 1 || B[i] > 2*lenC {
			return false
		}
	}
	for i := 0; i < lenC; i++ {
		if C[i] < 1 || C[i] > 2*lenC {
			return false
		}
	}
	return true
}

/*
func Solution(A []int, B []int, C []int) int {
	if !IsValidInput(A, B, C) {
		return -1
	}
	counter := 0
	planks := map[int]int{}
	maxPlankNailed := 0
	for i := 0; i < len(C); i++ {
		nail := C[i]
		for j := 0; j < len(A); j++ {
			limLow := A[j]
			limHi := B[j]
			if nail >= limLow && nail <= limHi {
				if j < maxPlankNailed {
					return -1
				}
				if len(planks) == 0 {
					planks[j] = i
					counter++
					maxPlankNailed = j
				} else if len(planks) < j+1 {
					planks[j] = i
					counter++
					maxPlankNailed = j
				} else if len(planks) == j+1 {
					planks[j] = i
					maxPlankNailed = j
					break
				} else {
					return -1
				}
			}
		}
		if counter == len(A) {
			break
		}
	}
	if counter == 0 {
		return -1
	}
	return counter
}*/
