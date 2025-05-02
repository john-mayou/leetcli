Given a `api_logs.jsonl` file (JSON Lines format), extract a sorted list of unique user IDs where the `"status"` is `"error"`.
Use a combination of `jq`, `sort`, `uniq`, and optionally `awk` or `grep`.

### Exmaple `api_logs.jsonl`

```
{ "user_id": 42, "endpoint": "/login", "status": "success" }
{ "user_id": 17, "endpoint": "/purchase", "status": "error" }
{ "user_id": 42, "endpoint": "/update", "status": "error" }
{ "user_id": 17, "endpoint": "/login", "status": "error" }
{ "user_id": 99, "endpoint": "/logout", "status": "success" }
```

### Expected:

```
17
42
```
