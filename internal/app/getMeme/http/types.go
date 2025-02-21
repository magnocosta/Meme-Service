package http

type request struct {
  Lat   string
  Lon  string
  Query string
}

func (d request) GetLat()   string { return d.Lat   }
func (d request) GetLon()   string { return d.Lon   }
func (d request) GetQuery() string { return d.Query }

type item struct {
  URL string `json:"url"`
}

type response struct {
  Items []item `json:"items"`
}

type responseError struct {
  message string
}
