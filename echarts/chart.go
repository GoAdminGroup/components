package echarts

import (
	"bytes"
	"encoding/json"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/datasets"
	"github.com/go-echarts/go-echarts/templates"
	"html/template"
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

	buf := new(bytes.Buffer)

	if line, ok := c.Content.(*charts.Line); ok {
		_ = line.Render(buf)
	}

	if bar, ok := c.Content.(*charts.Bar); ok {
		_ = bar.Render(buf)
	}

	if bar3d, ok := c.Content.(*charts.Bar3D); ok {
		_ = bar3d.Render(buf)
	}

	if box, ok := c.Content.(*charts.BoxPlot); ok {
		_ = box.Render(buf)
	}

	if chart3d, ok := c.Content.(*charts.Chart3D); ok {
		_ = chart3d.Render(buf)
	}

	if rc, ok := c.Content.(*charts.RectChart); ok {
		_ = rc.Render(buf)
	}

	if ef, ok := c.Content.(*charts.EffectScatter); ok {
		_ = ef.Render(buf)
	}

	if fu, ok := c.Content.(*charts.Funnel); ok {
		_ = fu.Render(buf)
	}

	if ga, ok := c.Content.(*charts.Gauge); ok {
		_ = ga.Render(buf)
	}

	if geo, ok := c.Content.(*charts.Geo); ok {
		_ = geo.Render(buf)
	}

	if graph, ok := c.Content.(*charts.Graph); ok {
		_ = graph.Render(buf)
	}

	if hm, ok := c.Content.(*charts.HeatMap); ok {
		_ = hm.Render(buf)
	}

	if kl, ok := c.Content.(*charts.Kline); ok {
		_ = kl.Render(buf)
	}

	if line3d, ok := c.Content.(*charts.Line3D); ok {
		_ = line3d.Render(buf)
	}

	if li, ok := c.Content.(*charts.Liquid); ok {
		_ = li.Render(buf)
	}

	if ma, ok := c.Content.(*charts.Map); ok {
		_ = ma.Render(buf)
	}

	if pa, ok := c.Content.(*charts.Parallel); ok {
		_ = pa.Render(buf)
	}

	if pi, ok := c.Content.(*charts.Pie); ok {
		_ = pi.Render(buf)
	}

	if ra, ok := c.Content.(*charts.Radar); ok {
		_ = ra.Render(buf)
	}

	if sk, ok := c.Content.(*charts.Sankey); ok {
		_ = sk.Render(buf)
	}

	if sc, ok := c.Content.(*charts.Scatter); ok {
		_ = sc.Render(buf)
	}

	if sc3d, ok := c.Content.(*charts.Scatter3D); ok {
		_ = sc3d.Render(buf)
	}

	if su, ok := c.Content.(*charts.Surface3D); ok {
		_ = su.Render(buf)
	}

	if tr, ok := c.Content.(*charts.ThemeRiver); ok {
		_ = tr.Render(buf)
	}

	if wc, ok := c.Content.(*charts.WordCloud); ok {
		_ = wc.Render(buf)
	}

	content := buf.String()

	content = strings.Replace(content, `<meta charset="utf-8">`, "", -1)
	content = strings.Replace(content, `<title>Awesome go-echarts</title>`, "", -1)
	content = strings.Replace(content, `<script src="https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"></script>`, "", -1)
	content = strings.Replace(content, `<script src="https://go-echarts.github.io/go-echarts-assets/assets/echarts-gl.min.js"></script>`, "", -1)
	content = strings.Replace(content, `<link href="https://go-echarts.github.io/go-echarts-assets/assets/bulma.min.css" rel="stylesheet">`, "", -1)
	content = strings.Replace(content, `<div class="select" style="margin-right:10px; margin-top:10px; position:fixed; right:10px;"></div>`, "", - 1)
	content = strings.Replace(content, `container`, "echarts-container", - 1)

	return template.HTML(content)
}
