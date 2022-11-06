package sort

//MergeSort
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[0:middle]
	right := arr[middle:]
	return merge(MergeSort(left), MergeSort(right))
}

//将排好序的两个切片合并，例如1、3和2、4、5
func merge(left []int, right []int) []int {
	var result []int
	//for循环退出后result为1、2、3，left切片为空，right切片变成4、5
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	//之后拼接剩余切片内容
	for len(left) != 0 {
		result = append(result, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		result = append(result, right[0])
		right = right[1:]
	}

	return result
}

/*
场景应用：
设想你有一个20GB的文件，每行一个字符串，说明如何对这个文件进行排序。内存肯定没有20GB大，所以不可能采用传统排序法。
但是可以将文件分成许多块，每块xMB,针对每个快各自进行排序，存回文件系统。然后将这些块逐一合并，最终得到全部排好序的文件。

外排序的一个例子是外归并排序（External merge sort），它读入一些能放在内存内的数据量，
在内存中排序后输出为一个顺串（即是内部数据有序的临时文件），处理完所有的数据后再进行归并。

比如，要对900MB的数据进行排序，但机器上只有100MB的可用内存时，外归并排序按如下方法操作：
1、	读入100 MB的数据至内存中，用某种常规方式（如快速排序、堆排序、归并排序等方法）在内存中完成排序。
2、	将排序完成的数据写入磁盘。
3、	重复步骤1和2直到所有的数据都存入了不同的100 MB的块（临时文件）中。在这个例子中，有900 MB数据，单个临时文件大小为100 MB，
	所以会产生9个临时文件。
4、	读入每个临时文件（顺串）的前10 MB（ = 100 MB / (9块 + 1)）的数据放入内存中的输入缓冲区，最后的10 MB作为输出缓冲区。
   （实践中，将输入缓冲适当调小，而适当增大输出缓冲区能获得更好的效果。）
5、	执行九路归并算法，将结果输出到输出缓冲区。一旦输出缓冲区满，将缓冲区中的数据写出至目标文件，清空缓冲区。
   	一旦9个输入缓冲区中的一个变空，就从这个缓冲区关联的文件，读入下一个10M数据，除非这个文件已读完。
   	这是“外归并排序”能在主存外完成排序的关键步骤 -- 因为“归并算法”(merge algorithm)对每一个大块只是顺序地做一轮访问(进行归并)，
   	每个大块不用完全载入主存。
*/
