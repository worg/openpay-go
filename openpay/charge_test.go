package openpay

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestCharge(t *testing.T) {
	var chargeID = ``

	Convey(`Charge Client Test`, t, func() {
		client := NewClient(``, ``, false)

		Convey(`Add a store charge`, func() {
			dueDate := timeStamp(time.Now().Add(time.Hour).Format(time.RFC3339))
			charge, err := client.Charges.Create(Charge{
				Method:      `store`,
				Amount:      100,
				Description: `In store charge`,
				OrderID:     `CHRG-` + time.Now().String(),
				DueDate:     &dueDate,
			})

			chargeID = charge.ID
			So(err, ShouldBeNil)
			So(charge.Amount, ShouldEqual, 100)
			So(charge.PaymentMethod, ShouldNotBeNil)
		})

		Convey(`Retrieve a single charge`, func() {
			charge, err := client.Charges.Get(chargeID)

			So(err, ShouldBeNil)
			So(charge.Amount, ShouldEqual, 100)
		})
	})
}
