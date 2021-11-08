package dfs

func generateParenthesis(n int) []string {
	paths := make([]string, 0)

	var dfs func(path string, left, right int)
	dfs = func(path string, left, right int) {
		if left > right {
			return
		}

		if len(path) == 2*n {
			paths = append(paths, path)
			return
		}
		if left > 0 {
			dfs(path+"(", left-1, right)
		}
		if right > 0 {
			dfs(path+")", left, right-1)
		}
	}

	dfs("", n, n)
	return paths
}
