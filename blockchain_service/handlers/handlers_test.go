package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	//"math/rand"
)

type Wallet struct {
	CurrencyCode   string  
	PinCode 		string 
}


func Benchmark(b *testing.B) {
	httpposturl := "http://localhost:8091/create/"
	var1, _ := json.Marshal(Wallet{
		CurrencyCode: "",
		PinCode: "123456",}) 
	r1, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(var1))

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "caractères spéciaux password",
			args: args{w: httptest.NewRecorder(), r: r1},
		},
		
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
				CreateNewPlayer(tt.args.w, tt.args.r)
		}
}
}

func TestCreateNewPlayer(t *testing.T) {
	httpposturl := "http://localhost:8091/create/"
	var1, _ := json.Marshal(Wallet{
		CurrencyCode: "",
		PinCode: "123456",}) 
	r1, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(var1))

	var2, _ := json.Marshal(Wallet{
		CurrencyCode: "ethereum", 
		PinCode:  ""})
	r2, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(var2))

	var3, _ := json.Marshal(Wallet{
		CurrencyCode: "ethereum",
		PinCode:  "123456"})
	r3, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(var3))

	var4, _ := json.Marshal(Wallet{
		CurrencyCode: "Bob",
		PinCode:  "123456"})
	r4, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(var4))

	var5, _ := json.Marshal(Wallet{
		CurrencyCode: "Bob",
		PinCode:  "123456"})
	r5, _ := http.NewRequest("GET", httpposturl, bytes.NewBuffer(var5))

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{w: httptest.NewRecorder(), r: r1},
		},
		{
			name: "test2",
			args: args{w: httptest.NewRecorder(), r: r2},
		},
		{
			name: "test3",
			args: args{w: httptest.NewRecorder(), r: r3},
		},
		{
			name: "test4",
			args: args{w: httptest.NewRecorder(), r: r4},
		},
		{
			name: "test4",
			args: args{w: httptest.NewRecorder(), r: r5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateNewPlayer(tt.args.w, tt.args.r)
		})
	}
}
