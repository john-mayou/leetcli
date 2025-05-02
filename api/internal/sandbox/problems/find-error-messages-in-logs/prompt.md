Given a `logs.txt` file, extract all the lines that contain the word "error" (case-insensitive) using `grep`.

### Example `logs.txt`:

```
INFO: Server started
ERROR: Failed to connect to DB
INFO: Error occurred, retrying API request
WARNING: High memory usage
```

### Expected:

```
ERROR: Failed to connect to DB
INFO: Error occured, retrying API request
```
