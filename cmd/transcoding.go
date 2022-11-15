package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os/exec"
	"strconv"
)

var transcodingCmd = &cobra.Command{
	Use:   "transcoding",
	Short: "transcoding",
	Long:  "transcoding",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("transcoding -- haha")
		execCmd := exec.Command("ffprobe", "./307478996-1-208.mp4", "-show_streams", "-select_streams", "v", "-print_format", "json")
		var stdout, stderr bytes.Buffer
		execCmd.Stdout = &stdout
		execCmd.Stderr = &stderr
		err := execCmd.Run()
		outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
		if err != nil {
			log.Fatalf("execCmd.Run() failed with %s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(transcodingCmd)
}

type probeFormat struct {
	Duration string `json:"duration"`
}

type probeData struct {
	Format probeFormat `json:"format"`
}

func probeDuration(a string) (float64, error) {
	pd := probeData{}
	err := json.Unmarshal([]byte(a), &pd)
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseFloat(pd.Format.Duration, 64)
	if err != nil {
		return 0, err
	}
	return f, nil
}
