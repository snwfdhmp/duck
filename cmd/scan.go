package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/karrick/godirwalk"
	"github.com/snwfdhmp/duck/pkg/data"
	"github.com/snwfdhmp/duck/pkg/projects"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	addMode    bool
	doctorMode bool
	repairMode bool
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			color.Red("Please give a path.")
			return
		}

		lookupPath, err := filepath.Abs(args[0])
		if err != nil {
			color.Red("Cannot process path '" + args[0] + "': " + err.Error())
			return
		}

		//load projects from data
		dataHasError := false

		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 1, 2, 1, ' ', 0)

		addedCount := 0
		projectCount := 0

		dataPath := filepath.Dir(data.Path)

		if repairMode {
			doctorMode = true
		}

		err = godirwalk.Walk(lookupPath, &godirwalk.Options{
			Callback: func(path string, de *godirwalk.Dirent) error {
				if de.Name() != ".duck" || !de.IsDir() || path == dataPath {
					return nil
				}
				fmt.Printf("Scanning for duck projects from " + lookupPath + " : " + strconv.Itoa(projectCount) + " projects found\r")
				path = filepath.Dir(path)
				projectCount += 1
				name := filepath.Base(path)
				found := false

				output := ""

				if !dataHasError {
					for _, p := range data.Projects.Keys() {
						if p.Value() == path {
							found = true
							if p.Name() == name {
								output += Green("\tknown")
							} else {
								output += Magenta("\tknown under '" + p.Name() + "'")
							}
							break
						}
					}
					if !found {
						output += Red("\tunkwown")
						if addMode {
							err := data.AddProject(name, path, false, false)
							if err != nil {
								output += Red(", cannot add: " + err.Error())
							} else {
								output = Yellow("\tadded")
								addedCount += 1
							}
						}
					}
					if doctorMode {
						healty, err := projects.IsHealthy(path)
						var doctor string
						if err != nil {
							doctor = Red("\tcannot doctor: " + err.Error())
						} else if !healty {
							doctor = Magenta("\tsick")
							if repairMode {
								RepairProject(path)
								doctor = Yellow("\trepaired")
							}
						} else {
							doctor = Green("\thealthy")
						}
						output += doctor
					}
				}
				fmt.Fprintln(w, "- "+Yellow(name)+"\t"+Blue(path)+output)
				return nil
			},
		})

		if err != nil {
			color.Red("Could not scan projects: " + err.Error())
		}

		//CLEAR
		fmt.Printf("\n")
		w.Flush()

		msg := "Found " + Yellow(strconv.Itoa(projectCount)) + " projects"
		if addedCount > 0 {
			msg += " and added " + Yellow(strconv.Itoa(addedCount)) + " of them."
		}

		fmt.Println(msg)
	},
}

func ScanProjects(from string) ([]string, error) {
	var paths []string
	return paths, afero.Walk(fs, from, func(path string, info os.FileInfo, err error) error {
		if info.Name() == ".duck" && info.IsDir() {
			paths = append(paths, filepath.Dir(path))
		}
		return nil
	})
}

func init() {
	projectsCmd.AddCommand(scanCmd)
	scanCmd.Flags().BoolVarP(&addMode, "add", "a", false, "add missing projects")
	scanCmd.Flags().BoolVarP(&doctorMode, "doctor", "d", false, "pass each project to doctor")
	scanCmd.Flags().BoolVarP(&repairMode, "repair", "r", false, "repair each project judged not healthy by doctor")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// scanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// scanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
