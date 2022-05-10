package main

// func main() {
// 	n, m := 0, 0
// 	fmt.Scanf("%d %d\n", &n, &m)

// 	fmt.Println(solution(n, m))
// }

// 获得字典树的某个分支有多少个节点。
func getCountOfPre(pre int, n int) int {
	return 0
}

func solution(n, m int) int {
	// 从1开始计算。
	pre := 1
	for m != 0 {
		count := getCountOfPre(pre, n)
		if count >= m {
			// 搜索分支节点的m-1个节点。
			m--
			if m == 0 {
				break
			}
			// 进入前缀树的下一层。
			pre = pre * 10
		} else {
			m -= count
			// 在同一层的另一个分支。
			pre++
		}
	}

	return pre
}
