package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/iancoleman/strcase"
)

func copyFile(src, dst string) error {
	fmt.Printf("Copying file %s to %s\n", src, dst)
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

	return nil
}

func extractGlfwConstants() []string {
	out := []string{}

	content, err := os.ReadFile("thirdparty/glfw/include/GLFW/glfw3.h")
	if err != nil {
		panic(err)
	}

	renameConstant := func(name string) string {
		name = strcase.ToCamel(name)
		name = strings.ReplaceAll(name, "Opengl", "OpenGL")
		if name == "OpenGLForwardCompat" {
			return "OpenGLForwardCompatible"
		}
		return name
	}

	knownConstants := map[string]string{}

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#define GLFW_") {
			line = strings.TrimSpace(line)
			line = strings.ReplaceAll(line, "\t", " ")
			for {
				before := line
				line = strings.ReplaceAll(line, "  ", " ")
				if before == line {
					break
				}
			}

			parts := strings.Split(line, " ")
			_, err := strconv.Atoi(parts[2])
			if err == nil || strings.HasPrefix(parts[2], "0x") {
				knownConstants[parts[1]] = parts[2]
				// out = append(out, fmt.Sprintf("const %s = %s", renameConstant(strings.TrimPrefix(parts[1], "GLFW_")), parts[2]))
				out = append(out, fmt.Sprintf("const %s = %s", renameConstant(strings.TrimPrefix(parts[1], "GLFW_")), "C."+parts[1]))
			} else if knownConstants[parts[2]] != "" {
				// out = append(out, fmt.Sprintf("const %s = %s", renameConstant(strings.TrimPrefix(parts[1], "GLFW_")), knownConstants[parts[2]]))
				out = append(out, fmt.Sprintf("const %s = %s", renameConstant(strings.TrimPrefix(parts[1], "GLFW_")), "C."+parts[1]))
			}
		}
	}

	return out
}

func generateHeaders() {
	copyFile("thirdparty/glfw/include/GLFW/glfw3.h", "dist/include/GLFW/glfw3.h")
	copyFile("thirdparty/glfw/include/GLFW/glfw3native.h", "dist/include/GLFW/glfw3native.h")
	copyFile("thirdparty/glad/include/glad.h", "dist/include/glad/glad.h")
	copyFile("thirdparty/cimgui/cimgui.h", "dist/include/cimgui/cimgui.h")
	copyFile("thirdparty/cimgui/generator/output/cimgui_impl.h", "dist/include/cimgui/cimgui_impl.h")
}

func generateGlfwConstants() {
	preamble := "package glfw\n"
	preamble += "\n"
	preamble += "// #cgo CFLAGS: -I../dist/include\n"
	preamble += "// #include \"GLFW/glfw3.h\"\n"
	preamble += "import \"C\"\n"
	preamble += "\n"

	glfwConstants := extractGlfwConstants()
	err := os.WriteFile("glfw/constants.go", []byte(preamble+strings.Join(glfwConstants, "\n")), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Written %d constants to glfw/key_constants.go\n", len(glfwConstants))
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
	err := os.WriteFile("imgui/constants.go", []byte(preamble+strings.Join(imguiConstants, "\n")), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Written %d constants to imgui/constants.go\n", len(imguiConstants))
}

func main() {
	generateHeaders()
	generateGlfwConstants()
	generateImguiConstants()
}
