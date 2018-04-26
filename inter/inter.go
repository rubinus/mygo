package inter

import "fmt"

type People struct {
	Name string
	Age int
}

type Student struct {
	Class string
}

type (
	fib interface {
		do()
	}
	handfunc func(int) error
)

func (*handfunc) do() {
	panic("implement me")
}

func (s Student) WhichClass() {
	fmt.Println("my class : ", s.Class)
}

func (p People) Say() {
	fmt.Println("my name : ", p.Name)
}

func (p People) Sing() {
	fmt.Println("my age : ", p.Age)
}

func (p People) Call(num int)  {
	fmt.Println("my number : ", num)
}


type PAS interface {
	Say()
	Sing()
}

type ABC interface {
	WhichClass()
}

type Ainter interface {
	PAS
	ABC
}

type DEF struct {
	Name string
	Age int
	Class string
}

func (def DEF) Say() {
	fmt.Println("DEF my Age : ", def.Age)
}

func (def DEF) Sing() {
	fmt.Println("DEF my Name : ", def.Name)
}

func (def DEF) WhichClass() {
	fmt.Println("DEF my Class : ", def.Class)
}
