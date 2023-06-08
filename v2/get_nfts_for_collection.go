package nft

type GetNFTsForCollection struct {
	Api
}

func NewGetNFTsForCollection(c *Client) *GetNFTsForCollection {
	return &GetNFTsForCollection{
		Api{c: c},
	}
}

func (a *GetNFTsForCollection) Do(p *GetNFTsForCollectionParam) (*GetNFTsForCollectionResponse, error) {
	req, resp := a.request("getNFTsForCollection", MethodGet, p), &GetNFTsForCollectionResponse{}
	if err := a.c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetNFTsForCollectionParam struct {
	ContractAddress     string `url:"contractAddress"`
	WithMetadata        bool   `url:"withMetadata"`
	StartToken          string `url:"startToken,omitempty"`
	Limit               int    `url:"limit,omitempty"`
	TokenUriTimeoutInMs int    `url:"tokenUriTimeoutInMs,omitempty"`
}

type GetNFTsForCollectionResponse struct {
	Nfts      []*Nft `json:"nfts"`
	NextToken string `json:"nextToken"`
}
