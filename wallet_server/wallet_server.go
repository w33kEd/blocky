package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"text/template"

	"github.com/w33ked/go-blocky/utils"
	"github.com/w33ked/go-blocky/wallet"
)

const tempDir = "templates"

type WalletServer struct {
	port    uint16
	gateway string
}

func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

func (ws *WalletServer) Port() uint16 {
	return ws.port
}

func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(tempDir, "index.html"))
		t.Execute(w, "")
	default:
		log.Printf("Err: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		myWallet := wallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		w.Write(m[:])
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERR: Invalid HTTP MEthod")
	}
}

func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(req.Body)
		var t wallet.TransactionRequest
		err := decoder.Decode(&t)
		if err != nil {
			log.Printf("ERR! %v", err)
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		if !t.Validate() {
			log.Printf("ERR! Missing Fields!")
			io.WriteString(w, string(utils.JsonStatus("fail")))
			return
		}

		fmt.Println(*t.SenderPublicKey)
		fmt.Println(*t.SenderBlockchainAddress)
		fmt.Println(*t.SenderPrivateKey)
		fmt.Println(*t.RecipientBlockchainAddress)
		fmt.Println(*t.Value)
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERR! Invalid HTTP Method")
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
