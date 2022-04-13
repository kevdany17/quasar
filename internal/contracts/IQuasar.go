package contracts

type IQuasar interface {
	GetLocation(distance ...float32) (x, y float32)
	GetMessage(messages ...[]string) (msg string)
}
