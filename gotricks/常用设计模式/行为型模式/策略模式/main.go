package main

import "fmt"

/*
北京是一个四季分明的城市，每个季节天气情况都有明显特点；我们定义一个显示天气情况的季节接口，具体的四季实现，
都会保存一个城市和天气情况的映射表，城市对象会包含季节接口，随着四季的变化，天气情况也随之变化；
*/
func main() {

	wxPay := &WxPay{}
	px := NewPayCtx(wxPay)
	px.Pay()
	// 假设现在发现微信支付没钱，改用三方支付进行支付
	thPay := &ThirdPay{}
	px.setPayBehavior(thPay)
	px.Pay()

	//Beijing := NewCity("北京", "四季分明")
	//
	//Beijing.SetSeason(NewSpring())
	//fmt.Println(Beijing)
	//
	//Beijing.SetSeason(NewSummer())
	//fmt.Println(Beijing)
	//
	//Beijing.SetSeason(NewAutumn())
	//fmt.Println(Beijing)
	//
	//Beijing.SetSeason(NewWinter())
	//fmt.Println(Beijing)
}

/**
------------------------------------------------------------------------------------
*/

type PayBehavior interface {
	OrderPay(px *PayCtx)
}

type WxPay struct{}

func (*WxPay) OrderPay(px *PayCtx) {
	fmt.Printf("Wx支付加工支付请求 %v\n", px.payParams)
	fmt.Println("正在使用Wx支付进行支付")
}

// 三方支付
type ThirdPay struct{}

func (*ThirdPay) OrderPay(px *PayCtx) {
	fmt.Printf("三方支付加工支付请求 %v\n", px.payParams)
	fmt.Println("正在使用三方支付进行支付")
}

type PayCtx struct {
	// 提供支付能力的接口实现
	payBehavior PayBehavior
	// 支付参数
	payParams map[string]interface{}
}

func (px *PayCtx) setPayBehavior(p PayBehavior) {
	px.payBehavior = p
}

func (px *PayCtx) Pay() {
	px.payBehavior.OrderPay(px)
}

func NewPayCtx(p PayBehavior) *PayCtx {
	// 支付参数，Mock数据
	params := map[string]interface{}{
		"appId": "234fdfdngj4",
		"mchId": 123456,
	}
	return &PayCtx{
		payBehavior: p,
		payParams:   params,
	}
}

/**
------------------------------------------------------------------------------------
*/

// Season 季节的策略接口，不同季节表现得天气不同
type Season interface {
	ShowWeather(city string) string // 显示指定城市的天气情况
}

type spring struct {
	weathers map[string]string // 存储不同城市春天气候
}

func NewSpring() *spring {
	return &spring{
		weathers: map[string]string{"北京": "干燥多风", "昆明": "清凉舒适"},
	}
}

func (s *spring) ShowWeather(city string) string {
	return fmt.Sprintf("%s的春天，%s;", city, s.weathers[city])
}

type summer struct {
	weathers map[string]string // 存储不同城市夏天气候
}

func NewSummer() *summer {
	return &summer{
		weathers: map[string]string{"北京": "高温多雨", "昆明": "清凉舒适"},
	}
}

func (s *summer) ShowWeather(city string) string {
	return fmt.Sprintf("%s的夏天，%s;", city, s.weathers[city])
}

type autumn struct {
	weathers map[string]string // 存储不同城市秋天气候
}

func NewAutumn() *autumn {
	return &autumn{
		weathers: map[string]string{"北京": "凉爽舒适", "昆明": "清凉舒适"},
	}
}

func (a *autumn) ShowWeather(city string) string {
	return fmt.Sprintf("%s的秋天，%s;", city, a.weathers[city])
}

type winter struct {
	weathers map[string]string // 存储不同城市冬天气候
}

func NewWinter() *winter {
	return &winter{
		weathers: map[string]string{"北京": "干燥寒冷", "昆明": "清凉舒适"},
	}
}

func (w *winter) ShowWeather(city string) string {
	return fmt.Sprintf("%s的冬天，%s;", city, w.weathers[city])
}

// City 城市
type City struct {
	name    string
	feature string
	season  Season
}

// NewCity 根据名称及季候特征创建城市
func NewCity(name, feature string) *City {
	return &City{
		name:    name,
		feature: feature,
	}
}

// SetSeason 设置不同季节，类似天气在不同季节的不同策略
func (c *City) SetSeason(season Season) {
	c.season = season
}

// String 显示城市的气候信息
func (c *City) String() string {
	return fmt.Sprintf("%s%s，%s", c.name, c.feature, c.season.ShowWeather(c.name))
}
