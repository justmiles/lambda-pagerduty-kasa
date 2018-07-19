package main

type Request struct {
	Messages []struct {
		Event string `json:"event"`
	} `json:"messages"`
}
