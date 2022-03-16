package meilele

import (
	"regexp"

	"github.com/nekoimi/fetcher/engine"
	"github.com/nekoimi/fetcher/fetcher"
)

const goodTitleRegex = `<title>[^<]*【.*】(.*?)[a-zA-Z0-9-]*【价格 图片 尺寸 品牌】<\/title>`
const goodPriceRegex = `<strong id="JS_effect_price" class="gnum">([0-9]+)\.00<\/strong>`
const goodBannersRegex = `<img src="\/\/img[0-9]+\.mllres\.com/[^>]*\.[jpg|png|jpeg]+"[^>]*width="86"[^>]*height="57"[^>]*data\-show_img="(\/\/image\.meilele\.com\/[^>]*\.[jpg|png|jpeg]+)"[^>]*data\-zoom_img="\/\/image\.meilele\.com\/[^>]*\.[jpg|png|jpeg]+"\/>`
const goodBrandRegex = `<td>商品品牌[^<]*<a href="\/brand\-[0-9]+\/" target="_blank" class="red">(.*?)<\/a><\/td>`
const goodDetailImagesRegex = `<div[^>]*class="img"[^>]*data\-source="handledpictures"[^>]*style="margin\-bottom:10px;">[^<]*<img src="\/\/image\.meilele\.com\/images\/blank\.gif" data\-src="(\/\/image\.meilele\.com\/images\/[0-9]+/[0-9]+\.jpg)"[^>]*\/><\/div>`

// 商品详情解析器
func ParseGoodDetails(content []byte) engine.ParseResult {
	html := fetcher.RegexpReplaceEmpty(content, clearHtmlRegex)
	goods := engine.Item{}
	// Title
	titleRe := regexp.MustCompile(goodTitleRegex)
	titleRes := titleRe.FindStringSubmatch(html)
	if len(titleRes) >= 2 {
		goods.Title = titleRes[1]
	}
	// // Price
	// priceRe := regexp.MustCompile(goodPriceRegex)
	// priceRes := priceRe.FindStringSubmatch(html)
	// if len(priceRes) >= 2 {
	// 	goods.Price, _ = strconv.ParseFloat(priceRes[1], 64)
	// }
	// // Banners
	// bannersRe := regexp.MustCompile(goodBannersRegex)
	// bannersRes := bannersRe.FindStringSubmatch(html)
	// if len(bannersRes) >= 2 {
	// 	goods.Banners = []string{bannersRes[1]}
	// }
	// // Brand
	// brandRe := regexp.MustCompile(goodBrandRegex)
	// brandRes := brandRe.FindStringSubmatch(html)
	// if len(brandRes) >= 2 {
	// 	goods.Brand = brandRes[1]
	// }
	// Details
	detailRe := regexp.MustCompile(goodDetailImagesRegex)
	detailsRes := detailRe.FindAllStringSubmatch(html, -1)
	var images []string
	for _, image := range detailsRes {
		images = append(images, image[1])
	}
	goods.Details = images

	result := engine.ParseResult{}
	result.Items = append(result.Items, goods)
	return result
}
