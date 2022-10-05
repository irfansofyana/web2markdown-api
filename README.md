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
    "url": "https://golangbyexample.com/json-response-body-http-go/"
}'
```

It will return a JSON API response:

```json
{
    "url": "https://golangbyexample.com/json-response-body-http-go/",
    "title": "Return JSON body in HTTP response in Go (Golang) - Welcome To Golang By Example",
    "data": "- [Overview](#Overview \"Overview\")\n- [Example](#Example \"Example\")\n\n## **Overvie** w\n\n**Write** method of the ResponseWriter interface in **net/http** package can be used to set the JSON body in an HTTP response\n\nIn GO a response is represented by the **ResponseWriter** Interface.  Here is the link to the interface –\n\n[https://golang.org/pkg/net/http/#ResponseWriter](https://golang.org/pkg/net/http/#ResponseWriter)\n\nResponseWriter interface is used by an HTTP handler to construct an HTTP response. It provides three functions to set the response parameters\n\n- Header – For writing response header\n\n- Write(\\[\\]byte) – For writing response body\n\n- WriteHeader(statusCode int) – For writing the http status code\n\n**Write** function can be used to set the response body. It takes a slice of bytes as input. Also, there is a **Header** function. This function can be used to set the content type of the response body using the Content-Type header. For eg in the case of the JSON response body, we need to set the Content-Type header as **“application/json”.**\n\n```\nw.Header().Set(\"Content-Type\", \"application/json\")\n```\n\nAlso, note that **WriteHeader** function can be used to set the HTTP status code for the response\n\n## **Example**\n\nLet’s see an example of sending http status code and JSON response body\n\nBelow is the program for the same\n\n```\npackage main\n\nimport (\n\t\"encoding/json\"\n\t\"log\"\n\t\"net/http\"\n)\n\nfunc main() {\n\thandler := http.HandlerFunc(handleRequest)\n\thttp.Handle(\"/example\", handler)\n\thttp.ListenAndServe(\":8080\", nil)\n}\n\nfunc handleRequest(w http.ResponseWriter, r *http.Request) {\n\tw.WriteHeader(http.StatusCreated)\n\tw.Header().Set(\"Content-Type\", \"application/json\")\n\tresp := make(map[string]string)\n\tresp[\"message\"] = \"Status Created\"\n\tjsonResp, err := json.Marshal(resp)\n\tif err != nil {\n\t\tlog.Fatalf(\"Error happened in JSON marshal. Err: %s\", err)\n\t}\n\tw.Write(jsonResp)\n\treturn\n}\n```\n\nIn the above program, this is how we create a JSON response. We use the **json.Marshal** function to convert the **map\\[string\\]string** into json bytes.\n\n```\nresp := make(map[string]string)\nresp[\"message\"] = \"Status Created\"\njsonResp, err := json.Marshal(resp)\nif err != nil {\n\tlog.Fatalf(\"Error happened in JSON marshal. Err: %s\", err)\n}\nw.Write(jsonResp)\n```\n\nIt then uses the **Write** function to return the JSON response body. The above code returns the below JSON response body back in response\n\n```\n{\"message\":\"Status Created\"}\n```\n\nAlso, we are using the **WriteHeader** function to specify the 201 http status code.\n\nRun the above program. It will start a server on 8080 port on your local machine. Now make the below curl call to the server\n\n```\ncurl -v -X POST http://localhost:8080/example\n```\n\nBelow will be the output\n\n```\n* Connected to localhost (::1) port 8080 (#0)\n> POST /example HTTP/1.1\n> Host: localhost:8080\n> User-Agent: curl/7.54.0\n> Accept: */*\n>\n< HTTP/1.1 201 Created\n< Date: Sat, 10 Jul 2021 10:40:33 GMT\n< Content-Length: 28\n< Content-Type: text/plain; charset=utf-8\n<\n* Connection #0 to host localhost left intact\n{\"message\":\"Status Created\"}\n```\n\nAs you can see from the output, it will correctly return the **201** status code along with the JSON body.\n\nAlso, check out our Golang advance tutorial Series - [Golang Advance Tutorial](https://golangbyexample.com/golang-comprehensive-tutorial/)\n\n[go](https://golangbyexample.com/tag/go/)\n[golang](https://golangbyexample.com/tag/golang/)"
}
```