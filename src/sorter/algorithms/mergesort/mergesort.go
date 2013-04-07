//mergesort.go
package mergesort

func MergeSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}

	mid := len(values) / 2

	left := values[:mid]
	right := values[mid:]

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

//merge 操作，数组必须为有序数组
func merge(left, right []int) (sorted []int) {
	i := 0
	j := 0

	for (i < len(left)) && (j < len(right)) {
		if left[i] <= right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	if i <= len(left)-1 {
		for ; i < len(left); i++ {
			sorted = append(sorted, left[i])
		}
	} else {
		for ; j < len(right); j++ {
			sorted = append(sorted, right[j])
		}
	}

	return sorted
}

//func merge(left ,right []int)(sorted []int){
//	sorted=make([]int)
//	i,j:=0
//	for ;i<len(left)&&j<len(right);{
//		if left[i]<=right[j]{
//			sorted=append(sorted,left[i])
//			i++
//		}
//		else{
//			sorted=append(sorted,right[j])
//			j++
//		}
//	}

//	if(i<len(left)-1){
//		for;i<len(left);i++{
//			sorted=append(sorted,left[i])
//		}
//	}
//	else{
//		for ;j<len(right);j++{
//			sorted=append(sorted,right[j])
//		}
//	}
//	return sorted

//}
