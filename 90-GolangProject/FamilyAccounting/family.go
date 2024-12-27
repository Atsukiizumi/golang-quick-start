package main

type Billing struct {
	AccType    string
	Accbalance float64
	RevenueNum float64
	Comment    string
}

type Family struct {
	Billing []*Billing
	balance float64
	flag    int
}

func NewFamily() *Family {
	family := &Family{
		Billing: make([]*Billing, 0, 10),
		balance: 10000.0,
		flag:    0,
	}

	return family
}

// 感觉有点多余
func (this *Family) UpdateBalance(money float64) float64 {
	this.balance = this.balance + money
	return this.balance
}

func (this *Family) AddBalance(acctype string, revenuenum float64, comment string) {
	switch acctype {
	case "收入":
		this.Billing = append(this.Billing, &Billing{AccType: acctype, Accbalance: this.UpdateBalance(revenuenum), RevenueNum: revenuenum, Comment: comment})
		break
	case "支出":
		revenuenum *= -1
		this.Billing = append(this.Billing, &Billing{AccType: acctype, Accbalance: this.UpdateBalance(revenuenum), RevenueNum: revenuenum, Comment: comment})
		break
	default:
		break
	}
}
