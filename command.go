package main

//import object
import (
    "fmt"
    //"./client"
    "strings"
)

// The CommandFunc type is an adapter to allow the use of
// ordinary functions as Command functions.
type CommandFunc func(client *Client)

func (comm CommandFunc) Process(client *Client) {
    comm(client)
}

var commandMap = map[string] CommandFunc {
    "set": Set, 
    "get": Get,
}

func Set(client *Client) {
    fmt.Printf("\n\nSettings: %s %s", client.Request.Argv[1], client.Request.Argv[2])
    client.Db.Set(client.Request.Argv[1], client.Request.Argv[2])
    // Reply
    client.Response.Ok()
}

func Get(client *Client) {
    buf := client.Db.Get(client.Request.Argv[1])
    if (len(buf) == 0) {
        client.Response.Nil()
        return
    }
    
    fmt.Printf("\n\nGET RESPONSE: %s", string(buf))
    client.Response.SendBulk(buf)
    //_, _ = client.Conn.Write([]byte("+OK\r\n"))
    //client.Conn.Close()
}

func Unknown(client *Client) {
    // Send err
}

func CommandFromRequest(r *Request) func(client *Client) {
    key := strings.ToLower(string(r.Argv[0]))
    comm, ok := commandMap[key]
    if !ok {
        // Unknown command. 
        return Unknown
    }

    return comm
}