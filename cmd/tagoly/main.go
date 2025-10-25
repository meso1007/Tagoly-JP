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
	// 1. 設定読み込み
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
	scope, scopeList := generator.DetectScopeWithListImproved(files)
	if len(scopeList) > 1 {
		fmt.Println("Multiple scopes detected:", scopeList)
		promptScope := &survey.Select{
			Message:  "Select scope:",
			Options:  scopeList,
			PageSize: 10,
			Default:  scope,
		}
		survey.AskOne(promptScope, &scope)
	}

	if scope == "root" {
		scope = ""
	}

	// 4. タグ選択
	tag := prompt.SelectCommitType(cfg.CustomTags)

	// 5. メッセージ入力
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
