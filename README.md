# ASCII Server

### What you will need

* Go 1.18
* Docker

### How to run?
```Bash
make run
```

### How to test?
```bash
make test
```

### How to query?
```
curl localhost:8080/canvas/62a22d7b8a95a01c6aedfb0f
```

If you want to add more canvases you can add them in the "installCanvases" function
in the installers.go file.
