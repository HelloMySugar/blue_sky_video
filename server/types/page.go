package types

type Pager struct {
	Limit  int `param:"<in:query> <name:limit>"`
	Offset int `param:"<in:query> <name:offset>"`
}
