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
        console.log(`button contents: ${title}`)
        console.log(`script contents: ${script}`)
        
        $(`#button-${this.name}`).html(title)
        $(`#button-${this.name}`).attr("onClick", callback)
        // $(`#script-${this.name}`).html(script)
        // eval($(`#script-${this.name}`).html())
        // $.getScript($(`#script-${this.name}`).html())
    }
}
