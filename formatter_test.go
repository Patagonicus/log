package log_test

import (
	"fmt"
	"net"
	"time"

	"github.com/Patagonicus/log"
	"github.com/Patagonicus/log/mock"
)

func ExampleFormatter() {
	formatter := log.NewTextFormatter(log.FormatterConfig{
		Clock: mock.NewClock(mustParse("2006-01-02T15:04:05Z"), time.Second),
	})

	b, err := formatter.Format(log.Message{
		log.InfoLevel,
		"application started succesfully",
		nil,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	b, err = formatter.Format(log.Message{
		log.DebugLevel,
		"client connected",
		log.Fields{
			"user": "foobar",
			"ip":   net.IPv4(127, 0, 0, 1),
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	// Output:
	// 2006-01-02T15:04:05Z info: application started succesfully
	// 2006-01-02T15:04:06Z debug: client connected "ip"="127.0.0.1" "user"="foobar"
}

func mustParse(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}
