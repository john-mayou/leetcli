Given a `users.json` array, extract all usernames using `jq`.

### Example `users.json`

```
[
  { "id": 1, "username": "alice", "email": "alice@example.com" },
  { "id": 2, "username": "bob", "email": "bob@example.com" },
  { "id": 3, "username": "charlie", "email": "charlie@example.com" }
]
```

### Expected:

```
alice
bob
charlie
```
