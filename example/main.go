package main

import (
	"fmt"
	"log"

	"github.com/gravityblast/etherscan"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	c, err := etherscan.NewClient(etherscan.Mainnet, "")
	checkErr(err)

	resp, err := c.Account("0x0000000000000000000000000000000000000000")
	checkErr(err)

	fmt.Printf("%+v\n", resp)
}
