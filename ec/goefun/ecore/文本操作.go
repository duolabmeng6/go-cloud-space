package ecore

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/grand"
)

// 调用格式： 〈整数型〉 取文本长度 （文本型 文本数据） - 系统核心支持库->文本操作
// 英文名称：len
// 取文本型数据的长度，不包含结束0。本命令为初级命令。
// 参数<1>的名称为“文本数据”，类型为“文本型（text）”。参数值指定欲检查其长度的文本数据。
//
// 操作系统需求： Windows、Linux
func E取文本长度(value string) int64 {
	return E到整数(len([]rune(value)))
}

// 调用格式： 〈文本型〉 取文本左边 （文本型 欲取其部分的文本，整数型 欲取出字符的数目） - 系统核心支持库->文本操作
// 英文名称：left
// 返回一个文本，其中包含指定文本中从左边算起指定数量的字符。本命令为初级命令。
// 参数<1>的名称为“欲取其部分的文本”，类型为“文本型（text）”。
// 参数<2>的名称为“欲取出字符的数目”，类型为“整数型（int）”。
//
// 操作系统需求： Windows、Linux
func E取文本左边(欲取其部分的文本 string, 欲取出字符的数目 int64) string {
	if E取文本长度(欲取其部分的文本) < 欲取出字符的数目 {
		欲取出字符的数目 = E取文本长度(欲取其部分的文本)
	}
	return string([]rune(欲取其部分的文本)[:欲取出字符的数目])
}

//调用格式： 〈文本型〉 取文本右边 （文本型 欲取其部分的文本，整数型 欲取出字符的数目） - 系统核心支持库->文本操作
//英文名称：right
//返回一个文本，其中包含指定文本中从右边算起指定数量的字符。本命令为初级命令。
//参数<1>的名称为“欲取其部分的文本”，类型为“文本型（text）”。
//参数<2>的名称为“欲取出字符的数目”，类型为“整数型（int）”。
//
//操作系统需求： Windows、Linux

func E取文本右边(欲取其部分的文本 string, 欲取出字符的数目 int64) string {
	l := E取文本长度(欲取其部分的文本)
	lpos := l - 欲取出字符的数目
	if lpos < 0 {
		lpos = 0
	}
	return string([]rune(欲取其部分的文本)[lpos:l])
}

//取文本中间

// 调用格式： 〈文本型〉 字符 （字节型 欲取其字符的字符代码） - 系统核心支持库->文本操作
// 英文名称：chr
// 返回一个文本，其中包含有与指定字符代码相关的字符。本命令为初级命令。
// 参数<1>的名称为“欲取其字符的字符代码”，类型为“字节型（byte）”。
//
// 操作系统需求： Windows、Linux
func E字符(字节型 int8) string {
	return string(byte(字节型))
}

// 调用格式： 〈整数型〉 取代码 （文本型 欲取字符代码的文本，［整数型 欲取其代码的字符位置］） - 系统核心支持库->文本操作
// 英文名称：asc
// 返回文本中指定位置处字符的代码。如果指定位置超出文本长度，返回0。本命令为初级命令。
// 参数<1>的名称为“欲取字符代码的文本”，类型为“文本型（text）”。
// 参数<2>的名称为“欲取其代码的字符位置”，类型为“整数型（int）”，可以被省略。1为首位置，2为第2个位置，如此类推。如果本参数被省略，默认为首位置。
//
// 操作系统需求： Windows、Linux
func E取代码(欲取字符代码的文本 string) int {
	for _, char := range []rune(欲取字符代码的文本) {
		return int(char)
	}
	return 0
}

// 调用格式： 〈整数型〉 寻找文本 （文本型 被搜寻的文本，文本型 欲寻找的文本，［整数型 起始搜寻位置］，逻辑型 是否不区分大小写） - 系统核心支持库->文本操作
// 英文名称：InStr
// 返回一个整数值，指定一文本在另一文本中最先出现的位置，位置值从 1 开始。如果未找到，返回-1。本命令为初级命令。
// 参数<1>的名称为“被搜寻的文本”，类型为“文本型（text）”。
// 参数<2>的名称为“欲寻找的文本”，类型为“文本型（text）”。
// 参数<3>的名称为“起始搜寻位置”，类型为“整数型（int）”，可以被省略。位置值从 1 开始。如果本参数被省略，默认为 1 。
// 参数<4>的名称为“是否不区分大小写”，类型为“逻辑型（bool）”，初始值为“假”。为真不区分大小写，为假区分。
//
// 操作系统需求： Windows、Linux
func E寻找文本(被搜寻的文本 string, 欲寻找的文本 string) int {
	return strings.Index(被搜寻的文本, 欲寻找的文本)
}

