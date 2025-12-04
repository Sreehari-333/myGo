var highest = struct {
	sync.RWMutex
	data map[int]int // itemID -> highest bid amount
}{data: make(map[int]int)}


var rooms = struct {
	sync.RWMutex
	data map[int][]*websocket.Conn
}{data: make(map[int][]*websocket.Conn)}


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func wsHandler(c *gin.Context) {
	itemId, _ := strconv.Atoi(c.Param("itemId"))
	conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)

	rooms.Lock()
	rooms.data[itemId] = append(rooms.data[itemId], conn)
	rooms.Unlock()

	// Keep connection open
	for {
		if _, _, err := conn.ReadMessage(); err != nil {
			conn.Close()
			return
		}
	}
}


type BidRequest struct {
	UserID int `json:"userId"`
	Amount int `json:"amount"`
}

func placeBid(c *gin.Context) {
	itemId, _ := strconv.Atoi(c.Param("id"))

	var body BidRequest
	c.BindJSON(&body)

	// check highest bid
	highest.Lock()
	current := highest.data[itemId]

	if body.Amount <= current {
		highest.Unlock()
		c.JSON(400, gin.H{"error": "Bid must be higher"})
		return
	}

	// update highest bid
	highest.data[itemId] = body.Amount
	highest.Unlock()

	// send update to all websocket clients
	broadcast(itemId, body.Amount)

	c.JSON(200, gin.H{
		"status":     "Bid accepted",
		"new_highest": body.Amount,
	})
}


func broadcast(itemId int, amount int) {
	rooms.RLock()
	connections := rooms.data[itemId]
	rooms.RUnlock()

	msg, _ := json.Marshal(gin.H{
		"type":   "new_bid",
		"amount": amount,
	})

	for _, conn := range connections {
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func main() {
	r := gin.Default()

	r.GET("/ws/:itemId", wsHandler)
	r.POST("/items/:id/bid", placeBid)

	r.Run(":8080")
}



