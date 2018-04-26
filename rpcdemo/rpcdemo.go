package rpcdemo

import "errors"

type DemoService struct {
}

type Args struct {
	A, B int
}

func (DemoService) Div(args Args, result *float64) error {
	if args.B == 0 {
		return errors.New("分母不能是0")
	}
	*result = float64(args.A) / float64(args.B)
	return nil
}
