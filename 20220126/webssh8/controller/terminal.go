package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"webssh/core"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// TermWs 获取终端ws
func TermWs(c *gin.Context, timeout time.Duration) *ResponseBody {

	fmt.Println("============= c := ", c)
	// FIRST get basic info

	responseBody := ResponseBody{Msg: "success"}
	defer TimeCost(time.Now(), &responseBody)
	sshInfo := c.DefaultQuery("sshInfo", "")
	cols := c.DefaultQuery("cols", "150")
	rows := c.DefaultQuery("rows", "35")
	col, _ := strconv.Atoi(cols)
	row, _ := strconv.Atoi(rows)
	sshClient, err := core.DecodedMsgToSSHClient(sshInfo)
	fmt.Println("**************** =>", sshClient)
	if err != nil {
		fmt.Println(err)
		responseBody.Msg = err.Error()
		return &responseBody
	}

	// Upgrade to websocket
	wsConn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		responseBody.Msg = err.Error()
		return &responseBody
	}

	// GenerateClient
	err = sshClient.GenerateClient()
	if err != nil {
		wsConn.WriteMessage(1, []byte(err.Error()))
		wsConn.Close()
		fmt.Println(err)
		responseBody.Msg = err.Error()
		return &responseBody
	}

	// InitTerminal
	sshClient.InitTerminal(wsConn, row, col)

	// Connect
	sshClient.Connect(wsConn, timeout)

	return &responseBody
}
