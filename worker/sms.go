package worker

import (
	"log"

	"github.com/streadway/amqp"
)

func consumeSMSQueue(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf(" [SMS] %s", d.Body)
	}
}

//ProcessSMS Realiza o processamento da fila de SMS
func ProcessSMS() {

	// Realiza a conexão com o servidor RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Falha ao se conectar ao RabbitMQ")
	defer conn.Close()

	// Estabelece um canal de comunicação
	ch, err := conn.Channel()
	failOnError(err, "Falha ao abrir o canal de comunicação")
	defer ch.Close()

	// Recupera a instância da fila
	q, err := ch.QueueDeclare(
		"sms", // name
		true,  // durable
		false, // delete when usused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Falha ao declarar a fila")

	// Consome dados da fila
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Falha ao registrar o consumer")

	forever := make(chan bool)

	go consumeSMSQueue(msgs)

	log.Printf(" [SMS] Aguardando chegada na fila de SMS")

	<-forever
	Semaphore.Done()
}
