package nft

type GetNFTs struct {
	Api
}

func NewGetNFTs(c *Client) *GetNFTs {
	return &GetNFTs{
		Api{c: c},
	}
}

func (a *GetNFTs) Do(p *GetNFTsParam) (*GetNFTsResponse, error) {
	req, resp := a.request("getNFTs", MethodGet, p), &GetNFTsResponse{}
	if err := a.c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetNFTsParam struct {
	Owner               string   `url:"owner"`
	ContractAddresses   []string `url:"contractAddresses[],omitempty"`
	WithMetadata        bool     `url:"withMetadata"`
	ExcludeFilters      []string `url:"excludeFilters[],omitempty"`
	IncludeFilters      []string `url:"includeFilters[],omitempty"`
	SpamConfidenceLevel string   `url:"spamConfidenceLevel,omitempty"`
	TokenUriTimeoutInMs int      `url:"tokenUriTimeoutInMs,omitempty"`
	OrderBy             string   `url:"orderBy,omitempty"`
	PageKey             string   `url:"pageKey,omitempty"`
	PageSize            int      `url:"pageSize,omitempty"`
}

type GetNFTsResponse struct {
	OwnedNfts  []*Nft `json:"ownedNfts"`
	TotalCount int64  `json:"totalCount"`
	BlockHash  string `json:"blockHash"`
}
