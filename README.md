[![GoDoc](https://godoc.org/github.com/nickvanw/gifgo?status.svg)](https://godoc.org/github.com/nickvanw/gifgo)
# gifgo

gifgo is a [giphy](https://giphy.com) API client. Using it can be nice and easy:


```
client, _ := gifgo.New()

query := "New York City"
searchRequest := gifgo.SearchReq{Query: query}
results, _ := gifgo.Search(searchRequest)
// use the results, luke!
```

A much more complete example can be found in the giphy command line in `cmd/giphy`.