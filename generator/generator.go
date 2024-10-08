package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type EnumT struct {
	Name      string `json:"name"`
	Value     string `json:"value"`
	CalcValue int    `json:"calc_value"`
}

type structsAndEnumsT struct {
	Enums map[string][]EnumT `json:"enums"`
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

	generateConstants(&structsAndEnums)
	generateTypedefs(typedefsDict)
	generateDefinitions(definitions)
	generateWrappersHeaders(definitions)
	generateWrappersSources(definitions)
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
	typedefsContent.WriteString("// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -Idist/include\n")
	typedefsContent.WriteString("// #include \"cimgui/cimgui.h\"\n")
	typedefsContent.WriteString("import \"C\"\n")
	typedefsContent.WriteString("\n")

	sortedNames := make([]string, 0, len(typedefsDict))
	for name := range typedefsDict {
		sortedNames = append(sortedNames, name)
	}
	sort.Strings(sortedNames)

	blacklist := []string{
		"const_iterator", "iterator", "value_type", "STB_TexteditState",
		"ImVec1", "ImVec2", "ImVec2ih", "ImVec4", // These are replaced by mgl32.Vec2 and mgl32.Vec4.
	}

	for _, name := range sortedNames {
		if slices.Contains(blacklist, name) {
			continue
		}

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
	case "char*":
		return "*C.char"
	case "bool*":
		return "*C.bool"
	case "void*":
		return "unsafe.Pointer"
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
		return "[3]float32"
	case "float[4]":
		return "[4]float32"
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

	// FIXME: Not sure about these.
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
	}

	switch cType {
	case "ImVec2":
		return "mgl32.Vec2"
	case "ImVec4":
		return "mgl32.Vec4"
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
	switch s {
	case "type":
		return "_type"
	}
	return s
}

func generateDefinitions(definitions definitionsT) {
	output := strings.Builder{}
	output.WriteString("package imgui\n")
	output.WriteString("\n")
	output.WriteString("// #cgo CFLAGS: -DCIMGUI_DEFINE_ENUMS_AND_STRUCTS=1 -Idist/include\n")
	output.WriteString("// #include \"cimgui/cimgui.h\"\n")
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
		"igNewFrame",                   // Needed to be overridden to control the pool memory allocator.
		"igGetAllocatorFunctions",      // FIXME: Wants a void**, why?
		"igImFormatStringToTempBuffer", // FIXME: Wants a char**, why?
		"igImTextStrFromUtf8",          // FIXME: Wants a char**, why?
		"igAddDrawListToDrawDataEx",    // FIXME: Wants a ImVector_ImDrawListPtr*, why?
		"igDockBuilderCopyDockSpace",   // FIXME: Wants a Vector_const_charPtr*, why?
		"igDockBuilderCopyNode",        // FIXME: Wants a Vector_ImGuiID*, why?
	}

	whitelist := []string{
		"igCreateContext",
		"igDestroyContext",
		"igRender",
		"igGetDrawData",
		"igShowDemoWindow",
		"igBegin",
		"igEnd",
		"igText",
		"igButton",
		"igCheckbox",
		"igDockSpaceOverViewport",
		"igUpdatePlatformWindows",
		"igRenderPlatformWindowsDefault",
		"igText",
	}

	for _, def := range sortedDefinitions {
		if slices.Contains(blacklist, def.OverloadedName) {
			continue
		}

		if !strings.HasPrefix(def.OverloadedName, "ig") {
			continue
		}

		if !slices.Contains(whitelist, def.OverloadedName) {
			continue
		}

		// fmt.Println("=>", def.OverloadedName)

		renamedName := strings.ReplaceAll(def.OverloadedName, "ig", "")
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

		// FIXME: Skip function pointer parameters for now.
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
					output.WriteString("_fmt string, _args ...interface{}")
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
			if hasVariadic {
				output.WriteString("\t_vfmt := fmt.Sprintf(_fmt, _args...)\n")
			}

			if hasReturnType {
				output.WriteString(fmt.Sprintf("\treturn (%s)(", cToGoType(def.Ret)))
			} else {
				output.WriteString("\t")
			}

			if hasVariadic {
				output.WriteString(fmt.Sprintf("C.wrap_%s(", def.OverloadedName))
			} else {
				output.WriteString(fmt.Sprintf("C.%s(", def.OverloadedName))
			}

			for i, arg := range def.ArgsT {
				if i > 0 {
					output.WriteString(", ")
				}

				isString := cToGoType(arg.Type) == "string"
				isVec := strings.HasPrefix(cToGoType(arg.Type), "mgl32.")
				nextArgIsVariadic := len(def.ArgsT) > i+1 && def.ArgsT[i+1].Name == "..."

				output.WriteString(fmt.Sprintf("(%s)(", cToCgoType(arg.Type)))

				renamedVar := safeIdentifier(arg.Name)
				if nextArgIsVariadic {
					renamedVar = "_vfmt"
				}

				if isString {
					output.WriteString(fmt.Sprintf("stringPool.StoreCString(%s)", renamedVar))
				} else if isVec {
					n := strings.TrimPrefix(cToGoType(arg.Type), "mgl32.")
					output.WriteString(fmt.Sprintf("mgl%sToIm%s(%s)", n, n, renamedVar))
				} else {
					output.WriteString(renamedVar)
				}

				output.WriteString(")")

				if nextArgIsVariadic {
					break
				}
			}

			if hasReturnType {
				output.WriteString(")")
			}

			output.WriteString(")\n")
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
	output.WriteString("#include \"cimgui/cimgui.h\"\n")
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
	output.WriteString("#include \"cimgui/cimgui.h\"\n\n")

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
