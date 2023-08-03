package file

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"pubsub/msg"
	"sync"
	"time"
)

type CSVMessageWriter struct {
	mutex *sync.Mutex
}

func NewCSVMessageWriter() MessageWriter {
	return &CSVMessageWriter{
		mutex: &sync.Mutex{},
	}

}

func (csvw *CSVMessageWriter) Write(message msg.MessageHolder, publisherID, subscriberID, topic string) {

	csvw.mutex.Lock()
	defer csvw.mutex.Unlock()

	// Errors only being logged because this operation should not block the messages
	filename := fmt.Sprintf("%s.csv", topic)

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	info, err := file.Stat()
	if err != nil {
		log.Println(err)
	}

	if info.Size() == 0 {
		headers := []string{"publisherID", "subscriberID", "timestamp", "message"}
		err = writer.Write(headers)
		if err != nil {
			log.Println(err)
		}
	}

	record := []string{
		publisherID,
		subscriberID,
		time.Now().Format(time.RFC3339),
		message.GetStringBody(),
	}

	err = writer.Write(record)

	if err != nil {
		log.Println(err)
	}
}
