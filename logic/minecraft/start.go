package minecraft

import (
	java "core-system/logic/java"
	s "core-system/utils/system"
)

func StartMinecraft() {
	s.Logger.Printf("Starting minecraft instance...")

	_ = java.IsInstalled()
}
