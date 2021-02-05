
var levelList = ['INFO', 'WARN', 'DEBUG', 'TRACE', 'ERROR']
var PREVIOUS_QUERY = ""

function formatQuery(service, limit) {
    // hard-coding queries for now
    // because I can't be bothered to figure out a progromatic
    // way to generate them
    if (levelList.length == 0) {
        // return `@service:${service} LIMIT ${limit}`
        return ""
    } else if (levelList.length == 1) {
        return `( @level:${levelList[0]} AND @service:${service} ) LIMIT ${limit}`
    } else if (levelList.length == 2) {
        return `( ( @level:${levelList[0]} OR @level:${levelList[1]} ) AND @service:${service} ) LIMIT ${limit}`
    } else if (levelList.length == 3) {
        return `( ( ( @level:${levelList[0]} OR @level:${levelList[1]} ) OR @level:${levelList[2]} ) AND @service:${service} ) LIMIT ${limit}`
    } else if (levelList.length == 4) {
        return `( ( ( ( @level:${levelList[0]} OR @level:${levelList[1]} ) OR @level:${levelList[2]} ) OR @level:${levelList[3]} ) AND @service:${service} ) LIMIT ${limit}`
    } else if (levelList.length == 5) {
        return `( ( ( ( ( @level:${levelList[0]} OR @level:${levelList[1]} ) OR @level:${levelList[2]} ) OR @level:${levelList[3]} ) OR @level:${levelList[4]} ) AND @service:${service} ) LIMIT ${limit}`
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
    PREVIOUS_QUERY = formatQuery(service, lineLimit)
    refreshLogs(basePath)
}

function query(basePath) {
    queryString = $("#query").val()
    PREVIOUS_QUERY = queryString
    $.ajax({
        type: "POST",
        url: basePath + "/api/logs/query",
        data: JSON.stringify({"query": queryString}),
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

function changeService(basePath, service) {
    $("#services_button").text(service)
    toggleDropdown('services_dropdown')
    lineLimit = $("#line_limit").val()
    queryString = formatQuery(service, lineLimit)
    PREVIOUS_QUERY = queryString
    $.ajax({
        type: "POST",
        url: basePath + "/api/logs/query",
        data: JSON.stringify({"query": queryString}),
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
        $.ajax({
            type: "POST",
            url: basePath + "/api/logs/query",
            data: JSON.stringify({"query": PREVIOUS_QUERY}),
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