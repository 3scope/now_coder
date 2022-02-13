package divide_and_conquer

/*
 * 分治法类似于回溯法，与递归结合十分紧密。
 *
 * 使用分治法求解可以得到一个N叉树，递归树的度取决于分治法每次所分的子问题数量。
 *
 * 			 root
 *		  /   ｜   \
 *	   Left	Middle Right
 *
 * 使用分治法的三个步骤：
 * 1. 确定递归函数的参数以及返回值。
 * 2. 确定终止条件，对简单问题进行处理。
 * 3. 写程序划分原参数的相应子集，进行递归（将结果合并[可选]）。
 */

func DivideAndConquer() {
	// Simple problems, no divide and conquer.
	// Usually the end of the recursion.
	// Small problems can be solved directly.
	if true {
		// Directly solve.
		return
	}
	// To divide into subset.
	// Divide origin problem into smaller subproblems.
	for {
		// Recursive function.
		DivideAndConquer()
		return
	}
}

func Dichotomy() {
	// Small problems can be solved directly.
	if true {
		// Directly solve.
		return
	} else {
		Dichotomy() // [Left, Middle]
		Dichotomy() // [Middle, Right]
		return
	}
}
