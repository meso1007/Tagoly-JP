// internal/generator/scope.go
package generator

import "strings"

// DetectScopeWithList は複数スコープのリストも返す
func DetectScopeWithList(files []string) (string, []string) {
	counts := map[string]int{}
	for _, f := range files {
		parts := strings.Split(f, "/")
		var scope string
		if len(parts) < 2 {
			scope = "root"
		} else {
			switch parts[0] {
			case "internal":
				scope = parts[1]
			case "cmd":
				scope = parts[1]
			default:
				scope = parts[0]
			}
		}
		counts[scope]++
	}

	// スコープを配列に
	scopes := []string{}
	for s := range counts {
		scopes = append(scopes, s)
	}

	// 単一スコープならそのまま、複数なら "multiple" とスコープ一覧
	if len(scopes) == 1 {
		return scopes[0], scopes
	}
	return "multiple", scopes
}
