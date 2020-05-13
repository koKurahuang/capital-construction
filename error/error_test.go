package error

import (
	"fmt"
	"github.com/koKurahuang/capital-construction/config"
	"os"
	"testing"
)

func TestGojson(t *testing.T) {
	os.Setenv("ZUES_ERROR_JSON", "C:\\GOPATH\\src\\github.com\\koKurahuang\\capital-construction\\error\\errors.json")
	os.Setenv("ZUES_ERROR_LANG", "ZH")
	os.Setenv("ZUES_RULE","C:\\GOPATH\\src\\github.com\\koKurahuang\\capital-construction\\config")
	config.Initialize()
	Initilize()
	err := New("9999", "0001", 123456, 789)
	fmt.Print(err)
}

func TestPrint(t *testing.T) {
	a := []interface{}{"1", "2"}
	fmt.Printf("one %s, two %s", a...)
}

func TestCallStack( t *testing.T) {
	os.Setenv("ZUES_ERROR_JSON", "C:\\GOPATH\\src\\github.com\\koKurahuang\\capital-construction\\error\\errors.json")
	os.Setenv("ZUES_ERROR_LANG", "ZH")
	os.Setenv("ZUES_RULE","C:\\GOPATH\\src\\github.com\\koKurahuang\\capital-construction\\config")
	config.Initialize()
	Initilize()

	err :=  c()
	fmt.Println(err)
}

func a() error {
	return New("9999", "0001", 123456, 789)
}

func b() error {
	return a()
}

func c() error {
	return b()
}