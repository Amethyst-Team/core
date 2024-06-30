package minecraft

import (
	"bufio"
	"io"
	"log"
	"os/exec"
	"path"
	"sync"

	ws "core-system/logic/websockets"
	"core-system/structs"
	"core-system/utils/system"
)

var cmdStdin io.WriteCloser
var mu sync.Mutex
var enableWebSocket bool

// StartMinecraft starts a Minecraft server process and sets up WebSocket communication.
// It returns an ApiError if the WebSocket route is already enabled.
//
// Parameters:
// w http.ResponseWriter: The response writer for HTTP requests.
// r *http.Request: The HTTP request object.
//
// Return:
// structs.ApiError: An ApiError object containing a message and a status code.
func StartMinecraft() structs.ApiError {
	// currently hardcoded till versions come out
	wsRoute := "/minecraft-ws"

	currentPwd := system.Pwd
	mcServerPath := "server-data/minecraft/server"
	serverInstallationPath := path.Join(currentPwd, mcServerPath)
	mcJar := path.Join(serverInstallationPath, "server.jar")

	command := "java -Xmx4096M -Xms4096M -jar " + mcJar + " nogui"
	cmd := exec.Command("powershell", "-Command", command)
	cmd.Dir = serverInstallationPath

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalf("Failed to get stdout pipe: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalf("Failed to get stderr pipe: %v", err)
	}

	cmdStdin, err = cmd.StdinPipe()
	if err != nil {
		log.Fatalf("Failed to get stdin pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start command: %v", err)
	}

	log.Printf("Java process started with pid %d\n", cmd.Process.Pid)

	go broadcastOutput(bufio.NewScanner(stdout), wsRoute)
	go broadcastOutput(bufio.NewScanner(stderr), wsRoute)

	mu.Lock()
	defer mu.Unlock()

	if !enableWebSocket {
		ws.AddWSRoute(wsRoute, ws.HandleConnections)
		enableWebSocket = true
		log.Println("WebSocket route /ws enabled")

		return structs.ApiError{
			Message: "WebSocket route enabled",
			Code:    200,
		}
	} else {
		return structs.ApiError{
			Message: "WebSocket route already enabled",
			Code:    400,
		}
	}
}

// broadcastOutput reads from the provided bufio.Scanner and broadcasts each line to the WebSocket route.
// It also logs each line to the console.
//
// Parameters:
// scanner *bufio.Scanner: The scanner to read from.
// wsRoute string: The WebSocket route to broadcast messages to.
func broadcastOutput(scanner *bufio.Scanner, wsRoute string) {
	// Loop through each line in the scanner
	for scanner.Scan() {
		// Get the text of the current line
		msg := scanner.Text()

		// Broadcast the message to the WebSocket route
		ws.BroadcastMessage(wsRoute, msg)

		// Log the message to the console
		log.Print(msg)
	}

	// If there was an error reading from the scanner, log it
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading output: %v", err)
	}
}

// SendCommandToMinecraft sends a command to the Minecraft server process.
// It writes the command to the stdin pipe of the Java process, effectively sending it to the server.
//
// Parameters:
// command string: The command to send to the Minecraft server.
//
// Return:
// error: An error if there was an issue writing to the stdin pipe of the Java process.
//
//	Returns nil if the command was successfully written.
func SendCommandToMinecraft(command string) error {
	// Lock the mutex to ensure that only one goroutine can access the cmdStdin pipe at a time.
	mu.Lock()
	defer mu.Unlock()

	// Write the command to the stdin pipe of the Java process.
	// Append a newline character to the command to ensure it is sent as a complete command.
	if _, err := cmdStdin.Write([]byte(command + "\n")); err != nil {
		// If there was an error writing to the pipe, return the error.
		return err
	}

	// If the command was successfully written to the pipe, return nil.
	return nil
}
