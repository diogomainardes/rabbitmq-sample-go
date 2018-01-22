package worker

import (
	"fmt"
	"log"
	"sync"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

//Semaphore Semáforo utilizado pelas goroutines
var Semaphore sync.WaitGroup
