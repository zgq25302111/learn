package main

type Ia interface {
	fa()
}

type Ib = interface {
	fb()
}

type Ic interface {
	fa()
	fb()
}

type Id = interface {
	Ia // embed Ia
	Ib // embed Ib
}

type Ie interface {
	Ia // embed Ia
	fb()
}

func main() {

}