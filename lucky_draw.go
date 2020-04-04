package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/gomodule/redigo/redis"
)

// GetAllAwardBatch get all the AwardBatchs
func GetAllAwardBatch() []AwardBatch{

	var awardBatches []AwardBatch

	// get redis conn
	conn, err := GetConn()

	if err != nil {
		fmt.Printf("get conn error, %v \n", err)
		return awardBatches
	}

	defer conn.Close()

	awardInfoKey := getAwardInfoKey()
	values , err := redis.Values(conn.Do("ZRANGE",awardInfoKey,0,-1,"WITHSCORES"))

	if err != nil || len(values) == 0 {
		fmt.Println("get all award redis error ", err)
	}

	for index, value := range values {

		if index % 2 == 0 {

			awardName , ok := value.([]byte)

			if !ok {
				fmt.Println("value type error : ", value)
				continue
			}
			awardBatches = append(awardBatches, AwardBatch{
				name:string(awardName),
			})

		} else {

			lastUpdateTimeStr , ok := value.([]byte)
			if !ok {
				fmt.Println("time type error : ", lastUpdateTimeStr)
				continue
			}

			lastUpdateTime ,err := strconv.ParseInt(string(lastUpdateTimeStr), 10, 64)

			if err != nil {
				fmt.Println("time type error", err)
				continue
			}

			awardBatches[index/2].SetUpdateTime(lastUpdateTime)
		}
	}

	// fill totalAmount
	for index, awardBatch := range awardBatches {
		awardBatches[index].SetTotalAmount(Conf.AwardMap[awardBatch.GetName()])
	}

	return awardBatches
}

// Determine if the user has won prize
func GetAward(username string) *AwardBatch {

	awardBatch := GetAwardBatch()

	if awardBatch == nil {
		return awardBatch
	}

	// update lastUpdateTime and balance
	conn, err := GetConn()

	if err != nil {
		fmt.Printf("get conn is nil , %v \n", err)
		return nil
	}

	defer conn.Close()

	conn.Send("WATCH", getAwardBalanceKey())
	conn.Send("MULTI")
	conn.Send("ZADD", getAwardInfoKey(),time.Now().Unix(), awardBatch.GetName())
	conn.Send("HSET", getAwardBalanceKey(), awardBatch.GetName(), awardBatch.totalBalance - 1)
	conn.Send("EXEC")

	err = conn.Flush()
	if err != nil {
		fmt.Println("redis error, " , err)
		return nil
	}

	fmt.Println("congratulations , you won ", awardBatch.GetName() )

	awardTime := time.Unix(awardBatch.GetUpdateTime(), 0).Format("2006-01-02 15:04:05")
//	userName := req.Form.Get("user_name")
	if err := SaveRecords(awardBatch.GetName() , awardTime, username); err != nil {
		fmt.Printf("save records error, %v", err)
	}

	return awardBatch

}

// GetAwardBatch implemented a specific lucky draw algorithm
func GetAwardBatch() *AwardBatch {

	awardBatch, err := RandomGetAwardBatch()

	if awardBatch == nil || err != nil {
		fmt.Println("sorry, you didn't win the prize.")
		return nil
	}

	// Determine if the prize release time point has been reached
	startTime , _ := ParseStringToTime(Conf.Award.StartTime)
	endTime , _ := ParseStringToTime(Conf.Award.EndTime)
	totalAmount := awardBatch.totalAmount
	totalBalance := awardBatch.totalBalance
	lastUpdateTime := awardBatch.GetUpdateTime()
	random := rand.New(rand.NewSource(lastUpdateTime))

	detaTime := (endTime - startTime) / awardBatch.totalAmount

	// calculate when the next award will be released
	releaseTime := startTime + (totalAmount - totalBalance) * detaTime + int64(random.Int()) % detaTime

	fmt.Println("relaseTime : ", time.Unix(releaseTime, 0).Format("2006-01-02 15:04:05"))

	if time.Now().Unix() < releaseTime {
		// If you do not reach the point of release, you will not win
		fmt.Println("sorry, you didn't win the prize")
		return nil
	}

	return awardBatch
}

// RandomGetAwardBatch choose a random award from the award pool
func RandomGetAwardBatch() (*AwardBatch, error) {

	conn, err := GetConn()
	if err != nil {
		fmt.Printf("get conn is nil , %v \n", err)
		return nil, err
	}

	defer conn.Close()

	retMap, err := redis.Int64Map(conn.Do("HGETALL", getAwardBalanceKey()))

	if err != nil || retMap == nil {
		fmt.Println("redis HGETALL award error", err)
		return nil, err
	}

	totalBalance := int64(0)
	for _, value := range retMap {
		totalBalance += value
	}

	fmt.Println("retMap : ", retMap)

	if totalBalance == 0 {
		fmt.Println("total balance is 0")
		return nil, errors.New("total balance is 0")
	}

	fmt.Println("totalBalance :", totalBalance)

	awardBatches := GetAllAwardBatch()

	for index , awardBatch := range awardBatches {
		awardBatches[index].totalBalance = retMap[awardBatch.GetName()]
	}

	fmt.Println("awardBatches :", awardBatches)

	random := rand.New(rand.NewSource(totalBalance))

	num := random.Int63n(totalBalance)

	for _ , awardBatch := range awardBatches {

		// The awards have been drawn
		if awardBatch.GetTotalBalance() <= 0 {
			fmt.Println("The prizes have been drawn ...")
			continue
		}

		num = num - awardBatch.GetTotalBalance()
		if num < 0 {
			return &awardBatch, nil
		}
	}

	return nil, nil
}


// InitAwardPool initializes the award pool
func InitAwardPool() error {

	conn, err := GetConn()
	if err != nil {
		fmt.Printf("get conn is nil , %v \n", err)
		return err
	}

	defer conn.Close()

	conn.Send("ZADD", getAwardInfoKey(), time.Now().Unix(), "A")
	conn.Send("ZADD", getAwardInfoKey(), time.Now().Unix(), "B")
	conn.Send("ZADD", getAwardInfoKey(), time.Now().Unix(), "C")

	conn.Send("HSET", getAwardBalanceKey(), "A", Conf.AwardMap["A"])
	conn.Send("HSET", getAwardBalanceKey(), "B", Conf.AwardMap["B"])
	conn.Send("HSET", getAwardBalanceKey(), "C", Conf.AwardMap["C"])
	conn.Flush()

	for i := 0 ; i < 6; i++ {
		_ , err := conn.Receive()

		if err != nil {
			fmt.Println("conn send error", err)
		}
	}

	return nil
}





