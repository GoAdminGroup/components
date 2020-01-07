package echarts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/datasets"
	"github.com/go-echarts/go-echarts/templates"
	"html/template"
	"regexp"
	"strings"
)

type Chart struct {
	Content interface{}
}

func NewChart() *Chart {

	templates.BaseTpl = List["base"]
	templates.ChartTpl = List["chart"]
	templates.HeaderTpl = List["header"]
	templates.PageTpl = List["page"]
	templates.RoutersTpl = List["routes"]

	_ = json.Unmarshal([]byte(coorJson), &datasets.Coordinates)
	_ = json.Unmarshal([]byte(mapfileName), &datasets.MapFileNames)

	return new(Chart)
}

func (c *Chart) SetContent(content interface{}) *Chart {
	c.Content = content
	return c
}

func (c *Chart) GetTemplate() (*template.Template, string) {
	return nil, "echarts"
}

func (c *Chart) GetAssetList() []string {
	return AssetsList
}

func (c *Chart) GetAsset(name string) ([]byte, error) {
	return Asset(name[1:])
}

func (c *Chart) IsAPage() bool {
	return false
}

func (c *Chart) GetName() string {
	return "echarts"
}

func (c *Chart) GetContent() template.HTML {

	buf, chartId := c.getContent()

	content := buf.String()

	content = strings.Replace(content, `<meta charset="utf-8">`, "", -1)
	content = strings.Replace(content, `<title>Awesome go-echarts</title>`, "", -1)
	content = strings.Replace(content, `<script src="https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"></script>`, "", -1)
	content = strings.Replace(content, `<script src="https://go-echarts.github.io/go-echarts-assets/assets/maps/china.js"></script>`, "", -1)
	content = strings.Replace(content, `<script src="https://go-echarts.github.io/go-echarts-assets/assets/echarts-gl.min.js"></script>`, "", -1)
	content = strings.Replace(content, `<script src="https://go-echarts.github.io/go-echarts-assets/assets/echarts-liquidfill.min.js"></script>`, "", -1)
	content = strings.Replace(content, `<link href="https://go-echarts.github.io/go-echarts-assets/assets/bulma.min.css" rel="stylesheet">`, "", -1)
	content = strings.Replace(content, `<div class="select" style="margin-right:10px; margin-top:10px; position:fixed; right:10px;"></div>`, "", - 1)
	content = strings.Replace(content, `container`, "echarts-container", - 1)

	return template.HTML(content) + template.HTML(fmt.Sprintf(resizeJS, chartId))
}

func (c *Chart) GetOptions() template.JS {

	buf, _ := c.getContent()

	reg, _ := regexp.Compile(`{([\s\S]*)};`)

	match := reg.FindString(buf.String())

	return template.JS(match[:len(match)-1])
}

func (c *Chart) getContent() (*bytes.Buffer, string) {
	buf := new(bytes.Buffer)
	chartId := ""

	if line, ok := c.Content.(*charts.Line); ok {
		chartId = line.ChartID
		_ = line.Render(buf)
	}

	if bar, ok := c.Content.(*charts.Bar); ok {
		chartId = bar.ChartID
		_ = bar.Render(buf)
	}

	if bar3d, ok := c.Content.(*charts.Bar3D); ok {
		chartId = bar3d.ChartID
		_ = bar3d.Render(buf)
	}

	if box, ok := c.Content.(*charts.BoxPlot); ok {
		chartId = box.ChartID
		_ = box.Render(buf)
	}

	if chart3d, ok := c.Content.(*charts.Chart3D); ok {
		chartId = chart3d.ChartID
		_ = chart3d.Render(buf)
	}

	if rc, ok := c.Content.(*charts.RectChart); ok {
		chartId = rc.ChartID
		_ = rc.Render(buf)
	}

	if ef, ok := c.Content.(*charts.EffectScatter); ok {
		chartId = ef.ChartID
		_ = ef.Render(buf)
	}

	if fu, ok := c.Content.(*charts.Funnel); ok {
		chartId = fu.ChartID
		_ = fu.Render(buf)
	}

	if ga, ok := c.Content.(*charts.Gauge); ok {
		chartId = ga.ChartID
		_ = ga.Render(buf)
	}

	if geo, ok := c.Content.(*charts.Geo); ok {
		chartId = geo.ChartID
		_ = geo.Render(buf)
	}

	if graph, ok := c.Content.(*charts.Graph); ok {
		chartId = graph.ChartID
		_ = graph.Render(buf)
	}

	if hm, ok := c.Content.(*charts.HeatMap); ok {
		chartId = hm.ChartID
		_ = hm.Render(buf)
	}

	if kl, ok := c.Content.(*charts.Kline); ok {
		chartId = kl.ChartID
		_ = kl.Render(buf)
	}

	if line3d, ok := c.Content.(*charts.Line3D); ok {
		chartId = line3d.ChartID
		_ = line3d.Render(buf)
	}

	if li, ok := c.Content.(*charts.Liquid); ok {
		chartId = li.ChartID
		_ = li.Render(buf)
	}

	if ma, ok := c.Content.(*charts.Map); ok {
		chartId = ma.ChartID
		_ = ma.Render(buf)
	}

	if pa, ok := c.Content.(*charts.Parallel); ok {
		chartId = pa.ChartID
		_ = pa.Render(buf)
	}

	if pi, ok := c.Content.(*charts.Pie); ok {
		chartId = pi.ChartID
		_ = pi.Render(buf)
	}

	if ra, ok := c.Content.(*charts.Radar); ok {
		chartId = ra.ChartID
		_ = ra.Render(buf)
	}

	if sk, ok := c.Content.(*charts.Sankey); ok {
		chartId = sk.ChartID
		_ = sk.Render(buf)
	}

	if sc, ok := c.Content.(*charts.Scatter); ok {
		chartId = sc.ChartID
		_ = sc.Render(buf)
	}

	if sc3d, ok := c.Content.(*charts.Scatter3D); ok {
		chartId = sc3d.ChartID
		_ = sc3d.Render(buf)
	}

	if su, ok := c.Content.(*charts.Surface3D); ok {
		chartId = su.ChartID
		_ = su.Render(buf)
	}

	if tr, ok := c.Content.(*charts.ThemeRiver); ok {
		chartId = tr.ChartID
		_ = tr.Render(buf)
	}

	if wc, ok := c.Content.(*charts.WordCloud); ok {
		chartId = wc.ChartID
		_ = wc.Render(buf)
	}

	return buf, chartId
}

const resizeJS = `<script>
	window.addEventListener('resize',function () {
		myChart_%s.resize();
	})
</script>`
