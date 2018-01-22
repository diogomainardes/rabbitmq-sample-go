package producer

import (
	"math/rand"

	"github.com/streadway/amqp"
)

//SendMessageCheckout manda mensagem aleatorias para o checkout
func SendMessageCheckout() {
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
		"checkout", // name
		true,       // durable
		false,      // delete when usused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Falha ao declarar a fila")

	body := "Envio de mensagens para o checkout"

	var numberOfPublishes = rand.Intn(1000000)
	for i := 0; i < numberOfPublishes; i++ {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Falha ao declarar a fila")
	}

}

//SendMessageSMS manda mensagem aleatorias para o SMS
func SendMessageSMS() {
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

	body := "Envio de mensagens para o SMS"

	var numberOfPublishes = rand.Intn(1000000)
	for i := 0; i < numberOfPublishes; i++ {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Falha ao declarar a fila")
	}

}

//SendMessageEmails manda mensagem aleatorias para o checkout
func SendMessageEmails() {
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

	body := "Envio de mensagens para o Emails"

	var numberOfPublishes = rand.Intn(1000000)
	for i := 0; i < numberOfPublishes; i++ {
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		failOnError(err, "Falha ao declarar a fila")
	}

}
