// Generated by generate/main.go. DO NOT EDIT.

package golines

import (
	dst "github.com/dave/dst"
	"log"
)

func NodeToGraphNode(node dst.Node) *GraphNode {
	gNode := &GraphNode{Node: node, Edges: []*GraphEdge{}}
	var child *GraphNode
	switch n := node.(type) {
	case *dst.ArrayType:
		gNode.Type = "ArrayType"
		if n.Len != nil {
			child = NodeToGraphNode(n.Len)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Len"})
		}
		if n.Elt != nil {
			child = NodeToGraphNode(n.Elt)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Elt"})
		}
	case *dst.AssignStmt:
		gNode.Type = "AssignStmt"
		if n.Lhs != nil {
			for _, obj := range n.Lhs {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Lhs"})
			}
		}
		if n.Rhs != nil {
			for _, obj := range n.Rhs {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Rhs"})
			}
		}
	case *dst.BadDecl:
		gNode.Type = "BadDecl"
	case *dst.BadExpr:
		gNode.Type = "BadExpr"
	case *dst.BadStmt:
		gNode.Type = "BadStmt"
	case *dst.BasicLit:
		gNode.Type = "BasicLit"
		gNode.Value = n.Value
	case *dst.BinaryExpr:
		gNode.Type = "BinaryExpr"
		if n.X != nil {
			child = NodeToGraphNode(n.X)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "X"})
		}
		if n.Y != nil {
			child = NodeToGraphNode(n.Y)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Y"})
		}
	case *dst.BlockStmt:
		gNode.Type = "BlockStmt"
		if n.List != nil {
			for _, obj := range n.List {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "List"})
			}
		}
	case *dst.BranchStmt:
		gNode.Type = "BranchStmt"
		if n.Label != nil {
			child = NodeToGraphNode(n.Label)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Label"})
		}
	case *dst.CallExpr:
		gNode.Type = "CallExpr"
		if n.Fun != nil {
			child = NodeToGraphNode(n.Fun)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Fun"})
		}
		if n.Args != nil {
			for _, obj := range n.Args {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Args"})
			}
		}
	case *dst.CaseClause:
		gNode.Type = "CaseClause"
		if n.List != nil {
			for _, obj := range n.List {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "List"})
			}
		}
		if n.Body != nil {
			for _, obj := range n.Body {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Body"})
			}
		}
	case *dst.ChanType:
		gNode.Type = "ChanType"
		if n.Value != nil {
			child = NodeToGraphNode(n.Value)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Value"})
		}
	case *dst.CommClause:
		gNode.Type = "CommClause"
		if n.Comm != nil {
			child = NodeToGraphNode(n.Comm)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Comm"})
		}
		if n.Body != nil {
			for _, obj := range n.Body {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Body"})
			}
		}
	case *dst.CompositeLit:
		gNode.Type = "CompositeLit"
		if n.Type != nil {
			child = NodeToGraphNode(n.Type)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Type"})
		}
		if n.Elts != nil {
			for _, obj := range n.Elts {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Elts"})
			}
		}
	case *dst.DeclStmt:
		gNode.Type = "DeclStmt"
		if n.Decl != nil {
			child = NodeToGraphNode(n.Decl)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Decl"})
		}
	case *dst.DeferStmt:
		gNode.Type = "DeferStmt"
		if n.Call != nil {
			child = NodeToGraphNode(n.Call)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Call"})
		}
	case *dst.Ellipsis:
		gNode.Type = "Ellipsis"
		if n.Elt != nil {
			child = NodeToGraphNode(n.Elt)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Elt"})
		}
	case *dst.EmptyStmt:
		gNode.Type = "EmptyStmt"
	case *dst.ExprStmt:
		gNode.Type = "ExprStmt"
		if n.X != nil {
			child = NodeToGraphNode(n.X)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "X"})
		}
	case *dst.Field:
		gNode.Type = "Field"
		if n.Names != nil {
			for _, obj := range n.Names {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Names"})
			}
		}
		if n.Type != nil {
			child = NodeToGraphNode(n.Type)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Type"})
		}
		if n.Tag != nil {
			child = NodeToGraphNode(n.Tag)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Tag"})
		}
	case *dst.FieldList:
		gNode.Type = "FieldList"
		if n.List != nil {
			for _, obj := range n.List {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "List"})
			}
		}
	case *dst.File:
		gNode.Type = "File"
		if n.Name != nil {
			child = NodeToGraphNode(n.Name)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Name"})
		}
		if n.Decls != nil {
			for _, obj := range n.Decls {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Decls"})
			}
		}
		if n.Imports != nil {
			for _, obj := range n.Imports {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Imports"})
			}
		}
	case *dst.ForStmt:
		gNode.Type = "ForStmt"
		if n.Init != nil {
			child = NodeToGraphNode(n.Init)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Init"})
		}
		if n.Cond != nil {
			child = NodeToGraphNode(n.Cond)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Cond"})
		}
		if n.Post != nil {
			child = NodeToGraphNode(n.Post)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Post"})
		}
		if n.Body != nil {
			child = NodeToGraphNode(n.Body)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Body"})
		}
	case *dst.FuncDecl:
		gNode.Type = "FuncDecl"
		if n.Recv != nil {
			child = NodeToGraphNode(n.Recv)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Recv"})
		}
		if n.Name != nil {
			child = NodeToGraphNode(n.Name)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Name"})
		}
		if n.Type.Params != nil {
			child = NodeToGraphNode(n.Type.Params)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Params"})
		}
		if n.Type.Results != nil {
			child = NodeToGraphNode(n.Type.Results)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Results"})
		}
		if n.Body != nil {
			child = NodeToGraphNode(n.Body)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Body"})
		}
	case *dst.FuncLit:
		gNode.Type = "FuncLit"
		if n.Type != nil {
			child = NodeToGraphNode(n.Type)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Type"})
		}
		if n.Body != nil {
			child = NodeToGraphNode(n.Body)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Body"})
		}
	case *dst.FuncType:
		gNode.Type = "FuncType"
		if n.Params != nil {
			child = NodeToGraphNode(n.Params)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Params"})
		}
		if n.Results != nil {
			child = NodeToGraphNode(n.Results)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Results"})
		}
	case *dst.GenDecl:
		gNode.Type = "GenDecl"
		if n.Specs != nil {
			for _, obj := range n.Specs {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Specs"})
			}
		}
	case *dst.GoStmt:
		gNode.Type = "GoStmt"
		if n.Call != nil {
			child = NodeToGraphNode(n.Call)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Call"})
		}
	case *dst.Ident:
		gNode.Type = "Ident"
		gNode.Value = n.Name
	case *dst.IfStmt:
		gNode.Type = "IfStmt"
		if n.Init != nil {
			child = NodeToGraphNode(n.Init)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Init"})
		}
		if n.Cond != nil {
			child = NodeToGraphNode(n.Cond)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Cond"})
		}
		if n.Body != nil {
			child = NodeToGraphNode(n.Body)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Body"})
		}
		if n.Else != nil {
			child = NodeToGraphNode(n.Else)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Else"})
		}
	case *dst.ImportSpec:
		gNode.Type = "ImportSpec"
		if n.Name != nil {
			child = NodeToGraphNode(n.Name)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Name"})
		}
		if n.Path != nil {
			child = NodeToGraphNode(n.Path)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Path"})
		}
	case *dst.IncDecStmt:
		gNode.Type = "IncDecStmt"
		if n.X != nil {
			child = NodeToGraphNode(n.X)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "X"})
		}
	case *dst.IndexExpr:
		gNode.Type = "IndexExpr"
		if n.X != nil {
			child = NodeToGraphNode(n.X)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "X"})
		}
		if n.Index != nil {
			child = NodeToGraphNode(n.Index)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Index"})
		}
	case *dst.InterfaceType:
		gNode.Type = "InterfaceType"
		if n.Methods != nil {
			child = NodeToGraphNode(n.Methods)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Methods"})
		}
	case *dst.KeyValueExpr:
		gNode.Type = "KeyValueExpr"
		if n.Key != nil {
			child = NodeToGraphNode(n.Key)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Key"})
		}
		if n.Value != nil {
			child = NodeToGraphNode(n.Value)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Value"})
		}
	case *dst.LabeledStmt:
		gNode.Type = "LabeledStmt"
		if n.Label != nil {
			child = NodeToGraphNode(n.Label)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Label"})
		}
		if n.Stmt != nil {
			child = NodeToGraphNode(n.Stmt)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Stmt"})
		}
	case *dst.MapType:
		gNode.Type = "MapType"
		if n.Key != nil {
			child = NodeToGraphNode(n.Key)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Key"})
		}
		if n.Value != nil {
			child = NodeToGraphNode(n.Value)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Value"})
		}
	case *dst.Package:
		gNode.Type = "Package"
	case *dst.ParenExpr:
		gNode.Type = "ParenExpr"
		if n.X != nil {
			child = NodeToGraphNode(n.X)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "X"})
		}
	case *dst.RangeStmt:
		gNode.Type = "RangeStmt"
		if n.Key != nil {
			child = NodeToGraphNode(n.Key)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Key"})
		}
		if n.Value != nil {
			child = NodeToGraphNode(n.Value)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Value"})
		}
		if n.X != nil {
			child = NodeToGraphNode(n.X)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "X"})
		}
		if n.Body != nil {
			child = NodeToGraphNode(n.Body)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Body"})
		}
	case *dst.ReturnStmt:
		gNode.Type = "ReturnStmt"
		if n.Results != nil {
			for _, obj := range n.Results {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Results"})
			}
		}
	case *dst.SelectStmt:
		gNode.Type = "SelectStmt"
		if n.Body != nil {
			child = NodeToGraphNode(n.Body)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Body"})
		}
	case *dst.SelectorExpr:
		gNode.Type = "SelectorExpr"
		if n.X != nil {
			child = NodeToGraphNode(n.X)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "X"})
		}
		if n.Sel != nil {
			child = NodeToGraphNode(n.Sel)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Sel"})
		}
	case *dst.SendStmt:
		gNode.Type = "SendStmt"
		if n.Chan != nil {
			child = NodeToGraphNode(n.Chan)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Chan"})
		}
		if n.Value != nil {
			child = NodeToGraphNode(n.Value)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Value"})
		}
	case *dst.SliceExpr:
		gNode.Type = "SliceExpr"
		if n.X != nil {
			child = NodeToGraphNode(n.X)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "X"})
		}
		if n.Low != nil {
			child = NodeToGraphNode(n.Low)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Low"})
		}
		if n.High != nil {
			child = NodeToGraphNode(n.High)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "High"})
		}
		if n.Max != nil {
			child = NodeToGraphNode(n.Max)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Max"})
		}
	case *dst.StarExpr:
		gNode.Type = "StarExpr"
		if n.X != nil {
			child = NodeToGraphNode(n.X)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "X"})
		}
	case *dst.StructType:
		gNode.Type = "StructType"
		if n.Fields != nil {
			child = NodeToGraphNode(n.Fields)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Fields"})
		}
	case *dst.SwitchStmt:
		gNode.Type = "SwitchStmt"
		if n.Init != nil {
			child = NodeToGraphNode(n.Init)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Init"})
		}
		if n.Tag != nil {
			child = NodeToGraphNode(n.Tag)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Tag"})
		}
		if n.Body != nil {
			child = NodeToGraphNode(n.Body)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Body"})
		}
	case *dst.TypeAssertExpr:
		gNode.Type = "TypeAssertExpr"
		if n.X != nil {
			child = NodeToGraphNode(n.X)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "X"})
		}
		if n.Type != nil {
			child = NodeToGraphNode(n.Type)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Type"})
		}
	case *dst.TypeSpec:
		gNode.Type = "TypeSpec"
		if n.Name != nil {
			child = NodeToGraphNode(n.Name)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Name"})
		}
		if n.Type != nil {
			child = NodeToGraphNode(n.Type)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Type"})
		}
	case *dst.TypeSwitchStmt:
		gNode.Type = "TypeSwitchStmt"
		if n.Init != nil {
			child = NodeToGraphNode(n.Init)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Init"})
		}
		if n.Assign != nil {
			child = NodeToGraphNode(n.Assign)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Assign"})
		}
		if n.Body != nil {
			child = NodeToGraphNode(n.Body)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Body"})
		}
	case *dst.UnaryExpr:
		gNode.Type = "UnaryExpr"
		if n.X != nil {
			child = NodeToGraphNode(n.X)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "X"})
		}
	case *dst.ValueSpec:
		gNode.Type = "ValueSpec"
		if n.Names != nil {
			for _, obj := range n.Names {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Names"})
			}
		}
		if n.Type != nil {
			child = NodeToGraphNode(n.Type)
			gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Type"})
		}
		if n.Values != nil {
			for _, obj := range n.Values {
				child = NodeToGraphNode(obj)
				gNode.Edges = append(gNode.Edges, &GraphEdge{Dest: child, Relationship: "Values"})
			}
		}
	default:
		log.Println("Unrecognized type")
	}
	return gNode
}
