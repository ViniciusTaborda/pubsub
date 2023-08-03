package file

import (
	"encoding/csv"
	"os"
	"pubsub/msg"
	"reflect"
	"testing"
)

func TestCSVMessageWriter(t *testing.T) {
	testCases := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "Write creates a new file with headers",
			test: func(t *testing.T) {
				csvw := NewCSVMessageWriter()
				message := &msg.GenericMessageHolder{Body: "test message"}
				csvw.Write(message, "pub1", "sub1", "test")

				file, err := os.Open("test.csv")
				if err != nil {
					t.Errorf("Expected file to exist, got error: %v", err)
				}
				defer file.Close()

				info, err := file.Stat()
				if err != nil {
					t.Errorf("Expected to get file info, got error: %v", err)
				}

				if info.Size() == 0 {
					t.Errorf("Expected file to be non-empty, got empty file")
				}

				//Cleaning up the logger file
				err = os.Remove("test.csv")
				if err != nil {
					t.Errorf("Expected to remove file, got error: %v", err)
				}
			},
		},
		{
			name: "Write appends a new record to an existing file",
			test: func(t *testing.T) {
				csvw := NewCSVMessageWriter()
				message1 := &msg.GenericMessageHolder{Body: "test message 1"}
				csvw.Write(message1, "pub1", "sub1", "test")

				file, err := os.Open("test.csv")
				if err != nil {
					t.Errorf("Expected file to exist, got error: %v", err)
				}
				defer file.Close()

				info1, err := file.Stat()
				if err != nil {
					t.Errorf("Expected to get file info, got error: %v", err)
				}

				message2 := &msg.GenericMessageHolder{Body: "test message 2"}
				csvw.Write(message2, "pub2", "sub2", "test")

				info2, err := file.Stat()
				if err != nil {
					t.Errorf("Expected to get file info, got error: %v", err)
				}

				if info2.Size() <= info1.Size() {
					t.Errorf("Expected file size to increase, got no increase")
				}

				//Cleaning up the logger file
				err = os.Remove("test.csv")
				if err != nil {
					t.Errorf("Expected to remove file, got error: %v", err)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.test)
	}
}

func TestCSVMessageWriterHeaders(t *testing.T) {
	csvw := NewCSVMessageWriter()
	message := &msg.GenericMessageHolder{Body: "test message"}
	csvw.Write(message, "pub1", "sub1", "test")

	file, err := os.Open("test.csv")
	if err != nil {
		t.Errorf("Expected file to exist, got error: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		t.Errorf("Expected to read headers, got error: %v", err)
	}

	expectedHeaders := []string{"publisherID", "subscriberID", "timestamp", "message"}
	if !reflect.DeepEqual(headers, expectedHeaders) {
		t.Errorf("Expected headers to be %v, got %v", expectedHeaders, headers)
	}

	//Cleaning up the logger file
	err = os.Remove("test.csv")
	if err != nil {
		t.Errorf("Expected to remove file, got error: %v", err)
	}
}

func TestCSVMessageWriterFileMissing(t *testing.T) {
	csvw := NewCSVMessageWriter()
	message := &msg.GenericMessageHolder{Body: "test message"}
	csvw.Write(message, "pub1", "sub1", "test")

	file, err := os.Open("test.csv")
	if err != nil {
		t.Errorf("Expected file to exist, got error: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		t.Errorf("Expected to read headers, got error: %v", err)
	}

	expectedHeaders := []string{"publisherID", "subscriberID", "timestamp", "message"}
	if !reflect.DeepEqual(headers, expectedHeaders) {
		t.Errorf("Expected headers to be %v, got %v", expectedHeaders, headers)
	}

	//Cleaning up the logger file
	err = os.Remove("test.csv")
	if err != nil {
		t.Errorf("Expected to remove file, got error: %v", err)
	}
}
