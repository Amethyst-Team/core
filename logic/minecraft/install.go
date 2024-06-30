package minecraft

import (
	s "core-system/utils/system"
)

const mcUrl = "http://piston-data.mojang.com/v1/objects/450698d1863ab5180c25d7c804ef0fe6369dd1ba/server.jar"
const mcServerPath = "server-data\\minecraft\\server"

func InstallMinecraft() (err error) {
	filePath, err := s.DownloadFile("minecraft", mcUrl, false)

	if err != nil {
		panic(err)
	}

	s.CreateFolder(mcServerPath)
	s.MoveFile(filePath, mcServerPath)

	return err
}
