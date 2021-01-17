package cmd

import (
	"eoc/questapi"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"time"
)

type QuestApiParseCommand struct {
	logger  *logrus.Logger
	command *cobra.Command
	parser  *questapi.ParseService
}

func (g *QuestApiParseCommand) Command() *cobra.Command {
	return g.command
}

// new instance of command
func NewQuestApiParseCommand(
	logger *logrus.Logger,
	parser *questapi.ParseService,
) *QuestApiParseCommand {
	i := &QuestApiParseCommand{
		parser: parser,
		logger: logger,
		command: &cobra.Command{
			Use:   "quest:parse",
			Short: "Parses EQEmu/Server Quest API source files for documentation",
		},
	}

	i.command.Args = i.Validate
	i.command.Run = i.Handle

	return i
}

// Handle implementation of the Command interface
func (c *QuestApiParseCommand) Handle(_ *cobra.Command, args []string) {
	start := time.Now()

	methods := c.parser.Parse(true)

	fmt.Printf("# Perl\n")
	for methodType, methods := range methods.PerlApi {
		for _, method := range methods {
			returnType := ""
			if method.ReturnType != "" {
				returnType = fmt.Sprintf(" # %v", method.ReturnType)
			}

			fmt.Printf("%v::%v;%v\n", methodType, method.Method, returnType)
		}
	}

	fmt.Printf("\n\n")

	fmt.Printf("# Lua\n")
	for methodType, methods := range methods.LuaApi {
		for _, method := range methods {
			returnType := ""
			if method.ReturnType != "" {
				returnType = fmt.Sprintf(" -- %v", method.ReturnType)
			}

			fmt.Printf("%v::%v%v\n", methodType, method.Method, returnType)
		}
	}


	fmt.Printf("took %v\n", time.Since(start))
}

// Validate
func (g *QuestApiParseCommand) Validate(_ *cobra.Command, args []string) error {
	// Validate

	return nil
}
