package main

import (
	"sort"

	"github.com/dave/dst/gendst/data"
	"github.com/dave/jennifer/jen"
)

const (
	dstPath = "github.com/dave/dst"
)

// genNodeToGraphNode generates go code that converts a dst node to a golines library
// graph node. The latter can then be used to generate a dot file for debugging
// purposes.
func genNodeToGraphNode() error {
	f := jen.NewFile("golines")
	f.HeaderComment("Generated by generate/main.go. DO NOT EDIT.")

	// Sort the keys in the dst map so that output is in stable order
	dataKeys := []string{}

	for key := range data.Info {
		dataKeys = append(dataKeys, key)
	}

	sort.Strings(dataKeys)

	f.Func().Id("NodeToGraphNode").Params(
		jen.Id("node").Qual(dstPath, "Node"),
	).Op("*").Id("GraphNode").BlockFunc(
		func(g *jen.Group) {
			g.Id("gNode").Op(":=").Op("&").Id("GraphNode").Values(
				jen.Id("Node").Op(":").Id("node"),
				jen.Id("Edges").Op(":").Index().Op("*").Id("GraphEdge").Values(),
			)

			g.Var().Id("child").Op("*").Id("GraphNode")

			g.Switch(jen.Id("n").Op(":=").Id("node").Assert(jen.Id("type"))).BlockFunc(
				func(g *jen.Group) {
					for _, key := range dataKeys {
						parts := data.Info[key]

						g.Case(jen.Op("*").Qual(dstPath, key)).BlockFunc(
							func(g *jen.Group) {
								g.Id("gNode").Dot("Type").Op("=").Lit(key)

								for _, part := range parts {
									switch p := part.(type) {
									case data.Node:
										g.If(p.Field.Get("n").Op("!=").Nil()).Block(
											jen.Id("child").Op("=").Id("NodeToGraphNode").
												Call(
													p.Field.Get("n"),
												),
											jen.Id("gNode").Dot("Edges").Op("=").Append(
												jen.Id("gNode").Dot("Edges"),
												jen.Op("&").Id("GraphEdge").Values(
													jen.Id("Dest").Op(":").Id("child"),
													jen.Id("Relationship").Op(":").Lit(p.Name),
												),
											),
										)
									case data.List:
										g.If(p.Field.Get("n").Op("!=").Nil()).Block(
											jen.For(
												jen.List(jen.Id("_"), jen.Id("obj")).
													Op(":=").Range().Add(p.Field.Get("n")),
											).Block(
												jen.Id("child").Op("=").Id("NodeToGraphNode").
													Call(
														jen.Id("obj"),
													),
												jen.Id("gNode").Dot("Edges").Op("=").Append(
													jen.Id("gNode").Dot("Edges"),
													jen.Op("&").Id("GraphEdge").Values(
														jen.Id("Dest").Op(":").Id("child"),
														jen.Id("Relationship").Op(":").Lit(p.Name),
													),
												),
											),
										)
									case data.String:
										g.Id("gNode").Dot("Value").Op("=").Add(
											p.ValueField.Get("n"),
										)
									}
								}
							},
						)
					}
					g.Default().Block(
						jen.Qual("log", "Println").Call(
							jen.Lit("Unrecognized type"),
						),
					)
				},
			)

			g.Return().Id("gNode")
		},
	)

	f.Save("../graph_generated.go")

	return nil
}
