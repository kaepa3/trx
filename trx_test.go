package trx

import (
	"log"
	"testing"
)

func TestParse(t *testing.T) {
	obj, err := Parse("test.trx")
	if err != nil {
		t.Errorf("Parse() , want %v", err)
	}
	log.Printf("%v", obj)
}
