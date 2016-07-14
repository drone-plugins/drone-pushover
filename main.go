package main

import (
	"fmt"
	"os"
	"time"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
	"github.com/drone/drone-go/template"
	"github.com/gregdel/pushover"
)

var (
	buildCommit string
)

func main() {
	fmt.Printf("Drone Pushover Plugin built from %s\n", buildCommit)

	system := drone.System{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("system", &system)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if vargs.Retry == 0 {
		vargs.Retry = 60 * time.Second
	}

	if vargs.Expire == 0 {
		vargs.Expire = 3600 * time.Second
	}

	if vargs.Token == "" {
		fmt.Println("Please provide a app token")
		os.Exit(1)
	}

	if vargs.User == "" {
		fmt.Println("Please provide a user token")
		os.Exit(1)
	}

	client := pushover.New(vargs.Token)

	resp, err := client.SendMessage(
		&pushover.Message{
			Title:      BuildTitle(system, repo, build, vargs.Title),
			Message:    BuildBody(system, repo, build, vargs.Body),
			URL:        fmt.Sprintf("%s/%s/%d", system.Link, repo.FullName, build.Number),
			URLTitle:   "Link to the Build",
			DeviceName: vargs.Device,
			Sound:      vargs.Sound,
			Priority:   vargs.Priority,
			Retry:      vargs.Retry,
			Expire:     vargs.Expire,
			Timestamp:  time.Now().Unix(),
		},
		pushover.NewRecipient(
			vargs.User,
		),
	)

	fmt.Println(resp)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// BuildTitle renders the Pushover title from a template.
func BuildTitle(system drone.System, repo drone.Repo, build drone.Build, tmpl string) string {
	payload := &drone.Payload{
		System: &system,
		Repo:   &repo,
		Build:  &build,
	}

	if tmpl == "" {
		tmpl = defaultTitle
	}

	msg, err := template.RenderTrim(tmpl, payload)

	if err != nil {
		return err.Error()
	}

	return msg
}

// BuildBody renders the Pushover body from a template.
func BuildBody(system drone.System, repo drone.Repo, build drone.Build, tmpl string) string {
	payload := &drone.Payload{
		System: &system,
		Repo:   &repo,
		Build:  &build,
	}

	if tmpl == "" {
		tmpl = defaultBody
	}

	msg, err := template.RenderTrim(tmpl, payload)

	if err != nil {
		return err.Error()
	}

	return msg
}
