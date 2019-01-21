package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/olivere/elastic"
)

type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

func main() {
	// Create a client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		// Handle error
	}
	esversion, err := client.ElasticsearchVersion("http://localhost:9200")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Create an index
	//_, err = client.CreateIndex("tweets").Do(context.Background())
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}

	// Add a document to the index
	tweet := Tweet{User: "olivere", Message: "Take Five"}
	_, err = client.Index().
		Index("tweets").
		Type("doc").
		Id("1").
		BodyJson(tweet).
		Refresh("wait_for").
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	// Search with a term query
	termQuery := elastic.NewTermQuery("user", "olivere")
	searchResult, err := client.Search().
		Index("tweets").            // search in index "tweets"
		Query(termQuery).           // specify the query
		Sort("user.keyword", true). // sort by "user" field, ascending
		From(0).Size(10).           // take documents 0-9
		Pretty(true).               // pretty print request and response JSON
		Do(context.Background())    // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Each is a convenience function that iterates over hits in a search result.
	// It makes sure you don't need to check for nil values in the response.
	// However, it ignores errors in serialization. If you want full control
	// over iterating the hits, see below.
	var ttyp Tweet
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	}
	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d tweets\n", searchResult.TotalHits())

	// Here's how you iterate through results with full control over each step.
	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Tweet (could also be just a map[string]interface{}).
			var t Tweet
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with tweet
			fmt.Printf("Tweet by %s: %s\n", t.User, t.Message)
		}
	} else {
		// No hits
		fmt.Print("Found no tweets\n")
	}

	// Delete the index again
	//_, err = client.DeleteIndex("tweets").Do(context.Background())
	//if err != nil {
	//	// Handle error
	//	panic(err)
	//}
}
