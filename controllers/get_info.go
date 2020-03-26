package make_message

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	gitlab_info "notification-center/struct"
	"notification-center/utils"
)

func GetInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	info := gitlab_info.Job{}
	_ = json.NewDecoder(r.Body).Decode(&info)
	log.Print(info)

	if info.BuildStatus == "failed" {
		title := "*OPS, BUILD FAILED!*\n"
		text := title + fmt.Sprintf("*Project*: [%s](%s)\n", info.ProjectName, info.Repository.GitHTTPURL)
		text += fmt.Sprintf("*Branch*: %s\n", info.Ref)
		text += fmt.Sprintf("*Commit*: [%s](%s)\n", info.Commit.Sha, info.Repository.Homepage+"/commit/"+info.Commit.Sha)
		text += fmt.Sprintf("*Message*: %s\n", info.Commit.Message)
		text += fmt.Sprintf("*Started by*: %s\n", info.User.Name)
		text += fmt.Sprintf("*Author*: %s  %s\n", info.Commit.AuthorName, info.Commit.AuthorEmail)
		text += fmt.Sprintf("*Please visit the* [job page](%s) *to know more information*", info.Repository.Homepage+"/-/jobs/"+info.BuildID)

		utils.SendMessage(text)
	}
}
