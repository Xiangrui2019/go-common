package ecode

// 0作为OK正常
// 小于0为程序异常, 即500错误
// 大于0为业务提示、用户出错提示, 即400或其它
var (
	// 成功
	OK = New(0)
	// 找不到记录
	NotFound = New(1)
	// 服务器一般性异常
	ServerException = New(-2)
	// 服务器数据库操作性异常
	DatabaseException = New(-3)
)
