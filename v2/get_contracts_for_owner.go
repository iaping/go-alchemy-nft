package nft

type GetContractsForOwner struct {
	Api
}

func NewGetContractsForOwner(c *Client) *GetContractsForOwner {
	return &GetContractsForOwner{
		Api{c: c},
	}
}

func (a *GetContractsForOwner) Do(p *GetContractsForOwnerParam) (*GetContractsForOwnerResponse, error) {
	req, resp := a.request("getContractsForOwner", MethodGet, p), &GetContractsForOwnerResponse{}
	if err := a.c.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type GetContractsForOwnerParam struct {
	Owner               string   `url:"owner"`
	WithMetadata        bool     `url:"withMetadata"`
	ExcludeFilters      []string `url:"excludeFilters[],omitempty"`
	IncludeFilters      []string `url:"includeFilters[],omitempty"`
	SpamConfidenceLevel string   `url:"spamConfidenceLevel,omitempty"`
	OrderBy             string   `url:"orderBy,omitempty"`
	PageKey             string   `url:"pageKey,omitempty"`
	PageSize            int      `url:"pageSize,omitempty"`
}

type GetContractsForOwnerResponse struct {
	Contracts  []*Contract `json:"contracts"`
	TotalCount int64       `json:"totalCount"`
}
