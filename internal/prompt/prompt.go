package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// コミットタイプ構造体
type CommitType struct {
	Key   string
	Label string
}

// 標準のコミットタイプ
var defaultCommitTypes = []CommitType{
	{"feat", "新しい機能"},
	{"fix", "バグ修正"},
	{"docs", "ドキュメント"},
	{"refactor", "コード改善"},
	{"style", "UI / スタイル修正"},
	{"chore", "メンテナンス"},
}

// カスタムタグにも対応したコミットタイプ選択
func SelectCommitType(customTags []string) string {
	options := []string{}

	// 標準タグ
	for _, c := range defaultCommitTypes {
		options = append(options, fmt.Sprintf("%s (%s)", c.Key, c.Label))
	}

	// カスタムタグ
	for _, ct := range customTags {
		options = append(options, fmt.Sprintf("%s (カスタム)", ct))
	}

	var selected string
	prompt := &survey.Select{
		Message:  "🔧 コミットタイプを選択してください:",
		Options:  options,
		PageSize: 10,
	}
	survey.AskOne(prompt, &selected)

	// 選択結果をキーとして返す
	for _, c := range defaultCommitTypes {
		if selected == fmt.Sprintf("%s (%s)", c.Key, c.Label) {
			return c.Key
		}
	}
	for _, ct := range customTags {
		if selected == fmt.Sprintf("%s (カスタム)", ct) {
			return ct
		}
	}

	return "feat"
}

// コミットメッセージ入力
func InputCommitMessage() string {
	var message string
	prompt := &survey.Input{
		Message: "📝 コミットメッセージを入力してください:",
	}
	survey.AskOne(prompt, &message)
	return message
}

// 確認ダイアログ
func ConfirmCommit(final string) bool {
	var confirm bool
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("💬 作成されたメッセージ:\n%s\n\n✅ この内容でコミットしますか？", final),
		Default: true,
	}
	survey.AskOne(prompt, &confirm)
	return confirm
}
