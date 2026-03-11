package main

type Audit struct {
	CreatedAt string
	UpdatedAt string
}

type Article struct {
	Title string
	Audit
}
