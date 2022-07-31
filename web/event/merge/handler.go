package merge

import (
	"encoding/json"
	"fmt"
	"papillon/event"
)

var (
	Notifier []string
	Name     string
)

//MergeRequest Handle
func Handle(buf *[]byte) (*event.MessageNotify, error) {
	var mr MergeRequest

	err := json.Unmarshal(*buf, &mr)
	if err != nil {
		return nil, err
	}

	//currently supports three type of state of MR
	content := ""
	switch mr.ObjectAttributes.Action {
	case "open":
		content = mr.genContent()
	case "update":
		content = mr.genContent()
	case "merge":
		content = mr.genContent()
	}

	return &event.MessageNotify{
		MsgType:  "text",
		Notifier: mr.GenNotifier(),
		Content:  content,
	}, nil
}

// 将获取MR的assignee分离是为了保证和其他事件的一致性
// generates notifier depending on user.Name
func (mr *MergeRequest) GenNotifier() []string {
	Notifier = make([]string, 0)

	for _, user := range mr.Assignees {
		Notifier = append(Notifier, user.Name)
	}

	return Notifier
}

// 获取具体结构化的内容
// reformat Message content to read
func (mr *MergeRequest) genContent() string {
	content := fmt.Sprintf(
		"%v\nMerge Request %v\nProject %v\n%v by %v\n",
		mr.ObjectAttributes.URL,
		mr.ObjectAttributes.Action,
		mr.Project.Name,
		mr.ObjectAttributes.Description,
		mr.User.Name,
	)

	return content
}
