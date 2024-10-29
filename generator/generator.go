package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/iancoleman/strcase"
)

type EnumT struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	CalcValue int    `json:"calc_value"`
}

type FieldT struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type structsAndEnumsT struct {
	Enums   map[string][]EnumT  `json:"enums"`
	Structs map[string][]FieldT `json:"structs"`
}

type typedefsDictT map[string]string

type definitionsT map[string][]definitionT

type definitionT struct {
	Args  string `json:"args"`
	ArgsT []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"argsT"`
	Name           string `json:"cimguiname"`
	OverloadedName string `json:"ov_cimguiname"`
	Ret            string `json:"ret"`
}

func main() {
	content, err := os.ReadFile("thirdparty/cimgui/generator/output/structs_and_enums.json")
	if err != nil {
		panic(err)
	}

	// Parse the JSON file.
	var structsAndEnums structsAndEnumsT
	err = json.NewDecoder(bytes.NewReader(content)).Decode(&structsAndEnums)
	if err != nil {
		panic(err)
	}

	content, err = os.ReadFile("thirdparty/cimgui/generator/output/definitions.json")
	if err != nil {
		panic(err)
	}

	// Parse the JSON file.
	var definitions definitionsT
	err = json.NewDecoder(bytes.NewReader(content)).Decode(&definitions)
	if err != nil {
		panic(err)
	}

	content, err = os.ReadFile("thirdparty/cimgui/generator/output/typedefs_dict.json")
	if err != nil {
		panic(err)
	}

	// Parse the JSON file.
	var typedefsDict map[string]string
	err = json.NewDecoder(bytes.NewReader(content)).Decode(&typedefsDict)
	if err != nil {
		panic(err)
	}

	copyFile("thirdparty/cimgui/cimgui.h", "dist/include/cimgui/cimgui.h")
	copyFile("thirdparty/cimgui/cimgui_impl.h", "dist/include/cimgui/cimgui_impl.h")

	generateConstants(&structsAndEnums)
	generateTypedefs(typedefsDict)
	generateDefinitions(definitions)
	generateWrappersHeaders(definitions)
	generateWrappersSources(definitions)
}

func copyFile(src, dst string) error {
	_ = os.MkdirAll(filepath.Dir(dst), 0750)
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = srcFile.WriteTo(dstFile)
	if err != nil {
		return err
	}

	fmt.Printf("Copied file %s to %s\n", src, dst)

	return nil
}

func generateConstants(structsAndEnums *structsAndEnumsT) {
	constantsContent := strings.Builder{}
	constantsContent.WriteString("package imgui\n")
	constantsContent.WriteString("\n")

	flattenedEnums := make([]EnumT, 0, len(structsAndEnums.Enums))
	for _, enum := range structsAndEnums.Enums {
		flattenedEnums = append(flattenedEnums, enum...)
	}

	sort.Slice(flattenedEnums, func(i, j int) bool {
		return flattenedEnums[i].Name < flattenedEnums[j].Name
	})

	for _, field := range flattenedEnums {
		renamedName := strings.ReplaceAll(field.Name, "ImGui", "")
		renamedValue := strings.ReplaceAll(field.Value, "ImGui", "")
		renamedValue = strings.ReplaceAll(renamedValue, "~", "^")
		renamedValue = strings.ReplaceAll(renamedValue, "(int)", "")

		constantsContent.WriteString(fmt.Sprintf("const %s = %s\n", renamedName, renamedValue))
	}

	err := os.WriteFile("imgui_constants.go", []byte(constantsContent.String()), 0644)
	if err != nil {
		panic(err)
	}
}

func isImLike(s string) bool {
	if strings.HasPrefix(s, "Im") {
		r, _ := utf8.DecodeRuneInString(s[2:])
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func generateTypedefs(typedefsDict typedefsDictT) {
	typedefsContent := strings.Builder{}
	typedefsContent.WriteString("package imgui\n")
	typedefsContent.WriteString("\n")
	typedefsContent.WriteString("// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS 1\n")
	typedefsContent.WriteString("// #include \"dist/include/cimgui/cimgui.h\"\n")
	typedefsContent.WriteString("import \"C\"\n")
	typedefsContent.WriteString("\n")

	// sortedNames := make([]string, 0, len(structsAndEnums.Structs))
	// for name := range structsAndEnums.Structs {
	// 	sortedNames = append(sortedNames, name)
	// }
	// sort.Strings(sortedNames)

	sortedNames := make([]string, 0, len(typedefsDict))
	for name := range typedefsDict {
		sortedNames = append(sortedNames, name)
	}
	sort.Strings(sortedNames)

	blacklist := []string{
		"const_iterator", "iterator", "value_type", "STB_TexteditState",
		"ImVec1", "ImVec2", "ImVec2ih", "ImVec4", // These are specially handled and replaced by mgl32.Vec2 and mgl32.Vec4.
	}

	for _, name := range sortedNames {
		if slices.Contains(blacklist, name) {
			continue
		}

		// s := structsAndEnums.Structs[name]
		// _ = s

		// Use `go tool cgo -godefs` to generate the Go struct from the C struct?

		// reflect.TypeOf()
		// typedefsContent.WriteString(fmt.Sprintf("type %s struct {\n", name))
		// for _, field := range s {
		// typedefsContent.WriteString(fmt.Sprintf("\t%s %s\n", safeIdentifier(field.Name), cToGoType(field.Type)))
		// }
		// typedefsContent.WriteString("}\n\n")

		typedefsContent.WriteString(fmt.Sprintf("type %s C.%s\n", cToGoType(name), name))
	}

	err := os.WriteFile("imgui_typedefs.go", []byte(typedefsContent.String()), 0644)
	if err != nil {
		panic(err)
	}
}

func cToCgoType(cType string) string {
	originalCType := cType

	if strings.HasPrefix(cType, "const ") {
		cType = strings.TrimSpace(strings.TrimPrefix(cType, "const "))
		return cToCgoType(cType)
	}

	switch cType {
	case "char":
		return "C.char"
	case "bool":
		return "C.bool"
	case "int":
		return "C.int"
	case "float":
		return "C.float"
	case "float*":
		return "*C.float"
	case "double":
		return "C.double"
	case "double*":
		return "*C.double"
	case "int*":
		return "*C.int"
	case "unsigned int*":
		return "*C.uint"
	case "unsigned int":
		return "C.uint"
	case "char*":
		return "*C.char"
	case "bool*":
		return "*C.bool"
	case "void*":
		return "unsafe.Pointer"
	case "size_t":
		return "C.size_t"
	case "size_t*":
		return "*C.size_t"
	case "int[2]":
		return "*C.int"
	case "int[3]":
		return "*C.int"
	case "int[4]":
		return "*C.int"
	case "float[2]":
		return "*C.float"
	case "float[3]":
		return "*C.float"
	case "float[4]":
		return "*C.float"
	case "char[5]":
		return "*C.char"
	case "char* const[]":
		return "**C.char"
	case "unsigned char":
		return "C.uchar"
	case "unsigned char*":
		return "*C.uchar"
	case "unsigned char[256]":
		return "*C.uchar"
	}

	if isImLike(cType) || strings.HasPrefix(cType, "ig") {
		if strings.HasSuffix(cType, "**") {
			cType = "**C." + strings.TrimSuffix(cType, "**")
		} else if strings.HasSuffix(cType, "*") {
			cType = "*C." + strings.TrimSuffix(cType, "*")
		} else {
			cType = "C." + cType
		}
		return cType
	}

	panic(fmt.Errorf("unknown C type to convert to Cgo: %s", originalCType))
}

func cToGoType(cType string) string {
	originalCType := cType

	if strings.HasPrefix(cType, "const ") {
		cType = strings.TrimSpace(strings.TrimPrefix(cType, "const "))
		return cToGoType(cType)
	}

	switch cType {
	case "int":
		return "int"
	case "int*":
		return "*int"
	case "int[2]":
		return "[2]int"
	case "int[3]":
		return "[3]int"
	case "int[4]":
		return "[4]int"
	case "float":
		return "float32"
	case "float*":
		return "*float32"
	case "float[2]":
		return "[2]float32"
	case "float[3]":
		return "*mgl32.Vec3"
	case "float[4]":
		return "*mgl32.Vec4"
	case "double":
		return "float64"
	case "double*":
		return "*float64"
	case "char":
		return "byte"
	case "bool":
		return "bool"
	case "bool*":
		return "*bool"
	case "char*":
		return "string"
	case "unsigned int*":
		return "*uint"
	case "unsigned int":
		return "uint"
	case "void*":
		return "unsafe.Pointer"
	case "size_t":
		return "uint"
	case "size_t*":
		return "*uint"
	case "unsigned char[256]":
		return "[256]byte"
	case "unsigned char*":
		return "[]byte"
	case "unsigned char":
		return "byte"
	case "char* const[]":
		return "[]string"
	case "char[5]":
		return "[5]byte"

		// case "short":
		// 	return "int16"
		// case "unsigned short":
		// 	return "uint16"
		// case "bool(*)(ImFontAtlas* atlas)":
		// 	return "func(atlas *ImFontAtlas) bool"
		// case "void(*)(ImGuiContext* ctx,ImGuiDockNode* node,ImGuiTabBar* tab_bar)":
		// 	return "func(ctx *ImGuiContext, node *ImGuiDockNode, tab_bar *ImGuiTabBar)"
	}

	switch cType {
	case "ImVec2*":
		return "*mgl32.Vec2"
	case "ImVec2":
		return "mgl32.Vec2"
	case "ImVec4":
		return "mgl32.Vec4"
	case "ImVec4*":
		return "*mgl32.Vec4"
	}

	if strings.HasPrefix(cType, "ImGui") {
		cType = strings.TrimPrefix(cType, "ImGui")
		if strings.HasSuffix(cType, "**") {
			cType = "**" + strings.TrimSuffix(cType, "**")
		}
		if strings.HasSuffix(cType, "*") {
			cType = "*" + strings.TrimSuffix(cType, "*")
		}
		return cType
	}

	if isImLike(cType) {
		cType = cType[2:]
		if strings.HasSuffix(cType, "*") {
			cType = "*" + strings.TrimSuffix(cType, "*")
		}
		return cType
	}

	panic(fmt.Errorf("unknown C type to convert go Go: %s", originalCType))
}

func safeIdentifier(s string) string {
	s = strcase.ToLowerCamel(s)

	switch s {
	case "type":
		return "ty"
	}

	return s
}

func generateDefinitions(definitions definitionsT) {
	output := strings.Builder{}
	output.WriteString("package imgui\n")
	output.WriteString("\n")
	output.WriteString("// #define CIMGUI_DEFINE_ENUMS_AND_STRUCTS 1\n")
	output.WriteString("// #include \"dist/include/cimgui/cimgui.h\"\n")
	output.WriteString("// #include \"imgui_wrappers.h\"\n")
	output.WriteString("// #include <stdbool.h>\n")
	output.WriteString("import \"C\"\n")
	output.WriteString("import \"fmt\"\n")
	output.WriteString("import \"unsafe\"\n")
	output.WriteString("import \"github.com/go-gl/mathgl/mgl32\"\n")
	output.WriteString("\n")

	sortedDefinitions := make([]definitionT, 0, len(definitions))
	for _, definition := range definitions {
		sortedDefinitions = append(sortedDefinitions, definition...)
	}
	sort.Slice(sortedDefinitions, func(i, j int) bool {
		return sortedDefinitions[i].OverloadedName < sortedDefinitions[j].OverloadedName
	})

	blacklist := []string{
		"igNewFrame",                 // Needed to be overridden to control the pool memory allocator.
		"igGetAllocatorFunctions",    // Wont be tweaking the allocator functions from Go.
		"igAddDrawListToDrawDataEx",  // TODO: Wants an ImVector_ImDrawListPtr*
		"igDockBuilderCopyDockSpace", // TODO: Wants a ImVector_const_charPtr*
		"igDockBuilderCopyNode",      // TODO: Wants a ImVector_ImGuiID*
	}

	for _, def := range sortedDefinitions {
		if slices.Contains(blacklist, def.OverloadedName) {
			continue
		}

		if !strings.HasPrefix(def.OverloadedName, "ig") {
			continue
		}

		if strings.HasPrefix(def.OverloadedName, "igIm") {
			continue
		}

		// fmt.Println("=>", def.OverloadedName)

		renamedName := strings.TrimPrefix(def.OverloadedName, "ig")
		if isImLike(renamedName) {
			renamedName = renamedName[2:]
		}

		// Skip the debug functions.
		if strings.HasPrefix(renamedName, "Debug") {
			continue
		}

		// Skip va_list functions.
		hasVaList := false
		for _, arg := range def.ArgsT {
			if arg.Type == "va_list" {
				hasVaList = true
				break
			}
		}

		if hasVaList {
			continue
		}

		// Skip functions with function pointer parameters for now.
		hasFunctionPointer := false
		for _, arg := range def.ArgsT {
			if strings.Contains(arg.Type, "(*)") {
				hasFunctionPointer = true
				break
			}
		}

		if hasFunctionPointer {
			continue
		}

		output.WriteString(fmt.Sprintf("func %s(", renamedName))

		// Parameters.
		hasVariadic := false
		{
			for i, arg := range def.ArgsT {
				if i > 0 {
					output.WriteString(", ")
				}

				isNextArgumentVariadic := len(def.ArgsT) > i+1 && def.ArgsT[i+1].Name == "..."

				if isNextArgumentVariadic {
					output.WriteString("vfmt string, vargs ...interface{}")
					hasVariadic = true
					break
				} else {
					output.WriteString(safeIdentifier(arg.Name))
					output.WriteString(" ")
					output.WriteString(cToGoType(arg.Type))
				}
			}
		}

		output.WriteString(") ")

		hasReturnType := def.Ret != "void" && def.Ret != ""

		// Return type.
		{
			if hasReturnType {
				output.WriteString(cToGoType(def.Ret))
				output.WriteString(" ")
			}
		}

		output.WriteString("{\n")

		// Body.
		{
			for i, arg := range def.ArgsT {
				isNextArgumentVariadic := len(def.ArgsT) > i+1 && def.ArgsT[i+1].Name == "..."

				output.WriteString(fmt.Sprintf("\ta%d := ", i))

				expr := safeIdentifier(arg.Name)

				cType := arg.Type
				goType := cToGoType(cType)
				cgoType := cToCgoType(cType)

				if isNextArgumentVariadic {
					expr = "fmt.Sprintf(vfmt, vargs...)"
				}

				switch goType {
				case "string":
					expr = fmt.Sprintf("stringPool.StoreCString(%s)", expr)
				case "mgl32.Vec2":
					expr = fmt.Sprintf("mglVec2ToImVec2(%s)", expr)
				case "mgl32.Vec4":
					expr = fmt.Sprintf("mglVec4ToImVec4(%s)", expr)
				case "*mgl32.Vec2":
					expr = fmt.Sprintf("(%s)(unsafe.Pointer(&%s[0]))", cgoType, expr)
				case "*mgl32.Vec4":
					expr = fmt.Sprintf("(%s)(unsafe.Pointer(&%s[0]))", cgoType, expr)
				default:
					if strings.HasPrefix(goType, "[") {
						expr = fmt.Sprintf("&%s[0]", expr)
					}
					if strings.HasPrefix(cgoType, "*") {
						expr = fmt.Sprintf("unsafe.Pointer(%s)", expr)
					}
					expr = fmt.Sprintf("(%s)(%s)", cgoType, expr)
				}

				output.WriteString(fmt.Sprintf("%s\n", expr))

				if isNextArgumentVariadic {
					break
				}
			}

			output.WriteString("\t")
			if hasReturnType {
				output.WriteString("call := ")
			}

			if hasVariadic {
				output.WriteString(fmt.Sprintf("C.wrap_%s(", def.OverloadedName))
			} else {
				output.WriteString(fmt.Sprintf("C.%s(", def.OverloadedName))
			}

			for i := range def.ArgsT {
				if i > 0 {
					output.WriteString(", ")
				}

				output.WriteString(fmt.Sprintf("a%d", i))

				isNextArgumentVariadic := len(def.ArgsT) > i+1 && def.ArgsT[i+1].Name == "..."
				if isNextArgumentVariadic {
					break
				}
			}

			output.WriteString(")\n")

			if hasReturnType {
				expr := "call"

				goType := cToGoType(def.Ret)

				if goType == "string" {
					expr = fmt.Sprintf("C.GoString(%s)", expr)
				} else if goType == "*mgl32.Vec4" {
					expr = fmt.Sprintf("(*mgl32.Vec4)(unsafe.Pointer(%s))", expr)
				} else {
					expr = fmt.Sprintf("(%s)(%s)", goType, expr)
				}

				output.WriteString(fmt.Sprintf("\tr := %s\n", expr))
				output.WriteString("\treturn r\n")
			}
		}

		output.WriteString("}\n\n")
	}

	err := os.WriteFile("imgui_functions.go", []byte(output.String()), 0644)
	if err != nil {
		panic(err)
	}
}

func wrapperPrototypeForDefinition(def definitionT) string {
	prototype := strings.Builder{}
	prototype.WriteString(fmt.Sprintf("%s wrap_%s(", def.Ret, def.OverloadedName))

	for i, arg := range def.ArgsT {
		if i > 0 {
			prototype.WriteString(", ")
		}

		nextArgIsVariadic := len(def.ArgsT) > i+1 && def.ArgsT[i+1].Name == "..."

		prototype.WriteString(fmt.Sprintf("%s %s", arg.Type, arg.Name))

		if nextArgIsVariadic {
			break
		}
	}

	prototype.WriteString(")")
	return prototype.String()
}

func generateWrappersSources(definitions definitionsT) {
	output := strings.Builder{}
	output.WriteString("#define CIMGUI_DEFINE_ENUMS_AND_STRUCTS 1\n")
	output.WriteString("#include \"dist/include/cimgui/cimgui.h\"\n")
	output.WriteString("\n")

	sortedDefinitions := make([]definitionT, 0, len(definitions))
	for _, definition := range definitions {
		sortedDefinitions = append(sortedDefinitions, definition...)
	}
	sort.Slice(sortedDefinitions, func(i, j int) bool {
		return sortedDefinitions[i].OverloadedName < sortedDefinitions[j].OverloadedName
	})

	for _, def := range sortedDefinitions {
		isVariadic := false
		for _, arg := range def.ArgsT {
			if arg.Name == "..." {
				isVariadic = true
				break
			}
		}

		if !isVariadic {
			continue
		}

		output.WriteString(wrapperPrototypeForDefinition(def))
		output.WriteString(" {\n")
		output.WriteString("\t")

		if def.Ret != "void" && def.Ret != "" {
			output.WriteString("return ")
		}

		output.WriteString(fmt.Sprintf("%s(", def.OverloadedName))
		for i, arg := range def.ArgsT {
			if i > 0 {
				output.WriteString(", ")
			}

			nextArgIsVariadic := len(def.ArgsT) > i+1 && def.ArgsT[i+1].Name == "..."

			output.WriteString(arg.Name)

			if nextArgIsVariadic {
				break
			}
		}

		output.WriteString(");\n")
		output.WriteString("}\n\n")
	}

	err := os.WriteFile("imgui_wrappers.c", []byte(output.String()), 0644)
	if err != nil {
		panic(err)
	}
}

func generateWrappersHeaders(definitions definitionsT) {
	output := strings.Builder{}
	output.WriteString("#define CIMGUI_DEFINE_ENUMS_AND_STRUCTS 1\n")
	output.WriteString("#include \"dist/include/cimgui/cimgui.h\"\n")
	output.WriteString("\n")

	sortedDefinitions := make([]definitionT, 0, len(definitions))
	for _, definition := range definitions {
		sortedDefinitions = append(sortedDefinitions, definition...)
	}
	sort.Slice(sortedDefinitions, func(i, j int) bool {
		return sortedDefinitions[i].OverloadedName < sortedDefinitions[j].OverloadedName
	})

	for _, def := range sortedDefinitions {
		isVariadic := false
		for _, arg := range def.ArgsT {
			if arg.Name == "..." {
				isVariadic = true
				break
			}
		}

		if !isVariadic {
			continue
		}

		output.WriteString(wrapperPrototypeForDefinition(def))
		output.WriteString(";\n")
	}

	err := os.WriteFile("imgui_wrappers.h", []byte(output.String()), 0644)
	if err != nil {
		panic(err)
	}
}
