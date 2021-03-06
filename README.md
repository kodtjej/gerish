# Gerish - for all your API-testing error, timeout, slowish needs
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/xonvanetta/gerish)]()
[![build status](https://gitlab.com/xonvanetta/gerish/badges/master/build.svg)](https://gitlab.com/xonvanetta/gerish/commits/master)
[![coverage report](https://gitlab.com/xonvanetta/gerish/badges/master/coverage.svg)](https://gitlab.com/xonvanetta/gerish/commits/master)

## Installation
To use gerish, you need to install go and set all the environment variables.
When that is done you just run:
``` 
$ go get gitlab.com/xonvanetta/gerish
```

Then you can run gerish just by typing:
```
$  gerish slow
``` 
## Commands

- slow
  - When you need a server that responds really slowly.
- faulty
  - When you need a server that responds only with errors.
- redirect
  - When you want the server to respond with a temporary redirect.
- connectionfail
  - When you want a server giving you "Connection reset by peer".
- healthy
  - When you want a healthy http server that returns an optional body.

## Flags
All of these commands have their own flags which you can list as so:
``` 
$ gerish slow -h
``` 

### --interval
All commands have this flag in common, by specifying this flag and an integer value, you can set how often the command should return an error.