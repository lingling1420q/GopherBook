package wangyiyun

import (
	"context"
	"fmt"
	"github.com/wuxiaoxiaoshen/GopherBook/chapter7/assistance"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"

	"github.com/antchfx/htmlquery"
)

func WangYiYun(url string) {
	content, err := assistance.Selenium(url)
	if err != nil {
		return
	}
	doc, err := htmlquery.Parse(strings.NewReader(content))
	if err != nil {
		return
	}
	lists := htmlquery.Find(doc, `//div[@id="toplist"]/div/div/ul[@class="f-cb"]/li`)

	for _, i := range lists {
		pName := htmlquery.FindOne(i, `//p[@class="name"]`)
		if pName == nil {
			fmt.Println("nil")
			continue
		}
		href := htmlquery.InnerText(htmlquery.FindOne(pName, "/a/@href"))
		name := htmlquery.InnerText(pName)
		pFrequent := htmlquery.InnerText(htmlquery.FindOne(i, `//p[@class="s-fc4"]`))
		discoverUrl := fmt.Sprintf("%s%s", HOST, href)
		var topList ResultForWangYiYun
		topList = ResultForWangYiYun{
			Title:          name,
			Url:            discoverUrl,
			UpdateFrequent: pFrequent,
		}
		ch := make(chan *ResultForWangYiYun, 1)
		go func() {
			defer close(ch)
			RankDetail(&topList, ch)
		}()
		fmt.Println(<-ch)

	}

}

func RankDetail(result *ResultForWangYiYun, ok chan *ResultForWangYiYun) {
	content, err := assistance.Selenium(result.Url)
	if err != nil {
		log.Println(err)
		return
	}
	doc, err := htmlquery.Parse(strings.NewReader(content))
	if err != nil {
		log.Println(err)
		return
	}
	trs := htmlquery.Find(doc, `//tbody/tr`)
	for _, tr := range trs {
		tds := htmlquery.Find(tr, "/td")
		var rank Rank
		for _, i := range tds {
			title := htmlquery.FindOne(i, `//div[@class="ttc"]//b/@title`)
			if title != nil {
				rank.Title = htmlquery.InnerText(title)
			}
			href := htmlquery.FindOne(i, `//div[@class="ttc"//a/@href]`)
			if href != nil {
				rank.Url = fmt.Sprintf("%s%s", HOST, htmlquery.InnerText(href))
			}
			times := htmlquery.FindOne(i, `//span[@class="u-dur "]`)
			if times != nil {
				rank.Time = htmlquery.InnerText(times)
			}
			text := htmlquery.FindOne(i, `//div[@class="text"]`)
			if text != nil {
				rank.Author = htmlquery.InnerText(htmlquery.FindOne(text, "/@title"))
				if htmlquery.FindOne(text, "//span/a/@href") != nil {
					rank.AuthorUrl = fmt.Sprintf("%s%s", HOST, htmlquery.InnerText(htmlquery.FindOne(text, "//span/a/@href")))
				}
			}
		}
		result.Ranks = append(result.Ranks, rank)
	}
	fav := htmlquery.FindOne(doc, `//a[@id="toplist-fav"]/i`)
	replacer := strings.NewReplacer("(", "", ")", "")
	if fav != nil {
		result.Fav, _ = strconv.Atoi(replacer.Replace(htmlquery.InnerText(fav)))
	}
	share := htmlquery.FindOne(doc, `//a[@id="toplist-share"]/i`)
	if share != nil {
		result.Share, _ = strconv.Atoi(replacer.Replace(htmlquery.InnerText(share)))
	}
	comment := htmlquery.FindOne(doc, `//div[@class="btns f-cb"]//a[4]/i/span`)
	if comment != nil {
		result.Comment, _ = strconv.Atoi(htmlquery.InnerText(comment))
	}
	ok <- result
}

func WangYiYunByChromedp(url string) {
	var a = func(a *chromedp.ExecAllocator) {
		chromedp.Flag("disable-web-security", true)(a)
	}
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		a,
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()
	var response string
	var title string
	chromedp.Run(taskCtx, Tasks(url, &title, &response))
	fmt.Println(title, response)
}
func Tasks(url string, title *string, response *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Title(title),
		chromedp.WaitVisible(`#g_iframe`, chromedp.ByID),
		chromedp.Sleep(2 * time.Second),
		chromedp.Evaluate(`document.getElementById('g_iframe').contentWindow.document.body.outerHTML;`, &response),
	}
}
