package singleton

type singleton struct {
}

var instance *singleton

// Singletonインスタンスを取得するメソッドのみをsingletonパッケージからエクスポートできるようにし
// インスタンスの一意性を保つ
func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}
