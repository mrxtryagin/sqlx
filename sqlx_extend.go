package sqlx

import "time"

type SqlExecCallBack func(startAt time.Time, query string, args ...interface{})

// defaultSqlExecCallBack
// todo： 没有logger的问题
func defaultSqlExecCallBack(startAt time.Time, query string, args ...interface{}) {

}
