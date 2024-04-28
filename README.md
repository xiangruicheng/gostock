# gostock

## Instruction
### 1 Create DB
ddl.Create()

### 2 Init Data
* Init stock info 
```
datainit.InitStockInfo()
```
* Init all kline 
```
datainit.BatchInitKline()
```
* Init kline for an code
```
ainit.InitKline(model.StockInfoModel_TypeEtf, "159915", "SZ", config.Data.Xueqiu.InitNum)
```
