/*
Channels are the pipes that connect concurrent goroutines
We can send values into channels from one goroutine and receive those value into another goroutine
*/
//Example of using channels
/*
package main
import "fmt"
func main() {
	message := make(chan string)
	go func() { message <- "ping" }() //----> This is Called sender
	// Here in the example we can see that the sender and receiver both block the line until both are ready
	msg := <-message //----> This is called Receiver
	fmt.Println(msg)
}
*/

//Similarly If the channel is buffered we can send the
// value without having a corresponding receive

/*
when using channels as function parameters we can
specify if a channel is meant to only send or receive
values.

for Example:-
func ping(pings chan<-string,msg string){pings<-msg}
Here the ping function only accepts a channel for sending values
if pong <-chan string it means receiving a value
*/
/*
package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Hey This is the message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
*/

//Note:
//pings <-chan string ---> used for receiving
//pongs chan<- string ----> used for sending
//sending means := channel <- "some message"
//receiving means:= varName := <-channels

//Receiving data from multiple channel

package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 23
	}()

	// go func() {
	// 	chan2 <- "Sangam Silwal"
	// }()
	defer close(chan1)
	defer close(chan2)

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("value received from channel 1", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("value received from channel 1", chan2Val)
		}

	}

}
