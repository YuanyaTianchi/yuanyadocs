package main

import (
	"fmt"
)

const (
	攻增  增强类型 = "攻增"
	暴增  增强类型 = "暴增"
	伤增  增强类型 = "伤增"
	首伤增 增强类型 = "首伤增"
	额伤量 增强类型 = "额伤量"
	伤量增 增强类型 = "伤量增"
	蓄伤增 增强类型 = "蓄伤增"

	攻增2 增强类型 = "攻增2"

	背击增  增强类型 = "背击增"
	噩梦法增 增强类型 = "噩梦法增"
	支配技增 增强类型 = "支配技增"
	支配时增 增强类型 = "支配时增"

	default率 = 1
	default系 = 0
)

var (
	刀锋残影3 = 刻印{名字: "刀锋残影3", 增强集: []增强体{
		{类型: 攻增2, 率: default率, 系: 0.48},
	}}
	咒术人偶3 = 刻印{名字: "咒术人偶3", 增强集: []增强体{
		{类型: 攻增, 率: default率, 系: 0.16},
	}}
	重量强化3 = 刻印{名字: "重量强化3", 增强集: []增强体{
		{类型: 攻增, 率: default率, 系: 0.18},
	}}
	以太充能3 = 刻印{名字: "以太充能3", 增强集: []增强体{
		{类型: 攻增, 率: default率, 系: 0.15},
	}}
	肾上腺素3 = 刻印{名字: "肾上腺素3", 增强集: []增强体{
		{类型: 攻增, 率: default率, 系: 0.05},
		{类型: 暴增, 率: 0.15, 系: default系},
	}}
	精密短刀3 = 刻印{名字: "精密短刀3", 增强集: []增强体{
		{类型: 暴增, 率: 0.2, 系: -0.12},
	}}
	尖刺重锤3 = 刻印{名字: "尖刺重锤3", 增强集: []增强体{
		{类型: 暴增, 率: 0, 系: 0.5},
		{类型: 伤增, 率: 0.1, 系: -0.2},
	}}
	怨恨3 = 刻印{名字: "怨恨3", 增强集: []增强体{
		{类型: 首伤增, 率: default率, 系: 0.2},
	}}
	奇袭大师3 = 刻印{名字: "奇袭大师3", 增强集: []增强体{
		{类型: 伤量增, 率: 1, 系: 0.25},
	}}
	蓄力强化3 = 刻印{名字: "蓄力强化3", 增强集: []增强体{
		{类型: 蓄伤增, 率: default率, 系: 0.2},
	}}
	迅捷利刃3 = 刻印{名字: "迅捷利刃3", 增强集: []增强体{
		{类型: 额伤量, 率: default率, 系: 0.18},
	}}

	背击 = 刻印{名字: "背击", 增强集: []增强体{
		{类型: 背击增, 率: 1, 系: 0.05},
	}}

	噩梦2支配2绝地2 = 刻印{名字: "噩梦2+支配2+绝地2", 增强集: []增强体{
		{类型: 噩梦法增, 率: default率, 系: 0.12},
		{类型: 支配技增, 率: default率, 系: 0.10},
		{类型: 支配时增, 率: default率, 系: 0.10},
		{类型: 暴增, 率: 0, 系: 0.72},
	}}
	噩梦2支配4 = 刻印{名字: "噩梦2+支配4", 增强集: []增强体{
		{类型: 噩梦法增, 率: default率, 系: 0.12},
		{类型: 支配技增, 率: default率, 系: 0.25},
		{类型: 支配时增, 率: default率, 系: 0.10},
	}}
)

type 增强类型 string

type 刻印 struct {
	名字  string
	增强集 []增强体
}

type 增强体 struct {
	类型 增强类型
	率  float64
	系  float64
}

