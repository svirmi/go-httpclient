package main

import (
	"fmt"

	"github.com/vanng822/go-solr/solr"
)

func main() {
	si, _ := solr.NewSolrInterface("http://localhost:8983/solr", "collection1")

	query := solr.NewQuery()
	query.Q("*:*")
	s := si.Search(query)
	r, _ := s.Result(nil)
	fmt.Println(r.Results.Docs)
	fmt.Println("Hello World!")
}
