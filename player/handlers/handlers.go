package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"player/models"
	"regexp"
	"strings"
)

// function to add user
func CreateNewPlayer(w http.ResponseWriter, r *http.Request) {

	newUser := &player.User{}

	// vififaing metod used
	if r.Method == "POST" {

		//call to read file and check
		resBool, errStr := CheckFormValue(w, r)

		//vreifaing the response
		if resBool == false {
			io.WriteString(w, errStr)
			return
		}

		//split string joit previsly
		s := strings.Split(errStr, "-")

		newUser.Username = s[0]
		newUser.Password = s[1]
		newUser.PinCode = s[2]

		//open file
		file, err := os.OpenFile("user.json", os.O_RDWR, 0644)
		if err != nil {
			io.WriteString(w, "failed to open user.json file: "+err.Error())
			return
		}
		defer file.Close()

		//read file and unmarshall json file to slice of users
		b, _ := ioutil.ReadAll(file)
		var alUsrs player.AllUsers
		err = json.Unmarshal(b, &alUsrs.Users)

		max := 0

		//verifain if user aredady exist
		//generation of id(last id at the json file+1)
		for _, usr := range alUsrs.Users {
			if usr.Username == s[0] {
				io.WriteString(w, "this username aredy exist\n")
				fmt.Println("this username aredy exist")
				return
			}
			if usr.Id > max {
				max = usr.Id
			}
		}

		id := max + 1
		newUser.Id = id

		//appending newUser to slice of all Users and rewrite json file
		alUsrs.Users = append(alUsrs.Users, newUser)
		newUserBytes, err := json.MarshalIndent(&alUsrs.Users, "", " ")
		if err != nil {
			io.WriteString(w, "failed to marshal newUser: "+err.Error())
			return
		}
		ioutil.WriteFile("user.json", newUserBytes, 0666)

		//split url to get the currency code
		blo := strings.Split(r.URL.Path, "/")

		postBody, _ := json.Marshal(map[string]string{
			"CurrencyCode": blo[1],
			"PinCode":      s[1],
		})

		responseBody := bytes.NewBuffer(postBody)
		resp, err := http.Post("http://localhost:8091/create/", "application/json", responseBody)
		//Handle Error
		if err != nil {
			io.WriteString(w, "failed to post create wallet: "+err.Error())
			return
		}
		defer resp.Body.Close()

		//Read the response body
		body1, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			io.WriteString(w, "failed to read response body: "+err.Error())
			return
		}
		io.WriteString(w, "player addet successfully\n")
		CreateNewWallet(w, body1, id, s[2])


	} else {
		io.WriteString(w, "metod used is not POST")
	}

}

// read file and check
func CheckFormValue(w http.ResponseWriter, r *http.Request) (res bool, errStr string) {

	//read file and unmarshall json file to slice of users
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, "failed to read file: "+err.Error())
		return
	}

	var post player.User

	err = json.Unmarshal(reqBody, &post)
	if err != nil {
		io.WriteString(w, "failed to read response body: "+err.Error())
		return
	}

	//tests whether a pattern matches a string
	usr := regexp.MustCompile(`^[a-z0-9_]{3,100}$`)
	pas := regexp.MustCompile(`^.{6,32}$`)
	cod := regexp.MustCompile(`^[0-9]{6}$`)

	//tests if string are valid
	if post.Username == "" || post.Password == "" || post.PinCode == "" {
		return false, "the form is not complete"
	} else {
		if !usr.MatchString(post.Username) || !pas.MatchString(post.Password) || !cod.MatchString(post.PinCode) {
			return false, "you must respect the input form: username(3-100 characters lowercase, a-z,0-9 or _), password (6-32 characters), pincode(6 digits 0-9)"
		} else {
			//creationg slice of strings
			slice := []string{post.Username, post.Password, post.PinCode}
			//joyn slice par -
			result := strings.Join(slice, "-")
			return true, result
		}
	}
}

// create wallet
func CreateNewWallet(w http.ResponseWriter, body []byte, value int, pin string) {
	
	newWallet := &player.Wallet{}
	
	var wallet player.Wallet

	json.Unmarshal(body, &wallet)

	

	newWallet.WalletAddress = wallet.WalletAddress
	newWallet.CurrencyBalance = wallet.CurrencyBalance
	newWallet.CurrencyCode = wallet.CurrencyCode
/*
	for _, wall := range wallet.Wallets {
		newWallet.WalletAddress = wall.WalletAddress
		newWallet.CurrencyBalance = wall.CurrencyBalance
		newWallet.CurrencyCode = wall.CurrencyCode
	}*/

	newWallet.PlayerId = value
	newWallet.PinCode = pin

	//sb := string(body)
	//fmt.Printf(sb)

	wal, err := os.OpenFile("wallet.json", os.O_RDWR, 0644)
	if err != nil {
		io.WriteString(w, "failed to open wallet.json file: "+err.Error())
		//return
	}
	defer wal.Close()

	//read file and unmarshall json file to slice of users
	a, _ := ioutil.ReadAll(wal)
	var alWallet player.AllWallets
	err = json.Unmarshal(a, &alWallet.Wallets)


	alWallet.Wallets = append(alWallet.Wallets, newWallet)
	newWalletBytes, err := json.MarshalIndent(&alWallet.Wallets, "", " ")
	if err != nil {
		io.WriteString(w, "failed to marshal newWallet: "+err.Error())
		return
	}
	ioutil.WriteFile("wallet.json", newWalletBytes, 0666)
	io.WriteString(w, "wallet addet successfully\n")
}
