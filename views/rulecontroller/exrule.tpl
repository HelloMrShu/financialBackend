<div class="select-form">
    <form action="{{.url}}" method="get" id="ruleForm">
        <div class="form-group search-item">
            模板名称：<input type="text" value="{{.sname}}" name="sname" placeholder="请输入模板名称" />
        </div>
        <div class="form-group search-item">
            模板KEY：<input type="text" value="{{.stype}}" name="stype" placeholder="请输入模板KEY" />
        </div>
        <div class="form-group search-item">
            <input type="submit" value="查询" />
        </div>
    </form>
</div>

<div class="row rule-title">
    <div class="col-md-1">ID</div>
    <div class="col-md-1">结算方式</div>
    <div class="col-md-1">平台</div>
    <div class="col-md-1">媒体</div>
    <div class="col-md-4">模板名称</div>
    <div class="col-md-3">模板KEY</div>
    <div class="col-md-1">配置信息</div>
</div>

{{range .rules}}
<div class="row rule-line">
    <div class="col-md-1">{{.Id}}</div>
    <div class="col-md-1">不限</div>
    <div class="col-md-1">不限</div>
    <div class="col-md-1">不限</div>
    <div id="rn-{{.Id}}" class="col-md-4">{{.Name}}</div>
    <div class="col-md-3">{{.Special_type}}</div>
    <div class="col-md-1">
        <input id="rc-{{.Id}}" type="hidden" value="{{.Content}}" />
        <span class="expand" onclick="expandexchange('{{.Id}}')">[展开]</span>
    </div>
</div>
{{end}}

<div class="row">
    <div class="paginator">
        <div class="page-item">
            <a href="/ae/test?page=1">首页</a>
        </div>
        <div class="page-item">
            <a href="/ae/test?page={{.paginator.PrePage}}">上一页</a>
        </div>
        <div class="page-item">
            {{range $k, $v := .paginator.Ranges}}
            <a href='/ae/test?page={{$v}}' {{if eq $.paginator.Page $v}} class="active" {{end}}>    
                <span>{{ $v }}</span>
            </a>
            {{end}}
        </div>
        <div class="page-item">
            <a href='/ae/test?page={{.paginator.NextPage}}'>下一页</a>
        </div>
        <div class="page-item">
            <a href="/ae/test?page={{.paginator.TotalPages}}">尾页</a>
        </div>
        <div class="page-total">共&nbsp;{{.paginator.TotalPages}}&nbsp;页</div>
    </div>
</div>

<div class="modal fade" id="panel" tabindex="-1" role="dialog" aria-labelledby="exampleModalScrollableTitle" aria-hidden="true">
    <div class="modal-dialog modal-lg modal-dialog modal-dialog-scrollable" role="document">
        <div class="modal-content">
            <div class="modal-header">
                <h5 class="modal-title" id="exampleModalScrollableTitle"></h5>
                <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                <span aria-hidden="true">&times;</span>
                </button>
            </div>
            <pre>
                <div class="modal-body">
                    ...
                </div>

            </pre>
            
            <div class="modal-footer">
                <button type="button" class="btn btn-secondary" data-dismiss="modal">关闭</button>
            </div>
        </div>
    </div>
</div>