package main

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"strings"
)

type visitor struct {
	Methods map[string]ast.Node
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	st, ok := node.(*ast.TypeSpec)
	if !ok {
		return v
	}
	if st.Name.Name != "FullNode" {
		return nil
	}
	iface := st.Type.(*ast.InterfaceType)
	for _, m := range iface.Methods.List {
		if len(m.Names) > 0 {
			v.Methods[m.Names[0].Name] = m
		}
	}
	return v
}

func extractMethodDocs() map[string]string {
	pkg, err := build.Import("github.com/filecoin-project/lotus/api", "", 0)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, pkg.Root+"/api", nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ap := pkgs["api"]
	f := ap.Files[pkg.Root+"/api/api_full.go"]
	cmap := ast.NewCommentMap(fset, f, f.Comments)
	v := &visitor{make(map[string]ast.Node)}

	ast.Walk(v, ap)

	mdocs := make(map[string]string)
	for mn, node := range v.Methods {
		cs := cmap.Filter(node).Comments()
		if len(cs) == 0 {
			continue
		}
		last := cs[len(cs)-1].Text()
		if !strings.HasPrefix(last, "MethodGroup:") {
			mdocs[mn] = last
		}
	}
	return mdocs
}
