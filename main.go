package main

import (
	"fmt"
	"log"

	"github.com/diogomainardes/rabbitmq-sample-go/producer"
	"github.com/diogomainardes/rabbitmq-sample-go/worker"
)

func main() {
	var entrada string

	log.Printf("Inserindo mensagens na fila do checkout...")
	producer.SendMessageCheckout()
	log.Printf("Inserindo mensagens na fila do SMS...")
	producer.SendMessageSMS()
	log.Printf("Inserindo mensagens na fila do Emails...")
	producer.SendMessageEmails()

	log.Printf("Inserido mensagens com sucesso nas filas.")
	log.Printf("Pressione qualquer tecla para realizar o processamento.")
	fmt.Scanln(&entrada)

	worker.Semaphore.Add(3)
	go worker.ProcessCheckout()
	go worker.ProcessEmails()
	go worker.ProcessSMS()
	log.Printf("Aguardando mensagens...")

	worker.Semaphore.Wait()
}
