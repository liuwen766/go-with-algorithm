package main

func test(key string) ([]byte, error) {
	return []byte(key), nil
}

type DB struct{ url string }

func (db *DB) Query(sql string, args ...string) string {
	// ...
	return "hello"
}

func (db *DB) Get(key string) ([]byte, error) {
	// ...
	v := db.Query("SELECT NAME FROM TABLE WHEN NAME= ?", key)
	return []byte(v), nil
}

func main() {
	// 调用方式一  GetterFunc 类型的函数作为参数   匿名函数
	GetFromSource(GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	}), "hello")
	// 调用方式一 GetterFunc 类型的函数作为参数  普通函数
	GetFromSource(GetterFunc(test), "hello")

	// 调用方式二 实现了 Getter 接口的结构体作为参数
	// 这种方式适用于逻辑较为复杂的场景，如果对数据库的操作需要很多信息，地址、用户名、密码，还有很多中间状态需要保持，比如超时、重连、加锁等等。这种情况下，更适合封装为一个结构体作为参数。
	// 这样，既能够将普通的函数类型（需类型转换）作为参数，也可以将结构体作为参数，使用更为灵活，可读性也更好，这就是接口型函数的价值。
	GetFromSource(new(DB), "hello")

}

// Getter 定义接口
type Getter interface {
	Get(key string) ([]byte, error)
}

// GetterFunc 定义一个函数    接口型函数
type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

// GetFromSource 使用场景：从某数据源获取结果
func GetFromSource(getter Getter, key string) []byte {
	get, err := getter.Get(key)
	if err == nil {
		return get
	}
	return nil
}
