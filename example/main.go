package main

import (
	"fmt"
	"github.com/netroby/mw256enc"
)

func main() {

	originStr := `今年5月，Hyperloop One在位于拉斯维加斯的测试跑道上完成了超级高铁的首次“全真空条件”下的测试。当时的测试时间较短，高铁在真空管道中只跑了5.3秒，最高时速只有70英里/小时（约113公里/小时）。本月初，Hyperloop One宣布公司已经于7月29日在内华达州的一处沙漠中完成第二阶段的测试，最高时速达到310km/h。

而后者Hyperloop Transportation Technologies公司目前正和印度政府进行紧密的磋商，尝试在全球第二大人口国家提供超级高铁服务。然而这些显然还没有达到马斯克的预期，上个月有报道称马斯克非常严谨的使用“The Boring Company” 来挖掘隧道来支持Hyperloop项目的建设。

数日之后，马斯克自己在推文中证实已经获得了“口头[政府]批准”，挖掘一条从纽约出发、途径费城、巴尔的摩，最终到达华盛顿的地下超级高铁，将4小时的车程缩短到29分钟。

而现在，《连线》（Wired）和《彭博社》（Bloomberg）同时获得来自于马斯克的路线总规划，创建一家Boring Company Hyperloop公司来挖掘从纽约到华盛顿的地下通道，并会设立多个站点。

Boring Company发言人向连线透露：

在Boring Company，我们计划创建低成本、快速挖掘的通道来构建新型高速运输系统。由电力驱动的常规压力隧道能够达到125mph以上的运输速度。而对于类似纽约到华盛顿的长距离运输，在降压隧道中能够确保运输仓最高达到600mph以上的速度，也就是Hyperloop。
	`
	fmt.Printf("%s\n", originStr)
	fmt.Printf("Origin str len %d\n", len(originStr))

	encodedStr := mw256enc.Encode([]byte(originStr))
	fmt.Println(encodedStr)
	fmt.Printf("encodedStr len %d\n", len(encodedStr))
	decodedStr, _ := mw256enc.Decode(encodedStr)
	fmt.Printf("%s\n", decodedStr)

}
