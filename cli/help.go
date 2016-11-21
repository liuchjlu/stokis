package cli

import (
	"fmt"
)

func help() {
	var helpstring = `
	Usage: manage COMMAND [args...]
	Version: 0.1
	Author:liuchjlu
	Email:liucaihong@iie.ac.cn

	Comands:
		train		[localyaml path]		start the job of model-train.
		test		[localyaml path]		start the job of model-test.
		evaluate	[localyaml path]		start the service of result-evaluate
		getresult	[localyaml path]		get the results to local
		destroy 	[localyaml path]		destroy the current job

		help
	`
	fmt.Println(helpstring)
}
