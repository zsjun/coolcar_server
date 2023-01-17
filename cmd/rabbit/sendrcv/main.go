package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	// 物理通道
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	// 虚拟的通道
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	// 建立一个队列
	q, err := ch.QueueDeclare(
		"go_q1",
		true,  // durable
		false, // autoDelete
		false, // exlusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		panic(err)
	}
	// 
	go consume("c1", conn, q.Name)
	go consume("c2", conn, q.Name)

	i := 0
	// 发送数据
	for {
		i++
		err := ch.Publish(
			"", // exchange
			q.Name,
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				Body: []byte(fmt.Sprintf("message %d", i)),
			},
		)
		if err != nil {
			fmt.Println(err.Error())
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func consume(consumer string, conn *amqp.Connection, q string) {
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		q,
		consumer, // consumer
		true,     // autoAck
		false,    // exclusive
		false,    // noLocal
		false,    // noWait
		nil,      // args
	)
	if err != nil {
		panic(err)
	}

	for msg := range msgs {
		fmt.Printf("%s: %s\n", consumer, msg.Body)
	}
}
