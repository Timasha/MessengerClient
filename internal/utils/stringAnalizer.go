package utils

func FindArguments(startNum int, str string) (args []string) {
	var argNum int = 0
	args = append(args, "")
	for i := startNum + 1; i < len([]byte(str)); i++ {
		if []byte(str)[i] == 32 {
			args = append(args, "")
			argNum++
			continue
		}
		args[argNum] += string([]byte(str)[i])
	}
	return args
}
