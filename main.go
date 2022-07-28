// *****************************************************
// Задержка для сценария по файлу JSON
// *****************************************************
package main

import (
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	"os"
)


// ******************************************************
// Main process ....
// ******************************************************
func main() {
	Logo()
  ReadScenarioJson("scenario.json")	
	// ReadPostJson("post.json")

}

// **********************************************
// Read scenario json
// **********************************************
func ReadScenarioJson(namejsonfile string){
   	jsonFile, err := os.Open(namejsonfile)
	
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

    // Show data JSON
    var scenar []Scenario
    json.Unmarshal(jsonData, &scenar)
  

   // For range
   for _, r:=range scenar{
   	 fmt.Println("Pause", r.Pause, r.Title)
     fmt.Print(">")	
     time.Sleep(r.Pause * time.Second)
   }
}

// *****************************************************
// Read Json post file
// *****************************************************
func ReadPostJson(namejsonfile string){
	jsonFile, err := os.Open(namejsonfile)
	
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

  // Show data JSON
  var post Post
  json.Unmarshal(jsonData, &post)
  fmt.Println(post)
}


