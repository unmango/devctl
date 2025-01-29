package renovate

// A CustomManager maps to the [renovate custom manager] object
//
// [renovate custom manager]: https://docs.renovatebot.com/configuration-options/#custommanagers
type CustomManager struct {
	CustomType         string   `json:"customType,omitempty"`
	DataSourceTemplate string   `json:"datasourceTemplate,omitempty"`
	DepNameTemplate    string   `json:"depNameTemplate,omitempty"`
	FileMatch          []string `json:"fileMatch,omitempty"`
	MatchStrings       []string `json:"matchStrings,omitempty"`
}