// 调用格式： 〈整数型〉 倒找文本 （文本型 被搜寻的文本，文本型 欲寻找的文本，［整数型 起始搜寻位置］，逻辑型 是否不区分大小写） - 系统核心支持库->文本操作
// 英文名称：InStrRev
// 返回一个整数值，指定一文本在另一文本中最后出现的位置，位置值从 1 开始。如果未找到，返回-1。本命令为初级命令。
// 参数<1>的名称为“被搜寻的文本”，类型为“文本型（text）”。
// 参数<2>的名称为“欲寻找的文本”，类型为“文本型（text）”。
// 参数<3>的名称为“起始搜寻位置”，类型为“整数型（int）”，可以被省略。位置值从 1 开始。如果本参数被省略，默认为从被搜寻文本的尾部开始。
// 参数<4>的名称为“是否不区分大小写”，类型为“逻辑型（bool）”，初始值为“假”。为真不区分大小写，为假区分。
//
// 操作系统需求： Windows、Linux
func E倒找文本(被搜寻的文本 string, 欲寻找的文本 string) int {
	return strings.LastIndex(被搜寻的文本, 欲寻找的文本)
}

/*
调用格式： 〈文本型〉 到大写 （文本型 欲变换的文本） - 系统核心支持库->文本操作
英文名称：UCase
将文本中的小写英文字母变换为大写，返回变换后的结果文本。本命令为初级命令。
参数<1>的名称为“欲变换的文本”，类型为“文本型（text）”。

操作系统需求： Windows、Linux
*/
func E到大写(value string) string {
	return strings.ToUpper(value)
}
func E到小写(value string) string {
	return strings.ToLower(value)
}
func E到全角(value string) string {
	return sBCtoDBC(value)
}
func E到半角(value string) string {
	return dBCtoSBCNew(value)
}

// 调用格式： 〈文本型〉 删首空 （文本型 欲删除空格的文本） - 系统核心支持库->文本操作
// 英文名称：LTrim
// 返回一个文本，其中包含被删除了首部全角或半角空格的指定文本。本命令为初级命令。
// 参数<1>的名称为“欲删除空格的文本”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func E删首空(欲删除空格的文本 string) string {
	return strings.TrimLeft(欲删除空格的文本, " ")
}

// 调用格式： 〈文本型〉 删尾空 （文本型 欲删除空格的文本） - 系统核心支持库->文本操作
// 英文名称：RTrim
// 返回一个文本，其中包含被删除了尾部全角或半角空格的指定文本。本命令为初级命令。
// 参数<1>的名称为“欲删除空格的文本”，类型为“文本型（text）”。
//
// 操作系统需求： Windows、Linux
func E删尾空(欲删除空格的文本 string) string {
	return strings.TrimRight(欲删除空格的文本, " ")
}

func E删首尾空(内容 string) string {
	return strings.TrimSpace(内容)
}

// 删全部空
func E删全部空(内容 string) string {
	return strings.Join(strings.FieldsFunc(内容, unicode.IsSpace), "")
}

//文本替换
//调用格式： 〈文本型〉 文本替换 （文本型 欲被替换的文本，整数型 起始替换位置，整数型 替换长度，［文本型 用作替换的文本］） - 系统核心支持库->文本操作
//英文名称：ReplaceText
//将指定文本的某一部分用其它的文本替换。本命令为初级命令。
//参数<1>的名称为“欲被替换的文本”，类型为“文本型（text）”。
//参数<2>的名称为“起始替换位置”，类型为“整数型（int）”。替换的起始位置，1为首位置，2为第2个位置，如此类推。
//参数<3>的名称为“替换长度”，类型为“整数型（int）”。
//参数<4>的名称为“用作替换的文本”，类型为“文本型（text）”，可以被省略。如果本参数被省略，则删除文本中的指定部分。
//
//操作系统需求： Windows、Linux

