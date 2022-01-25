package branch_and_bound

/*
 * 分支限界法一般用于求解最优问题，其原理是列出所有可能“状态”，利用上下界来提高搜索效率。
 * 它类似于回溯法，但回溯法使用的是深度优先搜索，而分支限界法使用的是广度优先搜索。
 *
 * 使用分支限界法求解可以得到一个树形结构（状态树）。
 *
 * 			 root
 *		  /   ｜   \
 *	   Left	Middle Right
 *
 * 活节点：指下一步可以展开但还没有展开的节点。
 * 分支限界法最有代表性的是 FIFO 分支限界和 LC 分支限界，前者属于树的层次遍历(穷举)，后者是将深度优先和广度优先便利结合起来。
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
