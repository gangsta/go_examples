package main
// my comment

import (
    "fmt"
    "net"
    "os"
    "os/exec"
    "log"
)

const (
    CONN_HOST = ""
    CONN_PORT = "3033"
    CONN_TYPE = "tcp"
)

func main() {
    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
  // Make a buffer to hold incoming data.
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  reqLen, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }
  // Send a response back to person contacting us.
  conn.Write([]byte("Message received."))
  // Close the connection when you're done with it.
  go curlRequest()
  conn.Close()
  fmt.Println("Bytes:", reqLen)
}

func curlRequest() {
  cmd := exec.Command("curl", "-s", "-k", "-o", "-w \"%{http_code}\"", "-XPOST", "http://172.17.8.121:8086/query --data-urlencode \"q=CREATE DATABASE mydb\"")
//  cmd := exec.Command("ls", "-lah")
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  err := cmd.Run()
  if err != nil {
    log.Fatalf("cmd.Run() failed with %s\n", err)
  }
}

