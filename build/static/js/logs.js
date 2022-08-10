
var levelList = ['INFO', 'WARN', 'DEBUG', 'TRACE', 'ERROR']

function formatQuery(db_name, service, limit) {
    if (levelList.length == 0) {
        return ""
    } else {
        infoList = '"level" = ' + levelList[0]
        for (var i = 1; i < levelList.length; i++) {
            infoList = infoList + ' OR "level" = ' + levelList[i]
        }

        // console.log(infoList);
        // console.log(`SELECTBY (((service = ${service} AND level IN ${infoList}) ORDERDESC timestamp) LIMIT ${limit})`);
        return `GET RECORD ${dbName}.logs * | FILTER (${infoList}) AND "service" = ${service} | ORDERDSC "timestamp" | LIMIT ${limit}`
        // return `SELECTBY (((service = ${service} AND level IN ${infoList}) ORDERDESC timestamp) LIMIT ${limit})`
    }
}

function togglePausePlay() {
    if ($("#live_view").hasClass('play')) {
        $("#live_view").removeClass('play');
        $("#live_view").removeClass('fa-pause');
        $("#live_view").addClass('pause');
        $("#live_view").addClass('fa-play');
    } else {
        $("#live_view").removeClass('pause');
        $("#live_view").removeClass('fa-play');
        $("#live_view").addClass('play');
        $("#live_view").addClass('fa-pause');
    }
}

function filterLevel(basePath, level) {
    index = levelList.indexOf(level);
    if (index > -1) {
        levelList.splice(index, 1);
        $('#' + level).removeClass("branding-secondary")
        $('#' + level).addClass("whitetail-white")
        $('#' + level + "_check").removeClass("fa-check-square")
        $('#' + level + "_check").addClass("fa-square")
    } else {
        levelList.push(level)
        $('#' + level).removeClass("whitetail-white")
        $('#' + level).addClass("branding-secondary")
        $('#' + level + "_check").removeClass("fa-square")
        $('#' + level + "_check").addClass("fa-check-square")
    }
    service = $("#services_button").text()
    lineLimit = $("#line_limit").val()
    refreshLogs(basePath)
}

function changeService(basePath, service) {
    $("#services_button").html("Service Filter: " + service + '<i class="fas fa-caret-down" style="position: absolute; right: 10px;top: 12px;"></i>');
    toggleDropdown('services_dropdown')
    lineLimit = $("#line_limit").val()
    queryString = formatQuery(db_name, service, lineLimit)
    console.log(queryString)
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
                logOut = ''
                if (data['logs'] != null) {
                    for (var i = 0; i < data['logs'].length; i++) {
                        logOut += data['logs'][i] + "<br>"
                    }
                }
                $("#logs").html(logOut)
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
    if (service != "Service Filter") {
        lineLimit = $("#line_limit").val()
        realService = service.split(": ");
        console.log(realService)
        queryString = formatQuery(realService[1], lineLimit)
        console.log(queryString)
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
                    console.log("Got data")
                    logOut = ''
                    if (data['logs'] != null) {
                        for (var i = 0; i < data['logs'].length; i++) {
                            logOut += data['logs'][i] + "<br>"
                        }
                    }
                    $("#logs").html(logOut)
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

function updateServices(basePath) {
    $.ajax({
        type: "GET",
        url: basePath + "/api/indices/services",
        contentType:"application/json;",
        dataType:"json",
        success: function(data, status) {
            console.log(data)
            $("#services_dropdown").empty();
            data['services'].sort();
            for(var i = 0; i < data['services'].length; i++) {
                $("#services_dropdown").append("<button class=\"w3-bar-item w3-button branding-hover-primary\" onclick=\"changeService('" + basePath + "', '" + data['services'][i] + "')\">" + data['services'][i] + "</button>")
            }
        },
        error: function(data, status) {
            console.log(data)
        },
        async: false
    });
    toggleDropdown('services_dropdown')
}
