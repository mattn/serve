# serve

Serve static files as soon as you want!

## Usage

### Serve files under current directory as /

```
$ serve
2014/10/29 19:36:41 serving /home/mattn as / on :5000
```

### Serve files under /tmp as /

```
$ serve -r /tmp
2014/10/29 19:36:41 serving /tmp as / on :5000
```

### Serve files under /tmp as /foo

```
$ serve -r /tmp -p /tmp
2014/10/29 19:36:41 serving /tmp as /foo on :5000
```

### Serve files under /tmp as /foo on 192.168.0.3:80

```
$ serve -r /tmp -p /tmp -a 192.168.0.3:80
2014/10/29 19:36:41 serving /tmp as /foo on 192.168.0.3:80
```

## Requirements

golang

## Installation

```
$ go install github.com/mattn/serve@latest
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)

