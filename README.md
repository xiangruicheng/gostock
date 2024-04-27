# gostock

## Instruction
### 1 Create DB
1. Create MySQL Database,example "stock"
2. Create the table using sql statements in the doc directory
3. Modify config.yaml in the doc config

### 2 Init Data
* Init stock info
```
datainit.InitStockInfo()
```

* Init kline
```
datainit.BatchInitKline()
```