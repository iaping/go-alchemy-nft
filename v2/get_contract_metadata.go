package nft

type GetContractMetadata struct {
	Api
}

func NewGetContractMetadata(c *Client) *GetContractMetadata {
	return &GetContractMetadata{
		Api{c: c},
	}
}

func (a *GetContractMetadata) Do(p *GetContractMetadataParam) (*ContractWithMetadata, error) {
	req, resp := a.request("getContractMetadata", MethodGet, p), &ContractWithMetadata{}
	if err := a.c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetContractMetadataParam struct {
	ContractAddress string `url:"contractAddress"`
}
