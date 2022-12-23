package main

import "fmt"

/*
对于一段经历的描述，经历就可能有多种实现，比如旅游经历，探险经历这相当于第一层次的类结构，同时描述旅游经历或探险经历又包含多个维度，
比如如何到达目的地，在目的地开展了什么活动等，到达目的地有很多种方式，比如飞机、火车、汽车等；开展的活动又根据地点不同而不同，
海边可以冲浪，山地可以攀岩，荒漠可以徒步穿越等；这两个维度的变化点对于描述经历来说相当于第二层次类实现，通过接口被第一层次引用。

这里对于经历描述存在三个维度的变化，
1.经历本身的两个实现：旅游经历与探险经历。
2.交通方式的两个实现：飞机和汽车。
3.开展活动的三个实现：冲浪、攀岩与徒步穿越。

如果用一个类层次去实现就需要2*2*3=12个不同的实现类，如果用桥接模式仅需要2+2+3=7个不同的类，而且两种方式的加速度也不一样，
比如增加一个交通方式火车，非桥接模式需要增加2*3*3-12=6个实现类，桥接模式2+3+3-7=1个实现类；
桥接模式大大增加了类之间组合的灵活性。
*/
func main() {
	// 坐飞机去三亚度蜜月
	honeymoonTravel := NewTravelExperience("honeymoon", new(airplane), NewSeaside("SanyaYalongBay"))
	fmt.Println(honeymoonTravel.Describe())
	// 坐车去泰山毕业旅游
	graduationTrip := NewTravelExperience("graduationTrip", new(car), NewMountain("Tarzan"))
	fmt.Println(graduationTrip.Describe())

	// 野外生存培训后，坐车去罗布泊，徒步穿越
	desertAdventure := NewAdventureExperience("wilderness survival training", "adventure", new(car), NewDesert("Lop Nor"))
	fmt.Println(desertAdventure.Describe())
}

// Traffic 交通工具
type Traffic interface {
	Transport() string
}

// airplane 飞机
type airplane struct{}

// Transport 坐飞机
func (a *airplane) Transport() string {
	return "by airplane"
}

// car 汽车
type car struct{}

// Transport 坐汽车
func (t *car) Transport() string {
	return "by car"
}

// Location 地点
type Location interface {
	Name() string       // 地点名称
	PlaySports() string // 参与运动
}

// namedLocation 被命名的地点，统一引用此类型，声明名字字段及获取方法
type namedLocation struct {
	name string
}

// Name 获取地点名称
func (n namedLocation) Name() string {
	return n.name
}

// seaside 海边
type seaside struct {
	namedLocation
}

// NewSeaside 创建指定名字的海边，比如三亚湾
func NewSeaside(name string) *seaside {
	return &seaside{
		namedLocation: namedLocation{
			name: name,
		},
	}
}

// PlaySports 海边可以冲浪
func (s *seaside) PlaySports() string {
	return fmt.Sprintf("surfing")
}

// mountain 山
type mountain struct {
	namedLocation
}

// NewMountain 创建指定名字的山，比如泰山
func NewMountain(name string) *mountain {
	return &mountain{
		namedLocation: namedLocation{
			name: name,
		},
	}
}

// PlaySports 可以爬山
func (m *mountain) PlaySports() string {
	return fmt.Sprintf("climbing")
}

// desert 荒漠
type desert struct {
	namedLocation
}

// NewDesert 创建指定名字的荒漠，比如罗布泊
func NewDesert(name string) *desert {
	return &desert{
		namedLocation: namedLocation{
			name: name,
		},
	}
}

// PlaySports 荒漠可以徒步穿越
func (d *desert) PlaySports() string {
	return fmt.Sprintf("trekking")
}

// Experience 经历
type Experience interface {
	Describe() string // 描述经历
}

// travelExperience 旅游经历
type travelExperience struct {
	subject  string
	traffic  Traffic
	location Location
}

// NewTravelExperience 创建旅游经历，包括主题、交通方式、地点
func NewTravelExperience(subject string, traffic Traffic, location Location) *travelExperience {
	return &travelExperience{
		subject:  subject,
		traffic:  traffic,
		location: location,
	}
}

// Describe 描述旅游经历
func (t *travelExperience) Describe() string {
	return fmt.Sprintf("%s is to %s %s and %s", t.subject, t.location.Name(), t.traffic.Transport(), t.location.PlaySports())
}

// adventureExperience 探险经历
type adventureExperience struct {
	survivalTraining string
	travelExperience
}

// NewAdventureExperience 创建探险经历，包括探险需要的培训，其他的与路由参数类似
func NewAdventureExperience(training string, subject string, traffic Traffic, location Location) *adventureExperience {
	return &adventureExperience{
		survivalTraining: training,
		travelExperience: *NewTravelExperience(subject, traffic, location),
	}
}

// Describe 描述探险经历
func (a *adventureExperience) Describe() string {
	return fmt.Sprintf("after %s, %s", a.survivalTraining, a.travelExperience.Describe())
}
