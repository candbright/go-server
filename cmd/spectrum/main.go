package main

import (
	"github.com/candbright/go-server/internal/spectrum"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strconv"
)

func main() {
	engine := gin.New()
	engine.Use(cors.Default())
	engine.GET("/spectrum/random/get", func(c *gin.Context) {
		num := c.Query("num")
		if num == "" {
			c.JSON(400, errors.New("num not set"))
			return
		}
		numInt, err := strconv.Atoi(num)
		if err != nil {
			c.JSON(400, err)
			return
		}
		list := spectrum.RandomBy(spectrum.FourNotesRunMap,
			numInt,
			spectrum.ResetRules(
				spectrum.RuleSameFoot,
				spectrum.RuleReverse,
				spectrum.RuleDiagonal,
				spectrum.RuleNoRepeat,
			))
		arr := list.ToArray()
		c.JSON(200, arr)
	})
	_ = engine.Run(":18001")
}
