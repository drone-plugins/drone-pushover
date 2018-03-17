package main

import (
	"fmt"
	"time"

	"github.com/drone/drone-template-lib/template"
	"github.com/gregdel/pushover"
	"github.com/pkg/errors"
)

type (
	Repo struct {
		Owner string
		Name  string
	}

	Build struct {
		Tag      string
		Event    string
		Number   int
		Commit   string
		Ref      string
		Branch   string
		Author   string
		Pull     string
		Message  string
		DeployTo string
		Status   string
		Link     string
		Started  int64
		Created  int64
	}

	Config struct {
		Token    string
		User     string
		Title    string
		Message  string
		Device   string
		Sound    string
		Priority int
		Retry    time.Duration
		Expire   time.Duration
	}

	Job struct {
		Started int64
	}

	Plugin struct {
		Repo   Repo
		Build  Build
		Config Config
		Job    Job
	}
)

func (p Plugin) Exec() error {
	if p.Config.Token == "" {
		return errors.New("Please provide an app token")
	}

	if p.Config.User == "" {
		return errors.New("Please provide an user token")
	}

	title, err := template.RenderTrim(p.Config.Title, p)

	if err != nil {
		return errors.Wrap(err, "failed to render title")
	}

	message, err := template.RenderTrim(p.Config.Message, p)

	if err != nil {
		return errors.Wrap(err, "failed to render message")
	}

	client := pushover.New(p.Config.Token)

	resp, err := client.SendMessage(
		&pushover.Message{
			Title:      title,
			Message:    message,
			URL:        p.Build.Link,
			URLTitle:   "Link to the Build",
			DeviceName: p.Config.Device,
			Sound:      p.Config.Sound,
			Priority:   p.Config.Priority,
			Retry:      p.Config.Retry,
			Expire:     p.Config.Expire,
			Timestamp:  time.Now().Unix(),
		},
		pushover.NewRecipient(
			p.Config.User,
		),
	)

	fmt.Println(resp)

	if err != nil {
		return errors.Wrap(err, "failed to submit message")
	}

	return nil
}
