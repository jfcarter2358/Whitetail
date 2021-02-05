function query(basePath) {
    queryString = $("#query").val()
    document.getElementById("loader").style.display = "block";
    $.ajax({
        type: "POST",
        url: basePath + "/api/logs/query",
        data: JSON.stringify({"query": queryString}),
        contentType:"application/json;",
        dataType:"json",
        success: function(data, status) {
            $("#logs").html(data['logs'])
            document.getElementById("loader").style.display = "none";
        },
        error: function(data, status) {
            document.getElementById("loader").style.display = "none";
            alert(data)
        }
    });
}