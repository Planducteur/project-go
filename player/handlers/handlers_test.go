package handlers

import (
	"bytes"
	"encoding/json"
	//"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	Id       int    `json:"Id"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	PinCode  string `json:"PinCode"`
}

func Benchmark(b *testing.B) {

	httpposturl := "http://localhost/ethereum/wallets/create/"
	var1, _ := json.Marshal(User{
		Username: "Bob",
		Password: "123456%", // doit accepter les caractères spéciaux dans le mot de passe
		PinCode:  "123456"})
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

	for _, tt := range tests {
		for i := 0; i < b.N; i++ {
			CreateNewPlayer(tt.args.w, tt.args.r)
		}

	}

}

func TestCreateNewPlayer(t *testing.T) {

	httpposturl := "http://localhost/ethereum/wallets/create/"
	var1, _ := json.Marshal(User{
		Username: "Bob",
		Password: "123456%", // doit accepter les caractères spéciaux dans le mot de passe
		PinCode:  "123456"})
	r1, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(var1))

	var2, _ := json.Marshal(User{
		Username: "Bob#", // user contient des caractères spéciaux
		Password: "123456",
		PinCode:  "123456"})
	r2, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(var2))

	var3, _ := json.Marshal(User{
		Username: "Bob",
		Password: "12", // mot de passe trop court (6-32 characters)
		PinCode:  "123456"})
	r3, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(var3))

	var4, _ := json.Marshal(User{
		Username: "Bob",
		Password: "123456",
		PinCode:  "1234", // pin trop court (6 digits (0-9))
	})
	r4, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(var4))

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
		{
			name: "caractères spéciaux user",
			args: args{w: httptest.NewRecorder(), r: r2},
		},
		{
			name: "mot de passe trop court",
			args: args{w: httptest.NewRecorder(), r: r3},
		},
		{
			name: "pin trop court",
			args: args{w: httptest.NewRecorder(), r: r4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateNewPlayer(tt.args.w, tt.args.r)
		})
	}
}
