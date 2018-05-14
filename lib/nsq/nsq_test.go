package msg_queue

import (
	"testing"
	"github.com/nsqio/go-nsq"
	"log"
	log2 "git.zhuzi.me/zzjz/zhuzi-bootstrap/lib/log"
)

type Handler struct {
}

func (p *Handler) HandleMessage(message *nsq.Message) error {
	log.Print(string(message.Body))
	return nil
}

func TestPublish(t *testing.T) {
	bs := []byte("bs")

	err := Publish("test", bs)
	if err != nil {
		t.Fatal(err)
	}

	select {}
}

func TestListen(t *testing.T) {
	err := Listen("test", "chan-1", &Handler{})
	if err != nil {
		t.Fatal(err)
	}

}

func init() {
	log2.SetDebug(true)
	Init("127.0.0.1:4150", "127.0.0.1:4161", true)
}
