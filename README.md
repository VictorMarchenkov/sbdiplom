# daemon

This package provides a daemon developed as diploma to Skillbox course "GO-разработчик"

After some analysing we decide to use as base packages:

- github.com/takama/daemon
- https://github.com/haxii/daemon (a command line wrapper for https://github.com/takama/daemon)

# Directory structure

```
- cmd
- internal
- pkg
```

## To install the daemon on port 80

```sh
httpdaemon -s install -p 80
```

to start or stop the daemon

```sh
httpdaemon -s start
httpdaemon -s stop
```

to get status of the daemon

```sh
httpdaemon -s
httpdaemon -s status
```

to remove the daemon

```
httpdaemon -s remove
```

## heroku
[Getting Started on Heroku with Go](https://devcenter.heroku.com/articles/getting-started-with-go)


```
heroku apps
```

```
heroku apps:delete "app name"
```

Creates a new heroku-dyno and copies the master project branch to this dyno
```
heroku create
git push heroku main
```

You can display the token via the CLI:

```http request
heroku auth:token
8225ee0c-aa0f-42e5-955c-d602e899650d
```