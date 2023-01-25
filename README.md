## GraphQL prototype

### Request
```
curl -g \
-X POST \
-H "Content-Type: application/json" \
-d '{"query":"query{RootQuery {hello}}"}' \
localhost:8080/graphql
```

### Response
```
{
	"data": {
		"hello": "world"
	}
}
```