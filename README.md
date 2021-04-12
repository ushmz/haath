# haath

Simple command to extract URLs from Chrome BrowserHistory json file.

Please specify where your history file is exported with `t` option.

- e : Means extension. If your history file was exported by some chrome extension, type e as a value of `t` option.
- t : Means Takeout. If your history file was exported in [takeout.google.com](https://takeout.google.com), type t as a value of `t` option.

```sh
go run ./haath.go -t (e|t) ${PATH_TO_FILE}
```

Or build binary and execute.

```sh
go build ./haath.go
./haath -t (e|t) ${PATH_TO_FILE}
```
