package windows

import (
	s "core-system/utils/system"
	"os"
	"strings"
	"syscall"

	"golang.org/x/sys/windows"
)

// runAppAsElevated function is used to run the current application with elevated privileges.
// It uses the ShellExecute function from the windows package to achieve this.
// If the function fails to execute, it will print the error to the console.
func RunAppAsElevated() {
	// Define the verb for the ShellExecute function.
	// "runas" verb is used to run the application with elevated privileges.
	verb := "runas"

	// Get the path of the current executable.
	exe, _ := os.Executable()

	// Get the current working directory.
	cwd, _ := os.Getwd()

	// Join the arguments passed to the application into a single string.
	args := strings.Join(os.Args[1:], " ")

	// Convert the verb, executable path, working directory, and arguments to UTF-16 pointers.
	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	// Define the show command for the ShellExecute function.
	// SW_NORMAL is used to display the application window in its normal state.
	var showCmd int32 = 1

	// Call the ShellExecute function to run the application with elevated privileges.
	// If an error occurs, print it to the console.
	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		s.Logger.Println(err)
	}
}

// isAppAdmin checks if the current process is running with administrative privileges.
// It uses the IsElevated method from the windows package to determine this.
//
// The function returns a boolean value indicating whether the process is elevated (true) or not (false).
//
// Note: This function is specific to the Windows operating system.
func IsAppAdmin() bool {
	// Call the IsElevated method on the current process token to check if it is elevated.
	elevated := windows.GetCurrentProcessToken().IsElevated()

	// Return the result of the IsElevated method.
	return elevated
}
