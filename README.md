
# Gerish - for all your API-testing error, timeout, slowish needs
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/xonvanetta/gerish)]

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

- Slow
  - When you need a server that responds really slowly.
- Faulty
  - When you need a server that responds only with errors.
- Redirect
  - When you want the server to respond with a temporary redirect.

## Flags
All of these commands have their own flags which you can list as so:
``` 
$ gerish Slow -h
``` 

### --interval
All commands have this flag in common, by specifying this flag and an integer value, you can set how often the command should return an error.