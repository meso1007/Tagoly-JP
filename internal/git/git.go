package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ステージされているファイルがあるか確認
func HasStagedChanges() bool {
	cmd := exec.Command("git", "diff", "--cached", "--name-only")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return false
	}
	return out.Len() > 0
}

// 実際にgit commitを実行
func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to commit: %v", err)
	}
	return nil
}
func GetChangedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", "--cached")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	files := strings.Split(strings.TrimSpace(string(out)), "\n")
	return files, nil
}
