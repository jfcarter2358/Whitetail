function updateSettings(basePath) {
    data = {
        "primary_color": {
            "background": $('#bg1').val(),
            "text": $('#t1').val()
        },
        "secondary_color": {
            "background": $('#bg2').val(),
            "text": $('#t2').val()
        },
        "tertiary_color": {
            "background": $('#bg3').val(),
            "text": $('#t3').val()
        },
        "INFO_color": $('#INFO').val(),
        "WARN_color": $('#WARN').val(),
        "DEBUG_color": $('#DEBUG').val(),
        "TRACE_color": $('#TRACE').val(),
        "ERROR_color": $('#ERROR').val(),
    }

    $.ajax({
        type: "POST",
        url: basePath + "/api/settings/colors/update",
        data: JSON.stringify(data),
        contentType:"application/json;",
        dataType:"json",
        success: function(data, status) {
            window.location.reload(true);
        },
        error: function(data, status) {
            console.log(data)
        }
    });
}

function defaultSettings(basePath) {
    data = {}

    $.ajax({
        type: "POST",
        url: basePath + "/api/settings/colors/default",
        data: JSON.stringify(data),
        contentType:"application/json;",
        dataType:"json",
        success: function(data, status) {
            window.location.reload(true);
        },
        error: function(data, status) {
            console.log(data)
        }
    });
}

function defaultLogo(basePath) {
    data = {}

    $.ajax({
        type: "POST",
        url: basePath + "/api/settings/logo/default",
        data: JSON.stringify(data),
        contentType:"application/json;",
        dataType:"json",
        success: function(data, status) {
            window.location.reload(true);
        },
        error: function(data, status) {
            console.log(data)
        }
    });
}

function defaultIcon(basePath) {
    data = {}

    $.ajax({
        type: "POST",
        url: basePath + "/api/settings/icon/default",
        data: JSON.stringify(data),
        contentType:"application/json;",
        dataType:"json",
        success: function(data, status) {
            window.location.reload(true);
        },
        error: function(data, status) {
            console.log(data)
        }
    });
}