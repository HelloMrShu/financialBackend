<!DOCTYPE html>
<html>
  <head>
      <title>Beego</title>
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
      <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootswatch/3.3.5/cerulean/bootstrap.min.css">
      <link rel="stylesheet" type="text/css" href="https://cdn.bootcss.com/font-awesome/4.6.0/css/font-awesome.min.css">
      <link rel="stylesheet" href="http://netdna.bootstrapcdn.com/bootstrap/3.0.3/css/bootstrap-theme.min.css">
      <link type="text/css" rel="stylesheet" href="/static/css/layout.css" />
      <link type="text/css" rel="stylesheet" href="/static/css/rule.css" />
      <link type="text/css" rel="stylesheet" href="/static/css/index.css" />
      {{.HtmlHead}}
  </head>
  <body>
      {{template "components/nav.tpl" .}}

      <div class="container" id="content">
          {{.LayoutContent}}
      </div>
      
      <div class="container">
        <div class="footer">
              Copyright © 2020 Sohu All Rights Reserved
        </div>
      </div>
      
      <script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js"></script>
      {{.Scripts}}
  </body>
</html>