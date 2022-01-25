package backtracking

/*
 * 回溯与递归结合十分紧密，一般会在递归函数结束后进行回溯，例如：
 *
 * recursive()
 * // 回溯代码
 *
 * 具体的应用有：
 * - 子集的排列 / 组合
 * - 切割为满足要求的子集
 * - 求不同的子集
 * - 棋盘问题（八皇后）
 *
 * 回溯法都可以抽象成为一个树形结构（状态树）：
 *
 * 		  root
 *	    /      \
 *	Left  ...  Right
 *
 * 使用回溯法的三个步骤：
 * 1. 确定递归函数的参数以及返回值。
 * 2. 确定终止条件。
 * 3. 单层递归处理逻辑。
 */

// Generic code template.
func Backtracking() {
	// Termination condition.
	if true {
		// Collect results.
		return
	}
	// Iterate over a subset of the collection.
	for false {
		// Deal with subset.
		// Recursive function.
		// Backtracking code.
	}
}
