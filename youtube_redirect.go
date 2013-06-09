package main

import( 
  "net/http"
  "io"
  "time"
  "os"
)

type Relay struct{}

var client = http.Client{}

func LogEvent(entry string) {
  file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0600)
  if err != nil { panic(err) }
  file.WriteString(entry)
  err = file.Close()
  if err != nil { panic(err) }
}

func (relay Relay) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
  relativeUrl := request.URL
  urlString := "http://youtube.com" + relativeUrl.String()
  logString := request.RemoteAddr + urlString + time.Now().String()
  LogEvent(logString)
  response, _ := client.Get(urlString)
  io.Copy(writer, response.Body) 
  response.Body.Close()
}



func main() {
  relay := Relay{}
  http.Handle("/", relay)
  http.ListenAndServe("localhost:4000", nil)
}