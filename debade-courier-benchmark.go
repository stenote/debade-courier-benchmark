package main

import (
	"flag"
	"log"
	"runtime"
	"time"

	"github.com/pebbe/zmq4"
)

func main() {
	var (
		//flag
		path       string
		concurrent int
		number     int
		data       string
	)

	flag.StringVar(&path, "p", "tcp://0.0.0.0:3333", "0MQ 连接路径")
	flag.IntVar(&concurrent, "c", 10, "并发数")
	flag.IntVar(&number, "n", 100, "每秒发送数")
	flag.StringVar(&data, "d", "", "发送的数据(json字符串)")

	flag.Parse()

	if len(data) == 0 {
		log.Fatal("Wrong Data!")

	}

	log.Println(data)

	// 设置并发数
	runtime.GOMAXPROCS(concurrent)

	context, err := zmq4.NewContext()

	if err != nil {
		log.Fatalf("Cannot Create 0MQ Context: %v", err)
	}

	socket, err := context.NewSocket(zmq4.PUSH)

	if err != nil {
		log.Fatalf("Cannot Create 0MQ Socket: %v", err)
	}

	defer socket.Close()

	socket.Connect(path)

	log.Printf("Connect %s", path)

	log.Println("Running !")

	for {

		for i := 0; i < number; i++ {
			go socket.Send(data, zmq4.DONTWAIT)
		}

		time.Sleep(time.Second)
	}
}
