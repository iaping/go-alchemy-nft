# go-alchemy-nft
golang sdk for https://docs.alchemy.com/reference/nft-api-endpoints

## Installation

go get：
```shell
go get github.com/iaping/go-alchemy-nft/v2
```

## Example
All examples are in examples

## Rest api example code
```go
package main

import (
	"fmt"

	nft "github.com/iaping/go-alchemy-nft/v2"
)

func main() {
	client := nft.New("https://eth-mainnet.g.alchemy.com", "xxxxxxxxxxx", nil)
    resp, err := client.NewGetNFTs().Do(&nft.GetNFTsParam{
		Owner:        "0xe525FAE3fC6fBB23Af05E54Ff413613A6573CFf2",
		WithMetadata: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
```

## Help
Let me know if you have any problems.