package lib

// EnvironmentVariables ...
type EnvironmentVariables struct {
	DbFile                   string
	DatabaseName             string
	DatabaseAttachedFullPath string
	DefaultPort              string
	HasMultiplePort          bool
}

// IPResponse ...
type IPResponse struct {
	Origin string `json:"origin"`
}

// UserAgentResponse ...
type UserAgentResponse struct {
	UserAgent string `json:"user-agent"`
}

// HeadersResponse ...
type HeadersResponse struct {
	Headers map[string]string `json:"headers"`
}

// CookiesResponse ...
type CookiesResponse struct {
	Cookies map[string]string `json:"cookies"`
}

// JSONResponse ...
type JSONResponse interface{}

// GetResponse ...
type GetResponse struct {
	Args map[string][]string `json:"args"`
	HeadersResponse
	IPResponse
	URL string `json:"url"`
}

// PostResponse ...
type PostResponse struct {
	Args map[string][]string `json:"args"`
	Data JSONResponse        `json:"data"`
	Form map[string]string   `json:"form"`
	HeadersResponse
	IPResponse
	URL string `json:"url"`
}

// GzipResponse ...
type GzipResponse struct {
	HeadersResponse
	IPResponse
	Gzipped bool `json:"gzipped"`
}

// DeflateResponse ...
type DeflateResponse struct {
	HeadersResponse
	IPResponse
	Deflated bool `json:"deflated"`
}

// BasicAuthResponse ...
type BasicAuthResponse struct {
	Authenticated bool   `json:"authenticated"`
	User          string `json:"string"`
}

// WebCliController ...
type WebCliController struct {
	Port string
}

// APIDataModel ...
type APIDataModel struct {
	ID       int    `json:"id"`
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
	Body     string `json:"body"`
}

// EndpointModel ...
type EndpointModel struct {
	Key      string `json:"key"`
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
}

// Pair ...
type Pair struct {
	Key   string
	Value APIDataModel
}

// PairList ...
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value.ID > p[j].Value.ID }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// JsTreeDataModel ...
type JsTreeDataModel struct {
	ID       int               `json:"id"`
	Key      string            `json:"key"`
	Text     string            `json:"text"`
	Type     string            `json:"type"`
	Children []JsTreeDataModel `json:"children"`
}
