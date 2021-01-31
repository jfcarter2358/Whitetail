
var levelList = ['INFO', 'WARN', 'DEBUG', 'TRACE', 'ERROR']

function filterLevel(basePath, level) {
    index = levelList.indexOf(level);
    if (index > -1) {
        levelList.splice(index, 1);
        $('#' + level).removeClass("whitetail-green")
        $('#' + level).addClass("whitetail-white")
    } else {
        levelList.push(level)
        $('#' + level).removeClass("whitetail-white")
        $('#' + level).addClass("whitetail-green")
    }
    refreshLogs(basePath)
}

function getLogs(basePath) {
    service = $("#services_button").text()
    if (service != "Select a Service") {
        if ($("#live_view").is(':checked')) {
            keywordList = $("#keyword_list").val()
            lineLimit = $("#line_limit").val()
            $.ajax({
                type: "POST",
                url: basePath + "/api/logs/" + service,
                data: JSON.stringify({"limit": lineLimit,"keyword-list": keywordList,"log-levels":levelList}),
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
        data: JSON.stringify({"limit": lineLimit,"keyword-list": keywordList,"log-levels":levelList}),
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

function refreshLogs(basePath) {
    service = $("#services_button").text()
    if (service != "Select a Service") {
        keywordList = $("#keyword_list").val()
        lineLimit = $("#line_limit").val()
        $.ajax({
            type: "POST",
            url: basePath + "/api/logs/" + service,
            data: JSON.stringify({"limit": lineLimit,"keyword-list": keywordList,"log-levels":levelList}),
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