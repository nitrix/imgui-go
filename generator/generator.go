package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	cparser "github.com/nitrix/imgui-go/generator/cparser"
)

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

func generateHeaders() {
	copyFile("thirdparty/cimgui/cimgui.h", "dist/include/cimgui/cimgui.h")
	copyFile("thirdparty/cimgui/generator/output/cimgui_impl.h", "dist/include/cimgui/cimgui_impl.h")
}

func extractImguiConstants() []string {
	out := []string{}

	content, err := os.ReadFile("thirdparty/cimgui/cimgui.h")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	within := false
	num := 0

	for _, line := range lines {
		line := strings.TrimSpace(line)

		if !within && line == "typedef enum {" {
			within = true
			num = 0
		} else if within && strings.HasPrefix(line, "}") {
			within = false
		} else if within {
			line = strings.TrimSuffix(line, ",")
			line = strings.ReplaceAll(line, "(int)", "")
			line = strings.ReplaceAll(line, "ImGui", "")
			parts := strings.Split(line, "=")
			if len(parts) == 2 {
				k := parts[0]
				v := parts[1]
				k = strings.TrimSpace(k)
				v = strings.TrimSpace(v)
				v = strings.ReplaceAll(v, "~", "^")
				out = append(out, fmt.Sprintf("const %s = %s", k, v))
			} else {
				out = append(out, fmt.Sprintf("const %s = %d", parts[0], num))
				num++
			}
		}
	}

	return out
}

