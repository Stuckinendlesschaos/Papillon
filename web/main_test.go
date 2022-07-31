package main

import (
	"fmt"
	"os"
	"testing"
)

// func TestSendToWecom(t *testing.T) {
// 	wc := Wecom{
// 		MsgType: "markdown",
// 		Markdown: &Markdown{
// 			Content: "testo",
// 		},
// 	}
// 	fmt.Println(wc)
// 	sendToWecom(wc)
// }

func TestReadEnv(t *testing.T) {
	fmt.Println(os.Getenv("BOTADDR"))
}

//func TestGetUserId(t *testing.T) {
//	user := []merge_event.User{
//		merge_event.User{
//			Username: "fengxueming",
//		},
//		merge_event.User{
//			Username: "kuroame",
//		},
//	}
//	uid, _ := getUserId(user)
//	fmt.Println(uid)
//}

//func TestGenText(t *testing.T) {
//	w := genText(merge_event.Request{
//		ObjectAttributes: merge_event.ObjectAttributes{
//			URL:    "https://gitlab.com/fengxueming/test/merge_requests/1",
//			Action: "approved",
//			LastCommit: merge_event.LastCommit{
//				Title: "test",
//			},
//		},
//		Project: request.Project{
//			Name:   "test",
//			WebURL: "https://gitlab.com/fengxueming/test",
//		},
//		User: merge_event.User{
//			Name: "fengxueming",
//		},
//		Assignees: []merge_event.User{
//			merge_event.User{
//				Username: "fengxueming",
//			},
//			merge_event.User{
//				Username: "kuroame",
//			},
//		},
//	})
//	j, _ := json.Marshal(w)
//	fmt.Print(string(j))
//}
