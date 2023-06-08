package nft

type InvalidateContract struct {
	Api
}

func NewInvalidateContract(c *Client) *InvalidateContract {
	return &InvalidateContract{
		Api{c: c},
	}
}

func (a *InvalidateContract) Do(p *InvalidateContractParam) (*InvalidateContractResponse, error) {
	req, resp := a.request("invalidateContract", MethodGet, p), &InvalidateContractResponse{}
	if err := a.c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type InvalidateContractParam struct {
	ContractAddress string `url:"contractAddress"`
}

type InvalidateContractResponse struct {
	Success              bool `json:"contractAddress"`
	NumTokensInvalidated int  `json:"numTokensInvalidated"`
}
