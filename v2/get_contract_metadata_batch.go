package nft

type GetContractMetadataBatch struct {
	Api
}

func NewGetContractMetadataBatch(c *Client) *GetContractMetadataBatch {
	return &GetContractMetadataBatch{
		Api{c: c},
	}
}

func (a *GetContractMetadataBatch) Do(p *GetContractMetadataBatchParam) ([]*ContractWithMetadata, error) {
	req, resp := a.request("getContractMetadataBatch", MethodPost, p), []*ContractWithMetadata{}
	if err := a.c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetContractMetadataBatchParam struct {
	ContractAddresses []string `json:"contractAddresses"`
}
