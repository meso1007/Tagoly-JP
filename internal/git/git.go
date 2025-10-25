package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ステージされた変更があるか確認する
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

// 実際に `git commit` コマンドを実行する
func Commit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Stdout = nil
	cmd.Stderr = nil
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("コミットに失敗しました: %v", err)
	}
	return nil
}

// ステージされているファイル一覧を取得する
func GetChangedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--name-only", "--cached")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("変更ファイルの取得に失敗しました: %v", err)
	}
	files := strings.Split(strings.TrimSpace(string(out)), "\n")

	// 空の場合は空スライスを返す
	if len(files) == 1 && files[0] == "" {
		return []string{}, nil
	}

	return files, nil
}
