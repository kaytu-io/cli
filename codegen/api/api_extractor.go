package api

import (
	"github.com/iancoleman/strcase"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type SubParam struct {
	Name        string
	Description string
	Type        string
}

func (sub SubParam) TypeWithPackage() string {
	if sub.Type[:1] == "*" {
		return "*models." + sub.Type[1:]
	} else if sub.Type[:2] == "[]" {
		return "[]models." + sub.Type[2:]
	}
	return "models." + sub.Type
}

func (sub SubParam) TypeFlagFunc() string {
	switch sub.Type {
	case "int64":
		return "ReadInt64Flag"
	case "*int64":
		return "ReadInt64OptionalFlag"
	case "[]int64":
		return "ReadIntArrayFlag"
	case "string":
		return "ReadStringFlag"
	case "*string":
		return "ReadStringOptionalFlag"
	case "[]string":
		return "ReadStringArrayFlag"
	case "bool":
		return "ReadBooleanFlag"
	case "*bool":
		return "ReadBooleanOptionalFlag"
	default:
		return "ReadStructFlag"
	}
}

type ParamType struct {
	Name      string
	SubParams []SubParam
}

type API struct {
	MethodName         string
	ParamTypes         []ParamType
	ResponseType       string
	ParamsFilePath     string
	ResponseFilePath   string
	RequiresAuth       bool
	Client             string
	ShortDescription   string
	LongDescription    string
	ResponseHasPayload bool
}

func (api API) ClientVarName() string {
	return strcase.ToCamel(api.Client)
}

func findApiFiles(clientsFolder, methodName string) (paramsFile string, responseFile string) {
	err := filepath.Walk(clientsFolder, func(path string, info fs.FileInfo, err error) error {
		if info == nil || info.IsDir() || strings.HasSuffix(info.Name(), "_client.go") || !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		contentBytes, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		content := string(contentBytes)
		if strings.Contains(content, methodName+"Params") {
			if strings.HasSuffix(path, "_parameters.go") {
				if paramsFile != "" {
					panic("double finding? 1:" + paramsFile + " 2:" + path)
				}
				paramsFile = path
			} else {
				panic("invalid file")
			}
		}

		if strings.Contains(content, methodName+"Reader") {
			if strings.HasSuffix(path, "_responses.go") {
				if responseFile != "" {
					panic("double finding?")
				}
				responseFile = path
			} else {
				panic("invalid file")
			}
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	return
}

func extractAPIsFromClient(clientsFolder string, clientPath string) []API {
	var apis []API

	clientContentBytes, err := os.ReadFile(clientPath)
	if err != nil {
		panic(err)
	}

	clientContent := string(clientContentBytes)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", clientContent, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	var functions []*ast.Field

	packageName := f.Name.Name
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Name.Name == "ClientService" {
				if it, ok := x.Type.(*ast.InterfaceType); ok {
					functions = it.Methods.List
				}
			}
		}
		return true
	})

	//var packageName string
	for _, f := range functions {
		funcType := f.Type.(*ast.FuncType)

		requiresAuth := false

		var fieldTypes []string
		for _, field := range funcType.Params.List {
			var fieldType string
			switch field.Type.(type) {
			case *ast.StarExpr:
				fieldType = field.Type.(*ast.StarExpr).X.(*ast.Ident).Name
			case *ast.SelectorExpr:
				fieldType = field.Type.(*ast.SelectorExpr).Sel.Name
			case *ast.Ellipsis:
				fieldType = field.Type.(*ast.Ellipsis).Elt.(*ast.Ident).Name
			default:
				panic("invalid type")
			}
			if strings.Contains(fieldType, "ClientOption") {
				continue
			}
			if strings.Contains(fieldType, "ClientAuthInfoWriter") {
				requiresAuth = true
				continue
			}

			fieldTypes = append(fieldTypes, fieldType)
		}

		var response string
		if funcType.Results != nil {
			for _, resp := range funcType.Results.List {
				switch resp.Type.(type) {
				case *ast.StarExpr:
					response = resp.Type.(*ast.StarExpr).X.(*ast.Ident).Name
				case *ast.Ident:
					x := resp.Type.(*ast.Ident).Name
					if x == "error" {
						continue
					}
				default:
					panic("invalid response type")
				}
			}
		}

		methodName := f.Names[0].Name
		if strings.Contains(methodName, "SetTransport") {
			continue
		}

		paramsFile, responseFile := findApiFiles(clientsFolder, f.Names[0].Name)
		if paramsFile == "" || responseFile == "" {
			panic("files not found")
		}

		shortDesc, LongDesc := findDescription(methodName, clientPath)

		var paramTypes []ParamType
		for _, fieldType := range fieldTypes {
			subParams := extractSubParams(fieldType, paramsFile)
			paramTypes = append(paramTypes, ParamType{
				Name:      fieldType,
				SubParams: subParams,
			})
		}

		hasPayload := extractResponseHasPayload(response, responseFile)
		apis = append(apis, API{
			MethodName:         methodName,
			ParamTypes:         paramTypes,
			ResponseType:       response,
			ResponseHasPayload: hasPayload,
			ParamsFilePath:     paramsFile,
			ResponseFilePath:   responseFile,
			RequiresAuth:       requiresAuth,
			Client:             packageName,
			ShortDescription:   shortDesc,
			LongDescription:    LongDesc,
		})
	}

	return apis
}

func extractResponseHasPayload(response string, file string) bool {
	clientContentBytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	clientContent := string(clientContentBytes)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", clientContent, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	hasPayload := false
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Name.Name == "GetPayload" {
				hasPayload = true
			}
		}
		return true
	})
	return hasPayload
}

