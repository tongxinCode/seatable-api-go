package seatable_api

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

const (
	serverURL = "http://172.28.197.89"
	baseToken = "673ce1c6dc71487c63e9e9e99a9397805c125d78"
	tableName = "Table3"
)

var base *Base
var rowID string

func TestMain(m *testing.M) {
	var err error
	base, err = getBase()
	if err != nil {
		fmt.Printf("fail to get seatable API: %v", err)
		os.Exit(1)
	}
	code := m.Run()
	os.Exit(code)
}

func getBase() (*Base, error) {
	base := Init(baseToken, serverURL)
	err := base.Auth(false)
	return base, err
}

func TestGet(t *testing.T) {
	_, err := base.GetMetadata()
	if err != nil {
		t.Errorf("failed to get metadata: %v", err)
		t.FailNow()
	}
}

func TestPost(t *testing.T) {
	rowData := make(map[string]interface{})
	rowData["Name"] = "name1"
	rowData["age"] = 20
	ret, err := base.AppendRow(tableName, rowData)
	if err != nil {
		t.Errorf("failed to append row: %v", err)
	}
	rowID, _ = ret["_id"].(string)
}

func TestPut(t *testing.T) {
	rowData := make(map[string]interface{})
	rowData["Name"] = "name2"
	rowData["age"] = 10
	_, err := base.UpdateRow(tableName, rowID, rowData)
	if err != nil {
		t.Errorf("failed to update row: %v", err)
	}
}

// need API token
func TestUploadLocalFile(t *testing.T) {
	_, err := base.UploadLocalFile("testfile.md", "testfile.md", "", "file", false)
	if err != nil {
		t.Errorf("failed to upload local file: %v", err)
	}
}

// need API token
func TestUploadBytesFile(t *testing.T) {
	r := bytes.NewReader([]byte("hello world"))
	_, err := base.UploadBytesFile("hello.md", r, "", "file", false)
	if err != nil {
		t.Errorf("failed to upload bytes file: %v", err)
	}
}

func TestDelete(t *testing.T) {
	_, err := base.DeleteRow(tableName, rowID)
	if err != nil {
		t.Errorf("failed to delete row: %v", err)
	}
}
