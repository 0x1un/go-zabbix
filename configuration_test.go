package zabbix

import (
	"fmt"
	"testing"
)

// example: 导出所有的主机配置
func TestConfiguraExport(t *testing.T) {
	session := GetTestSession(t)
	hostParams := HostGetParams{}
	hostParams.OutputFields = []string{"hostid"}

	hosts, err := session.GetHosts(hostParams)
	if err != nil {
		t.Fatalf("Error getting Hosts: %v", err)
	}

	if len(hosts) == 0 {
		t.Fatal("No Hosts found")
	}
	hostIDS := make([]string, 0)
	for i, host := range hosts {
		if host.HostID == "" {
			t.Fatalf("Host %d returned in response body has no Host ID", i)
		}
		if len(host.HostID) == 0 {
			continue
		}
		hostIDS = append(hostIDS, host.HostID)
	}

	params := ConfigurationParamsRequest{
		Format: "json",
		Options: ConfiguraOption{
			Hosts: hostIDS,
		},
	}
	respData, err := session.ConfiguraExport(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(respData)
}
