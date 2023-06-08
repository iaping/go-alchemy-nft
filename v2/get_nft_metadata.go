package nft

type GetNFTMetadata struct {
	Api
}

func NewGetNFTMetadata(c *Client) *GetNFTMetadata {
	return &GetNFTMetadata{
		Api{c: c},
	}
}

func (a *GetNFTMetadata) Do(p *GetNFTMetadataParam) (*Nft, error) {
	req, resp := a.request("getNFTMetadata", MethodGet, p), &Nft{}
	if err := a.c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetNFTMetadataParam struct {
	ContractAddress     string `url:"contractAddress"`
	TokenId             string `url:"tokenId"`
	TokenType           string `url:"tokenType,omitempty"`
	TokenUriTimeoutInMs int    `url:"tokenUriTimeoutInMs,omitempty"`
	RefreshCache        bool   `url:"refreshCache"`
}