func DamagePrint(刻印集 ...刻印) {
	var msg string

	// 分组
	增强组集 := make(map[增强类型][]增强体, 3)
	for _, 刻印 := range 刻印集 {
		for _, 增强 := range 刻印.增强集 {
			增强组, ok := 增强组集[增强.类型]
			if !ok {
				增强组 = make([]增强体, 0, 3)
			}
			增强组 = append(增强组, 增强)
			增强组集[增强.类型] = 增强组
		}
		msg = fmt.Sprintf("%s+%s", msg, 刻印.名字)
	}

	// 计算伤害
	var 最终期望系 float64 = 1
	var 期望系集 []kv = make([]kv, 0, len(增强组集))
	for k, 增强组 := range 增强组集 {
		var 组期望系, 组增强率, 组增强系 float64 = 1, 0, 0
		for _, 增强 := range 增强组 {
			组增强率 += 增强.率
			组增强系 += 增强.系
		}
		if 组增强率 > 1 {
			组增强率 = 1
		}
		组期望系 += 组增强率 * 组增强系

		期望系集 = append(期望系集, kv{k: k, v: 组期望系})
		最终期望系 *= 组期望系
	}

	// fmt.Println(期望系集)
	fmt.Printf("%s+ 期望伤害系数: %f\n", msg, 最终期望系)
}

type kv struct {
	k 增强类型
	v float64
}

func DamagePrintWithOptions(暴击递增率 float64, 刻印集 ...刻印) {
	// 添加默认暴伤
	刻印集 = append(刻印集, 刻印{
		// 名字: fmt.Sprintf("%d%s默认暴伤", 100, "%"),
		增强集: []增强体{
			{类型: 暴增, 率: 0, 系: 1.55},
		},
	})

	// 其它暴击率
	var 其它暴率 float64 = 0
	for 其它暴率 <= 1 {
		新刻印集 := append(刻印集, 刻印{
			名字: fmt.Sprintf("%d%s其它暴击", int(其它暴率*100), "%"),
			增强集: []增强体{
				{类型: 暴增, 率: 其它暴率, 系: 0},
			},
		})
		DamagePrint(新刻印集...)

		其它暴率 += 暴击递增率
		if 其它暴率 >= 1 {
			break
		}
	}
	fmt.Println("-----by 天竾-----")
}

