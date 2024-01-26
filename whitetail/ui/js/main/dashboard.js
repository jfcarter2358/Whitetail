// import line.js
// import theme.js
// import modal.js
// import material.js

var graphs = []

function LoadDashboard() {
    let displayTable = $("#display-table")
    let body = $(displayTable).children('tbody')[0]
    let rows = $(body).children("tr")
    for (row of rows) {
        let cells = $(row).children('td')
        for (cell of cells) {
            let first = true
            let divs = $(cell).children('div')
            let graphDef = ""
            for (div of divs) {
                if (first) {
                    graphDef = $(div).text()
                    first = false
                    continue
                }
                console.log(`loading graph ${$(div).attr('id')}`)
            }
            if (graphDef.length > 0) {
                console.log(`graph def: ${graphDef}`)
                graphs.push(new Line(JSON.parse(graphDef)))
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
