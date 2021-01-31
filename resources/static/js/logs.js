function getLogs(basePath) {
    service = $("#services_button").text()
    if (service != "Select a Service") {
        if ($("#live_view").is(':checked')) {
            keywordList = $("#keyword_list").val()
            lineLimit = $("#line_limit").val()
            $.ajax({
                type: "POST",
                url: basePath + "/api/logs/" + service,
                data: JSON.stringify({"limit": lineLimit,"keyword-list": keywordList}),
                contentType:"application/json;",
                dataType:"json",
                success: function(data, status) {
                    $("#logs").html(data['logs'])
                },
                error: function(data, status) {
                    console.log(data)
                }
            });
        }
    }
}

function changeService(basePath, service) {
    $("#services_button").text(service)
    toggleDropdown('services_dropdown')
    keywordList = $("#keyword_list").val()
    lineLimit = $("#line_limit").val()
    $.ajax({
        type: "POST",
        url: basePath + "/api/logs/" + service,
        data: JSON.stringify({"limit": lineLimit,"keyword-list": keywordList}),
        contentType:"application/json;",
        dataType:"json",
        success: function(data, status) {
            $("#logs").html(data['logs'])
        },
        error: function(data, status) {
            console.log(data)
        }
    });
}