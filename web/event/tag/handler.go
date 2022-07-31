package tag

import (
	"encoding/json"
	"fmt"
	"papillon/event"
	"papillon/event/merge"
	gitlab "papillon/gitlab"
)

func Handle(buf *[]byte) (*event.MessageNotify, error) {
	var tr TagRequest
	err := json.Unmarshal(*buf, &tr)
	if err != nil {
		return nil, err
	}

	return &event.MessageNotify{
		MsgType:  "text",
		Notifier: tr.genNotifier(),
		Content:  tr.genContent(),
	}, nil
}

//generate Notifier for tagRequest
func (tr *TagRequest) genNotifier() []string {
	if merge.Notifier != nil {
		return merge.Notifier
	}

	g := &gitlab.GitLab{
		ProjectID: tr.Project.ID,
	}

	return g.GetNotifyUser()
}

//generate Content for releaseRequest
func (tr *TagRequest) genContent() string {
	content := fmt.Sprintf(
		"服务名：%v\n发布版本：%v\n操作人：%v\n%v\n",
		tr.Project.WebURL,
		tr.Ref[9:],
		tr.UserName,
		tr.Message,
	)

	return content
}
