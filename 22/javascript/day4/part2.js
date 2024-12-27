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

        for (let e of range1) {
            if (range2.includes(e)) {
                result++
                console.log(result) // I am not sure why the result variable is not uptating outside this scope
                break
            }
        }

    })

    console.log(result)
}

main()
