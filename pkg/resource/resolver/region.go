package resolver

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetRegions(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegions) (data *spec.GetRegionsData, code uint8, err error) {
	fmt.Println("DEBUG GetRegions")
	data = &spec.GetRegionsData{
		Regions: []spec.GetRegionData{
			spec.GetRegionData{
				Name: "hoge",
			},
		},
	}
	return
}
