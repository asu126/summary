package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var len1, len2 int = len(nums1), len(nums2)
	var middle, m_v1, m_v2 int = (len1 + len2) / 2, 0, 0
	var i, j int = 0, 0

	for i < len1 && j < len2 {
		if i+j == middle {
			if nums1[i] <= nums2[j] {
				m_v2 = nums1[i]
				i++
			} else {
				m_v2 = nums2[j]
				j++
			}
			break
		}

		if nums1[i] <= nums2[j] {
			m_v1 = nums1[i]
			i++
		} else {
			m_v1 = nums2[j]
			j++
		}

	}

	for i < len1 && i+j <= middle {
		if i+j == middle {
			m_v2 = nums1[i]
			break
		}
		m_v1 = nums1[i]
		i++
	}

	for j < len2 && i+j <= middle {
		if i+j == middle {
			m_v2 = nums2[j]
			break
		}
		m_v1 = nums2[j]
		j++
	}

	if (len1+len2)%2 == 0 {
		return ((float64)(m_v1 + m_v2)) / 2.0
	} else {
		return float64(m_v2)
	}

}

func main() {
	a1 := []int{1, 3, 5}
	a2 := []int{2, 4}
	fmt.Println(findMedianSortedArrays(a1, a2))
}
