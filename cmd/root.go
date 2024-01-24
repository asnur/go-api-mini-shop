package cmd

import (
	"fmt"
	"go-api-mini-shop/config"

	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize(initConfig)
}

var (
	rootCmd = &cobra.Command{
		Use:   "go_mini_shop",
		Short: "Go Mini Shop",
		Long:  `Go Mini Shop is a mini shop application built with Go Fiber and GORM`,
		Run: func(cmd *cobra.Command, args []string) {
			// Figlet Banner
			ascci := figlet4go.NewAsciiRender()
			banner, _ := ascci.Render("Go Mini Shop")
			fmt.Println(banner)
		},
	}
)

func Execute() error {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&ip, "ip", "i", "127.0.0.1", "IP Address")
	serverCmd.Flags().StringVarP(&port, "port", "p", "3000", "Port")

	return rootCmd.Execute()
}

func initConfig() {
	if err := config.InitializeConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
