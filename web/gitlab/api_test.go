package gitlab_api

import (
	"fmt"
	"testing"
)

func TestGetLastCommit(t *testing.T) {
	gitlab := new(GitLab)
	commit, err := gitlab.GetLastCommit()
	if err != nil {
		fmt.Errorf("get last commit info failed,err is %w", err)
	}

	fmt.Println(commit)
}
