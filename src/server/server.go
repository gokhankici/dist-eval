package main

import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"
)

type Msg struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
	Body struct {
		// common fields
		Type string `json:"type"`
		MsgId int `json:"msg_id"`
		InReplyTo int `json:"in_reply_to"`

		// raft init
		NodeId string `json:"node_id"`
		NodeIds []string `json:"node_ids"`

		// error code
		ErrCode int `json:"code"`
		ErrText int `json:"text"`

		// read / write / cas
		Key   string `json:"key"`
		Value string `json:"value"`
		From  string `json:"from"`
		To    string `json:"to"`

	} `json:"body"`
}

func msg_producer(c_in chan Msg, c_out chan bool) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		var m Msg

		b := []byte(scanner.Text())

		if err := json.Unmarshal(b, &m); err != nil {
			panic(err)
		}

		c_in <- m
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading stdin:", err)
	}

	c_out <- true
}

func msg_printer(c chan Msg, d1 chan bool, d2 chan bool) {
	for {
		select {
		case m := <-c:
			fmt.Fprintf(os.Stderr, "%+v\n", m)
		case <-d1:
			d2 <- true
			break
		}
	}
}

func main() {
	messages := make(chan Msg)
	done1    := make(chan bool)
	done2    := make(chan bool)

	go msg_producer(messages, done1)

	go msg_printer(messages, done1, done2)

	<-done2
}
