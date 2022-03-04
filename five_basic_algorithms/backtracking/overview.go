package backtracking

/*
 * 回溯法是一种暴力搜索的方式，与递归结合十分紧密，一般会在递归函数结束后进行回溯，例如：
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

/*
 * 关于递归中是否需要返回值的问题，以有关树的问题举例：
 * 1. 如果要搜索树中**一条**符合条件的路径，那么递归一定需要返回值，因为遇到符合条件的路径了就要及时返回（LeetCode 112）。
 * 2. 如果需要搜索整棵树且无须处理返回值，那么递归函数就不需要返回值（LeetCode 113）。
 * 3. 如果需要搜索整棵树且需要处理返回值，那么递归函数就需要返回值（LeetCode 236）。
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
