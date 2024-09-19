package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare(
		"order_queue", // 队列名称
		true,          // 持久化
		false,         // 自动删除
		false,         // 独占
		false,         // 无需等待
		nil,           // 额外属性
	)
	failOnError(err, "Failed to declare a queue")

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	// 创建消费者
	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者标识符
		true,   // 自动应答
		false,  // 独占
		false,  // 无需等待
		false,  // no local
		nil,    // 额外属性
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	<-forever
}
