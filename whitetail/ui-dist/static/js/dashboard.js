const Line = class {
    constructor(init) {
        // this.observer = init["observer"]
        // this.stream = init["stream"]
        this.plotName = init["name"]
        this.refresh = init["refresh"]
        this.x = init["x"]
        this.ys = init["ys"]
        this.xAxisLabel = init["x_axis_label"]
        this.yAxisLabels = init["y_axis_label"]
        this.yLabels = init["y_labels"]
        this.title = init["title"]
        this.colors = init["colors"]
        this.layout = []
        this.data = []
        this.width = init["width"]
        this.height = init["height"]
        this.source = init["source"]

        this.Build()
        setInterval(function() {
            this.Update()
        }.bind(this), this.refresh)

    }

    Build() {
        var self = this;
        $.ajax({
            type: "POST",
            url: "/api/v1/query",
            data: JSON.stringify({"query": this.source}),
            contentType:"application/json;",
            dataType:"json",
            success: function(data, status) {
                console.log(`name2: ${self.plotName}`)
                self.data = []
                let xs = []
                let ys = []
                for (let _ in self.ys) {
                    ys.push([])
                }
                for (let datum of data) {
                    xs.push(datum[self.x])
                    for (let idx in self.ys) {
                        ys[idx].push(datum[self.ys[idx]])
                    }
                }
                for (let idx in self.ys) {
                    let temp = {
                        x: xs,
                        y: ys[idx],
                        mode: 'lines',
                        line: {
                            color: self.colors[idx],
                            width: 1
                        },
                        name: self.yLabels[idx]
                    }
                    self.data.push(temp)
                }

                self.layout = {
                    title: self.title,
                    margin: {"t": 32, "b": 32, "l": 32, "r": 32},
                    // height: self.height,
                    // width: self.width,
                    autosize: true,
                    showlegend: true,
                    grid: { rows: 1, columns: 1 },
                    xaxis: {
                        title: self.xAxisLabel
                    },
                    yaxis: {
                        title: self.yAxisLabel
                    }
                };

                // Plotly.newPlot(`graph-${this.observer}-${this.stream}`, this.data, this.layout);
                Plotly.newPlot(`graph-${self.plotName}`, self.data, self.layout)
            },
            error: function(data, status) {
                console.log(data)
                $("#error-container").text(data.responseText)
                openModal('error-modal')
            }
        });
        // fetch(`/api/v1/basestation/${this.observer}/${this.stream}`)
        // .then((response) => response.json())
        // .then((rawData) => {
        //     console.log(rawData)
        //     let data = JSON.parse(rawData)
        //     this.data = []
        //     let xs = []
        //     let ys = []
        //     for (let y of this.ys) {
        //         ys.push([])
        //     }
        //     for (let datum of data) {
        //         xs.push(datum[this.x])
        //         for (let idx in this.ys) {
        //             ys[idx].push(datum[this.ys[idx]])
        //         }
        //     }
        //     for (let idx in this.ys) {
        //         let temp = {
        //             x: xs,
        //             y: ys[idx],
        //             mode: 'lines',
        //             line: {
        //                 color: this.colors[idx],
        //                 width: 1
        //             },
        //             name: this.yLabels[idx]
        //         }
        //         this.data.push(temp)
        //     }

        //     this.layout = {
        //         title: this.title,
        //         margin: {"t": 32, "b": 32, "l": 32, "r": 32},
        //         height: this.height,
        //         width: this.width,
        //         showlegend: true,
        //         grid: { rows: 1, columns: 1 },
        //         xaxis: {
        //             title: this.xAxisLabel
        //         },
        //         yaxis: {
        //             title: this.yAxisLabel
        //         }
        //     };

        //     Plotly.newPlot(`graph-${this.observer}-${this-stream}`, this.data, this.layout);
        // })
        // .catch((error) => {
        //     console.log(error);
        //     $("#error-container").text(error)
        //     openModal('error-modal')
        // });
    }

    Render(data) {
        // let data = JSON.parse(rawData)
        let xs = []
        let ys = []
        for (let y of this.ys) {
            ys.push([])
        }
        for (let datum of data) {
            xs.push(datum[this.x])
            for (let idx in this.ys) {
                ys[idx].push(datum[this.ys[idx]])
            }
        }
        for (let idx in this.ys) {
            this.data[idx] = {
                x: xs,
                y: ys[idx],
                mode: 'lines',
                line: {
                    color: this.colors[idx],
                    width: 1
                },
                name: this.yLabels[idx]
            }
        }

        // Plotly.redraw(`graph-${this.observer}-${this-stream}`);
        Plotly.redraw(`graph-${this.plotName}`)
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
                $("#error-container").html(data.responseText)
                openModal('error-modal')
            }
        });
    }
}

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
        contents += '<tr>'
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

var theme;

$(document).ready(function() {
    theme = localStorage.getItem('whitetail-theme');
    if (theme) {
        if (theme == 'light') {
            $('.dark').addClass('light').removeClass('dark');
        } else {
            $('.light').addClass('dark').removeClass('light');
        }
    } else {
        theme = 'light'
        localStorage.setItem('whitetail-theme', theme);
    }
})

function toggleTheme() {
    if (theme == 'light') {
        theme = 'dark'
        $('.light').addClass('dark').removeClass('light');
    } else {
        theme = 'light'
        $('.dark').addClass('light').removeClass('dark');
    }
    localStorage.setItem('whitetail-theme', theme);
}

function closeModal(modalID) {
    document.getElementById(modalID).style.display='none'
}

function openModal(modalID) {
    document.getElementById(modalID).style.display='block'
}
function toggleSidebar() {
    var sidebar = document.getElementById("sidebar")
    var page_darken = document.getElementById("page-darken")
    if (sidebar.className.indexOf("show") == -1) {
        sidebar.classList.add("show");
        sidebar.classList.remove("left-slide-out-300");
        void sidebar.offsetWidth;
        sidebar.classList.add("left-slide-in-300")
        $("#sidebar").css("left", "0px")

        page_darken.classList.remove("fade-out");
        void page_darken.offsetWidth;
        page_darken.classList.add("fade-in");
        $("#page-darken").css("opacity", "1")
    } else {
        sidebar.classList.remove("show");
        sidebar.classList.remove("left-slide-in-300");
        void sidebar.offsetWidth;
        sidebar.classList.add("left-slide-out-300")
        $("#sidebar").css("left", "-300px")

        page_darken.classList.remove("fade-in");
        void page_darken.offsetWidth;
        page_darken.classList.add("fade-out");
        $("#page-darken").css("opacity", "0")
    }
}

function toggleAccordion(id) {
    var x = document.getElementById(id);
    if (x.className.indexOf("w3-show") == -1) {
        x.className += " w3-show";
    } else {
        x.className = x.className.replace(" w3-show", "");
    }
}

const Button = class {
    constructor(init) {
        this.source = init["source"]
        this.name = init["name"]
        this.callback = init["callback"]
        this.x = init["x_coord"]
        this.y = init["y_coord"]
        this.rowSpan = init["row_span"]
        this.colSpan = init["col_span"]
        this.title = init["title"]
        this.js = init["js"]

        this.Render(this.title, this.js, this.callback)
    }

    Render(title, script, callback) {
        $(`#button-${this.name}`).html(title)
        $(`#button-${this.name}`).attr("onClick", callback)
    }
}


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
