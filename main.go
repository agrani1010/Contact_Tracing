package main
 
import (
  
  "fmt"
  "log"
  "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Contact Tracing")
   
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}
type User struct {
    Id int `json:"Id"`
    Name string `json:"Name"`
    Date_Of_Birth string `json:"Date_Of_Birth"`
    Phone_Number string `json:"Phone_Number"`
    Email_address string `json:"Email_address"`
}
