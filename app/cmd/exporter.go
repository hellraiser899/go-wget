var rootCmd = &cobra.Command{
    Use:   "gowget",
    Short: "gowget is a very fast cli tool to download stuff from web-services ",
    Long: `A Fast and Flexible `,
    Run: func(cmd *cobra.Command, args []string) {
      //TODO :go-wget
      //1 retrieve data from url 
      //2 support recursive directory
    },
  }
  
  func Execute() {
    if err := rootCmd.Execute(); err != nil {
      fmt.Fprintln(os.Stderr, err)
      os.Exit(1)
    }
  }
