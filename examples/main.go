package main

import (
	"fmt"
	nft "go-alchemy-nft/v2"
)

var client = nft.New("https://eth-mainnet.g.alchemy.com", "xxxxxx", nil)

func main() {
	testGetNFTs()
	testGetContractsForOwner()
	testIsHolderOfCollection()
	testGetOwnersForToken()
	testGetOwnersForCollection()
	testGetNFTMetadata()
	testGetContractMetadata()
	testInvalidateContract()
	testReingestContract()
	testSearchContractMetadata()
	testGetNFTsForCollection()
	testGetNFTMetadataBatch()
	testGetContractMetadataBatch()
}

func testGetNFTs() {
	resp, err := client.NewGetNFTs().Do(&nft.GetNFTsParam{
		Owner:        "0xe525FAE3fC6fBB23Af05E54Ff413613A6573CFf2",
		WithMetadata: true,
		// PageSize:     1,
		// ContractAddresses: []string{
		// },
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func testGetContractsForOwner() {
	resp, err := client.NewGetContractsForOwner().Do(&nft.GetContractsForOwnerParam{
		Owner:        "0xe525FAE3fC6fBB23Af05E54Ff413613A6573CFf2",
		WithMetadata: true,
		// PageSize:     1,
		// ContractAddresses: []string{
		// },
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func testIsHolderOfCollection() {
	resp, err := client.NewIsHolderOfCollection().Do(&nft.IsHolderOfCollectionParam{
		Wallet:          "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		ContractAddress: "0xe785E82358879F061BC3dcAC6f0444462D4b5330",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func testGetOwnersForToken() {
	resp, err := client.NewGetOwnersForToken().Do(&nft.GetOwnersForTokenParam{
		ContractAddress: "0x80321ac3ea82f848031f1a8db7b6a351f4cc0a97",
		TokenId:         "2655",
		// TokenId:         "0x0000000000000000000000000000000000000000000000000000000000000a5f",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func testGetOwnersForCollection() {
	resp, err := client.NewGetOwnersForCollection().Do(&nft.GetOwnersForCollectionParam{
		ContractAddress:   "0x80321ac3ea82f848031f1a8db7b6a351f4cc0a97",
		WithTokenBalances: false,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.(*nft.GetOwnersForCollectionResponse))

	resp, err = client.NewGetOwnersForCollection().Do(&nft.GetOwnersForCollectionParam{
		ContractAddress:   "0x80321ac3ea82f848031f1a8db7b6a351f4cc0a97",
		WithTokenBalances: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.(*nft.GetOwnersForCollectionWithTokenBalancesResponse))
}

func testGetNFTMetadata() {
	resp, err := client.NewGetNFTMetadata().Do(&nft.GetNFTMetadataParam{
		ContractAddress: "0x80321ac3ea82f848031f1a8db7b6a351f4cc0a97",
		TokenId:         "2655",
		// TokenId:         "0x0000000000000000000000000000000000000000000000000000000000000a5f",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func testGetContractMetadata() {
	resp, err := client.NewGetContractMetadata().Do(&nft.GetContractMetadataParam{
		ContractAddress: "0xe785E82358879F061BC3dcAC6f0444462D4b5330",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func testInvalidateContract() {
	resp, err := client.NewInvalidateContract().Do(&nft.InvalidateContractParam{
		ContractAddress: "0xe785E82358879F061BC3dcAC6f0444462D4b5330",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func testReingestContract() {
	resp, err := client.NewReingestContract().Do(&nft.ReingestContractParam{
		ContractAddress: "0xe785E82358879F061BC3dcAC6f0444462D4b5330",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func testSearchContractMetadata() {
	resp, err := client.NewSearchContractMetadata().Do(&nft.SearchContractMetadataParam{
		Query: "bored",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func testGetNFTsForCollection() {
	resp, err := client.NewGetNFTsForCollection().Do(&nft.GetNFTsForCollectionParam{
		ContractAddress: "0x80321ac3ea82f848031f1a8db7b6a351f4cc0a97",
		WithMetadata:    true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func testGetNFTMetadataBatch() {
	resp, err := client.NewGetNFTMetadataBatch().Do(&nft.GetNFTMetadataBatchParam{
		Tokens: []nft.TokenParam{
			{
				ContractAddress: "0x80321ac3ea82f848031f1a8db7b6a351f4cc0a97",
				TokenId:         "44",
			},
			{
				ContractAddress: "0x81Ae0bE3A8044772D04F32398bac1E1B4B215aa8",
				TokenId:         "5938",
			},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

func testGetContractMetadataBatch() {
	resp, err := client.NewGetContractMetadataBatch().Do(&nft.GetContractMetadataBatchParam{
		ContractAddresses: []string{
			"0x80321ac3ea82f848031f1a8db7b6a351f4cc0a97",
			"0x81Ae0bE3A8044772D04F32398bac1E1B4B215aa8",
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