func main() {
	fmt.Println("---气功下版本基础暴率20左右(副会心...?)，现版本35左右(副会心+巨浪)")
	// DamagePrintWithOptions(0.2, 迅捷利刃3, 重量强化3, 肾上腺素3, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 怨恨3, 重量强化3, 肾上腺素3, 尖刺重锤3)

	/*
		fmt.Println("国服下版本刀锋计算：")
		fmt.Println("基础爆伤增益1.55（头背套）")
		fmt.Println("500会心17.89暴击率，头背套17暴击率，背击10暴击，即其它暴击率44.89左右（按40算，会心不一定能吃到500），但是打本可能会有辅助")
		DamagePrintWithOptions(0.2, 刀锋残影3, 肾上腺素3, 奇袭大师3, 蓄力强化3)
		DamagePrintWithOptions(0.2, 刀锋残影3, 肾上腺素3, 奇袭大师3, 蓄力强化3, 精密短刀3)
		DamagePrintWithOptions(0.2, 刀锋残影3, 肾上腺素3, 奇袭大师3, 蓄力强化3, 尖刺重锤3)
		DamagePrintWithOptions(0.2, 刀锋残影3, 肾上腺素3, 奇袭大师3, 蓄力强化3, 咒术人偶3)
		DamagePrintWithOptions(0.2, 刀锋残影3, 肾上腺素3, 奇袭大师3, 蓄力强化3, 怨恨3)
		fmt.Println("精密肾上一起带，其它暴击高了以后暴击溢出，上限较低；实测刀锋残影攻击力与人偶攻击力非加算")
		fmt.Println("无溢出风险性价比选择精密短刀（很容易溢出）；有溢出风险选择尖刺重锤3（后续爆伤较高时会弱于人偶）；只考虑伤害无脑怨恨")
		fmt.Println("by Chihilis/天堇晴雨瑶桃竾")
		DamagePrintWithOptions(0.2, 刀锋残影3, 奇袭大师3, 蓄力强化3, 精密短刀3, 尖刺重锤3)
		DamagePrintWithOptions(0.2, 刀锋残影3, 奇袭大师3, 蓄力强化3, 怨恨3, 精密短刀3)
		DamagePrintWithOptions(0.2, 刀锋残影3, 奇袭大师3, 蓄力强化3, 怨恨3, 尖刺重锤3)
		fmt.Println("精密短刀3暴击率越低收益越高, 尖刺重锤3暴击率越高收益越高 两个一起带尖刺重锤会使收益变得更正向")
	*/

	// DamagePrintWithOptions(0.45, 噩梦2支配4, 重量强化3, 迅捷利刃3, 尖刺重锤3, 怨恨3)
	// DamagePrintWithOptions(0.45, 噩梦2支配2绝地2, 重量强化3, 迅捷利刃3, 奇袭大师3, 怨恨3)

	// // 迅捷减cd和觉醒刻印减cd乘算
	// fmt.Println("迅捷-31，觉醒套-18", float64(300)*OS(0.31)*OS(0.18))
	// fmt.Println("迅捷-31，觉醒套-18 + 主宰-15：", float64(300)*OS(0.31)*OS(0.18+0.15))
	// fmt.Println("迅捷-31，觉醒套-18, 觉醒2-25：", float64(300)*OS(0.31)*OS(0.18)*OS(0.25))
	// fmt.Println("迅捷-31，觉醒套-18 + 主宰-15, 觉醒1-10：", float64(300)*OS(0.31)*OS(0.18+0.15)*OS(0.1))
	// fmt.Println("迅捷-31，觉醒套-18 + 主宰-15, 觉醒2-25：", float64(300)*OS(0.31)*OS(0.18+0.15)*OS(0.25))

	// DamagePrintWithOptions(0.2, 肾上腺素3, 怨恨3, 重量强化3, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 迅捷利刃3, 怨恨3, 重量强化3, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 肾上腺素3, 奇袭大师3, 重量强化3, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 肾上腺素3, 奇袭大师3, 重量强化3, 迅捷利刃3)
	// DamagePrintWithOptions(0.2, 迅捷利刃3, 奇袭大师3, 重量强化3, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 怨恨3, 奇袭大师3, 重量强化3, 尖刺重锤3)

	// DamagePrintWithOptions(0.2, 怨恨3, 奇袭大师3, 肾上腺素3, 咒术人偶3)

	// DamagePrintWithOptions(0.2, 怨恨3, 奇袭大师3, 肾上腺素3, 咒术人偶3)

	// DamagePrintWithOptions(0.2, 怨恨3, 精密短刀3, 肾上腺素3, 咒术人偶3)
	// DamagePrintWithOptions(0.2, 迅捷利刃3, 重量强化3, 肾上腺素3, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 迅捷利刃3, 重量强化3, 肾上腺素3, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 怨恨3, 重量强化3, 肾上腺素3, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 怨恨3, 重量强化3, 肾上腺素3, 咒术人偶3)
	// DamagePrintWithOptions(0.2, 怨恨3, 重量强化3, 精密短刀3, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 怨恨3, 精密短刀3, 肾上腺素3, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 精密短刀3, 重量强化3, 肾上腺素3, 尖刺重锤3)

	// DamagePrintWithOptions(0.2, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 精密短刀3)
	// DamagePrintWithOptions(0.2, 重量强化3)
	// DamagePrintWithOptions(0.2, 精密短刀3, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 重量强化3, 尖刺重锤3)
	// DamagePrintWithOptions(0.2, 重量强化3, 怨恨3)
	
	// 逆天、怨恨、肾上、尖刺、精密
	// 逆天、怨恨、肾上、尖刺、重量

	DamagePrintWithOptions(0.2, 怨恨3,肾上腺素3,尖刺重锤3,精密短刀3)
	DamagePrintWithOptions(0.2, 怨恨3,肾上腺素3,尖刺重锤3,重量强化3)
	DamagePrintWithOptions(0.2, 怨恨3,肾上腺素3,精密短刀3,重量强化3)
}

func OS(num float64) float64 {
	return float64(1) - num
}

//3000000*1.64 =4920000
//3000000*1.48*1.16 =5150400
