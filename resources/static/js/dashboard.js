const Line = class {
    constructor(init) {
        this.observer = init["observer"]
        this.stream = init["stream"]
        this.name = init["name"]
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

        this.Build()
        setTimeout(function() {
            this.Update()
        }, this.refresh)

    }

    Build() {
        fetch(`/api/v1/basestation/${this.observer}/${this.stream}`)
        .then((response) => response.json())
        .then((rawData) => {
            this.data = []
            let xs = []
            let ys = []
            for (y of this.ys) {
                ys.push([])
            }
            for (datum of rawData) {
                xs.push(datum[this.x])
                for (idx in this.ys) {
                    ys[idx].push(datum[this.ys[idx]])
                }
            }
            for (idx in this.ys) {
                let temp = {
                    x: xs,
                    y: ys[idx],
                    mode: 'lines',
                    line: {
                        color: this.colors[idx],
                        width: 1
                    },
                    name: this.yLabels[idx]
                }
                this.data.push(temp)
            }

            this.layout = {
                title: this.title,
                margin: {"t": 32, "b": 32, "l": 32, "r": 32},
                height: this.height,
                width: this.width,
                showlegend: true,
                grid: { rows: 1, columns: 1 },
                xaxis: {
                    title: this.xAxisLabel
                },
                yaxis: {
                    title: this.yAxisLabel
                }
            };

            Plotly.newPlot(`graph-${this.observer}-${this-stream}`, this.data, this.layout);
        })
        .catch((error) => {
            console.log(error);
            $("#error-container").text(error)
            openModal('error-modal')
        });
    }

    Render(rawData) {
        let xs = []
        let ys = []
        for (y of this.ys) {
            ys.push([])
        }
        for (datum of rawData) {
            xs.push(datum[this.x])
            for (idx in this.ys) {
                ys[idx].push(datum[this.ys[idx]])
            }
        }
        for (idx in this.ys) {
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

        Plotly.redraw(`graph-${this.observer}-${this-stream}`);
    }

    Update() {
        fetch(`/api/v1/basestation/${this.observer}/${this.stream}`)
        .then((response) => response.json())
        .then((rawData) => {
            this.RenderGraph(rawData)
        })
        .catch((error) => {
            console.log(error);
            $("#error-container").text(error)
            openModal('error-modal')
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

