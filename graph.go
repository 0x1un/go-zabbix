package zabbix

type GraphGetParameters struct {
	Templated              bool                   `json:"templated"`
	Inherited              bool                   `json:"inherited"`
	StartSearch            bool                   `json:"startSearch"`
	SearchByAny            bool                   `json:"searchByAny"`
	SearchWildcardsEnabled bool                   `json:"searchWildcardsEnabled"`
	CountOutput            bool                   `json:"countOutput"`
	Editable               bool                   `json:"editable"`
	ExcludeSearch          bool                   `json:"excludeSearch"`
	Preservekeys           bool                   `json:"preservekeys"`
	Limit                  int64                  `json:"limit"`
	ExpandName             string                 `json:"expandName"`
	Graphids               []string               `json:"graphids"`
	Groupids               []string               `json:"groupids"`
	SortField              []string               `json:"sortField"`
	Templateids            []string               `json:"templateids"`
	Hostids                []string               `json:"hostids"`
	Itemids                []string               `json:"itemids"`
	Sortorder              []string               `json:"sortorder"`
	SelectGroups           SelectQuery            `json:"selectGroups"`
	SelectTemplates        SelectQuery            `json:"selectTemplates"`
	SelectHosts            SelectQuery            `json:"selectHosts"`
	SelectItems            SelectQuery            `json:"selectItems"`
	SelectGraphDiscovery   SelectQuery            `json:"selectGraphDiscovery"`
	SelectDiscoveryRule    SelectQuery            `json:"selectDiscoveryRule"`
	Output                 SelectQuery            `json:"output"`
	Filter                 map[string]interface{} `json:"filter"`
	Search                 map[string]interface{} `json:"search"`
}

type GraphResult []struct {
	Graphid        string `json:"graphid"`
	Name           string `json:"name"`
	Width          string `json:"width"`
	Height         string `json:"height"`
	Yaxismin       string `json:"yaxismin"`
	Yaxismax       string `json:"yaxismax"`
	Templateid     string `json:"templateid"`
	ShowWorkPeriod string `json:"show_work_period"`
	ShowTriggers   string `json:"show_triggers"`
	Graphtype      string `json:"graphtype"`
	ShowLegend     string `json:"show_legend"`
	Show3D         string `json:"show_3d"`
	PercentLeft    string `json:"percent_left"`
	PercentRight   string `json:"percent_right"`
	YminType       string `json:"ymin_type"`
	YmaxType       string `json:"ymax_type"`
	YminItemid     string `json:"ymin_itemid"`
	YmaxItemid     string `json:"ymax_itemid"`
	Flags          string `json:"flags"`
}

func (c *Session) GraphGet(params GraphGetParameters) (GraphResult, error) {
	var respData GraphResult
	if err := c.Get("graph.get", params, &respData); err != nil {
		return respData, err
	} else {
		return respData, err
	}
}
