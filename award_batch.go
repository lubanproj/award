package main

// AwardBatch defines the information for a prize
type AwardBatch struct {

	// award name
	name string

	// total remaining prizes
	totalBalance int64

	// total number of prizes
	totalAmount int64

	// time of last prize release
	updateTime int64

}

func (award *AwardBatch) GetName() string {
	return award.name
}

func (award *AwardBatch) GetTotalBalance() int64 {
	return award.totalBalance
}

func (award *AwardBatch) GetTotalAmount() int64 {
	return award.totalAmount
}

func (award *AwardBatch) GetUpdateTime() int64 {
	return award.updateTime
}

func (award *AwardBatch) SetTotalBalance(totalBalance int64)  {
	award.totalBalance = totalBalance
}

func (award *AwardBatch) SetTotalAmount(totalAmount int64)  {
	award.totalAmount = totalAmount
}

func (award *AwardBatch) SetUpdateTime(updateTime int64)  {
	award.updateTime = updateTime
}

func (award *AwardBatch) SetName(name string)  {
	award.name = name
}
