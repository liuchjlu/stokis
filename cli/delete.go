package cli

import ()

func DeleteJob() {
	commandname := "kubectl"
	params := []string{"delete", "-f", YamlPath}
	if err := cmd.ExecCommand(commandname, SysConfig.LogAddress, params); err != nil {
		log.Fatalf("cli.CreateTrain(): %+v\n", err)
	}
	log.Infoln("The Current job has bean deleted.")

}
