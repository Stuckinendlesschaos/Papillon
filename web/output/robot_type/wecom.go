package robot_type

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"papillon/conf"
	"papillon/event"
	"papillon/output"
)

type Wecom struct {
	MsgType  string   `json:"msgtype"`
	Notifier []string `json:"notifier"`
	Text     Text     `json:"text"`
}

type Markdown struct {
	Content string `json:"content,omitempty"`
}

type Text struct {
	Content             string    `json:"content,omitempty"`
	MentionedList       *[]string `json:"mentioned_list,omitempty"`
	MentionedMobileList *[]string `json:"mentioned_mobile_list,omitempty"`
}

type WcReturn struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaID   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

//WecomRobot is embedded Factory
type WecomRobot struct {
	wecom_robot *Wecom
}

//WeComRobot creates robot for Wecom
func (wr *WecomRobot) CreateRobot(msg *event.MessageNotify) output.Robot {
	wr.wecom_robot = &Wecom{
		msg.MsgType,
		msg.Notifier,
		Text{Content: msg.Content},
	}

	return wr.wecom_robot
}

//obtains reformatted msg and ready for sending
func (w *Wecom) SendSpecifiedAddress() error {
	userIdLists, err := getUserId(w.Notifier)
	if err != nil {
		return err
	}

	// 根据获得的assignees转化成WeCom能识别的userid
	// Not using mentioned list because it hides the users whose userid is either invalid or not registered
	for _, ass := range userIdLists {
		w.Text.Content = fmt.Sprintf(w.Text.Content+"\n<@%v>", ass)
	}

	// Make HTTP request
	client := &http.Client{}

	//发送到Wecom机器人的逻辑
	//converting structure to json
	body, err := json.Marshal(w)
	if err != nil {
		return err
	}
	//完整的HTTP通信逻辑
	//POST method for sending msg to Wecom
	req, err := http.NewRequest("POST", conf.Config.WeCom.DefaultAddr, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	//Set Request Header
	req.Header.Add("Content-Type", "application/json")
	//* send regine for real and obtain response
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handles wecom's return
	var wr WcReturn
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &wr)
	if err != nil {
		return err
	}

	return nil

}

//find necessary userid with username
func getUserId(users []string) (uid []string, err error) {
	for _, u := range users {
		if v, ok := conf.Usermap[u]; ok {
			uid = append(uid, v)
		} else {
			//* indicate that we have not this new user info
			uid = append(uid, u)
		}
	}

	return uid, nil
}
