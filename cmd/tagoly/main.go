package main

import (
	"fmt"
	"log"
	"os/exec"

	"tagoly/internal/config"
	"tagoly/internal/generator"
	"tagoly/internal/git"
	"tagoly/internal/prompt"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	// 1. 設定ファイル読み込み (.tagolyrc)
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// 2. ステージ済みファイル取得
	files, err := git.GetChangedFiles()
	if err != nil {
		log.Fatal(err)
	}

	// 3. スコープ検出
	scope, scopeList := generator.DetectScopeWithList(files)
	if scope == "multiple" {
		fmt.Println("Multiple scopes detected:", scopeList)
		// 複数の場合は選択させる
		promptScope := &survey.Select{
			Message:  "Select scope:",
			Options:  scopeList,
			PageSize: 10,
		}
		survey.AskOne(promptScope, &scope)
	}
	if scope == "root" {
		scope = ""
	}

	// 4. タグ選択（標準 + カスタム）
	tag := prompt.SelectCommitType(cfg.CustomTags)

	// 5. コミットメッセージ入力
	message := prompt.InputCommitMessage()

	// 6. コミットメッセージ生成
	var finalMessage string
	if scope != "" {
		finalMessage = fmt.Sprintf("%s(%s): %s", tag, scope, message)
	} else {
		finalMessage = fmt.Sprintf("%s: %s", tag, message)
	}

	// 7. 確認して git commit
	if prompt.ConfirmCommit(finalMessage) {
		cmd := exec.Command("git", "commit", "-m", finalMessage)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Committed!")
	} else {
		fmt.Println("Commit canceled.")
	}
}
