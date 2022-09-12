package designpatterns

type Factory interface {
	Create() Duck
}

type FactoryA struct {
}

func (a *FactoryA) Create() Duck {
	return &RedDuck{}
}

type FactoryB struct {
}

func (b *FactoryB) Create() Duck {
	return &GreenDuck{}
}

// func main() {
// 	// 创建对应的工厂对象。
// 	a := &FactoryA{}
// 	a.Create()
// }
