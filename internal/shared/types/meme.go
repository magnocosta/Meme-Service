package types

type Meme struct {
  URL string  `json:"url"` 
}

func (m Meme) GetURL() string { return m.URL }
