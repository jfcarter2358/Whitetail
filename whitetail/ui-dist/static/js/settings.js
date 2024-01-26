function updateSettings() {
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
        url: "/api/settings/colors/update",
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

function defaultSettings() {
    data = {}

    $.ajax({
        type: "POST",
        url: "/api/settings/colors/default",
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

function defaultLogo() {
    data = {}

    $.ajax({
        type: "POST",
        url: "/api/settings/logo/default",
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

function defaultIcon() {
    data = {}

    $.ajax({
        type: "POST",
        url: "/api/settings/icon/default",
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
