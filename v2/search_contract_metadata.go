package nft

type SearchContractMetadata struct {
	Api
}

func NewSearchContractMetadata(c *Client) *SearchContractMetadata {
	return &SearchContractMetadata{
		Api{c: c},
	}
}

func (a *SearchContractMetadata) Do(p *SearchContractMetadataParam) ([]*SearchContractMetadataResponse, error) {
	req, resp := a.request("searchContractMetadata", MethodGet, p), []*SearchContractMetadataResponse{}
	if err := a.c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type SearchContractMetadataParam struct {
	Query string `url:"query"`
}

type SearchContractMetadataResponse struct {
	Address          string            `json:"address"`
	ContractMetadata *ContractMetadata `json:"contractMetadata"`
}
