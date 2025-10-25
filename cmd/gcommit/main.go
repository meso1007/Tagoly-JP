package main

import (
	"fmt"
	"gcommit/internal/git"
	"gcommit/internal/prompt"
)

func main() {
	if !git.HasStagedChanges() {
		fmt.Println("⚠️  No staged changes found. Use 'git add' first.")
		return
	}

	commitType := prompt.SelectCommitType()
	message := prompt.InputCommitMessage()
	final := fmt.Sprintf("%s: %s", commitType, message)

	if prompt.ConfirmCommit(final) {
		err := git.Commit(final)
		if err != nil {
			fmt.Println("❌ Commit failed:", err)
			return
		}
		fmt.Println("✅ Commit done!")
	} else {
		fmt.Println("❌ Commit cancelled.")
	}
}
