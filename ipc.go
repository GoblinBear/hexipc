package hexipc

import (
    "fmt"
    "net"
    
    "beson"
    "beson/types"
)

const socketFile string = "/home/bear/hex.sock"
var connection net.Conn = nil

func NewHexIpc() error {
    c, err := net.Dial("unix", socketFile)
    if err != nil {
        fmt.Println(err)
        return err
    }
    
    connection = c
    return nil
}

func WriteBuf(data []byte) error {
    // encode data


    // write data to buffer
    _, err := connection.Write(data)
    if err != nil {
        fmt.Println(err)
        return err
    }

    return nil
}

func ReadBuf() error {
    // read data from buffer
    buf := make([]byte, 32)
    n, err := connection.Read(buf)
    if err != nil {
        return err
    }

    // decode data
    anchor, value := beson.Deserialize(buf, 0)
    fmt.Println("anchor =", anchor)
    fmt.Println("value =", value.(*types.String).Get())
    
    println("Client got:", buf)
    println("Client got:", string(buf[0:n]))

    return nil
}
