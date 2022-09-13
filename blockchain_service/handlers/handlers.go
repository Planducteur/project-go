package handlers

import (
	"blockchain_service/models"
	"encoding/json"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	//"os"
	"strings"
	"time"
)

// read file and check
func CheckFormValue(w http.ResponseWriter, r *http.Request) (res bool, errStr string) {
	
	//read file and unmarshall json file to slice of users
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, "failed to read file: "+err.Error())
		return
	}

	var post wallet.Wallet

	err = json.Unmarshal(reqBody, &post)
	if err != nil {
		io.WriteString(w, "failed to read response body: "+ err.Error())
		return
	}

	//tests if string are valid
	if post.CurrencyCode == "" {
		return false, "the form is not complete"
	} else {
		//creationg slice of strings
		myslice := []string{post.CurrencyCode}
		//joyn slice par -
		result := strings.Join(myslice, "-")
		return true, result

	}

}

// create random address
const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

// function to add Wallet
func CreateNewPlayer(w http.ResponseWriter, r *http.Request) {

	newWallet := &wallet.Wallet{}

// vififaing metod used
	if r.Method == "POST" {
		resBool, errStr := CheckFormValue(w, r)
		//vreifaing the response
		if resBool == false {
			io.WriteString(w, errStr)
			return
		}

//split string joit previsly
		s := strings.Split(errStr, "-")
		if s[0] == "ethereum" {
			newWallet.CurrencyCode = "ETH"
		} else {
			newWallet.CurrencyCode = "ATH"
		}
		newWallet.WalletAddress = "0x" + String(45)
		newWallet.CurrencyBalance = 0

		newWalletBytes, err := json.Marshal(newWallet)

		if err != nil {
			newWalletBytes, _ := json.MarshalIndent(err, "", "")
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(newWalletBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(newWalletBytes)
		if err != nil {
			io.WriteString(w, "Unable to write response :"+err.Error())
			return
		}

	} else {
		io.WriteString(w, "metod used is not POST")
	}

}
