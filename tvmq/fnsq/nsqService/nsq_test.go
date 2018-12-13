package nsqService

import (
	"fmt"
	"strconv"
	"testing"

	"code.tvmining.com/tvplay/tvmq/config"
)

func TestProducer(t *testing.T) {
	for i := 0; i < 10000; i++ {
		in := Producer(config.NsqHostProt, "chat", strconv.Itoa(i))
		fmt.Println(<-in)
	}

}
