const Table = class {
    constructor(init) {
        this.source = init["source"]
        this.name = init["name"]
        this.x = init["x_coord"]
        this.y = init["y_coord"]
        this.rowSpan = init["row_span"]
        this.colSpan = init["col_span"]
        this.refresh = init["refresh"]
        this.title = init["title"]

        this.Update()
        setInterval(function() {
            this.Update()
        }.bind(this), this.refresh)

    }

    // Build() {
    //     $.ajax({
    //         type: "POST",
    //         url: "/api/v1/query",
    //         data: JSON.stringify({"query": this.source}),
    //         contentType:"application/json;",
    //         dataType:"json",
    //         success: function(data, status) {
    //             let contents = ''
    //             contents += '<tr>'
    //             for (let [key, _] of Object.entries(data[0])) {
    //                 contents += `<th class="whitetail-text-brown">${key}</td>`
    //             }
    //             contents += '</tr>'
    //             for (let datum of data) {
    //                 contents += '<tr>'
    //                 for (let [_, val] of Object.entries(datum)) {
    //                     contents += `<td>${val}</td>`
    //                 }
    //                 contents += '</tr>'
    //             }
                
    //             $(`#table-${this.name}`).text(contents)
    //         },
    //         error: function(data, status) {
    //             console.log(data)
    //             $("#error-container").text(data.responseText)
    //             openModal('error-modal')
    //         }
    //     });
    // }

    Render(data) {
        let contents = ''
        contents += '<tr style="color:#000!important;background-color:#8F7E4F!important;">'
        for (let [key, _] of Object.entries(data[0])) {
            contents += `<th>${key}</th>`
        }
        contents += '</tr>'
        for (let datum of data) {
            contents += '<tr>'
            for (let [_, val] of Object.entries(datum)) {
                contents += `<td>${val}</td>`
            }
            contents += '</tr>'
        }

        $(`#table-${this.name}`).html(contents)
    }

    Update() {
        $.ajax({
            type: "POST",
            url: "/api/v1/query",
            data: JSON.stringify({"query": this.source}),
            contentType:"application/json;",
            dataType:"json",
            success: function(data, status) {
                this.Render(data)
            }.bind(this),
            error: function(data, status) {
                console.log(data)
                $("#error-container").text(data.responseText)
                openModal('error-modal')
            }
        });
    }
}
