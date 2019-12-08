package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/FengGeSe/nest/util"
	"github.com/spf13/cobra"
)

// flags
type initOptions struct {
	Path  string `flag:"path" shorthand:"p" default:"${PWD}" desc:"路径"`
	Force bool   `flag:"force" shorthand:"f" desc:"删除同名目录并创建"`
}

func (opts initOptions) IsValidated() error {
	if opts.Path == "" {
		return fmt.Errorf("路径不能为空")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)

	if err := util.SetFlagsByStruct(initCmd.Flags(), initOptions{}); err != nil {
		panic(err)
	}
}

type Project struct {
	Name string
}

var initCmd = &cobra.Command{
	Use:   "init [name]",
	Short: "Initialize a go-kit application",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 || args[0] == "" {
			return fmt.Errorf("项目名称不能为空")
		}
		opts := &initOptions{}
		if err := util.SetValuesFromFlags(cmd.Flags(), opts); err != nil {
			return err
		}
		return opts.IsValidated()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		// 获得绝对路径
		opts := &initOptions{}
		util.SetValuesFromFlags(cmd.Flags(), opts)
		if !filepath.IsAbs(opts.Path) {
			if absPath, err := filepath.Abs(opts.Path); err == nil {
				opts.Path = absPath
			} else {
				return nil
			}
		}
		opts.Path = filepath.Join(opts.Path, name)

		// 项目信息
		project := Project{
			Name: name,
		}

		if err := InitProject(opts.Path, project, opts.Force); err != nil {
			return err
		}

		cmd.Printf("project(%s) created in %s \n", project.Name, opts.Path)
		return nil
	},
}

// 在absPath目录下初始化项目结构
func InitProject(absPath string, project Project, force bool) error {
	// 目录存在报错 不存在创建
	if isExist, err := PathExists(absPath); err != nil {
		return err
	} else if isExist && !force {
		return fmt.Errorf("path(%s) is exist", absPath)
	} else if isExist && force {
		if err := os.RemoveAll(absPath); err != nil {
			return err
		}
	}

	if err := os.MkdirAll(absPath, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