func extractParamType(expr ast.Expr) string {
	switch expr.(type) {
	case *ast.Ident:
		return expr.(*ast.Ident).Name
	case *ast.ArrayType:
		return "[]" + extractParamType(expr.(*ast.ArrayType).Elt)
	case *ast.SelectorExpr:
		return expr.(*ast.SelectorExpr).Sel.Name
	case *ast.StarExpr:
		return "*" + extractParamType(expr.(*ast.StarExpr).X)
	default:
		panic("invalid type")
	}
}

func extractSubParams(fieldType string, file string) []SubParam {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", string(content), parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var subParams []SubParam
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			fname := x.Name.Name
			if x.Recv == nil || len(x.Recv.List) == 0 {
				return true
			}

			recParamType := extractParamType(x.Recv.List[0].Type)
			if strings.HasPrefix(recParamType, "*") {
				recParamType = recParamType[1:]
			}
			if recParamType != fieldType {
				return true
			}

			if !strings.HasPrefix(fname, "Set") {
				return true
			}
			if fname == "SetHTTPClient" || fname == "SetContext" || fname == "SetTimeout" || fname == "SetDefaults" {
				return true
			}

			if len(x.Type.Params.List) != 1 {
				panic("param list must only have one item")
			}

			ptype := x.Type.Params.List[0].Type
			stype := extractParamType(ptype)

			doc := extractParamDocument(f, x.Type.Params.List[0].Names[0].Name)
			subParams = append(subParams, SubParam{
				Name:        fname[3:],
				Description: strings.TrimSpace(doc),
				Type:        stype,
			})
		}
		return true
	})
	return subParams
}

func extractParamDocument(file *ast.File, name string) string {
	var doc string
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			switch t := x.Type.(type) {
			case *ast.StructType:
				for _, field := range t.Fields.List {
					fieldName := field.Names[0].Name
					if strings.ToLower(name) == strings.ToLower(fieldName) {
						doc = strings.TrimSpace(field.Doc.Text())
						if strings.HasPrefix(doc, fieldName) {
							doc = strings.TrimSpace(strings.Replace(doc, fieldName, "", 1))
							if doc[:1] == "." {
								doc = strings.TrimSpace(doc[1:])
							}
						}
					}
				}
			default:
				panic("unknown type")
			}
		}
		return true
	})
	return doc
}

func findDescription(funcName string, clientPath string) (string, string) {
	codeBytes, err := os.ReadFile(clientPath)
	if err != nil {
		panic(err)
	}
	code := string(codeBytes)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", code, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var comments *ast.CommentGroup
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Name.Name == funcName {
				comments = x.Doc
			}
		}
		return true
	})

	// Print the comments we found
	shortDesc := ""
	longDesc := comments.Text()
	respLines := strings.Split(longDesc, "\n")
	if strings.HasPrefix(strings.TrimSpace(respLines[0]), funcName) {
		shortDesc = strings.TrimSpace(respLines[0])
		shortDesc = strings.Replace(shortDesc, funcName, "", 1)
		shortDesc = strings.TrimSpace(shortDesc)

		respLines = respLines[1:]
		longDesc = strings.Join(respLines, "\n")
	}

	if len(shortDesc) > 0 {
		shortDesc = strings.ToUpper(shortDesc[:1]) + shortDesc[1:]
	}
	longDesc = strings.TrimSpace(longDesc)
	return shortDesc, longDesc
}

func ExtractAPIs(root string) []API {
	clientsFolder := root + "/pkg/api/kaytu/client"

	var apis []API
	err := filepath.Walk(clientsFolder, func(path string, info fs.FileInfo, err error) error {
		if info == nil || info.IsDir() || !strings.HasSuffix(info.Name(), "_client.go") {
			return nil
		}

		apis = append(apis, extractAPIsFromClient(clientsFolder, path)...)
		return nil
	})

	if err != nil {
		panic(err)
	}

	return apis
}
