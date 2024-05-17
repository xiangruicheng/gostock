package indicator

func MacdBatchRun() {
	new(Macd).BatchRun()
}

func MacdRun(code string) {
	new(Macd).Run(code)
}
