package route

import (
	"context"
	"github.com/candbright/go-server/internal/spectrum/core"
	"github.com/candbright/go-server/pkg/rest"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"strconv"
)

func init() {
	RegisterRoute(func(engine *gin.Engine) {
		engine.POST("/spectrum/random/get", rest.H(getRandom))
	})
}

func getRandom(c *gin.Context, ctx context.Context) error {
	num := c.Query("num")
	if num == "" {
		return errors.New("num not set")
	}
	numInt, err := strconv.Atoi(num)
	if err != nil {
		return err
	}
	mode := c.Query("mode")
	if mode == "" {
		return errors.New("mode not set")
	}
	modeInt, err := strconv.Atoi(mode)
	if err != nil {
		return err
	}

	var list *core.List[int]
	switch modeInt {
	case 0:
		list = core.RandomBy(core.FourNotesRunMap,
			numInt,
			core.ResetRules(
				core.RuleSameFoot,
				core.RuleReverse,
				core.RuleDiagonal,
				core.RuleNoRepeat,
			))
	case 1:
		numInt = numInt * 2
		list = core.RandomBy(core.TwoNotesMap,
			numInt,
			core.ResetRules(
				core.RuleSameFoot,
				core.RuleReverse,
				core.RuleDiagonal,
			))
	case 2:
		numInt = numInt * 2
		list = core.RandomBy(core.TwoNotesMap,
			numInt,
			core.ResetRules(
				core.RuleSameFoot,
				core.RuleReverse,
			))
	}
	if list == nil {
		return errors.New("list not set")
	}
	arr := list.ToArray()
	return rest.Json(arr)
}
