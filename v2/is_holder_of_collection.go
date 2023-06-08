package nft

type IsHolderOfCollection struct {
	Api
}

func NewIsHolderOfCollection(c *Client) *IsHolderOfCollection {
	return &IsHolderOfCollection{
		Api{c: c},
	}
}

func (a *IsHolderOfCollection) Do(p *IsHolderOfCollectionParam) (*IsHolderOfCollectionResponse, error) {
	req, resp := a.request("isHolderOfCollection", MethodGet, p), &IsHolderOfCollectionResponse{}
	if err := a.c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type IsHolderOfCollectionParam struct {
	Wallet          string `url:"wallet"`
	ContractAddress string `url:"contractAddress"`
}

type IsHolderOfCollectionResponse struct {
	IsHolderOfCollection bool `json:"isHolderOfCollection"`
}
