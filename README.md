# web2markdown-api

API to convert any content from the web to API.

## Motivation

Do you have any service/system that require to scrap a content from the web and convert them nicely to markdown easily? then you might find `web2markdown-api` is useful. 

With this API, you can convert any web content that you want into a markdown beautifully.

## Usage

The API is deployed using Vercel serverless function. Here is the example on how you can use the API

```bash
curl --location --request POST 'https://web2markdown-api.vercel.app/api/convert' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "https://kirandev.com/http-post-golang"
}'
```

It will return a JSON API response:

```JSON
{
    "url": "https://kirandev.com/http-post-golang",
    "title": "Make an HTTP POST request in Go",
    "data": "January 23, 2022\n\nWe will make a POST request to an endpoint with a JSON body and display the results on the console.\n\nThe endpoint will accept `id`, `title`, `body`, `userId` and create a new `post`.\n\nCreate a new folder called `http-request`.\n\n```\nmkdir http-request\n\ncd http-request\n\ntouch main.go\n```\n\nOpen the `main.go` and import the necessary packages.\n\n```\npackage main\n\nimport (\n\t\"bytes\"\n\t\"encoding/json\"\n\t\"fmt\"\n\t\"net/http\"\n)\n```\n\nCreate a struct that models the data received from the API.\n\n```\ntype Post struct {\n\tId     int    `json:\"id\"`\n\tTitle  string `json:\"title\"`\n\tBody   string `json:\"body\"`\n\tUserId int    `json:\"userId\"`\n}\n```\n\nCreate a POST request using the method `http.NewRequest`.\n\n```\nfunc main() {\n    // HTTP endpoint\n\tposturl := \"https://jsonplaceholder.typicode.com/posts\"\n\n    // JSON body\n\tbody := []byte(`{\n\t\t\"title\": \"Post title\",\n\t\t\"body\": \"Post description\",\n\t\t\"userId\": 1\n\t}`)\n\n    // Create a HTTP post request\n\tr, err := http.NewRequest(\"POST\", posturl, bytes.NewBuffer(body))\n\tif err != nil {\n\t\tpanic(err)\n\t}\n}\n```\n\nSet the HTTP request header.\n\n```\nr.Header.Add(\"Content-Type\", \"application/json\")\n```\n\nCreate a client and make the post request using the method `client.Do`\n\n```\nclient := &http.Client{}\nres, err := client.Do(r)\nif err != nil {\n\tpanic(err)\n}\n\ndefer res.Body.Close()\n```\n\nLet’s decode the JSON response using `json.NewDecoder` function that takes in the response body and a decode function that takes in a variable of type `Post`.\n\n```\npost := &Post{}\nderr := json.NewDecoder(res.Body).Decode(post)\nif derr != nil {\n\tpanic(derr)\n}\n```\n\nPanic if the HTTP status code not equals to `201`.\n\n```\nif res.StatusCode != http.StatusCreated {\n\tpanic(res.Status)\n}\n```\n\nFinally, print the newly created post on the console.\n\n```\nfmt.Println(\"Id:\", post.Id)\nfmt.Println(\"Title:\", post.Title)\nfmt.Println(\"Body:\", post.Body)\nfmt.Println(\"UserId:\", post.UserId)\n```\n\nHere is the complete working code.\n\n```\npackage main\n\nimport (\n\t\"bytes\"\n\t\"encoding/json\"\n\t\"fmt\"\n\t\"net/http\"\n)\n\ntype Post struct {\n\tId     int    `json:\"id\"`\n\tTitle  string `json:\"title\"`\n\tBody   string `json:\"body\"`\n\tUserId int    `json:\"userId\"`\n}\n\nfunc main() {\n\tposturl := \"https://jsonplaceholder.typicode.com/posts\"\n\n\tbody := []byte(`{\n\t\t\"title\": \"Post title\",''\n\t\t\"body\": \"Post description\",\n\t\t\"userId\": 1\n\t}`)\n\n\tr, err := http.NewRequest(\"POST\", posturl, bytes.NewBuffer(body))\n\tif err != nil {\n\t\tpanic(err)\n\t}\n\n\tr.Header.Add(\"Content-Type\", \"application/json\")\n\n\tclient := &http.Client{}\n\tres, err := client.Do(r)\n\tif err != nil {\n\t\tpanic(err)\n\t}\n\n\tdefer res.Body.Close()\n\n\tpost := &Post{}\n\tderr := json.NewDecoder(res.Body).Decode(post)\n\tif derr != nil {\n\t\tpanic(derr)\n\t}\n\n\tif res.StatusCode != http.StatusCreated {\n\t\tpanic(res.Status)\n\t}\n\n\tfmt.Println(\"Id:\", post.Id)\n\tfmt.Println(\"Title:\", post.Title)\n\tfmt.Println(\"Body:\", post.Body)\n\tfmt.Println(\"UserId:\", post.UserId)\n}\n```\n\n* * *"
}
```

## Usecase

This API can be use for many use case. One of them can be read further in the example-usecase [here](./example-usecase/README.md).