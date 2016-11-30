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
		evaluate	[localyaml path]		start the service of result-evaluate.
		results 	[localyaml path]		get the results to local.
		delete	 	[localyaml path]		destroy the current job.
		logs		[localyaml path]		get logs from standard out.

		help
	`
	fmt.Println(helpstring)
}
