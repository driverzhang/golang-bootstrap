package msg_queue

import (
	"github.com/nsqio/go-nsq"
	"time"
	"strings"
	"git.zhuzi.me/zzjz/zhuzi-bootstrap/lib/log"
	"os"
	"os/signal"
	"syscall"
	"sync"
)

var producer *nsq.Producer

var addrNsqLookups []string

var logLevel nsq.LogLevel

// 在调用Publish和Listen之前需要Init
// addrNsqp: 单个nsq地址, addrNsqLookupp:lookup地址 可以有多个,用逗号隔开
// todo 改成结构体, 有很多个参数, 如重试时间, 重试次数, 重连时间
func Init(addrNsq, addrNsqLookup string, debug bool) {
	addrNsqLookups = strings.Split(addrNsqLookup, ",")
	p, err := nsq.NewProducer(addrNsq, nsq.NewConfig())
	if err != nil {
		panic(err)
	}

	logLevel = nsq.LogLevelWarning
	if debug {
		logLevel = nsq.LogLevelInfo
	}

	p.SetLogger(log.NsqLogger(), logLevel)
	producer = p

	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		GracefulStop()
	}()
}

func Publish(topic string, body []byte, delay ...time.Duration) (err error) {
	if len(delay) == 0 {
		err = producer.Publish(topic, body)
	} else {
		err = producer.DeferredPublish(topic, delay[0], body)
	}
	return
}

var consumers []*nsq.Consumer
// 一个程序里会listen多个topic, 所以不阻塞
func Listen(topic string, channel string, handler nsq.Handler) (err error) {
	c, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
	if err != nil {
		return
	}
	c.AddHandler(handler)
	c.SetLogger(log.NsqLogger(), logLevel)
	err = c.ConnectToNSQLookupds(addrNsqLookups)
	if err != nil {
		panic(err)
	}

	consumers = append(consumers, c)
	return
}

func GracefulStop() {
	producer.Stop()

	var wg sync.WaitGroup
	for _, c := range consumers {
		wg.Add(1)
		go func() {
			c.Stop()
			wg.Done()
		}()
	}

	wg.Wait()
}
