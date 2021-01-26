/*
Copyright © 2021 Diederik van den Burger

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"generatorik/templates"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
)

func outputDir(flagValue string) string {
	var output string
	path, _ := os.Getwd()

	if flagValue != "" {
		output = flagValue
	} else {
		output = path
	}

	return output
}

// componentCmd represents the component command
var componentCmd = &cobra.Command{
	Use:   "component [componentname]",
	Short: "Generates a React component",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		outputFlag, _ := cmd.Flags().GetString("directory")
		outputDirectory := outputDir(outputFlag)

		var componentName string
		if len(args) == 0 {
			path, _ := os.Getwd()
			split := strings.Split(path, "/")
			componentName = split[len(split)-1]
		} else {
			componentName = args[0]
		}

		_, err := os.Stat(outputDirectory)
		if os.IsNotExist(err) {
			fmt.Println("Error: output directory does not exist")
			return
		}

		if noTest, _ := cmd.Flags().GetBool("no-test"); noTest == false {
			templates.WriteTest(templates.TestProps{
				ComponentName: componentName,
			}, path.Join(outputDirectory, strings.ToLower(componentName)+".spec.tsx"))
		}

		if noStory, _ := cmd.Flags().GetBool("no-story"); noStory == false {
			templates.WriteStory(templates.StoryProps{
				ComponentName: componentName,
			}, path.Join(outputDirectory, strings.ToLower(componentName)+".stories.mdx"))
		}

		noStyle, _ := cmd.Flags().GetBool("no-style")
		if noStyle == false {
			templates.WriteStyles(templates.StylesProps{
				ComponentName: componentName,
			}, path.Join(outputDirectory, strings.ToLower(componentName)+".styles.tsx"))
		}

		if noIndex, _ := cmd.Flags().GetBool("no-index"); noIndex == false {
			templates.WriteIndex(templates.IndexProps{
				ComponentName: componentName,
			}, path.Join(outputDirectory, "index.ts"))
		}

		hasProps, _ := cmd.Flags().GetBool("has-props")
		var componentProps = templates.ComponentProps{
			ComponentName: componentName,
			HasProps:      hasProps,
			HasStyles:     !noStyle,
		}

		templates.WriteComponent(
			componentProps,
			path.Join(outputDirectory, componentName+".tsx"),
		)
	},
}

func init() {
	rootCmd.AddCommand(componentCmd)

	componentCmd.Flags().StringP("directory", "d", "", "Output directory for the generated files")
	componentCmd.Flags().BoolP("has-props", "p", false, "Generate an interface for the component")

	componentCmd.Flags().Bool("no-test", false, "Do not generate a test file")
	componentCmd.Flags().Bool("no-story", false, "Do not generate a storybook file")
	componentCmd.Flags().Bool("no-style", false, "Do not generate a styles file")
	componentCmd.Flags().Bool("no-index", false, "Do not generate an index file")
}
