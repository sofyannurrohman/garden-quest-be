package transaction

import "garden-quest/user"

type CreateTransactionInput struct {
	Amount    int `json:"amount" binding:"required"`
	PackageID int `json:"package_id" binding:"required"`
	User      user.User
}
