package module

var Transactions []Item = []Item{}

type SummaryTransaction struct {
	Order int
	Name  string
	Price int
}

var SummaryTransactions []SummaryTransaction = []SummaryTransaction{}

func checkItem(transaction []SummaryTransaction, Name string) (bool, int) {
	for i := range transaction {
		if transaction[i].Name == Name {
			return true, i
		}
	}
	return false, 0
}

func SortTransactions() {
	for _, item := range Transactions {
		check, index := checkItem(SummaryTransactions, item.Name)
		if len(SummaryTransactions) == 0 {
			SummaryTransactions = append(SummaryTransactions, SummaryTransaction{
				Order: 1,
				Name:  item.Name,
				Price: item.Price,
			})
		} else if check {
			SummaryTransactions[index].Order = SummaryTransactions[index].Order + 1
		} else {
			SummaryTransactions = append(SummaryTransactions, SummaryTransaction{
				Order: 1,
				Name:  item.Name,
				Price: item.Price,
			})
		}
	}
}

func AddTransactions() {
	Transactions = append(Transactions, Cart...)
}
