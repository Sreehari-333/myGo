package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/stream", func(c *gin.Context) {
		filePath := "videos/sample1.mp4"

		file, err := os.Open(filePath)

		if err != nil {
			c.String(http.StatusInternalServerError, "could not open file")
			return
		}
		defer file.Close()

		fileInfo, _ := file.Stat()
		fileSize := fileInfo.Size()

		rangeHeader := c.GetHeader("Range")

		if rangeHeader == "" {
			c.Header("Content-Type", "video/mp4")
			c.Header("Content-Length", fmt.Sprintf("%d", fileSize))
			http.ServeContent(c.Writer, c.Request, filePath, fileInfo.ModTime(), file)
			return
		}

		rangeHeader = strings.TrimPrefix(rangeHeader, "bytes=")
		parts := strings.Split(rangeHeader, "-")

		start, _ := strconv.ParseInt(parts[0], 10, 64)

		var end int64
		if parts[1] != "" {
			end, _ = strconv.ParseInt(parts[1], 10, 64)
		} else {
			end = fileSize - 1
		}

		if end >= fileSize {
			end = fileSize - 1
		}

		contentLength := end - start + 1

		c.Header("Content-Type", "video/mp4")
		c.Header("Content-Length", fmt.Sprintf("%d", contentLength))
		c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
		c.Status(http.StatusPartialContent)

		file.Seek(start, 0)
		buf := make([]byte, contentLength)

		_, err = file.Read(buf)

		if err != nil {
			c.String(http.StatusInternalServerError, "could not read file chunk")
			return
		}
		c.Writer.Write(buf)

	})

	r.Run(":8080")

}
