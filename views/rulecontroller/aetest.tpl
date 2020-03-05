
<div class="row rule-title">
    <div class="col-md-1">ID</div>
    <div class="col-md-1">结算方式</div>
    <div class="col-md-1">平台</div>
    <div class="col-md-3">模板名称</div>
    <div class="col-md-2">媒体</div>
    <div class="col-md-4">配置</div>
</div>
{{range .rules}}
<div class="row">
    <div class="col-md-1">{{.Id}}</div>
    <div class="col-md-1">{{.Bidmode}}</div>
    <div class="col-md-1">{{.Platform}}</div>
    <div class="col-md-3">{{.Name}}</div>
    <div class="col-md-2">{{.Media}}</div>
    <div class="col-md-4"></div>
</div>
{{end}}

{{template "../components/pagination.tpl" .}}