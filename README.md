nginx-beanstalklog-replay
=========================

Replays http requests in [beanstalkd](http://kr.github.io/beanstalkd/) submitted via
[nginx-beanstalklog-module](https://github.com/r4um/nginx-beanstalklog-module) to [gor](https://github.com/buger/gor) replay server

```shell
$ go get -u http://github.com/kr/beanstalk
$ go build nginx-beanstalklog-replay.go
```

Usage

```shell
$ nginx-beanstalklog-replay -h
Usage of nginx-beanstalklog-replay:
  -b="127.0.0.1:11300": beanstalkd server address
  -f=3: sleep seconds after fatal errors
  -r="localhost:28020": gor replay server address
  -s=0: beanstalkd reserve timeout
  -t="nginx-log": beanstalkd tube to watch
```
