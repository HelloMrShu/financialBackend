function expandexchange(id) {
    rnId = 'rn-' + id
    rn = "#" + rnId
    var name = $(rn).text()

    rcId = 'rc-'+id
    rc = "#"+rcId
    var content = $(rc).val()
    content = JSON.stringify(JSON.parse(content), null, 2);
     
    $(".modal-title").html("模板：" + name)
    $(".modal-body").html("<pre>"+ content + "</pre>")
    $("#panel").modal("show")
}

