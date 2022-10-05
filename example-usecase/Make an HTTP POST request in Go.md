January 23, 2022

We will make a POST request to an endpoint with a JSON body and display the results on the console.

The endpoint will accept `id`, `title`, `body`, `userId` and create a new `post`.

Create a new folder called `http-request`.

```
mkdir http-request

cd http-request

touch main.go
```

Open the `main.go` and import the necessary packages.

```
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)
```

Create a struct that models the data received from the API.

```
type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}
```

Create a POST request using the method `http.NewRequest`.

```
func main() {
    // HTTP endpoint
	posturl := "https://jsonplaceholder.typicode.com/posts"

    // JSON body
	body := []byte(`{
		"title": "Post title",
		"body": "Post description",
		"userId": 1
	}`)

    // Create a HTTP post request
	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
}
```

Set the HTTP request header.

```
r.Header.Add("Content-Type", "application/json")
```

Create a client and make the post request using the method `client.Do`

```
client := &http.Client{}
res, err := client.Do(r)
if err != nil {
	panic(err)
}

defer res.Body.Close()
```

Let’s decode the JSON response using `json.NewDecoder` function that takes in the response body and a decode function that takes in a variable of type `Post`.

```
post := &Post{}
derr := json.NewDecoder(res.Body).Decode(post)
if derr != nil {
	panic(derr)
}
```

Panic if the HTTP status code not equals to `201`.

```
if res.StatusCode != http.StatusCreated {
	panic(res.Status)
}
```

Finally, print the newly created post on the console.

```
fmt.Println("Id:", post.Id)
fmt.Println("Title:", post.Title)
fmt.Println("Body:", post.Body)
fmt.Println("UserId:", post.UserId)
```

Here is the complete working code.

```
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId int    `json:"userId"`
}

func main() {
	posturl := "https://jsonplaceholder.typicode.com/posts"

	body := []byte(`{
		"title": "Post title",''
		"body": "Post description",
		"userId": 1
	}`)

	r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}

	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	post := &Post{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		panic(derr)
	}

	if res.StatusCode != http.StatusCreated {
		panic(res.Status)
	}

	fmt.Println("Id:", post.Id)
	fmt.Println("Title:", post.Title)
	fmt.Println("Body:", post.Body)
	fmt.Println("UserId:", post.UserId)
}
```

* * *