# etherscan-go
etherscan.io API client in Go

```go
// main.go

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
	c, err := etherscan.NewClient(etherscan.Mainnet, "YOUR_API_KEY")
	checkErr(err)

	resp, err := c.Account("0x0000000000000000000000000000000000000000")
	checkErr(err)

	fmt.Printf("%+v\n", resp)
}
```

```bash
go run main.go
&{Status:1 Message:OK Result:+7641879211751631214533}
```
