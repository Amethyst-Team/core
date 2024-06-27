package minecraft

import (
	s "core-system/utils/system"
)

const mcUrl = "http://piston-data.mojang.com/v1/objects/450698d1863ab5180c25d7c804ef0fe6369dd1ba/server.jar"

func InstallMinecraft() (err error) {
	_, err = s.DownloadFile("minecraft", mcUrl, false)

	return err
}
