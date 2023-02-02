## GraphQL prototype

### Start Server
```
go run main.go
```

### Open GraphiQL
```
http://localhost:8080/graphql
```

### Example query
```
query {
  transactions(offset: 5, limit: 5) {
    merchant
    description
  }
}
```

### Expected Response
```
{
  "data": {
    "transactions": [
      {
        "description": "MacBook Pro 16",
        "merchant": "Apple"
      },
      {
        "description": "Nest Hub Max",
        "merchant": "Google"
      },
      {
        "description": "Galaxy Watch 3",
        "merchant": "Samsung"
      },
      {
        "description": "Xbox Series X",
        "merchant": "Microsoft"
      },
      {
        "description": "Echo Dot (4th Gen)",
        "merchant": "Amazon"
      }
    ]
  }
}
```