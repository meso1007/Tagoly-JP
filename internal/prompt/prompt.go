package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// コミットタイプの構造体
type CommitType struct {
	Key   string
	Label string
}

// 選択肢リスト
var CommitTypes = []CommitType{
	{"feat", "New feature"},
	{"fix", "Bug fix"},
	{"docs", "Documentation"},
	{"refactor", "Code improvement"},
	{"style", "UI / CSS"},
	{"chore", "Maintenance"},
}

// コミットタイプを選択
func SelectCommitType() string {
	options := []string{}
	for _, c := range CommitTypes {
		options = append(options, fmt.Sprintf("%s (%s)", c.Key, c.Label))
	}

	var selected string
	prompt := &survey.Select{
		Message: "🔧 Select commit type:",
		Options: options,
	}
	survey.AskOne(prompt, &selected)

	// "feat (New feature)" → "feat" 部分だけ抽出
	for _, c := range CommitTypes {
		if selected == fmt.Sprintf("%s (%s)", c.Key, c.Label) {
			return c.Key
		}
	}
	return "feat"
}

// メッセージ入力
func InputCommitMessage() string {
	var message string
	prompt := &survey.Input{
		Message: "📝 Enter commit message:",
	}
	survey.AskOne(prompt, &message)
	return message
}

// 確認
func ConfirmCommit(final string) bool {
	var confirm bool
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("💬 Result: %s\n✅ Commit this message?", final),
		Default: true,
	}
	survey.AskOne(prompt, &confirm)
	return confirm
}
