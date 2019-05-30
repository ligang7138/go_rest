package beanstalk

import (
	"fmt"
	"github.com/kr/beanstalk"
	//beanstalk "github.com/beanstalkd/go-beanstalk"
	"time"
)

func Producer(fname, tubeName string) {
	if fname == "" || tubeName == "" {
		return
	}

	c, err := beanstalk.Dial("tcp", "127.0.0.1:11300")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Tube.Name = tubeName
	c.TubeSet.Name[tubeName] = true
	fmt.Println(fname, " [Producer] tubeName:", tubeName, " c.Tube.Name:", c.Tube.Name)

	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("for %s %d", tubeName, i)
		c.Put([]byte(msg), 30, 0, 120*time.Second)
		fmt.Println(fname, " [Producer] beanstalk put body:", msg)
		//time.Sleep(1 * time.Second)
	}

	fmt.Println("Producer() end.")
}