// 调用格式： 〈文本型〉 子文本替换 （文本型 欲被替换的文本，文本型 欲被替换的子文本，［文本型 用作替换的子文本］，［整数型 进行替换的起始位置］，［整数型 替换进行的次数］，逻辑型 是否区分大小写） - 系统核心支持库->文本操作
// 英文名称：RpSubText
// 返回一个文本，该文本中指定的子文本已被替换成另一子文本，并且替换发生的次数也是被指定的。本命令为初级命令。
// 参数<1>的名称为“欲被替换的文本”，类型为“文本型（text）”。
// 参数<2>的名称为“欲被替换的子文本”，类型为“文本型（text）”。
// 参数<3>的名称为“用作替换的子文本”，类型为“文本型（text）”，可以被省略。如果本参数被省略，默认为空文本。
// 参数<4>的名称为“进行替换的起始位置”，类型为“整数型（int）”，可以被省略。参数值指定被替换子文本的起始搜索位置。如果省略，默认从 1 开始。
// 参数<5>的名称为“替换进行的次数”，类型为“整数型（int）”，可以被省略。参数值指定对子文本进行替换的次数。如果省略，默认进行所有可能的替换。
// 参数<6>的名称为“是否区分大小写”，类型为“逻辑型（bool）”，初始值为“真”。为真区分大小写，为假不区分。
//
// 操作系统需求： Windows、Linux
func E子文本替换(欲被替换的文本 string, 欲被替换的子文本 string, 用作替换的子文本 string) string {
	return strings.Replace(欲被替换的文本, 欲被替换的子文本, 用作替换的子文本, -1)
}

// 调用格式： 〈文本型〉 取空白文本 （整数型 重复次数） - 系统核心支持库->文本操作
// 英文名称：space
// 返回具有指定数目半角空格的文本。本命令为初级命令。
// 参数<1>的名称为“重复次数”，类型为“整数型（int）”。
//
// 操作系统需求： Windows、Linux
func E取空白文本(重复次数 int) string {
	var str string
	for i := 0; i < 重复次数; i++ {
		str = str + " "
	}
	return str
}

// 调用格式： 〈文本型〉 取重复文本 （整数型 重复次数，文本型 待重复文本） - 系统核心支持库->文本操作
// 英文名称：string
// 返回一个文本，其中包含指定次数的文本重复结果。本命令为初级命令。
// 参数<1>的名称为“重复次数”，类型为“整数型（int）”。
// 参数<2>的名称为“待重复文本”，类型为“文本型（text）”。该文本将用于建立返回的文本。如果为空，将返回一个空文本。
//
// 操作系统需求： Windows、Linux
func E取重复文本(重复次数 int, 待重复文本 string) string {
	var str string
	for i := 0; i < 重复次数; i++ {
		str = str + 待重复文本
	}
	return str
}

// 调用格式： 〈文本型数组〉 分割文本 （文本型 待分割文本，［文本型 用作分割的文本］，［整数型 要返回的子文本数目］） - 系统核心支持库->文本操作
// 英文名称：split
// 将指定文本进行分割，返回分割后的一维文本数组。本命令为初级命令。
// 参数<1>的名称为“待分割文本”，类型为“文本型（text）”。如果参数值是一个长度为零的文本，则返回一个空数组，即没有任何成员的数组。
// 参数<2>的名称为“用作分割的文本”，类型为“文本型（text）”，可以被省略。参数值用于标识子文本边界。如果被省略，则默认使用半角逗号字符作为分隔符。如果是一个长度为零的文本，则返回的数组仅包含一个成员，即完整的“待分割文本”。
// 参数<3>的名称为“要返回的子文本数目”，类型为“整数型（int）”，可以被省略。如果被省略，则默认返回所有的子文本。
//
// 操作系统需求： Windows、Linux
func E分割文本(待分割文本 string, 用作分割的文本 string) []string {
	return strings.Split(待分割文本, 用作分割的文本)
}

func E格式化文本(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a...)
}

// 文本截取函数
func StrCut(内容 string, 表达式 string) string {
	args := strings.Split(表达式, "$")
	if len(args) == 2 {
		return E文本_取出中间文本(内容, args[0], args[1])
	}
	if len(args) == 1 {
		return E文本_取出中间文本(内容, args[0], "")
	}
	return ""
}

