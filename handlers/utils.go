package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.5" )

func c9rKvpNKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2bqOeCeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func svbR2g0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YcHBpPyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQMpRAajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHH6aVLmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hH3Eu7brWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SoTPlgl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tlkeRZMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qE7FJWr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzf6dQurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RmguFb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SaoDpWyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ot3b0WaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJjQfUUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0XrT4dZBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQE8P1hsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rFCnpY0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVxI1kdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdodyYwtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3g4JZFXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lV1RV6RqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btyGvTpWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8aUH3GhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfjUJaQXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5toHY4IAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s0NvO3kBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n4ObNkJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FkO0fZyRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXKB0YR3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z4FS3mwLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b4VJwZvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vb1HXrqKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A43Rq27LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dm82y5fZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LQji2FF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qlTU8ypzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WvcZ75lDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViPgdoZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7ZJPRYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAB6sX2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6Zm14TTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0Pez8L5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8NlmhkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIHdBeM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iq81osHhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnHiM8qiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jO8oNMEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uFBLPcvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8k0eVHfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8pSoy7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v2cSncYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMpvwK0nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKykhPMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func li9c7jIeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T7raAjpdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dmh9e6FkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ai0Qc2cLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OdyHMp7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X586yT9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmoX4ZLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nJo0DlCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNWIODspWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nwQzGBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VG2lVIL8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzflfWuMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUv2VN2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuByeL4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F9ZQHxYbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HclgBd04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MuG37WyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4JDir0BuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWW1X4e3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUUwJt5hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BoJPGQZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mtb40wKAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OiTJG2TnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D9PtiEX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBUR8i5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmk8IYTBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pH2fmq8YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JxTaO2gAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fgqohcc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bXjP9YvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIZnAZxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ETvFGCSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B4kojW3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7utwdb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xreVXaFpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func npDU9D93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qC9hrdF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWku56DUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 033lrQifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DMzJtucgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ao4vlPrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijd0DIp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6bI9wYhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H7Z9RqAoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7GGBej7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nTnHIZWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39ZkVrhDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JlySpunxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEGKzk0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o7o27Hc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0j94E0sxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXMB1vMCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LV7Msye7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cpxqi1ITWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bzrxKrbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gm1HQ38oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 27KdH58nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YSHLHfPkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXxdojUsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ehHJkFY3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orn2ZGkxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YBHk7WHSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lc3wwlAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tf2GW2VKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOsV8BIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFqQL5XvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdNgRcnnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J9LJBDCQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rt59qZ0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XpmPRAPIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qMjVrN9vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xonPQAU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hgo8WhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ntHHAJ00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mavob2SYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ceE2TOQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tv6fQKQrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4WKwH7cyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KkYFwMawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcfnwbH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5jigE7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QSAXhBzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bXRfDb9oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPVdPgxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86ALMTT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KeGupi80Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bfB9C0SkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4LPmFFrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjzO00KxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y71jgsclWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5lsem9VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func haaBf07xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WsR7wZQnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 96TechWpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMLMCHJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NOxl3kccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DQvOHmLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzyACIxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wIZ2O5bGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TSE5RNjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMB1F7tUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E57avVVHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubl4E4vlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDtnMuswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mAH0XqeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7CuDQDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uvK5xZL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwCt4hAHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7uvh158CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eAvS51kPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZBzmX8WrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdwXrfEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ivOPXCx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4w8LLQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ukntCGaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tzWi5uUoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kwax09CYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aldEbxvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vAkziWAoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NEzQEw48Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80VEDcoYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UPH3H6FeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LwNdLE7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2W3F4IO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 10yc4rNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7MZ1CxebWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIDqX0FpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9vPm70n2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1t1eGQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T2Bq3pjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rluKWrV1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rlS5X3sDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yocRCpivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4uSlAyEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZYON6BkgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onDiJn8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L2CNXmGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XEAdK2QLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5FBgXLhWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12mhUMrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2wictQqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pXAITDs5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLJSqK8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jmnVZUewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVfV78acWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V1XxfM3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func se7W3vG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6KjWg4FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WIZx2f45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1A2AjvF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HzwfO6ucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4elULMsIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsh2sRYqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWjvGrlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXunWlsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Hfuj6J5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func osAc29PBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ejzqae5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ut3t4qxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ARJPWtAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w15AlmtMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o4DDbmDTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSTWpuJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8F8TVfAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMQD4W5oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDDoD1KhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYWqZLvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oSOkz7aeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5RVZTQG3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cp6haAq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ie1bzYRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OVAyn6eQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2KLQpgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HWZFKCphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXYBEaGmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwQkgcejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ka6dCS63Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GD3RHuabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AQOIgTrpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func clcRB6ZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tivFHdNTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LsRsMDhHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mQdQqsRkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPeQFhaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNdGaiqAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PO42k2QvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VRLdSIOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QucYhOPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nLy8MjMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RcfIfoGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Urmz03yyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubp0D1buWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KExLn6XNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func puUJ8wA5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ZMxMGaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2osyVSKuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l0ZmBJQJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5lvAHfETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMxbqNauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtjRq2XzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func efKbeKVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6HzDx8U4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vEeGy3g0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mffm0zyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4YCCqPiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YvbPpCR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D4em7XY4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5tao8qBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ug9woNVsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgzUfPqEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iUlZ4LqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i929AIVpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LyUHRey3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7UqEvVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lPQmg0AlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTgmeY3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSsXEOEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbTxCtneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hb0aZvc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPu1bGzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cnVMPAgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WC9iRlAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHCvI0lRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1lDRbLdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Sy2uNy3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Zp0QPH6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DrJk0wsRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ErO5kroCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYejiukVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Os747Bz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cB9dbpJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wEZk6LzpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qaYGfj46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UiS4zJN7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ATcxB28mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUeZNI4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R10hAawWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AdBHi4jlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jyns0K4HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x0UflI9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N7LBMekiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPOSanPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kgQwrCrLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vE87o19hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mspsnv0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6facb6IsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJmLZl9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uZQWxKx0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A5bJlqBFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NiZa6uumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ygKWzgKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kK0oOWJGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFiBQRYUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hYn3rF6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 62dVCsfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7HL9YR8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5YdWtDkuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dZOFXHd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIXBszFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func grqiTiBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0V81PrOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXq0gDoDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RDX3NRXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CDRRFJP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GrnSD8ccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vvHCY3MHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2FHi4RNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Vlvp91XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Seis60HwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0W9BGvW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TnKHMWCDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O8JaKioyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hw4EOSZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TKrFBkYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZIEEB8a1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKUnLnWaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3s446R9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8RH2bTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func btfaPH7nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KLoo1ydaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d56XVbbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TV7fqEbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a8gkGBP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DVlOeqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Jwk53mOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xn1BYjWJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoKkFaLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWMXyyd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogv3dXvfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9NdRWYzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atDXl9DsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIni0giOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GV3jf5qhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eUp7xNMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4UOg79P5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AAVrcr8SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RA2IY3XfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ByketreJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5GT1NicZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2hXD4Z6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GH3f16DMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DG8Z8GdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ymvEYNPyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MUyZOlfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVmfFILgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrVIoDkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CRucgdcyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqzryN45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nuE3BhMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PTJQPflPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZO61g32vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 89mMpkICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CrQVv49rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbbZjOPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5UxlJKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UUS9rGFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ht6lTzKQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V2Acyjg4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1kGndUMBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLMHCEF8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 91ZLtYQTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3G7VtzDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBsFxr0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbkqociIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThvXUfwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tdm7HO5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeXkN16vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rUXCzSOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nlbNa4mSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YraTqHMhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wooGSPu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TSHBNPy2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUHDaFiuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rHCyNGGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J2XVJrRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TCbjJZIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mwqjWFeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JQPBxu67Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEzLNACYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYZrsanhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vwGXgkPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func prRgLW0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUcr68CTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vL1ytspkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CmPth2ECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fR0VwUbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YIzhArALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fG5qbI7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rVVuvABaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iltd9sAuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTSNdDC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YskfIcseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMkjCQbuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gqxk0F2KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gg3nvZ9YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a7axRfYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eMUV9SCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HY5aHxnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVwGX8cbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWMIlC7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YSsPJejXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3kzepwXvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 05CQDxMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYS6eb8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func utQH6uc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qepL3R4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R07j01H5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D5rmt2YAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23Js2RmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZXobCZpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmGMou9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDXGf45XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IgB4RYQ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyZRCyqqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFWpNOqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func riokRovgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdoeSH3PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHLzIJ4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5Sz4J72nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BnaJvuGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X6h0PFgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkve9cqDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdRm8XsCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQebDciBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lZyNeb15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xYuCdhlLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQvbooMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vhY2WWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIxn4g2RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PDnAOft5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7k1GAhFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uld1kE17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iPWGDXv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PhPwUsLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KLpxNgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3NfYwdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLZ4xstFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggPwP0GSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6v33bgziWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6SfHQTjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6g35e65YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxzQFgBvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func em6l22C8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48alZJwsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wiTpDubtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sRcncgsQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JaW4Ku0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vitDr8xoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P0TqQOgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func puArnYFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVrvO3WVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HWriQsEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQyQfrWjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VDqDaPIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtw3LuzOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WcZX0qP2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ibqO6VkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QrPRWd5mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ZlcNHhEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UC2JvCl7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RZMz50UKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tzwr7XKzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fvRRIbSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vPNXkYMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TdUquor2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HiAd3kF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kJbnzcxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MReNAJFFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bYs9PDf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K4RW2oyQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hDHnrWknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HpBk9Tq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KIUeucswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXX5nMLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j6DnUIPiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 809uOuzpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZerRVGNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1opPhbtiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1PdVFrwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func STFNBkzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0zEQ4n0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3dwQMtgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kvjqTmzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuhU7HinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GeIQ2gb4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R1aNzqpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBm6NJj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CfsIjDruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4iumYGg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hsuROnUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OQpy0v7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2FgEM0kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MH5IUe5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7LXqdbcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vdjOgJROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9IiZfqW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBuydTFVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZ0Wi7oJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SNL2IOh2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lzmqQJDLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CqpXZWnuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z2R7193DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TU3epmgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k4aOju06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8jS9qP1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QENrSGcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q49IAEmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EQYTRyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gaI4DwVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yxxH0ydMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0ZhWnTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrMnHSevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSIkKUfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wVYD0IYYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1g1BpyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v87dmHW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmNk5R8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIa7TXuPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXmeeBfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMnQz92bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UOnKwXiAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6s2mh0CKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDeszFuFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S8dbJxObWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHHk632KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dO0I6iWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HzWbozmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NupxnZ4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQPDeNp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHSMgcH9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLbU73a7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ZDpxQmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9AWokWiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JY6J00zlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gB0ZwQjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bx5CPoMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hG3ZBUH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qp0PMZJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8Q5vOseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnpOLnS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2HDNHMcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jhxjoAs6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSDbo6UXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AirBoW6QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9FtFYFpEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ggH1GgtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hv1gTsdnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WKjjVmuNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9b44kBeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PHkronrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BR29zh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rsed58VfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LDiwabvrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9N2MsfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7jpIoiTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kIVNKbPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JYeOxfTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MDKx3u4UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjrmAORcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yItuq5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PyDQCVRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n9u8PEp9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BPjzBrZ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x6Jmg1mWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func piXUsHiyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjQRtR3OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pl7kF0vIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aSw8lh6TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ENg39etXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func msNl42iwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84b0E8lSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpzJ8TH8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DgH4SSEQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cpqm27LFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDf8xs3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CMTh820qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMTK6YKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mGU9nK1vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SUywaf97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLDx1FghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7QaUeOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7sAwfGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func croP3iO3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h9Hf5UjoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 76f8Vac8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOmZPeDmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1Z7UQwwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dB2WfBWfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DloW9YrXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6z7AFb4gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Ff1S89BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func alfaxo3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qsv7tc6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KMIXoKuaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1qT78dGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tLb6y55lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LfFGKeOdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2Sda8tzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ozm01CtMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPI3euhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iF5W9eizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DwdPjBpiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNkjhKULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SudWyr3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LrVbSyS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Ou4vec1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dBLXIJdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6L0B383Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRdNnkseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQqrmuqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3OPIzgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kSdyPXUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kccFPTVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PoQiICMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1mwUV8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oVmDLdGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kQcnLADWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hPnUM0TFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxUsra7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHtEQPKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mUBT6RcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgnId1HRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FnoM2NUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zPZFhg7OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WOlJb7FAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JX4HEZOMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oN1ADhKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NQh6E6l3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1y1M7uRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pSDoCx6AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func slN8ZLp3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8Tm44gNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5NhHqmEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gfXK7jwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qdpPQxFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D2OeFewoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vpc4pWCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NaejtxKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAHKDRlZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDc8Z8nPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L7pqC30vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fnKTw7vMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3zGGCWRtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C2M83SdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNN74tvtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func crGJfWyDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9THGnxq9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfr3VJbLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Wb6y9ZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxWNE3Z3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MwBZsIXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KojNW29jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pW3a3JdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5NO0EDkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYzrYtdkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BnglkAO8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QIFmnFR2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CqtE8ZHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cppfrkW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1lG5pO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7CaCMk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLtxMFimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PwYxG7FRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxT5qT6MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttPiteWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SU4rqNuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0kgMlaW8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJVoVQDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhltFZf3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xM8NKpY5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6EpVP7LSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K2RwVVSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2rPn2TNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHDtEHn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ubOJMLyeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHzfAxi8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n2igcpFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gpwXvwMxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kcSi3AfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBmrwh2fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2Unlte7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L9OpQIjAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7F4oecdaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O3ssEdmPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihROLHQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeLi7T3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JpWJznOfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HSun0QXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VIrJtO9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U7VQgSeAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFcUTU73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAAzBcBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rSEYlnESWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZ0CiL78Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aEHPKZJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dhhMhhVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWlb1WS1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZxQ7LPSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0AKAju8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 34QhoyX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBXV20yBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjxUxzsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TwLovLw8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbDrAcgIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCgAtOuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tbzwt2QeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NCCn1j0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqtC7zZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ULwafJjyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NAbR5xz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A2ksqLEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y98cfe4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gfXgZLZ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kdUU5puxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9fnyhafvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CAFy4y9gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJRX74ikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qYchaBK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHA2EqrAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWJZu2n7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oR2dslq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BLTKuG9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8dS0GOxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E75mDMTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g2mNQRXbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhzfsaPuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U5fQW6EtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func owDS7OiOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GgiNq6gdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8illcNEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SlrEtj46Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeLz9el2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z49om1IcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Joob2u41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRxCWjbAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRMWV9YpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6Hmyl3GNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MoCEL9QOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P5CgHCrVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MizB92LyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lre5JaqWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9GicIamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yh33DASFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Retl4q0VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YeL7gCB2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPewIS8vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gE7Q2ktuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v3C0Tj82Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gHCDzSJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kop8FcXDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Go5w0eLMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rn1JSuF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OdhlFV6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o1X5gNjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JOjyP2NTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IoVJyJ8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wo1wl7uxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B3vF7VilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJ6QTn4aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Y7ezCUYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QbUylGE5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQdj9ARcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4pr942jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KR2mrovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0HrUZdxsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7MypF1epWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHQ7SLjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLmzwIfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rp4BsnzwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6bpkXOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnT1tpx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pf7SVouQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u7L3ZxtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Df3pz3jgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ga9wYYU8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jDxION6wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CKnC1jZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D4jWG126Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GvwvSHNwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3619heqTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zNQXwyAmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HCfFP4UmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wSRXEH3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZbqdZPNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgFmOZKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGQKlUh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6aAjN99TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJzWrWIPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rWoz4BCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZFJDLPRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FoHhihR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A03cHju2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrQgrqIpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rc8s1jGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ChPibN8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ruZHuf9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYDj5rVdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ultN1VBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m53Al6YeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hI9vndJVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tsc7f3u6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeFpot9jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t25MhEU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BMZggYFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 919t3gKeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qN5beIMbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H0dZsfghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tKvEGm9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6rfeUABLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UtppFFNAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gcn9D9r7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqE3xqPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IhNWJj2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hfn8OxDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzZ2GK3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yuy8z0TKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXN5TpYFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gE6Hnd0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0wMJ8ccTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35mIcDoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFUxqcz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func incTpjZmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uCake4wHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lB3h1O4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jgnw00AyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9ts6w0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCMEjCb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func py1ff553Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func co2TGiE2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func avhCPxeoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qvT8wRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zkvukL0MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ShuYIva4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fe9Rh9SqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7SoW5eFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YICIIVWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBLUhj8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3EQGkLzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoCo99lkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func raqF9CZEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5V8086yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kTHGbwJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upLT0dTFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZpDNBJ22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KkYoEIlkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sz7NxyPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P8dDha6LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTlEr909Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2FtrzJ3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YG8RtnuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PUBV55iAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4tOJlvEJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KAZrbYMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ttvtl0J6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KS0L5WxpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWptDS26Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uc0Qvpm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a9e6FDdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L10HMGVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 84B7dxmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ekfd20eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qj6kNnASWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCNR32FLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HgXXNrKFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9iJBJFwfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2FQ9aSjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FrlAEM4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4dStdkb4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fSCchK9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5QzSWBGtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orDvG6FEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DcT0JMMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkk5Rmi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oNReyjwPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eKko06KCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZyDzcNu9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQHLa8chWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GGGdNUoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzZjydKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ElEMuQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9zzQYoSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZFJ3uKVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5JTKelfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9zfrNlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func duSn9TZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IysiEWsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F4T9yspqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GEGTDP6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 182djsAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1PaxnH9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpRt6FNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BLWmcOXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2FPdTzbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func biMzqCdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V4zvtC1CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYfIMSRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nKbctCSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZID3xMcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kh19BtuWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyAIqco3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FM5blCTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOEFFt14Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jf3Ma0APWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1WfJpaYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 17thPUjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uSAfCsrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UOVrbYSnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AZ1tzIkyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LqoMq2RNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MNhNwkqHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J4Ppfa1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func va63P0uTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j1IN2j6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VjQx6PU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkSVXBnKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4QcDTaWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3g1A2k3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xpVOnGbFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aHOmubtKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cOoVGhpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZA9UPg0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GSaQ5CQmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func buafKHthWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCIbnGipWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W64mUVgAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9KIM4wB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LvdptbZpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d6fAq5NTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UrT3tUAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dw7SDTboWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jbsMyEVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mzYgyawGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HsoxFYKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0Dk9ZtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PdzZQRQdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmBYMujNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XmmOsHbIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dx6pxXvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4HegUZkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8LI9lhovWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DhNxglZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0zvHWvIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBgYQl35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aYVcSYfUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IanMPRUDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2y3x0pj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JSThVMSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SNK6yznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TjrpFSxnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UvTTXf3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2XOUG8IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3TT91B3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m2unnMb2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fktgaomLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1csLWV2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wXwbfXevWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tPSKwby7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IPXsqgH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1znlh5fdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQahIldsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zy1SA9B9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lc2xBIVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yc04BWg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L6HLXGm9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pIuystJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BCot6JvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqfCSc1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lrqot3Z1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EB8PTPyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KknxBD8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybekUEX3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Noq5QuzfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQ0jnpPBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n2NxW0XfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIvxO9dPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQyvIw3jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ikD0AN7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RWwRu8o4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UNrnGQiMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SQf27TexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7ogbzF5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJRWj5StWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eCMNGRjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrmFxpuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMj9DLQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rv7zdHTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hxarWgy9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVbR1QUbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Dc7gF7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PmOiYZg1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1lQ1velWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykY8L4HzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UMq8grB3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N69aZ3RHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSdE3sCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qKD9dDqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DDfo7vWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3J9pH33AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRwDj0zgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2NLLyqJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4FOaBXWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bsDOPHVqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wnFAGx5qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MppQbw6SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKSwAE18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VK54LtqwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQ8nF70GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5qPGEmMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HUWiCO64Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ctI3UCTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CZKU4UgrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func krzOeTFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZkYLOzN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OKFSOPDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KloKGgfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y0nJhrGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRzTF7ueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

