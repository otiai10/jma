package gens

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/otiai10/jma"
)

var (
	invalidcharsForIdentity = regexp.MustCompile("[\\s\\(\\)]")
)

type (
	SortableOffice struct {
		SortKey int
		jma.Area
	}
	SortableOffices []SortableOffice
)

func (so SortableOffices) Len() int {
	return len(so)
}
func (so SortableOffices) Less(i, j int) bool {
	return so[i].SortKey < so[j].SortKey
}
func (so SortableOffices) Swap(i, j int) {
	so[i], so[j] = so[j], so[i]
}

func OfficeAreaToAstValueSpec(pos token.Pos, identity string, office jma.Area) *ast.ValueSpec {
	children := []ast.Expr{}
	for _, ch := range office.Children {
		children = append(children, &ast.BasicLit{
			Kind:  token.STRING,
			Value: doublequote(ch),
		})
	}
	return &ast.ValueSpec{
		Names: []*ast.Ident{{Name: identity, NamePos: pos}},
		Values: []ast.Expr{
			&ast.CompositeLit{
				Type: ast.NewIdent("Area"),
				Elts: []ast.Expr{
					&ast.KeyValueExpr{
						Key: ast.NewIdent("Code"),
						Value: &ast.BasicLit{
							Kind:  token.STRING,
							Value: doublequote(office.Code),
						},
					},
					&ast.KeyValueExpr{
						Key: ast.NewIdent("Name"),
						Value: &ast.BasicLit{
							Kind:  token.STRING,
							Value: doublequote(office.Name),
						},
					},
					&ast.KeyValueExpr{
						Key: ast.NewIdent("NameEn"),
						Value: &ast.BasicLit{
							Kind:  token.STRING,
							Value: doublequote(office.NameEn),
						},
					},
					&ast.KeyValueExpr{
						Key: ast.NewIdent("NameEnLower"),
						Value: &ast.BasicLit{
							Kind:  token.STRING,
							Value: doublequote(strings.ToLower(office.NameEn)),
						},
					},
					&ast.KeyValueExpr{
						Key: ast.NewIdent("Kana"),
						Value: &ast.BasicLit{
							Kind:  token.STRING,
							Value: doublequote(office.Kana),
						},
					},
					&ast.KeyValueExpr{
						Key: ast.NewIdent("OfficeName"),
						Value: &ast.BasicLit{
							Kind:  token.STRING,
							Value: doublequote(office.OfficeName),
						},
					},
					&ast.KeyValueExpr{
						Key: ast.NewIdent("Parent"),
						Value: &ast.BasicLit{
							Kind:  token.STRING,
							Value: doublequote(office.Parent),
						},
					},
					&ast.KeyValueExpr{
						Key: ast.NewIdent("Children"),
						Value: &ast.CompositeLit{
							Type: &ast.ArrayType{
								Elt: &ast.BasicLit{
									Kind:  token.STRING,
									Value: "string",
								},
							},
							Elts: children,
						},
					},
				},
			},
		},
	}
}

func GenerateArea(inputpath, outputpath string) error {

	res, err := http.Get(inputpath)
	if err != nil {
		return fmt.Errorf("http request failed: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode >= 400 {
		return fmt.Errorf("got non-200 response from server: %v", err)
	}
	body := map[string]map[string]jma.Area{}
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return fmt.Errorf("failed to decode response from server: %v", err)
	}
	offices := SortableOffices{}
	for code, o := range body["offices"] {
		key, err := strconv.Atoi(code)
		if err != nil {
			return fmt.Errorf("failed to conv atoi: %s: %v", code, err)
		}
		o.Code = code
		offices = append(offices, SortableOffice{key, o})
	}
	sort.Sort(offices)

	outfile := "offices.go"
	output, err := os.Open(outfile)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer output.Close()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, outfile, output, parser.AllErrors)
	if err != nil {
		return fmt.Errorf("failed to parse file: %s: %v", outfile, err)
	}

	file := fset.File(1)

	// Target 1
	definition := f.Decls[0].(*ast.GenDecl)
	dline := fset.Position(definition.Lparen).Line
	definition.Specs = []ast.Spec{}

	// Target 2
	listed := f.Decls[1].(*ast.GenDecl).Specs[0].(*ast.ValueSpec).Values[0].(*ast.CompositeLit)
	lline := fset.Position(listed.Lbrace).Line
	listed.Elts = []ast.Expr{}

	for _, o := range offices {
		identity := invalidcharsForIdentity.ReplaceAllString(o.NameEn, "")
		dline += 1
		lline += 1

		spec := OfficeAreaToAstValueSpec(file.LineStart(dline), identity, o.Area)
		definition.Specs = append(definition.Specs, spec)

		expr := &ast.BasicLit{ValuePos: file.LineStart(lline), Kind: token.STRING, Value: identity}
		listed.Elts = append(listed.Elts, expr)
	}

	// {{{ FIXME:
	buf := bytes.NewBuffer(nil)
	if err := format.Node(buf, fset, f); err != nil {
		return fmt.Errorf("failed to format ast nodes: %v", err)
	}
	if err := os.RemoveAll("./offices.go"); err != nil {
		return fmt.Errorf("failed to remove a file: %v", err)
	}
	o, err := os.Create("./offices.go")
	if err != nil {
		return fmt.Errorf("failed to create a new file: %v", err)
	}
	if _, err := io.Copy(o, buf); err != nil {
		return fmt.Errorf("failed to dump ast nodes to the new file: %v", err)
	}
	// }}}

	return nil
}
