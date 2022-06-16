package godeng

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	pwd, _ := os.Getwd()
	fllepath := fmt.Sprintf("%s/configs/example.json", pwd)
	cfg, err := Parser(fllepath)
	if err != nil {
		t.Error(err)
	}
	log.Println(cfg)
}
