# RADIANT

### Description

Echo based web framework wrapper

### Prerequisites
   - Golang 1.17+ (tested)
   - A good editor like **Goland/VS Code**
   - Git
   - Below commands
```bash
    go get
    # if you like hot reloading
    go get -u github.com/cosmtrek/air
```

### Important before your run
  - Copy `config/conf.yaml.sample` to create `conf.yaml`
  - Never commit `config/conf.yaml`

### Run
```bash
# for hot reload
air run
# in traditional way
go run server.go
```

### Route

    static routing: http://localhost:8080/about
    dynamic routing: http://localhost:8080/bangaldesh/dhaka/bashundhara

### Note

    Default language: Engilsh(en-US)
    Set your browser language preference French(fr-FR) to get in French version.
