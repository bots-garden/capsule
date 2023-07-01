// Package main
package main

import (
	"strconv"
	"strings"

	"github.com/bots-garden/capsule-module-sdk"
	"github.com/valyala/fastjson"
)

var apiURL = capsule.GetEnv("GITLAB_API_URL")
var botToken = capsule.GetEnv("GITLAB_BOT_TOKEN")
var botName = capsule.GetEnv("GITLAB_BOT_NAME")

func addNoteToIssue(message string, projectID, issueIID int) capsule.HTTPResponse {

		headers := map[string]string {
			"Content-Type": "application/json; charset=utf-8",
			"PRIVATE-TOKEN": botToken,
		}

	_, err := capsule.HTTP(capsule.HTTPRequest{
		JSONBody: `{"body": "` + message + `"}`,
		URI:      apiURL + "/projects/" + strconv.Itoa(projectID) + "/issues/" + strconv.Itoa(issueIID) + "/notes",
		Method:   "POST",
		//Headers:  `{"Content-Type": "application/json; charset=utf-8", "PRIVATE-TOKEN": "` + botToken + `"}`,
		Headers: capsule.SetHeaders(headers), // transform a map[string]string to a JSON string
	})

	if err != nil {
		response := capsule.HTTPResponse{
			JSONBody:   `{"error":"` + err.Error() + `"}`,
			Headers:    `{"Content-Type": "application/json; charset=utf-8"}`,
			StatusCode: 500,
		}
		return response
	}
	response := capsule.HTTPResponse{
		JSONBody:   `{"message":"OK"}`,
		Headers:    `{"Content-Type": "application/json; charset=utf-8"}`,
		StatusCode: 200,
	}
	return response
}

func main() {

	capsule.SetHandleHTTP(func(param capsule.HTTPRequest) (capsule.HTTPResponse, error) {

		capsule.Print("📝: " + param.Body)
		capsule.Print("🔠: " + param.Method)
		capsule.Print("🌍: " + param.URI)
		capsule.Print("👒: " + param.Headers)

		// Headers
		// Get a map of headers (make a 0.0.5 release oh MDK)
		headers := capsule.GetHeaders(param.Headers)

		capsule.Print("🔐 X-Gitlab-Token: " + headers["X-Gitlab-Token"])
		capsule.Print("🔐 WEBHOOK_TOKEN: " + capsule.GetEnv("WEBHOOK_TOKEN"))

		capsule.Print("🦊 X-Gitlab-Instance: " + headers["X-Gitlab-Instance"])
		capsule.Print("📝 X-Gitlab-Event: " + headers["X-Gitlab-Event"])

		var p fastjson.Parser
		jsonObject, err := p.Parse(param.Body)
		if err != nil {
			capsule.Log("🔴" + err.Error())
		}
		// issue or note
		objectKind := jsonObject.GetStringBytes("object_kind")
		capsule.Print("Object : " + string(objectKind))

		// issue or note
		eventType := jsonObject.GetStringBytes("event_type")
		capsule.Print("Event  : " + string(eventType))

		// the user who created the issue or add a note
		userName := jsonObject.Get("user").GetStringBytes("username")
		capsule.Print("Handle : " + string(userName))

		// current project name
		projectName := jsonObject.Get("project").GetStringBytes("name")
		capsule.Print("Project: " + string(projectName))

		// current project ID
		projectID := jsonObject.Get("project").GetInt("id")
		capsule.Print("Project ID: " + strconv.Itoa(projectID))

		objectAttributes := jsonObject.Get("object_attributes")

		var issueID int
		var issueIID int
		var content []byte
		var issueTitle []byte

		if string(objectKind) == "issue" {
			content = objectAttributes.GetStringBytes("description")
			issueTitle = objectAttributes.GetStringBytes("title")
			issueID = objectAttributes.GetInt("id")
			issueIID = objectAttributes.GetInt("iid")
		}

		if string(objectKind) == "note" {
			content = objectAttributes.GetStringBytes("note")
			issue := jsonObject.Get("issue")
			issueTitle = issue.GetStringBytes("title")
			issueID = issue.GetInt("id")
			issueIID = issue.GetInt("iid")
		}

		capsule.Print("Title  : " + string(issueTitle))
		capsule.Print("Content (description or note): " + string(content))
		capsule.Print("Issue ID     : " + strconv.Itoa(issueID))
		capsule.Print("Issue IID    : " + strconv.Itoa(issueIID))

		if string(objectKind) == "note" {
			if strings.Contains(string(content), botName) {
				message := "🤗 Hey @" + string(userName)
				return addNoteToIssue(message, projectID, issueIID), nil
			}
		}

		if string(objectKind) == "issue" {
			if strings.Contains(string(content), botName) {
				message := "😃 Hello @" + string(userName)
				return addNoteToIssue(message, projectID, issueIID), nil
			}
		}

		return capsule.HTTPResponse{
			JSONBody:   `{"message":"this is not an issue or note event"}`,
			Headers:    `{"Content-Type": "application/json; charset=utf-8"}`,
			StatusCode: 200,
		}, nil
	})
}
