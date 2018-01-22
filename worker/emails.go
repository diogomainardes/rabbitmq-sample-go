package worker

import (
	"log"

	"github.com/streadway/amqp"
)

func consumeEmailQueue(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf(" [EMAILS] %s", d.Body)
	}
}

//ProcessEmails Realiza o processamento da fila de emails
func ProcessEmails() {

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
		"emails", // name
		true,     // durable
		false,    // delete when usused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
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

	go consumeEmailQueue(msgs)

	log.Printf(" [EMAILS] Aguardando chegada na fila de emails")

	<-forever
	Semaphore.Done()
}
