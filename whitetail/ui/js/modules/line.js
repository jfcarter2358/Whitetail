const Line = class {
    constructor(init) {
        // this.observer = init["observer"]
        // this.stream = init["stream"]
        this.plotName = init["name"]
        this.refresh = init["refresh"]
        this.x = init["x"]
        this.ys = init["ys"]
        this.xAxisLabel = init["x_axis_label"]
        this.yAxisLabel = init["y_axis_label"]
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
                let theme = localStorage.getItem('whitetail-theme');
                let bg_color = '#646464'
                let fg_color = '#ECEFF4'
                if (theme) {
                    if (theme == 'light') {
                        bg_color = '#ffffff'
                        fg_color = '#2E3440'
                    }
                }

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
                    title: {
                        text: `<b>${self.title}</b>`,
                        font: {
                            color: "#8F7E4F"
                        }
                    },
                    margin: {"t": 64, "b": 64, "l": 64, "r": 64},
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
                    },
                    font: {
                        color: fg_color
                    },
                    plot_bgcolor: bg_color,
                    paper_bgcolor: bg_color
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
