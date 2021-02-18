package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// https://gateway.marvel.com/v1/public/characters?nameStartsWith=iron&ts=1613547099&apikey=8138eacb2e1fa7025de15909bab7689c&hash=8458555f4c5c6af52ee556ee5840867e

func get_data(query string) string {
	timestamp := fmt.Sprint(time.Now().Unix())
	privateKey := "4b2cedb1907508e67ec68cae0fc260fae866b216"
	publicKey := "8138eacb2e1fa7025de15909bab7689c"

	md5Hash := getMD5Hash(timestamp + privateKey + publicKey)
	// characters?nameStartsWith=iron
	url := "https://gateway.marvel.com:443/v1/public/" + query + "&ts=" + timestamp + "&apikey=" + publicKey + "&hash=" + md5Hash
	fmt.Println(url)
	fmt.Println("Fetching from the api....")

	response, _ := http.Get(url)

	dataBytes, _ := ioutil.ReadAll(response.Body)
	data := string(dataBytes[:])
	return data
}
