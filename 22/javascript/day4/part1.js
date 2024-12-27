import readline from "readline"
import fs from "fs"

function createRange(start, end) {
    start = +start
    end = +end

    const range = []

    for (let i = start; i <= end; i++) {
        range.push(i)
    }

    return range
}

function main() {

    let result = 0

    const file = readline.createInterface({
        input: fs.createReadStream("./input.txt"),
        output: process.stdout,
        terminal: false
    })

    file.on("line", (line) => {

        const [rangeStr1, rangeStr2] = line.split(",")

        const range1 = createRange(...(rangeStr1.split("-")))
        const range2 = createRange(...(rangeStr2.split("-")))

        const cond1 = range1.every(e => range2.includes(e))
        const cond2 = range2.every(e => range1.includes(e))

        if (cond1 || cond2) {
            result++
            console.log(result) // I am not sure why the result variable is not uptating outside this scope
        }
    })

    console.log(result)
}

main()
