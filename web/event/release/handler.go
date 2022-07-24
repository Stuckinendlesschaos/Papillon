package release

import (
	"encoding/json"
	"fmt"
	"gitrabbit/event"
	"gitrabbit/event/merge"
	gitlab "gitrabbit/gitlab"
)

//处理Release的事件
//ReleaseRequest Handle
func Handle(buf *[]byte) (*event.MessageNotify, error) {
	var rr ReleaseRequest
	err := json.Unmarshal(*buf, &rr)
	if err != nil {
		return nil, err
	}

	return &event.MessageNotify{
		MsgType:  "text",
		Notifier: rr.genNotifier(),
		Content:  rr.genContent(),
	}, nil
}

//generate Notifier for releaseRequest
func (rr *ReleaseRequest) genNotifier() []string {
	if merge.Notifier != nil {
		return merge.Notifier
	}

	g := &gitlab.GitLab{
		ProjectID: rr.Project.ID,
	}

	return g.GetNotifyUser()
}

//generate Content for releaseRequest
func (rr *ReleaseRequest) genContent() string {
	content := fmt.Sprintf(
		"服务名：%v\n发布版本：%v\n发布时间：%v\n%v\n",
		rr.Project.WebURL,
		rr.Tag,
		rr.CreatedAt,
		rr.Description,
	)

	return content
}
