
var levelList = ['INFO', 'WARN', 'DEBUG', 'TRACE', 'ERROR']

function formatQuery(service, limit) {
    if (levelList.length == 0) {
        return ""
    } else {
        infoList = levelList[0]
        for (var i = 1; i < levelList.length; i++) {
            infoList = infoList + "," + levelList[i]
        }
        return `((service = ${service} AND level IN ${infoList}) ORDER_DESCEND timestamp) LIMIT ${limit}`
    }
}


function filterLevel(basePath, level) {
    index = levelList.indexOf(level);
    if (index > -1) {
        levelList.splice(index, 1);
        $('#' + level).removeClass("branding-secondary")
        $('#' + level).addClass("whitetail-white")
    } else {
        levelList.push(level)
        $('#' + level).removeClass("whitetail-white")
        $('#' + level).addClass("branding-secondary")
    }
    service = $("#services_button").text()
    lineLimit = $("#line_limit").val()
    refreshLogs(basePath)
}

function changeService(basePath, service) {
    $("#services_button").text(service)
    toggleDropdown('services_dropdown')
    lineLimit = $("#line_limit").val()
    queryString = formatQuery(service, lineLimit)
    if (queryString == "" ) {
        $("#logs").html("")
    } else {
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
                console.log(data)
            }
        });
    }
}

function refreshLogs(basePath) {
    service = $("#services_button").text()
    if (service != "Select a Service") {
        lineLimit = $("#line_limit").val()
        queryString = formatQuery(service, lineLimit)
        if (queryString == "" ) {
            $("#logs").html("")
        } else {
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
                    console.log(data)
                    document.getElementById("loader").style.display = "none";
                }
            });
        }
    }
}