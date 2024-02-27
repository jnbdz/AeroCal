package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var weightBalanceCmd = &cobra.Command{
	Use:   "weight-balance",
	Short: "Calculate and display weight and balance information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Weight and Balance Information:")
		// Access the weight balance configuration using Viper
		var weightBalances []struct {
			Aircraft string  `mapstructure:"aircraft"`
			Weight   float64 `mapstructure:"weight"`
			Arm      float64 `mapstructure:"arm"`
			Moment   float64 `mapstructure:"moment"`
		}

		// Unmarshal the weight balance config into the struct
		err := viper.UnmarshalKey("weight_balance", &weightBalances)
		if err != nil {
			fmt.Printf("Unable to decode into struct, %v", err)
		}

		// Process and display the data as needed
		for _, wb := range weightBalances {
			fmt.Printf("Aircraft: %s, Weight: %.2f, Arm: %.2f, Moment: %.2f\n",
				wb.Aircraft, wb.Weight, wb.Arm, wb.Moment)
		}
	},
}

func AddWeightBalanceCmd(root *cobra.Command) {
	root.AddCommand(weightBalanceCmd)
}