func generateImguiConstants() {
	preamble := "package imgui\n"
	preamble += "\n"

	imguiConstants := extractImguiConstants()
	err := os.WriteFile("constants.go", []byte(preamble+strings.Join(imguiConstants, "\n")), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Copied %d constants to constants.go\n", len(imguiConstants))
}

func convertGoTypeToCgoType(goType string) string {
	switch goType {
	case "int":
		return "C.int"
	case "bool":
		return "C.bool"
	case "*bool":
		return "*C.bool"
	case "float32":
		return "C.float"
	case "uint32":
		return "C.ImU32"
	case "unsafe.Pointer":
		return "*C.void"
	}

	if strings.HasPrefix(goType, "C.Im") || strings.HasPrefix(goType, "*C.Im") {
		return goType
	}

	panic("Unknown Go type to convert to Cgo type: " + goType)
}

func convertCTypeToGoType(cType string) string {
	cType = strings.TrimPrefix(cType, "const ")

	switch cType {
	case "int":
		return "int"
	case "float":
		return "float32"
	case "bool":
		return "bool"
	case "bool *":
		return "*bool"
	case "void *":
		return "unsafe.Pointer"
	case "char *":
		return "string"
	case "ImU32":
		return "uint32"
	case "ImVec2":
		return "mgl32.Vec2"
	case "ImVec3":
		return "mgl32.Vec3"
	case "ImVec4":
		return "mgl32.Vec4"
	}

	if strings.HasPrefix(cType, "ImGui") || strings.HasPrefix(cType, "ImDraw") || strings.HasPrefix(cType, "ImFont") {
		if strings.HasSuffix(cType, "*") {
			return "*C." + strings.TrimSpace(strings.TrimSuffix(cType, "*"))
		} else {
			return "C." + strings.TrimSpace(cType)
		}
	}

	panic("Unknown C type to convert to Go type: " + cType)
}

func generateWrapper(prototype cparser.Prototype) (string, string) {
	wrapperDeclParams := []string{}
	for j, param := range prototype.Parameters {
		if j < len(prototype.Parameters)-2 {
			wrapperDeclParams = append(wrapperDeclParams, fmt.Sprintf("%s %s", param.Ty, param.Name))
		} else {
			wrapperDeclParams = append(wrapperDeclParams, "const char *fmt")
			break
		}
	}

	wrapperDecl := fmt.Sprintf("void wrap_%s(%s)", prototype.Name, strings.Join(wrapperDeclParams, ", "))

	wrapperDef := strings.Builder{}
	wrapperDef.WriteString(fmt.Sprintf("void wrap_%s(", prototype.Name))
	for j, param := range prototype.Parameters {
		if j > 0 {
			wrapperDef.WriteString(", ")
		}

		if j < len(prototype.Parameters)-2 {
			wrapperDef.WriteString(fmt.Sprintf("%s %s", param.Ty, param.Name))
		} else {
			wrapperDef.WriteString("const char *fmt")
			break
		}
	}

	wrapperDef.WriteString(") {\n")

	if prototype.ReturnType != "void" {
		wrapperDef.WriteString("\treturn ")
	} else {
		wrapperDef.WriteString("\t")
	}

	wrapperDef.WriteString(fmt.Sprintf("%s(", prototype.Name))
	for j, param := range prototype.Parameters {
		if j > 0 {
			wrapperDef.WriteString(", ")
		}

		if j < len(prototype.Parameters)-2 {
			wrapperDef.WriteString(param.Name)
		} else {
			wrapperDef.WriteString("fmt")
			break
		}
	}

	wrapperDef.WriteString(");\n")
	wrapperDef.WriteString("}\n")

	return wrapperDecl, wrapperDef.String()
}

func generateImguiFunctions() {
	content, err := os.ReadFile("thirdparty/cimgui/cimgui.h")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")

	output := strings.Builder{}

	wrapperDecls := []string{}
	wrapperDefs := []string{}

	funcCount := 0

	for _, line := range lines {
		if strings.HasPrefix(line, "CIMGUI_API") {
			prototype := cparser.ParsePrototype(line)

			blacklist := []string{
				"igNewFrame",
			}

			whitelist := []string{
				"igShowDemoWindow",
				"igDockSpaceOverViewport",
				"igUpdatePlatformWindows",
				"igRenderPlatformWindowsDefault",
				"igGetDrawData",
				"igRender",
				"igCreateContext",
				"igDestroyContext",
				"igText",
				"igButton",
				"igCheckbox",
				"igBegin",
				"igEnd",
			}

			if slices.Contains(blacklist, prototype.Name) {
				continue
			}

			if slices.Contains(whitelist, prototype.Name) {
				// ================== Generate the function name ==================
				goName := strings.TrimPrefix(prototype.Name, "ig")
				output.WriteString(fmt.Sprintf("func %s(", goName))

				variadic := false

				// ================== Generate the parameters ==================
				for i, param := range prototype.Parameters {
					if i > 0 {
						output.WriteString(", ")
					}

					// Handle void prototype.
					if len(prototype.Parameters) == 1 && param.Name == "" {
						break
					}

					// Special case for variadic functions.
					if param.Ty == "const char *" && i < len(prototype.Parameters)-1 && prototype.Parameters[i+1].Ty == "..." {
						output.WriteString("fmt string, args ...interface{}")
						variadic = true

						wrapperDecl, wrapperDef := generateWrapper(prototype)
						wrapperDecls = append(wrapperDecls, wrapperDecl)
						wrapperDefs = append(wrapperDefs, wrapperDef)

						break
					}

					output.WriteString(param.Name)
					output.WriteString(" ")
					output.WriteString(convertCTypeToGoType(param.Ty))
				}

				output.WriteString(") ")

				// ================== Generate the return type ==================
				if prototype.ReturnType != "void" {
					output.WriteString(convertCTypeToGoType(prototype.ReturnType))
					output.WriteString(" ")
				}

				output.WriteString("{\n")

				// ================== Generate the function call ==================
				if variadic {
					output.WriteString("\ts := stringPool.StoreCString(gofmt.Sprintf(fmt, args...))\n")
				}

				output.WriteString("\t")

				parensReturn := false

				if prototype.ReturnType != "void" {
					output.WriteString("return ")
					goType := convertCTypeToGoType(prototype.ReturnType)
					cgoType := convertGoTypeToCgoType(goType)
					if goType != cgoType {
						output.WriteString(fmt.Sprintf("(%s)(", goType))
						parensReturn = true
					}
				}

				if variadic {
					output.WriteString(fmt.Sprintf("C.wrap_%s(", prototype.Name))
				} else {
					output.WriteString(fmt.Sprintf("C.%s(", prototype.Name))
				}

				for i, param := range prototype.Parameters {
					if i > 0 {
						output.WriteString(", ")
					}

					if len(prototype.Parameters) == 1 && param.Name == "" {
						break
					}

					if variadic && i == len(prototype.Parameters)-2 {
						output.WriteString("s")
						break
					}

					if param.Ty == "const char *" {
						output.WriteString(fmt.Sprintf("stringPool.StoreCString(%s)", param.Name))
						continue
					}

					if param.Ty == "ImVec2" || param.Ty == "const ImVec2" {
						output.WriteString(fmt.Sprintf("C.ImVec2{C.float(%s.X()), C.float(%s.Y())}", param.Name, param.Name))
						continue
					}

					if param.Ty == "ImVec3" || param.Ty == "const ImVec3" {
						output.WriteString(fmt.Sprintf("C.ImVec3{C.float(%s.X()), C.float(%s.Y()), C.float(%s.Z())}", param.Name, param.Name, param.Name))
						continue
					}

					if param.Ty == "ImVec4" || param.Ty == "const ImVec4" {
						output.WriteString(fmt.Sprintf("C.ImVec4{C.float(%s.X()), C.float(%s.Y()), C.float(%s.Z()), C.float(%s.W())}", param.Name, param.Name, param.Name, param.Name))
						continue
					}

					cgoType := convertGoTypeToCgoType(convertCTypeToGoType(param.Ty))
					if cgoType == "*C.void" {
						output.WriteString(param.Name)
					} else if cgoType == convertCTypeToGoType(param.Ty) {
						output.WriteString(param.Name)
					} else {
						output.WriteString(fmt.Sprintf("(%s)(%s)", cgoType, param.Name))
					}
				}
				output.WriteString(")")

				if parensReturn {
					output.WriteString(")")
				}

				output.WriteString("\n")
				output.WriteString("}\n\n")

				funcCount++
			}
		}
	}

	preamble := strings.Builder{}
	preamble.WriteString("package imgui\n")
	preamble.WriteString("\n")
	preamble.WriteString("// #include \"cimgui/cimgui.h\"\n")

	for _, wrapperDecl := range wrapperDecls {
		preamble.WriteString(fmt.Sprintf("// %s;\n", wrapperDecl))
	}

	preamble.WriteString("import \"C\"\n")
	preamble.WriteString("import gofmt \"fmt\"\n")
	preamble.WriteString("import \"github.com/go-gl/mathgl/mgl32\"\n")
	preamble.WriteString("import \"unsafe\"\n")
	preamble.WriteString("\n")

	final := strings.Builder{}
	final.WriteString(preamble.String())
	final.WriteString(output.String())

	err = os.WriteFile("imgui_functions.go", []byte(final.String()), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Generated %d imgui functions\n", funcCount)

	secondary := strings.Builder{}
	secondary.WriteString("#define CIMGUI_DEFINE_ENUMS_AND_STRUCTS 1\n")
	secondary.WriteString("#include \"cimgui/cimgui.h\"\n")
	secondary.WriteString("\n")

	for i, wrapperDef := range wrapperDefs {
		if i > 0 {
			secondary.WriteString("\n")
		}
		secondary.WriteString(wrapperDef)
	}

	err = os.WriteFile("imgui_wrappers.c", []byte(secondary.String()), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Generated %d imgui variadic wrappers\n", len(wrapperDefs))
}

func main() {
	generateHeaders()
	generateImguiConstants()
	generateImguiFunctions()
}
