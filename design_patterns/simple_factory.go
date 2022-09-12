package designpatterns

type Duck interface {
}

type RedDuck struct {
}

type GreenDuck struct {
}

func SimpleFactory(color string) Duck {
	if color == "Red" {
		return &RedDuck{}
	} else if color == "Green" {
		return &GreenDuck{}
	}
	return nil
}
