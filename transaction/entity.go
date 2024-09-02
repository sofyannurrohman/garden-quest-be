package transaction

import (
	"garden-quest/user"
	"time"

	"github.com/leekchan/accounting"
)

type Transaction struct {
	ID        int
	PackageID int
	UserID    int
	Amount    int
	Status    string
	User      user.User
	// PaymentURL string
	Package   Package
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Package struct {
	ID            int
	Name          string
	Type          string
	PlantTypeID   int
	WaterEnergyID int
	Price         int
}

func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
