你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
