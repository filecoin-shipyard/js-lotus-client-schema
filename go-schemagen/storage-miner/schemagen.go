package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"unicode"

	"github.com/filecoin-project/lotus/api"
)

var Methods = map[string]interface{}{}

type Visitor struct {
	Methods map[string]ast.Node
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	st, ok := node.(*ast.TypeSpec)
	if !ok {
		return v
	}

	// if st.Name.Name != "Common" {
	// if st.Name.Name != "FullNode" {
	if st.Name.Name != "StorageMiner" {
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

const noComment = "There are not yet any comments for this method."

func parseApiASTInfo() (map[string]string, map[string]string) {

	fset := token.NewFileSet()
	// apiDir := filepath.Join(os.Args[1], "api")
	apiDir := filepath.Join("../../lotus", "api")
	pkgs, err := parser.ParseDir(fset, apiDir, nil,
		parser.AllErrors|parser.ParseComments)
	if err != nil {
		fmt.Println("parse error: ", err)
	}

	ap := pkgs["api"]

	// f := ap.Files[filepath.Join(apiDir, "api_full.go")]
	// f := ap.Files[filepath.Join(apiDir, "api_common.go")]
	f := ap.Files[filepath.Join(apiDir, "api_storage.go")]

	cmap := ast.NewCommentMap(fset, f, f.Comments)

	v := &Visitor{make(map[string]ast.Node)}
	ast.Walk(v, pkgs["api"])

	groupDocs := make(map[string]string)
	out := make(map[string]string)
	for mn, node := range v.Methods {
		cs := cmap.Filter(node).Comments()
		if len(cs) == 0 {
			out[mn] = noComment
		} else {
			for _, c := range cs {
				if strings.HasPrefix(c.Text(), "MethodGroup:") {
					parts := strings.Split(c.Text(), "\n")
					groupName := strings.TrimSpace(parts[0][12:])
					comment := strings.Join(parts[1:], "\n")
					groupDocs[groupName] = comment

					break
				}
			}

			last := cs[len(cs)-1].Text()
			if !strings.HasPrefix(last, "MethodGroup:") {
				out[mn] = last
			} else {
				out[mn] = noComment
			}
		}
	}
	return out, groupDocs
}

type MethodGroup struct {
	GroupName string
	Header    string
	Methods   []*Method
}

type Method struct {
	Comment         string
	Name            string
	/*
	InputExample    string
	ResponseExample string
	*/
}

func methodGroupFromName(mn string) string {
	i := strings.IndexFunc(mn[1:], func(r rune) bool {
		return unicode.IsUpper(r)
	})
	if i < 0 {
		return ""
	}
	return mn[:i+1]
}

func main() {

	comments, groupComments := parseApiASTInfo()

	groups := make(map[string]*MethodGroup)

	// var api struct{ api.Common }
	// var api struct{ api.FullNode }
	var api struct{ api.StorageMiner }
	t := reflect.TypeOf(api)
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		groupName := methodGroupFromName(m.Name)

		g, ok := groups[groupName]
		if !ok {
			g = new(MethodGroup)
			g.Header = groupComments[groupName]
			g.GroupName = groupName
			groups[groupName] = g
		}

		/*
		var args []interface{}
		ft := m.Func.Type()
		for j := 2; j < ft.NumIn(); j++ {
			inp := ft.In(j)
			args = append(args, exampleValue(inp))
		}

		v, err := json.Marshal(args)
		if err != nil {
			panic(err)
		}

		outv := exampleValue(ft.Out(0))

		ov, err := json.Marshal(outv)
		if err != nil {
			panic(err)
		}
		*/

		g.Methods = append(g.Methods, &Method{
			Name:            m.Name,
			Comment:         comments[m.Name],
			// InputExample:    string(v),
			// ResponseExample: string(ov),
		})
	}

	var groupslice []*MethodGroup
	for _, g := range groups {
		groupslice = append(groupslice, g)
	}

	sort.Slice(groupslice, func(i, j int) bool {
		return groupslice[i].GroupName < groupslice[j].GroupName
	})

	methods := make(map[string]interface{})
	for _, g := range groupslice {
		// fmt.Printf("## %s\n", g.GroupName)
		// fmt.Printf("%s\n\n", g.Header)

		sort.Slice(g.Methods, func(i, j int) bool {
			return g.Methods[i].Name < g.Methods[j].Name
		})

		for _, m := range g.Methods {
			// fmt.Printf("### %s\n", m.Name)
			record := make(map[string]interface{})
			if m.Name == "ChainNotify" {
				record["subscription"] = true
			}
			methods[m.Name] = record

			// fmt.Printf("%s\n\n", m.Comment)

			// fmt.Printf("Inputs: `%s`\n\n", m.InputExample)
			// fmt.Printf("Response: `%s`\n\n", m.ResponseExample)
		}

	}
	output, err := json.Marshal(methods)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", output)
}
