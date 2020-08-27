package zabbix

import (
	"fmt"
	"testing"
)

func TestGraphGet(t *testing.T) {
	session := GetTestSession(t)
	gf := GraphGetParameters{}
	gf.Output = []string{"name"}
	// gf.Output = "extend"
	gf.Hostids = []string{"10673"}
	gf.SortField = []string{"name"}
	res, err := session.GraphGet(gf)
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range res {
		fmt.Printf("%d=%s\n", i, v.Graphid)
	}
}
