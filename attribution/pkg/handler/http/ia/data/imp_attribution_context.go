package data

import (
	"strconv"

	"github.com/TencentAd/attribution/attribution/pkg/data"
	"github.com/TencentAd/attribution/attribution/pkg/protocal/parse"
)

type ImpAttributionContext struct {
	*data.BaseContext

	ConvParseResult *parse.ConvParseResult
	CampaignId      int64 // 推广计划id
	CampaignIdStr   string

	EncryptData      []*IdSet // 广告主加密
	EncryptTwiceData []*IdSet // ams二次加密
	FirstDecryptData []*IdSet // 二次加密后，第一次解密
	IntersectData    []*IdSet // 归因交集
	FinalDecryptData []*IdSet // 明文数据

	OriginalIndex map[string][]int //原始用户ID => 下标
}

func NewImpAttributionContext(pr *parse.ConvParseResult) (*ImpAttributionContext, error) {
	campaignId := pr.CampaignId

	c := &ImpAttributionContext{
		BaseContext:     data.NewBaseContext(),
		ConvParseResult: pr,
		CampaignId:      campaignId,
		CampaignIdStr:   strconv.FormatInt(campaignId, 10),
	}

	c.buildOriginalIndex()
	return c, nil
}

func (c *ImpAttributionContext) GroupId() string {
	return c.CampaignIdStr
}

func (c *ImpAttributionContext) buildOriginalIndex() {
	c.OriginalIndex = make(map[string][]int)

	for i, convLog := range c.ConvParseResult.ConvLogs {
		userData := convLog.UserData
		c.OriginalIndex[userData.Imei] = append(c.OriginalIndex[userData.Imei], i)
		c.OriginalIndex[userData.Idfa] = append(c.OriginalIndex[userData.Idfa], i)
		c.OriginalIndex[userData.AndroidId] = append(c.OriginalIndex[userData.AndroidId], i)
		c.OriginalIndex[userData.Oaid] = append(c.OriginalIndex[userData.Oaid], i)
	}
}
