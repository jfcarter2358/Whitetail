function query(basePath) {
    queryString = $("#query").val()
    dbName = $("#db_name").text()
    queryString = queryString.replace("${db}", dbName)
    document.getElementById("loader").style.display = "block";
    $.ajax({
        type: "POST",
        url: basePath + "/api/logs/query",
        data: JSON.stringify({"query": queryString}),
        contentType:"application/json;",
        dataType:"json",
        success: function(data, status) {
            console.log(data)
            if (data['error'] != "") {
                document.getElementById("error_message").style.display = "block";
                $("#error").html(data['error'])
            } else {
                document.getElementById("error_message").style.display = "none";
                logOut = ''
                if (data['logs'] != null) {
                    for (var i = 0; i < data['logs'].length; i++) {
                        logOut += data['logs'][i] + "<br>"
                    }
                }
                $("#logs").html(logOut)
            }
            document.getElementById("loader").style.display = "none";
        },
        error: function(data, status) {
            console.log(data)
            document.getElementById("loader").style.display = "none";
            alert(data.statusText)
        }
    });
}
