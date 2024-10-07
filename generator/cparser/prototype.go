package parser

import (
	"strings"
)

type Parameter struct {
	Ty   string
	Name string
}

type Prototype struct {
	Name       string
	ReturnType string
	Parameters []Parameter
}

func formatType(ty string) string {
	ty = strings.ReplaceAll(ty, "*", " *")
	ty = strings.ReplaceAll(ty, "  ", " ")
	ty = strings.ReplaceAll(ty, "* *", "**")
	return ty
}

func ParsePrototype(line string) Prototype {
	// Extract the function name.
	parts := strings.Split(line, "(")
	parts = strings.Split(parts[0], " ")
	cName := parts[len(parts)-1]

	// Extract the return type.
	parts = strings.Split(line, "(")
	cReturnType := strings.TrimSuffix(parts[0], cName)
	cReturnType = strings.TrimPrefix(cReturnType, "CIMGUI_API")
	cReturnType = strings.TrimSpace(cReturnType)

	// Extract the parameters.
	parts = strings.Split(line, "(")
	parts = strings.Split(parts[1], ")")
	cParametersTxt := strings.Split(parts[0], ",")

	cParameters := make([]Parameter, len(cParametersTxt))
	for i, cParameterTxt := range cParametersTxt {
		cParameterTxt = strings.TrimSpace(cParameterTxt)
		parts = strings.Split(cParameterTxt, " ")
		cParameters[i].Name = parts[len(parts)-1]
		cParameters[i].Ty = strings.TrimSpace(strings.TrimSuffix(cParameterTxt, cParameters[i].Name))
	}

	// Patch nameless parameters.
	for i, cParameter := range cParameters {
		if cParameter.Ty == "" {
			cParameters[i].Ty = cParameter.Name
			cParameters[i].Name = ""
		}
	}

	// Format types.
	cReturnType = formatType(cReturnType)
	for i, cParameter := range cParameters {
		cParameters[i].Ty = formatType(cParameter.Ty)
	}

	return Prototype{
		Name:       cName,
		ReturnType: cReturnType,
		Parameters: cParameters,
	}
}
