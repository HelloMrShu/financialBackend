function expand() {
    var id = $()
    var content = $(".rule-content").val()

    $(".modal-body").html(content)
    $("#panel").modal("show")
}   

