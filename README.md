## GraphQL prototype

### Request
```
curl -g \
-X POST \
-H "Content-Type: application/json" \
-d '{"query":"query{merchant,description}"}' \
localhost:8080/graphql
```

### Response
```
{
	"data": {
		"description": "iPhone 10",
		"merchant": "Apple"
	}
}
```