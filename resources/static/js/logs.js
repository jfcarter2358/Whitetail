
var levelList = ['INFO', 'WARN', 'DEBUG', 'TRACE', 'ERROR']

function formatQuery(service, limit) {
    // hard-coding queries for now
    // because I can't be bothered to figure out a progromatic
    // way to generate them
    maxLimit = parseInt(limit) * 5
    if (levelList.length == 0) {
        return ""
    } else if (levelList.length == 1) {
        return `(((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[0]}) ORDER_ASCEND timestamp) LIMIT ${limit}`
    } else if (levelList.length == 2) {
        return `(((((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[0]}) LIMIT ${limit}) OR (((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[1]}) LIMIT ${limit})) ORDER_ASCEND timestamp) LIMIT ${limit}`
    } else if (levelList.length == 3) {
        return `((((((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[0]}) LIMIT ${limit}) OR (((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[1]}) LIMIT ${limit})) OR (((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[2]}) LIMIT ${limit})) ORDER_ASCEND timestamp) LIMIT ${limit}`
    } else if (levelList.length == 4) {
        return `(((((((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[0]}) LIMIT ${limit}) OR (((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[1]}) LIMIT ${limit})) OR (((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[2]}) LIMIT ${limit})) OR (((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[3]}) LIMIT ${limit})) ORDER_ASCEND timestamp) LIMIT ${limit}`
    } else if (levelList.length == 5) {
        return `((((((((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[0]}) LIMIT ${limit}) OR (((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[1]}) LIMIT ${limit})) OR (((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[2]}) LIMIT ${limit})) OR (((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[3]}) LIMIT ${limit})) OR (((service = ${service} LIMIT ${maxLimit}) AND level = ${levelList[4]}) LIMIT ${limit})) ORDER_ASCEND timestamp) LIMIT ${limit}`
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
    document.getElementById("loader").style.display = "block";
    refreshLogs(basePath)
}

function changeService(basePath, service) {
    $("#services_button").text(service)
    toggleDropdown('services_dropdown')
    lineLimit = $("#line_limit").val()
    queryString = formatQuery(service, lineLimit)
    console.log(queryString)
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

function refreshLogs(basePath) {
    service = $("#services_button").text()
    if (service != "Select a Service") {
        lineLimit = $("#line_limit").val()
        queryString = formatQuery(service, lineLimit)
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