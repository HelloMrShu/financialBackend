<div class="container" id="navbar">
    <nav>
      <ul class="nav nav-tabs nav-ul">
        <li role="presentation" {{ if compare .ap "index" }} class="active-menu"{{ end }}>
          <a href="/">首页</a>
        </li>
        <li role="presentation" {{ if compare .ap "ae_test" }} class="active-menu"{{ end }}>
          <a href="/ae/test">AE测试模板</a>
        </li>
        <li role="presentation" {{ if compare .ap "ae_prod" }} class="active-menu"{{ end }}>
          <a href="/ae/prod">AE线上模板</a>
        </li>
        <li role="presentation" {{ if compare .ap "ex_test" }} class="active-menu"{{ end }}>
          <a href="/ex/test">汇聚测试模板</a>
        </li>
        <li role="presentation" {{ if compare .ap "ex_prod" }} class="active-menu"{{ end }}>
          <a href="/ex/prod">汇聚线上模板</a>
        </li>
      </ul>
    </nav>
  </div>