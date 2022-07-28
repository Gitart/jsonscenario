package main

import (
	"testing"
	"os"
	"fmt"
  "encoding/json"
  "time"
  "io/ioutil"
)


// Test decode
func TestDecode(t *testing.T) {
	post, err := decode("post.json")

	if err != nil {
	 	 t.Error(err)
	}

	if post.Id != 1 {
		 t.Error("Wrong id, was expecting 1 but got", post.Id)
	}

	if post.Content != "Hello World!" {
		 t.Error("Wrong content, was expecting 'Hello World!' but got", post.Content)
	}
}


// Test Encode
func TestEncode(t *testing.T) {
	   t.Skip("Skipping encoding for now")
}

// Test long
func TestLongRunningMain(t *testing.T) {
	if testing.Short() {
		 t.Skip("Main test process....")
	}

	time.Sleep(10 * time.Second)
}


// Test long
func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		 t.Skip("Skipping long-running test in short mode")
	}

	time.Sleep(10 * time.Second)
}


// Decode
func decode(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
}

	defer jsonFile.Close()
	
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	return
}

// Parallel 
func TestParallel_1(t *testing.T) {
			t.Parallel()
			time.Sleep(1 * time.Second)
}

func TestParallel_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParallel_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}

// benchmarking the decode function
func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
	decode("post.json")
	}
}


// Bench
func unmarshal(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}

	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}
	
	json.Unmarshal(jsonData, &post)
   return
}

// Bench test
func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
	    unmarshal("post.json")
	}
}



