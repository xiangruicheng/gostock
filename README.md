# gostock
This system  used to capture stock quotation data and analyze it

## Initialization
* 1 Create DB
Create Database and create Table
```
go run main.go make:db
```
* 2 Init Stock info
Init Table `stock_info` data, It is the basis for other functions
```
go run main.go make:stock
```
* 3 Init kline
  Init Table `kline` data, days from `config.yaml` item `init-num` setting
```
go run main.go make:stock
```

* 4 Init people
```
go run main.go make:people
```
* 4 Init Block
```
go run main.go make:block
```

