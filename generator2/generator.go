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

	blacklist := []string{"const_iterator", "iterator", "value_type"}

	for _, name := range sortedNames {
		if slices.Contains(blacklist, name) {
			continue
		}
		renamedName := strings.ReplaceAll(name, "ImGui", "")

		// Special case for Im* types.
		if strings.HasPrefix(renamedName, "Im") {
			r, _ := utf8.DecodeRuneInString(renamedName[2:])
			if unicode.IsUpper(r) {
				renamedName = renamedName[2:]
			}
		}

		typedefsContent.WriteString(fmt.Sprintf("type %s C.%s\n", renamedName, name))
	}

	err := os.WriteFile("imgui_typedefs.go", []byte(typedefsContent.String()), 0644)
	if err != nil {
		panic(err)
	}
}

func generateDefinitions(definitions definitionsT) {
	output := strings.Builder{}
	output.WriteString("package imgui\n")
	output.WriteString("\n")
	output.WriteString("// #include \"cimgui/cimgui.h\"\n")
	output.WriteString("import \"C\"\n")
	output.WriteString("import \"unsafe\"\n")
	output.WriteString("\n")

	for _, definition := range definitions {
		for _, def := range definition {
			if strings.HasPrefix(def.OverloadedName, "ig") {
				fmt.Println("=>", def.OverloadedName)
			}
		}
	}
}
