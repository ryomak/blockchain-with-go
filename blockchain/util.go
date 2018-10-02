package blockchain

func (bc *Blockchain)GetAmount(nodeIdent string)int{
	amount := 0
	for _,v := range bc.Chain{
		for _,t := range v.Transactions{
			if t.Sender == nodeIdent{
				amount -= t.Amount
			}
			if t.Recipient == nodeIdent{
				amount += t.Amount
			}
		}
	}
	return amount
}