func E文本_取左边(被查找的文本 string, 欲寻找的文本 string) string {
	return E文本_取出中间文本(被查找的文本, "", 欲寻找的文本)
}
func E文本_取右边(被查找的文本 string, 欲寻找的文本 string) string {
	return E文本_取出中间文本(被查找的文本, 欲寻找的文本, "")
}

// 文本取出中间文本
func E文本_取出中间文本(内容 string, 左边文本 string, 右边文本 string) string {
	左边位置 := strings.Index(内容, 左边文本)
	if 左边位置 == -1 {
		return ""
	}
	左边位置 = 左边位置 + len(左边文本)
	内容 = string([]byte(内容)[左边位置:])

	var 右边位置 int
	if 右边文本 == "" {
		右边位置 = len(内容)
	} else {
		右边位置 = strings.Index(内容, 右边文本)
		if 右边位置 == -1 {
			return ""
		}
	}
	内容 = string([]byte(内容)[:右边位置])
	return 内容
}

// 子程序名：文本_取随机字母
// 取随机26个字母！
// 返回值类型：文本型
// 参数<1>的名称为“要取的字符数”，类型为“整数型”。注明：要取字符个数；。
// 参数<2>的名称为“字母类型”，类型为“整数型”，允许接收空参数数据。注明：可空：默认为小写字母；0=小写字母；1=大写字母；2=大小写字母混合；。
func E文本_取随机字母(要取的字符数 int, 字母类型 int) string {
	var str string
	if 字母类型 == 0 {
		str = "abcdefghijklmnopqrstuvwxyz"
	}
	if 字母类型 == 1 {
		str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if 字母类型 == 2 {
		str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	return grand.Str(str, 要取的字符数)
}

func E文本_取随机字母和数字(要取的字符数 int) string {
	return grand.Str("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", 要取的字符数)
}

func E文本_取随机数字(要取的字符数 int) string {
	return grand.Str("0123456789", 要取的字符数)
}

func E文本_删左边(欲处理文本 string, 删除长度 int64) string {
	return E取文本右边(欲处理文本, E取文本长度(欲处理文本)-删除长度)
}
func E文本_删右边(欲处理文本 string, 删除长度 int64) string {
	return E取文本左边(欲处理文本, E取文本长度(欲处理文本)-删除长度)
}

func E文本_删中间(欲处理文本 string, 起始位置 int64, 删除长度 int64) string {
	return E取文本左边(欲处理文本, 起始位置) + E文本_删左边(欲处理文本, 起始位置+删除长度)
}

func E文本_取出文本中汉字(s string) string {
	return E文本区分_只取汉子(s)
}

func E文本_逐字分割(s string) []string {
	r := []rune(s)
	strarr := []string{}
	for _, s := range r {
		strarr = append(strarr, string(s))
	}
	return strarr
}

func E文本_颠倒(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func E文本_取随机姓氏() string {
	百家姓 := "赵钱孙李周吴郑王冯陈褚卫蒋沈韩杨朱秦尤许何吕施张孔曹严华金魏陶姜戚谢邹喻柏水窦章云苏潘葛奚范彭郎鲁韦昌马苗凤花方俞任袁柳酆鲍史唐费廉岑薛雷贺倪汤滕殷罗毕郝邬安常乐于时傅皮卞齐康伍余元卜顾孟平黄和穆萧尹姚邵湛汪祁毛禹狄米贝明臧计伏成戴谈宋茅庞熊纪舒屈项祝董梁杜阮蓝闵席季麻强贾路娄危江童颜郭梅盛林刁钟徐邱骆高夏蔡田樊胡凌霍虞万支柯昝管卢莫柯房裘缪干解应宗丁宣贲邓郁单杭洪包诸左石崔吉钮龚程嵇邢滑裴陆荣翁荀羊于惠甄曲家封芮羿储靳汲邴糜松井段富巫乌焦巴弓牧隗山谷车侯宓蓬全郗班仰秋仲伊宫宁仇栾暴甘钭历戎祖武符刘景詹束龙叶幸司韶郜黎蓟溥印宿白怀蒲邰从鄂索咸籍赖卓蔺屠蒙池乔阳郁胥能苍双闻莘党翟谭贡劳逄姬申扶堵冉宰郦雍却璩桑桂濮牛寿通边扈燕冀浦尚农温别庄晏柴瞿阎充慕连茹习宦艾鱼容向古易慎戈廖庾终暨居衡步都耿满弘匡国文寇广禄阙东欧殳沃利蔚越夔隆师巩厍聂晁勾敖融冷訾辛阚那简饶空曾毋沙乜养鞠须丰巢关蒯相查后荆红游竺权逮盍益桓公"
	return grand.Str(百家姓, 1)
}
func E文本_自动补零(s string, len int) string {
	return E格式化文本("%0*d", len, E到整数(s))
}

//unicode的参数含义
//https://www.cnblogs.com/golove/p/3269099.html
//Golang学习 - unicode 包
//https://www.cnblogs.com/golove/p/3273585.html

func E文本_是否为小写字母(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}

func E文本_是否为大写字母(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}
func E文本_是否为字母(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) || unicode.IsUpper(r) {
			return true
		}
	}
	return false
}
func E文本_是否为数字(s string) bool {
	for _, r := range s {
		if unicode.IsNumber(r) {
			return true
		}
	}
	return false
}

func E文本_是否为汉字(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func E文本区分_只取字母(s string) string {

	str := ""
	for _, r := range s {
		if unicode.IsLower(r) || unicode.IsUpper(r) {
			str = str + string(r)
		}
	}
	return str
}

func E文本区分_只取数字(s string) string {
	str := ""
	for _, r := range s {
		if unicode.IsNumber(r) {
			str = str + string(r)
		}
	}
	return str
}

func E文本区分_只取汉子(s string) string {
	str := ""
	for _, r := range s {
		if unicode.Is(unicode.Scripts["Han"], r) {
			str = str + string(r)
		}
	}
	return str
}
func E文本区分_只取符号(s string) string {
	str := ""
	for _, r := range s {
		if unicode.IsSymbol(r) {
			str = str + string(r)
		}
	}
	return str
}

func E文本_首字母改大写(s string) string {
	if len(s) < 1 {
		return ""
	}
	strArry := []rune(s)
	if strArry[0] >= 97 && strArry[0] <= 122 {
		strArry[0] -= 32
	}
	return string(strArry)
}

func E取文本字数(value string) int {
	return utf8.RuneCountInString(value)
}

// 子程序名：判断文本
// 真 找到 假 找不到
// 返回值类型：逻辑型
// 参数<1>的名称为“与判断的文本”，类型为“文本型”。
// 参数<2>的名称为“关键字”，类型为“文本型”，允许接收空参数数据。
// 参数<3>的名称为“更多关键字”，类型为“文本型”，允许接收空参数数据，需要接收数组数据。
func E判断文本(内容 string, 关键字 ...interface{}) bool {
	for _, v := range 关键字 {
		str := E到文本(v)
		if strings.Index(内容, str) != -1 {
			return true
		}
	}
	return false
}

func E判断文本s(内容 string, 关键字 ...interface{}) string {
	for _, v := range 关键字 {
		str := E到文本(v)
		if strings.Index(内容, str) != -1 {
			return str
		}
	}
	return ""
}

func E判断文本前缀(s string, 前缀 string) bool {
	return strings.HasPrefix(s, 前缀)
}
func E判断文本后缀(s string, 后缀 string) bool {
	return strings.HasSuffix(s, 后缀)
}

func E文本_单词首字母大写(s string) string {
	return gstr.UcFirst(s)
}

func E文本_句子首字母大写(s string) string {
	return gstr.UcWords(s)
}

func E文本_自动换行(s string, 每行几个字符 int, 分隔符 string) string {
	return gstr.WordWrap(s, 每行几个字符, 分隔符)
}

// SimilarText计算两个字符串之间的相似度。参见http://php.net/manual/en/function.similar-text.php。
func E文本_相似文本(first, second string, percent *float64) int {
	return gstr.SimilarText(first, second, percent)
}

func E文本_随机文本(s string) string {
	return gstr.Shuffle(s)
}

func E文本_搜索切片文本(a []string, s string) int {
	return gstr.SearchArray(a, s)
}
