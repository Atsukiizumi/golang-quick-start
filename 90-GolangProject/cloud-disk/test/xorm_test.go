package test

import (
	"GolangProject/cloud-disk/core/define"
	"GolangProject/cloud-disk/core/models"
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"xorm.io/xorm"
)

func TestXormTest(t *testing.T) {
	engine, err := xorm.NewEngine("mysql", define.MySQLConn)
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.UserBasic, 0)
	err = engine.Find(&data)
	if err != nil {
		t.Fatal(err)
	}

	b, err := json.Marshal(data)
	if err != nil {
		t.Fatal(err)
	}

	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(dst.String())
}
