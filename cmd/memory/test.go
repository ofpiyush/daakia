package main

import (
	"net"
	"bufio"
	"fmt"
	"time"
)

//func logger(w http.ResponseWriter, r *http.Request) {
//	p, _ := ioutil.ReadAll(r.Body)
//	fmt.Println(string(p))
//	fmt.Fprintf(w, "ok")
//}
//func main() {
//	http.HandleFunc("/lala", logger)
//	http.ListenAndServe(":9000", nil)
//}


func main() {
	c, _ := net.Dial("tcp","127.0.0.1:3000")
	go read(c)
	c.Write([]byte("HAI Piyush\n"))

	i:= 0

	for {
		s := fmt.Sprintf("SAY hurr Bolo chala? %d \n",i)
		fmt.Print("Sending: ", s)
		fmt.Fprint(c,s)
		i++
		time.Sleep(1*time.Second)
	}

}
func read(c net.Conn) {
	r := bufio.NewReader(c)
	for {
		data,err:= r.ReadString('\n')
		if err !=nil {
			fmt.Println("error",err)
		}
		if len(data) >0 {
			fmt.Print("Server says: ", data)
		}

	}
}
