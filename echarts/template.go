package echarts

var List = map[string]string{
	"base":     `{{- define "base" }}
<div class="container">
    <div class="item" id="{{ .ChartID }}"
         style="width:{{ .InitOpts.Width }};height:{{ .InitOpts.Height }};"></div>
</div>
<script type="text/javascript">
    "use strict";
    let myChart___x__{{ .ChartID }}__x__ = echarts.init(document.getElementById('{{ .ChartID }}'), "{{ .Theme }}");
    let option___x__{{ .ChartID }}__x__ = {
        title: {{ .TitleOpts  }},
        tooltip: {{ .TooltipOpts }},
        legend: {{ .LegendOpts }},
    {{- if .HasGeo }}
        geo: {{ .GeoComponentOpts }},
    {{- end }}
    {{- if .HasRadar }}
        radar: {{ .RadarComponentOpts }},
    {{- end }}
    {{- if .HasParallel }}
        parallel: {{ .ParallelComponentOpts }},
        parallelAxis: {{ .ParallelAxisOpts }},
    {{- end }}
    {{- if .HasSingleAxis }}
        singleAxis: {{ .SingleAxisOpts }},
    {{- end }}
    {{- if .ToolboxOpts.Show }}
        toolbox: {{ .ToolboxOpts }},
    {{- end }}
    {{- if gt .DataZoomOptsList.Len 0 }}
        dataZoom:{{ .DataZoomOptsList }},
    {{- end }}
    {{- if gt .VisualMapOptsList.Len 0 }}
        visualMap:{{ .VisualMapOptsList }},
    {{- end }}
    {{- if .HasXYAxis }}
        xAxis: {{ .XAxisOptsList }},
        yAxis: {{ .YAxisOptsList }},
    {{- end }}
    {{- if .Has3DAxis }}
        xAxis3D: {{ .XAxis3D }},
        yAxis3D: {{ .YAxis3D }},
        zAxis3D: {{ .ZAxis3D }},
        grid3D: {{ .Grid3D }},
    {{- end }}
        series: [
        {{ range .Series }}
        {{- . }},
        {{ end -}}
        ],
    {{- if eq .Theme "white" }}
        color: {{ .Colors }},
    {{- end }}
    {{- if ne .BackgroundColor "" }}
        backgroundColor: {{ .BackgroundColor }}
    {{- end }}
    };
    myChart___x__{{ .ChartID }}__x__.setOption(option___x__{{ .ChartID }}__x__);

    {{- range .JSFunctions.Fns }}
		{{ . }}
	{{- end }}
</script>
{{ end }}`,
	"chart":    `{{- define "chart" }}
<!DOCTYPE html>
<html>
{{- template "header" . }}
<body>
{{- template "routers" . }}
{{- template "base" . }}
<style>
    .container {margin-top:30px; display: flex;justify-content: center;align-items: center;}
    .item {margin: auto;}
</style>
</body>
</html>
{{ end }}`,
	"header":   `{{ define "header" }}
<head>
    <meta charset="utf-8">
    <title>{{ .PageTitle }}</title>
{{- range .JSAssets.Values }}
    <script src="{{ . }}"></script>
{{- end }}
{{- range .CSSAssets.Values }}
    <link href="{{ . }}" rel="stylesheet">
{{- end }}
</head>
{{ end }}
`,
	"page":     `
{{- define "page" }}
<!DOCTYPE html>
<html>
{{- template "header" . }}
<body>
{{- template "routers" . }}
{{- range .Charts }}
    {{ template "base" . }}
    <br/>
{{- end }}
<style>
    .container {display: flex;justify-content: center;align-items: center;}
    .item {margin: auto;}
</style>
</body>
</html>
{{ end }}`,
	"routes":   `{{- define "routers" }}
<div class="select" style="margin-right:10px; margin-top:10px; position:fixed; right:10px;">
{{- if gt .Routers.Len 0}}
    <select onchange="window.location.href=this.options[this.selectedIndex].value">
    {{- range .Routers }}
        <option value="{{ .URL }}">{{ .Text }}</option>
    {{- end }}
    </select>
{{- end -}}
</div>
{{ end }}`,
	"timeline": `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>

</body>
</html>`,
}
