package nft

type GetNFTMetadataBatch struct {
	Api
}

func NewGetNFTMetadataBatch(c *Client) *GetNFTMetadataBatch {
	return &GetNFTMetadataBatch{
		Api{c: c},
	}
}

func (a *GetNFTMetadataBatch) Do(p *GetNFTMetadataBatchParam) ([]*Nft, error) {
	req, resp := a.request("getNFTMetadataBatch", MethodPost, p), []*Nft{}
	if err := a.c.Do(req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetNFTMetadataBatchParam struct {
	Tokens              []TokenParam `json:"tokens"`
	TokenUriTimeoutInMs int          `json:"tokenUriTimeoutInMs,omitempty"`
	RefreshCache        bool         `json:"refreshCache"`
}

type TokenParam struct {
	ContractAddress string `json:"contractAddress"`
	TokenId         string `json:"tokenId"`
	TokenType       string `json:"tokenType"`
}
