package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptionMd5 字符串加密
func EncryptionMd5(s string) string {
	// 创建md5哈希计算器
	ctx := md5.New()
	// 向计算器写入目标数据
	ctx.Write([]byte(s))
	// 输出十六进制字符串，Sum的nil表示不添加前缀
	return hex.EncodeToString(ctx.Sum(nil))
}
