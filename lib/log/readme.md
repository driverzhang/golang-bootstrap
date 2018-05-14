github.com/sirupsen/logrus的简单封装

使用方法见log_test.go

测试环境打印如下:
```
time="2018-05-10 13:59:20" level=error msg=error caller=" git.zhuzi.me/zzjz/zhuzi-bootstrap/lib/log/log_test.go:16:"
```
线上环境打印如下:
```
{"caller":"git.zhuzi.me/zzjz/zhuzi-bootstrap/lib/log/log_test.go:10","level":"error","msg":"error","time":"2018-05-10 13:59:20"}
```

线上环境使用json格式方便以后做日志收集