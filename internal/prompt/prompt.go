package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

type CommitType struct {
	Key   string
	Label string
}

// 標準タグ
var defaultCommitTypes = []CommitType{
	{"feat", "New feature"},
	{"fix", "Bug fix"},
	{"docs", "Documentation"},
	{"refactor", "Code improvement"},
	{"style", "UI / CSS"},
	{"chore", "Maintenance"},
}

// カスタムタグ対応
func SelectCommitType(customTags []string) string {
	options := []string{}

	for _, c := range defaultCommitTypes {
		options = append(options, fmt.Sprintf("%s (%s)", c.Key, c.Label))
	}
	for _, ct := range customTags {
		options = append(options, fmt.Sprintf("%s (Custom)", ct))
	}

	var selected string
	prompt := &survey.Select{
		Message:  "🔧 Select commit type:",
		Options:  options,
		PageSize: 10,
	}
	survey.AskOne(prompt, &selected)

	for _, c := range defaultCommitTypes {
		if selected == fmt.Sprintf("%s (%s)", c.Key, c.Label) {
			return c.Key
		}
	}
	for _, ct := range customTags {
		if selected == fmt.Sprintf("%s (Custom)", ct) {
			return ct
		}
	}

	return "feat"
}

func InputCommitMessage() string {
	var message string
	prompt := &survey.Input{
		Message: "📝 Enter commit message:",
	}
	survey.AskOne(prompt, &message)
	return message
}

func ConfirmCommit(final string) bool {
	var confirm bool
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("💬 Result: %s\n✅ Commit this message?", final),
		Default: true,
	}
	survey.AskOne(prompt, &confirm)
	return confirm
}
