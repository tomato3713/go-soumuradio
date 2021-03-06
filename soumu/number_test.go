package soumu

import (
	"context"
	"testing"
)

func TestGetNum(t *testing.T) {
	DebugEnable()
	cli, err := NewClient("")
	if err != nil {
		t.Fatalf("fault to make new client: %v", err)
	}

	opts := NewNumOptions(License, Amateur)
	got, err := cli.GetNum(context.Background(), opts)
	if err != nil {
		t.Fatalf("fault to getNum returned err: %v", err)
	}

	want := &Num{
		Musen: Musen{
			Count: "1200",
		},
		MusenInformation: MusenInformation{
			TotalCount:     "1200",
			LastUpdateDate: "2019/02/12",
		},
	}
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}
