package minecraft

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/gorilla/websocket"

	"core-system/utils/system"
)

var (
	logger    *log.Logger
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan string)
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// StartMinecraft starts the Minecraft server and sets up WebSocket broadcasting.
func StartMinecraft(w http.ResponseWriter, r *http.Request) {
	// Setup paths
	currentPwd := system.Pwd
	mcServerPath := "minecraft/server" // Adjust as per your server path
	serverInstallationPath := path.Join(currentPwd, mcServerPath)
	logsDir := path.Join(currentPwd, "logs/minecraft-server", time.Now().Format("20060102_150405"))
	os.MkdirAll(logsDir, os.ModePerm)
	logFilePath := path.Join(logsDir, "server.log")

	// Create log file
	logFile, err := os.Create(logFilePath)
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	defer logFile.Close()

	// Setup logger
	logger = log.New(logFile, "", log.LstdFlags)

	// Create necessary folders and files
	system.CreateFolder(serverInstallationPath)
	logger.Println("Installing Minecraft server...")
	eulaTxtPath := path.Join(serverInstallationPath, "eula.txt")
	system.CreateFile(eulaTxtPath, "eula=true")
	mcJar := path.Join(serverInstallationPath, "server.jar")

	// Start the Minecraft server
	command := "cd " + serverInstallationPath + "; java -Xmx1024M -Xms1024M -jar " + mcJar + " nogui"
	cmd := exec.Command("powershell", "-Command", command)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		logger.Fatalf("Failed to get stdout pipe: %v", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		logger.Fatalf("Failed to get stderr pipe: %v", err)
	}
	if err := cmd.Start(); err != nil {
		logger.Fatalf("Failed to start command: %v", err)
	}
	logger.Printf("Java process started with pid %d\n", cmd.Process.Pid)

	// Start goroutines to handle output
	go broadcastOutput(bufio.NewScanner(stdout))
	go broadcastOutput(bufio.NewScanner(stderr))

	// Dynamically register WebSocket handler
	http.HandleFunc("/ws", MinecraftHandler)

	// Return success response
	w.Write([]byte("Minecraft server started successfully"))
}

// MinecraftHandler handles WebSocket connections for Minecraft logs
func MinecraftHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			delete(clients, ws)
			break
		}
	}
}

// broadcastOutput reads scanner input and broadcasts to clients
func broadcastOutput(scanner *bufio.Scanner) {
	for scanner.Scan() {
		msg := scanner.Text()
		broadcast <- msg
		logger.Println(msg)
	}
	if err := scanner.Err(); err != nil {
		logger.Printf("Error reading output: %v", err)
	}
}
