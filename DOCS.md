Use this plugin for sending build status notifications via Pushover. You will
need to supply Drone with a Pushover user and application token. You can
override the default configuration with the following parameters:

* `user` - Your user or group token for Pushover
* `token` - Token of your registered application
* `device` - Recipient device or group name, optional
* `sound` - Play a sound, defaults to device default, optional
* `priority` - Priority of the message, defaults to `0`
* `retry` - Retry after `x` seconds on `emergency` priority, defaults to `60`
* `expire` - Expire after `x` seconds on `emergency` priority, defaults to `3600`

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
notify:
  pushover:
    user: uCBQgN8Ce8TBFHdEcEACMhkbEnJpxK
    token: aFeHUxHQdiBbragkajzwCNQESh4H6D
```

### Custom Messages

In some cases you may want to customize the body of the Pushover message you
can use custom templates. For the use case we expose the following additional
parameters:

* `body` - A handlebars template to create a custom payload body. For more
  details take a look at the [docs](http://handlebarsjs.com/) and the
  [Pushover docs](https://pushover.net/api#html).
* `title` - A handlebars template to create a custom payload body. For more
  details take a look at the [docs](http://handlebarsjs.com/) and the
  [Pushover docs](https://pushover.net/api#html).

Example configuration that generate a custom message:

```yaml
notify:
  pushover:
    user: uCBQgN8Ce8TBFHdEcEACMhkbEnJpxK
    token: aFeHUxHQdiBbragkajzwCNQESh4H6D
    body: >
      {{ repo.owner }}/{{ repo.name }}#{{ truncate build.commit 8 }}
      ({{ build.branch }})
      by {{ build.author }}
      in {{ duration build.started_at build.finished_at }} - {{ build.message }}
    title: >
      [{{ build.status }}]
      {{ repo.owner }}/{{ repo.name }}
      ({{ build.branch }} - {{ truncate build.commit 8 }})
```
