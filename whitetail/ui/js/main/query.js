// import theme.js
// import modal.js
// import material.js

function query() {
    queryString = $("#query").val()
    $("#spinner").css("display", "block")
    $("#page-darken").css("opacity", "1")
    $.ajax({
        type: "POST",
        url: "/api/v1/query",
        data: JSON.stringify({"query": queryString}),
        contentType:"application/json;",
        dataType:"json",
        success: function(data, status) {
            console.log(data)
            logOut = ''
            if (data != null) {
                for (var i = 0; i < data.length; i++) {
                    logOut += JSON.stringify(data[i]) + "<br>"
                }
            }
            $("#logs").html(logOut)
            $("#spinner").css("display", "none")
            $("#page-darken").css("opacity", "0")
        },
        error: function(data, status) {
            console.log(data)
            $("#spinner").css("display", "none")
            $("#page-darken").css("opacity", "0")
            $("#error-container").text(error)
            openModal('error-modal')
        }
    });
}
