package main

import (
	"time"
)

var (
	// defaultTitle is the default template for the message title
	defaultTitle = `[{{ build.status }}] {{ repo.owner }}/{{ repo.name }} ({{ build.branch }} - {{ truncate build.commit 8 }})`

	// defaultMessage is the default template for the message body
	defaultBody = `{{ repo.owner }}/{{ repo.name }}#{{ truncate build.commit 8 }} ({{ build.branch }}) by {{ build.author }} in {{ duration build.started_at build.finished_at }} - <i>{{ build.message }}</i>`
)

// Params are the parameters that the Pushover plugin can parse.
type Params struct {
	Token    string        `json:"token"`
	User     string        `json:"user"`
	Title    string        `json:"title"`
	Body     string        `json:"body"`
	Device   string        `json:"device"`
	Sound    string        `json:"sound"`
	Priority int           `json:"priority"`
	Retry    time.Duration `json:"retry"`
	Expire   time.Duration `json:"expire"`
}
