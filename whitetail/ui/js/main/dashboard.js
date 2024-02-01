// import line.js
// import stream.js
// import table.js
// import theme.js
// import modal.js
// import material.js
// import button.js

var graphs = []
var streams = []
var tables = []
var buttons = []

function LoadDashboard() {
    let displayTable = $("#display-table")
    let body = $(displayTable).children('tbody')[0]
    let rows = $(body).children("tr")
    for (let row of rows) {
        let cells = $(row).children('td')
        for (let cell of cells) {
            let divs = $(cell).children()
            let first = true
            let defs = []
            let ids = []
            for (let div of divs) {
                if (first) {
                    defs.push($(div).text())
                    first = false
                    continue
                }
                let id = $(div).attr('id')
                if (id.startsWith('script-')) {
                    continue
                }
                ids.push(id)
            }
            for (let i = 0; i < defs.length; i++) {
                let def = defs[i]
                let id = ids[i]
                if (id.startsWith('graph-')) { 
                    graphs.push(new Line(JSON.parse(def)))
                } else if (id.startsWith('stream-')) {
                    streams.push(new Stream(JSON.parse(def)))
                } else if (id.startsWith('table-')) {
                    tables.push(new Table(JSON.parse(def)))
                } else if (id.startsWith('button-')) {
                    buttons.push(new Button(JSON.parse(def)))
                }
            }
        }
    }
}

$( document ).ready(
    function()
    {
        LoadDashboard()
    }
);
