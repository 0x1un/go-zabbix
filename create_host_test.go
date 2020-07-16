package zabbix

import (
	"fmt"
	"testing"
)

func TestCreateHost(t *testing.T) {
	session := GetTestSession(t)
	params := CreateHostRequest{
		Host:        "the test host",
		VisibleName: "测试主机",
		Status:      StatusDisabled,
		Interfaces: []Interface{
			{
				Main:  1,
				Type:  TypeAgent,
				Port:  10050,
				Bulk:  1,
				Useip: 1,
				IP:    "172.19.2.1",
				DNS:   "",
			},
		},
		Groups: []Group{
			{
				GroupID: "54",
			},
		},
	}
	resp, err := session.CreateHost(params)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
