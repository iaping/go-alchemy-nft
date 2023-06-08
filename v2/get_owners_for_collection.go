package nft

type GetOwnersForCollection struct {
	Api
}

func NewGetOwnersForCollection(c *Client) *GetOwnersForCollection {
	return &GetOwnersForCollection{
		Api{c: c},
	}
}

func (a *GetOwnersForCollection) Do(p *GetOwnersForCollectionParam) (interface{}, error) {
	var resp interface{}
	if p.WithTokenBalances {
		resp = &GetOwnersForCollectionWithTokenBalancesResponse{}
	} else {
		resp = &GetOwnersForCollectionResponse{}
	}
	req := a.request("getOwnersForCollection", MethodGet, p)
	if err := a.c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetOwnersForCollectionParam struct {
	ContractAddress   string `url:"contractAddress"`
	WithTokenBalances bool   `url:"withTokenBalances"`
	Block             string `url:"block,omitempty"`
	PageKey           string `url:"pageKey,omitempty"`
}

type GetOwnersForCollectionResponse struct {
	OwnerAddresses []string `json:"ownerAddresses"`
}

type GetOwnersForCollectionWithTokenBalancesResponse struct {
	OwnerAddresses []struct {
		OwnerAddress  string          `json:"ownerAddress"`
		TokenBalances []*TokenBalance `json:"tokenBalances"`
	} `json:"ownerAddresses"`
}

type TokenBalance struct {
	TokenId string `json:"tokenId"`
	Balance int    `json:"balance"`
}
