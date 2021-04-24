package main

import (
	"math"
)

func quickSort(_intervals [][]int) [][]int {
	if len(_intervals) <= 1 {
		return _intervals
	}

	intervals := _intervals[:]
	pivot := 0
	l := 0
	r := len(intervals) - 1
	// fmt.Print("start", intervals, "\n")

	for l < r {
		// pivot left
		for index := r; index >= l; index-- {
			val := intervals[index][0]
			pivotValue := intervals[pivot][0]

			r = index
			if val < pivotValue {
				intervals[index], intervals[pivot] = intervals[pivot], intervals[index]

				pivot = index
				break
			}
		}

		// pivot right
		if l+1 >= r {
			break
		}
		for index := l; index <= r; index++ {
			val := intervals[index][0]
			pivotValue := intervals[pivot][0]

			l = index
			if val > pivotValue {
				intervals[index], intervals[pivot] = intervals[pivot], intervals[index]

				pivot = index
				break
			}
		}

	}

	// fmt.Print("intervals: ", intervals, pivot, "\n\n")
	leftGroup := quickSort(intervals[:pivot])
	rightGroup := [][]int{}

	if pivot+1 < len(intervals) {
		rightGroup = quickSort(intervals[pivot+1:])
	}

	group := [][]int{intervals[pivot]}
	group = append(leftGroup, group...)

	group = append(group, rightGroup...)

	return group
}

func merge(_intervals [][]int) [][]int {
	intervalsSorted := quickSort(_intervals)

	intervals := [][]int{}
	interval := intervalsSorted[0]

	for i := 1; i < len(intervalsSorted); i++ {
		curInterval := intervalsSorted[i]

		if interval[1] < curInterval[0] {
			intervals = append(intervals, interval)
			interval = curInterval
		} else {
			interval[0] = int(math.Min(float64(interval[0]), float64(curInterval[0])))
			interval[1] = int(math.Max(float64(interval[1]), float64(curInterval[1])))
		}
	}

	intervals = append(intervals, interval)

	return intervals
}

func main() {
	// intervals := [][]int{{1, 4}, {2, 2}, {5, 8}, {6, 7}, {3, 5}, {0, 0}}
	intervals := [][]int{{1, 4}, {0, 0}}

	// intervalsSorted := quickSort(intervals)
	// fmt.Print(intervalsSorted)

	merge(intervalsSorted)
}
