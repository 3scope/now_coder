package main

func Git(matrix []string, versionA int, versionB int) int {
	// write code here
	usedIndex := make(map[int]bool)
	return PreorderTraversal(matrix, &usedIndex, 0, versionA, versionB)
}

func PreorderTraversal(matrix []string, usedIndex *map[int]bool, root, versionA, versionB int) int {
	if root >= len(matrix) {
		return -1
	}

	if root == versionA || root == versionB {
		return root
	}

	(*usedIndex)[root] = true
	subResult := -1
	// To Store the count of finding nodes.
	count := 0
	children := matrix[root]
	for i := 0; i < len(children); i++ {
		if (*usedIndex)[i] {
			continue
		}
		if children[i] == '1' {
			flag := PreorderTraversal(matrix, usedIndex, i, versionA, versionB)
			if flag != -1 {
				subResult = flag
				count++
			}
		}
	}
	if count == 1 {
		return subResult
	} else if count == 2 {
		return root
	} else {
		return -1
	}
}
