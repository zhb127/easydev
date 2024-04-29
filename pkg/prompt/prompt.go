package prompt

type Prompt interface {
	Run() (interface{}, error)
}
