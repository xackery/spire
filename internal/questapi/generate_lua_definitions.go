package questapi

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (c *ParseService) generateLuaDefinitions() {
	luaDefinitions = "---@meta\n\n"

	//definition follows the pattern of:
	//method.Name: Description of Name
	//method.Name.Param: Description of Param
	if _, err := os.Stat("definition.txt"); err != nil {
		return
	}
	definitions := make(map[string]string)
	r, err := os.Open("definition.txt")
	if err == nil {
		fs := bufio.NewScanner(r)
		fs.Split(bufio.ScanLines)

		lineNumber := 0
		for fs.Scan() {
			lineNumber++
			line := fs.Text()
			idx := strings.Index(line, ":")
			if idx < 0 {
				continue
			}
			definitions[line[0:idx]] = line[idx+1:]
		}
		r.Close()
	}

	for name, methods := range luaMethods {

		luaDefinitions += fmt.Sprintf("---TODO: definition for %s\n", name)
		inheritence := ""
		if name == "Mob" {
			inheritence = " : Entity"
		}
		if name == "Client" {
			inheritence = " : Mob"
		}
		if name == "NPC" {
			inheritence = " : Mob"
		}

		luaDefinitions += fmt.Sprintf("---@class %s%s\n", name, inheritence)

		luaDefinitions += fmt.Sprintf("%s = {}\n\n", name)

		//luaDefinitions += fmt.Sprintf("---@field public %s # %s Binding\n", name, name)
		for _, method := range methods {

			argBuffer := ""
			paramBuffer := ""
			for _, param := range method.Params {
				if !strings.Contains(param, " ") {
					comment, ok := definitions[fmt.Sprintf("method.%s.%s", name, param)]
					if ok {
						paramBuffer += fmt.Sprintf("---@param %s # TODO: definition of parameter (and type this)\n", param)
					} else {
						paramBuffer += fmt.Sprintf("---@param %s # TODO: definition of parameter (and type this)\n", param)
					}
					argBuffer += fmt.Sprintf("%s, ", param)
					continue
				}
				typeName := param[0:strings.Index(param, " ")]
				switch typeName {
				case "int":
					typeName = "number"
				case "uint32_t":
					typeName = "number"
				case "uint32":
					typeName = "number"
				case "uint16":
					typeName = "number"
				case "int16":
					typeName = "number"
				}
				paramName := param[strings.Index(param, " ")+1:]

				paramBuffer += fmt.Sprintf("---@param %s %s # TODO: definition of parameter\n", paramName, typeName)
				argBuffer += fmt.Sprintf("%s, ", paramName)
			}
			if len(argBuffer) > 0 {
				argBuffer = argBuffer[0 : len(argBuffer)-2]
			}
			luaDefinitions += paramBuffer
			if name == "eq" {
				luaDefinitions += fmt.Sprintf("function %s.%s(%s) end\n\n", name, method.Method, argBuffer)
			} else {
				luaDefinitions += fmt.Sprintf("function %s:%s(%s) end\n\n", name, method.Method, argBuffer)
			}
		}
	}

	for _, event := range luaEvents {
		action := strings.TrimPrefix(event.EventIdentifier, "event_")
		action = strings.ReplaceAll(action, "_", "")
		luaDefinitions += fmt.Sprintf("---@class %sEvent%s\n", event.EntityType, action)

		fieldBuffer := ""

		fieldBuffer += fmt.Sprintf("---@field self %s # TODO: definition of field (and proper typing)\n", event.EntityType)
		for _, field := range event.EventVars {
			if !strings.Contains(field, " ") {
				typeName := "string"
				if strings.HasSuffix(field, "_id") {
					typeName = "number"
				}
				if strings.HasSuffix(field, "_version") {
					typeName = "number"
				}
				if strings.HasSuffix(field, "_level") {
					typeName = "number"
				}
				if field == "other" {
					typeName = "Client"
				}
				fieldBuffer += fmt.Sprintf("---@field %s %s # TODO: definition of field (and proper typing)\n", field, typeName)
				continue
			}
			typeName := field[0:strings.Index(field, " ")]
			switch typeName {
			case "int":
				typeName = "number"
			case "uint32_t":
				typeName = "number"
			case "uint32":
				typeName = "number"
			case "uint16":
				typeName = "number"
			case "int16":
				typeName = "number"
			}
			paramName := field[strings.Index(field, " ")+1:]

			fieldBuffer += fmt.Sprintf("---@field %s %s # TODO: definition of field\n", paramName, typeName)
		}
		luaDefinitions += fieldBuffer + "\n"

		luaDefinitions += fmt.Sprintf("--- %s is a %s event when %s occurs.\n", event.EventIdentifier, event.EntityType, action)
		luaDefinitions += fmt.Sprintf("---@param e %sEvent%s\n", event.EntityType, action)
		luaDefinitions += fmt.Sprintf("function %s(e) end\n\n", event.EventIdentifier)
	}
}
