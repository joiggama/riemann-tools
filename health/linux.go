package health

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/joiggama/riemann-tools/utils"
)

var CPU_USAGE = make([]int64, 4)

func CPU() int32 {
	regex, _ := regexp.Compile(`cpu (\s+(\d+)){4}`)

	file := utils.ReadFile("/proc/stat")
	data := strings.Trim(regex.FindAllString(file, -1)[0], "cpu  ")

	cpu_usage := make([]int64, 4)
	var usage int32

	for index, row := range strings.Split(data, " ") {
		cpu_usage[index], _ = strconv.ParseInt(row, 10, 64)
	}

	u2, n2, s2, i2 := cpu_usage[0], cpu_usage[1], cpu_usage[2], cpu_usage[3]

	if utils.ZeroedCollection(CPU_USAGE) == false {
		u1, n1, s1, i1 := CPU_USAGE[0], CPU_USAGE[1], CPU_USAGE[2], CPU_USAGE[3]

		used := (u2 + n2 + s2) - (u1 + n1 + s1)
		total := used + i2 - i1
		usage = int32((float64(used) / float64(total)) * 100)
	}

	CPU_USAGE = cpu_usage

	return usage
}

func Memory() int32 {
	data := strings.Trim(utils.ReadFile("/proc/meminfo"), "\n\n")

	regex, _ := regexp.Compile(`:?\s+`)

	mem_info := map[string]int64{}

	for _, row := range strings.Split(data, "\n") {
		metric := regex.Split(row, -1)
		key, value := metric[0], metric[1]
		mem_info[key], _ = strconv.ParseInt(value, 10, 64)
	}

	free := float64(mem_info["MemFree"] + mem_info["Buffers"] + mem_info["Cached"])
	used := int32((free / float64(mem_info["MemTotal"])) * 100)

	return used
}
