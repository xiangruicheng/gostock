package datainit

import "gostock/data/indicator"

func BatchUpdateMacd() {
	new(indicator.Macd).BatchRun()
}

func UpdateMacd(code string) {
	new(indicator.Macd).Run(code)
}
