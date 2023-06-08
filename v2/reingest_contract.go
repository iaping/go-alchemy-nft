package nft

type ReingestContract struct {
	Api
}

func NewReingestContract(c *Client) *ReingestContract {
	return &ReingestContract{
		Api{c: c},
	}
}

func (a *ReingestContract) Do(p *ReingestContractParam) (*ReingestContractResponse, error) {
	req, resp := a.request("reingestContract", MethodGet, p), &ReingestContractResponse{}
	if err := a.c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type ReingestContractParam struct {
	ContractAddress string `url:"contractAddress"`
}

type ReingestContractResponse struct {
	ContractAddress  string      `json:"contractAddress"`
	ReingestionState string      `json:"reingestionState"`
	Progress         interface{} `json:"progress"`
}
