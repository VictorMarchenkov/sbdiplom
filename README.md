# daemon

This package provides a service developed as diploma to Skillbox course "GO-разработчик"

After some analysing we decide to use as base packages:


# Directory structure

```
- cmd
- internal
- pkg
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
```
