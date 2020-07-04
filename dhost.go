package zabbix

// Retrieve discovered hosts by discovery rule response data
type DiscoveryGetHostResponse []struct {
	Dhostid   string `json:"dhostid"`
	Druleid   string `json:"druleid"`
	Dservices []struct {
		Dcheckid   string `json:"dcheckid"`
		Dhostid    string `json:"dhostid"`
		DNS        string `json:"dns"`
		Dserviceid string `json:"dserviceid"`
		IP         string `json:"ip"`
		Key        string `json:"key_"`
		Lastdown   string `json:"lastdown"`
		Lastup     string `json:"lastup"`
		Port       string `json:"port"`
		Status     string `json:"status"`
		Type       string `json:"type"`
		Value      string `json:"value"`
	} `json:"dservices"`
	Lastdown string `json:"lastdown"`
	Lastup   string `json:"lastup"`
	Status   string `json:"status"`
}

type DiscoveryGetHostParamsRequest struct {
	LimitSelect     int64       `json:"limitSelects"`
	CountOoutput    bool        `json:"countOutput"`
	Editable        bool        `json:"editable"`
	ExcludeSearch   bool        `json:"excludeSearch"`
	PreserveKeys    bool        `json:"preservekeys"`
	Dhostids        []string    `json:"dhostids"`
	Druleids        []string    `json:"druleids"`
	Dserviceids     []string    `json:"dserviceids"`
	SortField       []string    `json:"sortfield"`
	SelectDRules    SelectQuery `json:"selectDRules"`
	SelectDServices SelectQuery `json:"selectDServices"`
	Ouput           SelectQuery `json:"output"`
	GetParameters
}

func (c *Session) GetDHosts(params DiscoveryGetHostParamsRequest) (DiscoveryGetHostResponse, error) {
	respData := make(DiscoveryGetHostResponse, 0)
	err := c.Get("dhost.get", params, &respData)
	return respData, err
}
