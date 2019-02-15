package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "pushover plugin"
	app.Usage = "pushover plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Usage:  "pushover app token",
			EnvVar: "PUSHOVER_TOKEN,PLUGIN_TOKEN",
		},
		cli.StringFlag{
			Name:   "user",
			Usage:  "pushover user token",
			EnvVar: "PUSHOVER_USER,PLUGIN_USER",
		},
		cli.StringFlag{
			Name:   "title",
			Usage:  "title template",
			EnvVar: "PLUGIN_TITLE",
			Value:  "[{{ build.status }}] {{ repo.owner }}/{{ repo.name }} ({{ build.branch }} - {{ truncate build.commit 8 }})",
		},
		cli.StringFlag{
			Name:   "message",
			Usage:  "mssage template",
			EnvVar: "PLUGIN_MESSAGE",
			Value:  "{{ repo.owner }}/{{ repo.name }}#{{ truncate build.commit 8 }} ({{ build.branch }}) by {{ build.author }} - <i>{{ build.message }}</i>",
		},
		cli.StringFlag{
			Name:   "device",
			Usage:  "send to device",
			EnvVar: "PLUGIN_DEVICE",
		},
		cli.StringFlag{
			Name:   "sound",
			Usage:  "play a sound",
			EnvVar: "PLUGIN_SOUND",
		},
		cli.StringFlag{
			Name:   "priority",
			Usage:  "define priority",
			EnvVar: "PLUGIN_PRIORITY",
		},
		cli.DurationFlag{
			Name:   "retry",
			Usage:  "retry duration",
			EnvVar: "PLUGIN_RETRY",
			Value:  60 * time.Second,
		},
		cli.DurationFlag{
			Name:   "expire",
			Usage:  "expire duration",
			EnvVar: "PLUGIN_EXPIRE",
			Value:  3600 * time.Second,
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "commit.sha",
			Usage:  "git commit sha",
			EnvVar: "DRONE_COMMIT_SHA",
			Value:  "00000000",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:   "commit.branch",
			Value:  "master",
			Usage:  "git commit branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "commit.author",
			Usage:  "git author name",
			EnvVar: "DRONE_COMMIT_AUTHOR",
		},
		cli.StringFlag{
			Name:   "commit.pull",
			Usage:  "git pull request",
			EnvVar: "DRONE_PULL_REQUEST",
		},
		cli.StringFlag{
			Name:   "commit.message",
			Usage:  "commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.IntFlag{
			Name:   "build.number",
			Usage:  "build number",
			EnvVar: "DRONE_BUILD_NUMBER",
		},
		cli.StringFlag{
			Name:   "build.status",
			Usage:  "build status",
			Value:  "success",
			EnvVar: "DRONE_BUILD_STATUS",
		},
		cli.StringFlag{
			Name:   "build.link",
			Usage:  "build link",
			EnvVar: "DRONE_BUILD_LINK",
		},
		cli.Int64Flag{
			Name:   "build.started",
			Usage:  "build started",
			EnvVar: "DRONE_BUILD_STARTED",
		},
		cli.Int64Flag{
			Name:   "build.created",
			Usage:  "build created",
			EnvVar: "DRONE_BUILD_CREATED",
		},
		cli.StringFlag{
			Name:   "build.tag",
			Usage:  "build tag",
			EnvVar: "DRONE_TAG",
		},
		cli.StringFlag{
			Name:   "build.deployTo",
			Usage:  "environment deployed to",
			EnvVar: "DRONE_DEPLOY_TO",
		},
		cli.Int64Flag{
			Name:   "job.started",
			Usage:  "job started",
			EnvVar: "DRONE_JOB_STARTED",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Tag:      c.String("build.tag"),
			Number:   c.Int("build.number"),
			Event:    c.String("build.event"),
			Status:   c.String("build.status"),
			Commit:   c.String("commit.sha"),
			Ref:      c.String("commit.ref"),
			Branch:   c.String("commit.branch"),
			Author:   c.String("commit.author"),
			Pull:     c.String("commit.pull"),
			Message:  c.String("commit.message"),
			DeployTo: c.String("build.deployTo"),
			Link:     c.String("build.link"),
			Started:  c.Int64("build.started"),
			Created:  c.Int64("build.created"),
		},
		Job: Job{
			Started: c.Int64("job.started"),
		},
		Config: Config{
			Token:    c.String("token"),
			User:     c.String("user"),
			Title:    c.String("title"),
			Message:  c.String("message"),
			Device:   c.String("device"),
			Sound:    c.String("sound"),
			Priority: c.Int("priority"),
			Retry:    c.Duration("retry"),
			Expire:   c.Duration("expire"),
		},
	}

	return plugin.Exec()
}
