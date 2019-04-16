package hexipc

import (
    "fmt"
    "net"
    
    beson "github.com/GoblinBear/beson-go"
)

const socketFile string = "/home/bear/hex.sock"
var connection net.Conn = nil

// TODO: data pool

func NewHexIpc() error {
    c, err := net.Dial("unix", socketFile)
    if err != nil {
        fmt.Println(err)
        return err
    }
    
    connection = c
    return nil
}

func WriteBuf(data interface{}) error {
    // encode data
    ser := beson.Serialize(data)

    // write data to buffer
    _, err := connection.Write(ser)
    if err != nil {
        fmt.Println(err)
        return err
    }

    return nil
}

func ReadBuf() (interface{}, error) {
    // read data from buffer
    buf := make([]byte, 256)
    _, err := connection.Read(buf)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    // decode data
    _, value := beson.Deserialize(buf, 0)
    return value, nil
}
