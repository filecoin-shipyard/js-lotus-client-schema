// +build !windows

package main

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io/fs"
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
	if st.Name.Name != "Common" && st.Name.Name != "FullNode" && st.Name.Name != "StorageMiner" && st.Name.Name != "Worker" {
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
	pkgs, err := parser.ParseDir(fset, pkg.Root+"/api", func(fi fs.FileInfo) bool {
		return strings.HasPrefix(fi.Name(), "api_")
	}, parser.AllErrors|parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ap := pkgs["api"]
	mdocs := make(map[string]string)

	fs := []*ast.File{
		ap.Files[pkg.Root+"/api/api_common.go"],
		ap.Files[pkg.Root+"/api/api_full.go"],
		ap.Files[pkg.Root+"/api/api_storage.go"],
		ap.Files[pkg.Root+"/api/api_worker.go"],
	}
	v := &visitor{make(map[string]ast.Node)}

	for _, f := range fs {
		cmap := ast.NewCommentMap(fset, f, f.Comments)

		ast.Walk(v, ap)

		for mn, node := range v.Methods {
			cs := cmap.Filter(node).Comments()
			if len(cs) == 0 {
				continue
			}
			var last string
			for _, c := range cs {
				if c.Text() != "" {
					last = c.Text()
				}
			}
			if !strings.HasPrefix(last, "MethodGroup:") {
				mdocs[mn] = last
			}
		}
	}
	return mdocs
}
