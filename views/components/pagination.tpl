<div class="row">
    <div class="paginator">
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
    </div>
</div>