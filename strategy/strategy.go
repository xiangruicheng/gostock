package strategy

type Strategy interface {
	Run()
}

// hs300:  filter: -4 to -6  buy:close sell:2 day close
// cyb: filter:>-2 buy:close sell:2day open
