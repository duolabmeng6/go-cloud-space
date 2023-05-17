package ecore

import (
	"github.com/joho/godotenv"
	"os"
)

// env
func E加载环境变量_从文件(filepath string) bool {
	err := godotenv.Load(filepath)
	if err != nil {
		E调试输出("加载环境变量文件失败文件路径为:"+filepath, "错误信息为", err)
		return false
	}
	return true
}

func E加载环境变量_从内容到map(s string) (map[string]string, error) {
	myEnv, err := godotenv.Unmarshal(s)
	return myEnv, err
}

// 带有默认值的env读取函数
// E读环境变量 一样效果
func Env(环境变量名称 string, 默认值 string) string {
	v := E读环境变量(环境变量名称, 默认值)
	return v
}
func E读环境变量(环境变量名称 string, 默认值 string) string {
	var def string
	if len(默认值) > 1 {
		def = 默认值
	}
	e := os.Getenv(环境变量名称)
	if e == "" {
		return def
	} else {
		return e
	}
}
func E写环境变量(环境变量名称 string, 欲写入内容 string) bool {
	err := os.Setenv(环境变量名称, 欲写入内容)
	return err == nil
}
