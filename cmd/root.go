/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/miekg/dns"
	"github.com/spf13/cobra"
)

var (
	query  string
	qtype  string
	qclass string
)

var StringToClass = map[string]uint16{
	"IN":   dns.ClassINET,
	"CS":   dns.ClassCSNET,
	"CH":   dns.ClassCHAOS,
	"HS":   dns.ClassHESIOD,
	"NONE": dns.ClassNONE,
	"ANY":  dns.ClassANY,
}

var StringToType = map[string]uint16{
	"A":    dns.TypeA,
	"AAAA": dns.TypeAAAA,
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dns-message-base64url",
	Short: "Outputs the dns message in base64 url format",
	Long:  `For DoH requests, we have to pass the dns message in base64 url format. This application takes dns parameters and constructs the dns message and outputs the message in base64url format`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("Query:", query)
		//fmt.Println("Qtype", qtype)
		//fmt.Println("Qclass", qclass)
		if query[len(query)-1] != '.' {
			query = query + string('.')
		}
		questions := make([]dns.Question, 0)
		header := &dns.MsgHdr{
			Id:               dns.Id(),
			Response:         false,
			RecursionDesired: true,
		}
		query := &dns.Question{
			Name:   query,
			Qtype:  StringToClass[qclass],
			Qclass: StringToType[qtype],
		}
		questions = append(questions, *query)
		message := &dns.Msg{
			MsgHdr:   *header,
			Question: questions,
		}
		byteArr, err := message.Pack()
		if err != nil {
			fmt.Println("Error while creating the dns message:", err)
			return
		}
		//fmt.Println(byteArr)
		//fmt.Println(message.String())
		//encodedString := hex.EncodeToString(byteArr)
		//fmt.Println("Encoded Hex String: ", encodedString)
		base64_url := Encode(byteArr)
		fmt.Println(base64_url)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dns-message-base64url.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVar(&query, "query", ".", "Domain name of the dns request")
	rootCmd.Flags().StringVar(&qtype, "qtype", "A", "Type of the dns Query(Ex: A, AAAA, etc.)")
	rootCmd.Flags().StringVar(&qclass, "qclass", "IN", "Class of the dns query(Ex: IN, CH etc.)")
}
