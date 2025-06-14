package module

type SummaryTransaction struct {
	Order int
	Name  string
	Price int
}

type TransactionManager struct {
	transaction        []Item
	summaryTransaction []SummaryTransaction
}

func NewTransactionManager() *TransactionManager {
	return &TransactionManager{
		transaction:        []Item{},
		summaryTransaction: []SummaryTransaction{},
	}
}

func (t *TransactionManager) Add(items []Item) {
	t.transaction = append(t.transaction, items...)
}

func (t *TransactionManager) CheckSummary(name string) (bool, int) {
	for i, transaction := range t.summaryTransaction {
		if transaction.Name == name {
			return true, i
		}
	}
	return false, 0
}

func (t *TransactionManager) GetAll() []Item {
	return t.transaction
}

func (t *TransactionManager) GetSummary() []SummaryTransaction {
	return t.summaryTransaction
}

func (t *TransactionManager) SortTransaction() {
	for _, item := range t.transaction {
		check, index := t.CheckSummary(item.Name)
		if len(t.summaryTransaction) == 0 {
			t.summaryTransaction = append(t.summaryTransaction, SummaryTransaction{
				Order: 1,
				Name:  item.Name,
				Price: item.Price,
			})
		} else if check {
			t.summaryTransaction[index].Order = t.summaryTransaction[index].Order + 1
		} else {
			t.summaryTransaction = append(t.summaryTransaction, SummaryTransaction{
				Order: 1,
				Name:  item.Name,
				Price: item.Price,
			})
		}
	}
}
