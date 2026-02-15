# Fiber Boilerplate

## Migrate

> [!NOTE]
> Make sure you have `golang-migrate` installed.\
> See 👉 [How to install golang-migrate](https://github.com/golang-migrate/migrate)

### Create Migration

```bash
migrate create -ext sql -dir {path migrations} -seq {tabel name}_table
```

### Run Migration

Database URLs

Database connection strings are specified via URLs. The URL format is driver dependent but generally has the form: `dbdriver://username:password@host:port/dbname?param1=true&param2=false`

#### Up Migration
```bash
migrate -database {POSTGRESQL URL} -path {path migrations} up
```

#### Down Migration
```bash
migrate -database {POSTGRESQL URL} -path {path migrations} down
```

## Run The App

### Install Require App
```bash
go get
```

### Running Development With Air

> [!NOTE]
> Make sure you have `Air` installed.\
> See 👉 [How to install Air](https://github.com/air-verse/air)

#### Initialitation Air
```bash
air init
```

#### Run
```bash
air
```

### Running Production

#### Build The App
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
go build -ldflags="-w -s" -o ./bin/main ./main.go
```

#### Run The App
```bash
./bin/main
```