package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

var res []string

func generateParenthesis(n int) []string {
	res = make([]string, 0)
	dfs(n, 0, 0, "")
	return res
}

func dfs(n int, lc int, rc int, path string) {
	if lc == n && rc == n {
		res = append(res, path)
		return
	} else {
		if lc < n {
			dfs(n, lc + 1, rc, path + "(")
		}
		if rc < lc {
			dfs(n, lc, rc + 1, path + ")")
		}
	}

}

//func generateParenthesis(n int) []string {
//	res := []string{}
//	var dfs func(lRemain int, rRemain int, path string)
//	dfs = func(lMain, rMain int, path string) {
//		if 2*n == len(path) {
//			res = append(res, path)
//			return
//		}
//
//		if lMain > 0 {
//			dfs(lMain-1, rMain, path+"(")
//		}
//		if lMain < rMain {
//			dfs(lMain, rMain-1, path+")")
//		}
//	}
//
//	dfs(n, n, "")
//	return res
//}
