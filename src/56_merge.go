package main

import (
	"fmt"
	"sort"
)

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

type ByAge []Interval

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Start < a[j].Start }

func getMaxAndMin(a, b Interval) (c Interval) {
	c = a
	if b.Start < a.Start {
		c.Start = b.Start
	}

	if b.End > a.End {
		c.End = b.End
	}
	return c
}

func merge(intervals []Interval) []Interval {
	sort.Sort(ByAge(intervals))
	var length int = len(intervals)

	var result []Interval = make([]Interval, 0, length)
	var flag []bool = make([]bool, length)

	var tmp Interval
	for i := 0; i < length; i++ {
		if flag[i] {
			continue
		}
		tmp = intervals[i]
		for j := i + 1; j < length; j++ {
			if !flag[j] {
				if !((intervals[j].Start > tmp.End) || (intervals[j].End < tmp.Start)) {
					tmp = getMaxAndMin(tmp, intervals[j])
					flag[j] = true
				}
			}
		}
		result = append(result, tmp)
		// fmt.Println(tmp.Start, tmp.End)
	}
	return result
}

func main() {
	var a1 []Interval = []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(a1))
	fmt.Println("--------------")

	var a2 []Interval = []Interval{{1, 8}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(a2))
	fmt.Println("--------------")
	var a3 []Interval = []Interval{{1, 8}}
	fmt.Println(merge(a3))
	fmt.Println("--------------")

	var a4 []Interval = []Interval{{1, 4}, {0, 0}}
	fmt.Println(merge(a4))
	fmt.Println("--------------")
	// [[1,4],[0,0]]

	// [[2,3],[4,5],[6,7],[8,9],[1,10]]
	var a5 []Interval = []Interval{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}

	// fmt.Println(a5)
	fmt.Println(merge(a5))
	fmt.Println("--------------")

}
