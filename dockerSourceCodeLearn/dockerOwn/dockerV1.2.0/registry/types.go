package registry

type SearchResult struct{
	StarCount int	`json:"star_count"`
	IsOfficial	bool	`json:"is_official"`
	Name 	string	`json:"name"`
	IsTrusted	bool	`json:"is_trusted"`
	Description string `json:"description"`
}

type SearchResults struct{
	Quert string `json:"query"`
	NumResults int	`json:"num_results"`
	Results []SearchResult `json:"results"`
}

type RepositoryData struct{
	ImgList map[string]*ImgData
	Endpoints []string
	Tokens []string
}

type ImgData struct{
	ID string `json:"id"`
	CheckSum string `json:"checksum,omitempty"`
	CheckSumPayload string `json:"-"`
	Tag string `json:",omitempty"`
}

type RegistryInfo struct{
	Version string `json:"version"`
	Standalone bool `json:"standalone"`
}