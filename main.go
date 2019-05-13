package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/lyraproj/pcore/px"
	"github.com/lyraproj/pcore/types"
	"github.com/lyraproj/servicesdk/lang/go/lyra"
)

type In struct {
	Name   string
	Number int
}

type Out struct {
	Hash map[string]px.Value
}

func main() {
	// Workflow input is from Hiera and a constant (declared by annotations in the In struct).
	lyra.Serve(`lyra_sample`, nil, &lyra.Workflow{
		Parameters: &In{},
		Return:     &Out{},
		Steps: map[string]lyra.Step{
			`start`: &lyra.Action{
				Do: func(input struct{ Name string }) struct{ X int } {
					hclog.Default().Info("start", "Name", input.Name)
					return struct{ X int }{32}
				}},
			`end`: &lyra.Action{
				Do: func(input struct {
					Name   string
					Number int
					X      int
				}) *Out {
					hclog.Default().Info("end", "Name", input.Name, `Number`, input.Number, `X`, input.X)
					return &Out{
						map[string]px.Value{
							`hey`: types.WrapString(input.Name),
							`ho`:  types.WrapInteger(int64(input.Number + input.X))}}
				},
			},
		}})
}
