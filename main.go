package main

import (
	"log"

	"github.com/diogomainardes/rabbitmq-sample-go/worker"
)

func main() {
	worker.Semaphore.Add(3)

	go worker.ProcessCheckout()
	go worker.ProcessEmails()
	go worker.ProcessSMS()

	log.Printf("Aguardando mensagens...")

	worker.Semaphore.Wait()
}
