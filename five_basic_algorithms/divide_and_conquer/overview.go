package divide_and_conquer

/*
 * 分治法十分类似于回溯，与递归结合十分紧密。
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
 * 3. 划分问题规模，进行递归。
 */

func DivideAndConquer() {
	// Simple problems, no divide and conquer.
	// Usually the end of the recursion.
	// Small problems can be solved directly.
	if true {
		// Collect results.
		return
	}
	// Divide origin problem into smaller subproblems.
	for false {
		// Recursive function.
		DivideAndConquer()
	}
}
