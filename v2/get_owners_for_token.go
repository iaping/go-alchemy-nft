package nft

type GetOwnersForToken struct {
	Api
}

func NewGetOwnersForToken(c *Client) *GetOwnersForToken {
	return &GetOwnersForToken{
		Api{c: c},
	}
}

func (a *GetOwnersForToken) Do(p *GetOwnersForTokenParam) (*GetOwnersForTokenResponse, error) {
	req, resp := a.request("getOwnersForToken", MethodGet, p), &GetOwnersForTokenResponse{}
	if err := a.c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetOwnersForTokenParam struct {
	ContractAddress string `url:"contractAddress"`
	TokenId         string `url:"tokenId"`
	PageKey         string `url:"pageKey,omitempty"`
	PageSize        int    `url:"pageSize,omitempty"`
}

type GetOwnersForTokenResponse struct {
	Owners []string `json:"owners"`
}
