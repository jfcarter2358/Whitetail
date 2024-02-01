const Stream = class {
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
    //     var self = this;
    //     $.ajax({
    //         type: "POST",
    //         url: "/api/v1/query",
    //         data: JSON.stringify({"query": this.source}),
    //         contentType:"application/json;",
    //         dataType:"json",
    //         success: function(data, status) {
    //             let contents = ''
    //             for (let datum of data) {
    //                 contents += `${JSON.stringify(datum)}\n` 
    //             }
                
    //             $(`#stream-${self.name}`).text(contents)
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
        for (let datum of data) {
            contents += `<span>${JSON.stringify(datum)}</span><br>` 
        }

        $(`#stream-${this.name}`).html(contents)
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
