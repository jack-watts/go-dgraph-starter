package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	var data = struct {
		Status string `json:"status"`
		URL    string `json:"urls"`
	}{
		Status: "active",
		URL:    "/query",
	}

	writeJSON(w, data)
}

func (app *Application) query(w http.ResponseWriter, r *http.Request) {
	// create a new transaction
	txn := app.DG.NewReadOnlyTxn()

	// define our query
	const q = `{
		all(func: has(Product.reviews), first: 10) {
		  uid
		  expand(_all_)
		}
	  }`

	// pass our query as a transaction to get a response
	resp, err := txn.Query(context.Background(), q)
	if err != nil {
		log.Printf("error returning query response body: %s\n", err)
		return
	}

	var data interface{}

	// unmarshal the unstructured data to the 'data' interface.
	if err := json.Unmarshal(resp.GetJson(), &data); err != nil {
		log.Printf("%s\n", err)
	}

	writeJSON(w, data)

}
