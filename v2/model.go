package nft

import "time"

type Nft struct {
	Contract         *Contract         `json:"contract,omitempty"`
	Id               *Id               `json:"id,omitempty"`
	Balance          string            `json:"balance,omitempty"`
	Title            string            `json:"title,omitempty"`
	Description      string            `json:"description,omitempty"`
	TokenUri         map[string]string `json:"tokenUri,omitempty"`
	Media            []*Media          `json:"media,omitempty"`
	Metadata         *Metadata         `json:"metadata,omitempty"`
	TimeLastUpdated  time.Time         `json:"timeLastUpdated,omitempty"`
	ContractMetadata *ContractMetadata `json:"contractMetadata,omitempty"`
}

type Contract struct {
	ContractMetadata
	Address                string   `json:"address"`
	TotalBalance           int64    `json:"totalBalance,omitempty"`
	NumDistinctTokensOwned int64    `json:"numDistinctTokensOwned,omitempty"`
	IsSpam                 bool     `json:"isSpam"`
	TokenId                string   `json:"tokenId,omitempty"`
	Title                  string   `json:"title,omitempty"`
	TotalSupply            string   `json:"totalSupply,omitempty"`
	Media                  []*Media `json:"media,omitempty"`
}

type Id struct {
	TokenId       string         `json:"tokenId"`
	TokenMetadata *TokenMetadata `json:"tokenMetadata,omitempty"`
}

type TokenMetadata struct {
	TokenType string `json:"tokenType"`
}

type Media struct {
	Gateway   string `json:"gateway,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
	Raw       string `json:"raw,omitempty"`
	Format    string `json:"format,omitempty"`
	Bytes     int64  `json:"bytes,omitempty"`
}

type Metadata struct {
	Id               int64                    `json:"id,omitempty"`
	Used             bool                     `json:"used,omitempty"`
	Phase            string                   `json:"phase,omitempty"`
	TokenId          interface{}              `json:"tokenId,omitempty"`
	ImageUrl         string                   `json:"imageUrl,omitempty"`
	Name             string                   `json:"name,omitempty"`
	ContractId       int64                    `json:"contractId,omitempty"`
	Description      string                   `json:"description,omitempty"`
	Attributes       []map[string]interface{} `json:"attributes,omitempty"`
	ImageOriginalUrl string                   `json:"image_original_url,omitempty"`
	ExternalUrl      string                   `json:"external_url,omitempty"`
}

type ContractMetadata struct {
	Name                string   `json:"name,omitempty"`
	Symbol              string   `json:"symbol,omitempty"`
	TotalSupply         string   `json:"totalSupply,omitempty"`
	TokenType           string   `json:"tokenType,omitempty"`
	ContractDeployer    string   `json:"contractDeployer,omitempty"`
	DeployedBlockNumber int64    `json:"deployedBlockNumber,omitempty"`
	OpenSea             *OpenSea `json:"openSea,omitempty"`
}

type OpenSea struct {
	FloorPrice            float64 `json:"floorPrice,omitempty"`
	CollectionName        string  `json:"collectionName,omitempty"`
	SafelistRequestStatus string  `json:"safelistRequestStatus,omitempty"`
	ImageUrl              string  `json:"imageUrl,omitempty"`
	Description           string  `json:"description,omitempty"`
	ExternalUrl           string  `json:"externalUrl,omitempty"`
	DiscordUrl            string  `json:"discordUrl,omitempty"`
	LastIngestedAt        string  `json:"lastIngestedAt"`
}

type ContractWithMetadata struct {
	Address          string            `json:"address"`
	ContractMetadata *ContractMetadata `json:"contractMetadata"`
}
