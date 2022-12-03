package sf

import (
	"github.com/bwmarrin/snowflake"
	"time"
)

var node *snowflake.Node

//
//func Init(startTime string, machineID int64) (err error) {
//	var st time.Time
//	st, err = time.Parse("2006-01-02", startTime)
//	if err != nil {
//		return
//	}
//	snowflake.Epoch = st.UnixNano() / 1000000
//	node, err = snowflake.NewNode(machineID)
//
//	return
//}
//
//func GenId() int64 {
//	return node.Generate().Int64()
//}

func GenIdWithSnowFlake(startTime string, machineID int64) (err error, id int64) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err = snowflake.NewNode(machineID)
	if err != nil {
		return
	}
	id = node.Generate().Int64()

	return nil, id
}
