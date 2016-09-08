package main

import (
	"fmt"
	"testing"
)

func TestDescTable(t *testing.T) {
	conf := &DbConf{
		Host:     "rm-2ze75re47nby5u7mz.mysql.rds.aliyuncs.com",
		Port:     3306,
		UserName: "dsp_sys",
		Password: "Dsp_sys-pwd",
		DbName:   "dsp_sys",
		Charset:  "utf8",
	}

	table := descTable(conf, "ad_plan")
	fmt.Printf("%+x\n", table)
}
