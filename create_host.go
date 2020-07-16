package zabbix

const (
	TypeAgent = 1
	TypeSNMP
	TypeIPMI
	TypeJMX

	StatusEnabled  = 0
	StatusDisabled = 1
)

type Interface struct {
	Main  int    `json:"main"`
	Type  int    `json:"type"`
	Port  int    `json:"port"`
	Bulk  int    `json:"bulk"`
	Useip int    `json:"useip"`
	IP    string `json:"ip"`
	DNS   string `json:"dns"`
}

type Group struct {
	GroupID string `json:"groupid"`
}

type CreateHostRequest struct {
	Status      int         `json:"status"`
	Host        string      `json:"host"`
	VisibleName string      `json:"name"`
	Description string      `json:"description"`
	Interfaces  []Interface `json:"interfaces"`
	Groups      []Group     `json:"groups"`
}

type CreateHostResponse struct {
	HostIDs []string `json:"hostids"`
}

func (c *Session) CreateHost(params CreateHostRequest) (CreateHostResponse, error) {
	var respData CreateHostResponse
	return respData, c.Get("host.create", params, &respData)
}
