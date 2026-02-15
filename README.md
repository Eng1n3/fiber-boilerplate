# Fiber Boilerplate

## Migrate

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

