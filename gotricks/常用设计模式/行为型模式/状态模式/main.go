package main

import "fmt"

/*
IPhone手机充电就是一个手机电池状态的流转，一开始手机处于有电状态，插入充电插头后，继续充电到满电状态，
并进入断电保护，拔出充电插头后使用手机，由满电逐渐变为没电，最终关机；
*/
func main() {
	iPhone13Pro := NewIPhone("13 pro") // 刚创建的手机有部分电

	fmt.Println(iPhone13Pro.BatteryState()) // 打印部分电状态
	fmt.Println(iPhone13Pro.ConnectPlug())  // 插上电源插头，继续充满电
	fmt.Println(iPhone13Pro.ConnectPlug())  // 满电后再充电，会触发满电保护

	fmt.Println(iPhone13Pro.DisconnectPlug()) // 拔掉电源，使用手机消耗电量，变为有部分电
	fmt.Println(iPhone13Pro.DisconnectPlug()) // 一直使用手机，直到没电
	fmt.Println(iPhone13Pro.DisconnectPlug()) // 没电后会关机

	fmt.Println(iPhone13Pro.ConnectPlug()) // 再次插上电源一会，变为有电状态
}

// BatteryState 电池状态接口，支持手机充电线插拔事件
type BatteryState interface {
	ConnectPlug(iPhone *IPhone) string
	DisconnectPlug(iPhone *IPhone) string
}

// fullBatteryState 满电状态
type fullBatteryState struct{}

func (s *fullBatteryState) String() string {
	return "满电状态"
}

func (s *fullBatteryState) ConnectPlug(iPhone *IPhone) string {
	return iPhone.pauseCharge()
}

func (s *fullBatteryState) DisconnectPlug(iPhone *IPhone) string {
	iPhone.SetBatteryState(PartBatteryState)
	return fmt.Sprintf("%s,%s转为%s", iPhone.consume(), s, PartBatteryState)
}

// emptyBatteryState 空电状态
type emptyBatteryState struct{}

func (s *emptyBatteryState) String() string {
	return "没电状态"
}

func (s *emptyBatteryState) ConnectPlug(iPhone *IPhone) string {
	iPhone.SetBatteryState(PartBatteryState)
	return fmt.Sprintf("%s,%s转为%s", iPhone.charge(), s, PartBatteryState)
}

func (s *emptyBatteryState) DisconnectPlug(iPhone *IPhone) string {
	return iPhone.shutdown()
}

// partBatteryState 部分电状态
type partBatteryState struct{}

func (s *partBatteryState) String() string {
	return "有电状态"
}

func (s *partBatteryState) ConnectPlug(iPhone *IPhone) string {
	iPhone.SetBatteryState(FullBatteryState)
	return fmt.Sprintf("%s,%s转为%s", iPhone.charge(), s, FullBatteryState)
}

func (s *partBatteryState) DisconnectPlug(iPhone *IPhone) string {
	iPhone.SetBatteryState(EmptyBatteryState)
	return fmt.Sprintf("%s,%s转为%s", iPhone.consume(), s, EmptyBatteryState)
}

// 电池状态单例，全局统一使用三个状态的单例，不需要重复创建
var (
	FullBatteryState  = new(fullBatteryState)  // 满电
	EmptyBatteryState = new(emptyBatteryState) // 空电
	PartBatteryState  = new(partBatteryState)  // 部分电
)

// IPhone 已手机充电为例，实现状态模式
type IPhone struct {
	model        string       // 手机型号
	batteryState BatteryState // 电池状态
}

// NewIPhone 创建指定型号手机
func NewIPhone(model string) *IPhone {
	return &IPhone{
		model:        model,
		batteryState: PartBatteryState,
	}
}

// BatteryState 输出电池当前状态
func (i *IPhone) BatteryState() string {
	return fmt.Sprintf("iPhone %s 当前为%s", i.model, i.batteryState)
}

// ConnectPlug 连接充电线
func (i *IPhone) ConnectPlug() string {
	return fmt.Sprintf("iPhone %s 连接电源线,%s", i.model, i.batteryState.ConnectPlug(i))
}

// DisconnectPlug 断开充电线
func (i *IPhone) DisconnectPlug() string {
	return fmt.Sprintf("iPhone %s 断开电源线,%s", i.model, i.batteryState.DisconnectPlug(i))
}

// SetBatteryState 设置电池状态
func (i *IPhone) SetBatteryState(state BatteryState) {
	i.batteryState = state
}

func (i *IPhone) charge() string {
	return "正在充电"
}

func (i *IPhone) pauseCharge() string {
	return "电已满,暂停充电"
}

func (i *IPhone) shutdown() string {
	return "手机关闭"
}

func (i *IPhone) consume() string {
	return "使用中,消耗电量"
}
