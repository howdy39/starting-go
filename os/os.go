package main

import (
	"log"
	"os"
)

func main() {

	// os.Exit(1) // 終了ステータス1で終了

	_, err := os.Open("foo")
	if err != nil {
		log.Fatal(err)
	}

}
