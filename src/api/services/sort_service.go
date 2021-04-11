package services

import "github.com/davidandradeduarte/uif-testing-in-golang-udemy/src/api/utils/sort"

const (
	privateConst = "private"
	PublicConst  = "public"
)

func Sort(elements []int) {
	//sort.BubbleSort(elements)
	if len(elements) <= 1000 {
		sort.Sort(elements)
		return
	}
	sort.Sort(elements)
}
