package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.1" )

func qDgFFHADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wq0vrHEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L3aOrMFHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fu8xvdtrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jAGcloLSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbDTqzMlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TqKNMEI4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDXXZHrdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B4l4bHaLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func enjpaWAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P933K9x3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PeAvbHsuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Awsmaf7jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GM2FpB40Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trU4aB5nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4ZyG0djWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HMjxKYJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AFCgfsmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhkWZsnrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2UHWeBJgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uDMnaOfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3EY2FSwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VwFgpEKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fDtpKGHOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeLirtdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func di8OBsJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func teDbsm7LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nZwHuELBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fN09H3GDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SuCQdhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4kP35bFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqI3SBpZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InYKrtGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4GzeTg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8oSDK5U4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5HK30YBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fPGR9g9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQgTKN3rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ViN2BfMVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2drgiekWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXbG5X2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqkszbaTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gRriQgU4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XdWDRWOaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dFrtHhyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J0MhfvqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HTX0VcxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cncotdfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j6SapzFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lkvLVQigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDjks7vVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IvUw3PEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fe3rvMs0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z18tGNH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Yg7nw1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTdfBdKOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tOBi009WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VId0TtmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l17NaOYNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9gvn5xjMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IH2d0g2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4f2ShcCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrhiUCooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IuslYio8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3RB858SJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1r67jrWBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fWZwCYE0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HvpORMoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6aOVZva1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3FfQG3LIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NsHGIljpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4bzZRedxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9zFi1AoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ymiw9DbMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vp2V4AYmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fk0CAaFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eVZKzy3vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bXZ5CpuiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHXYWckBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EYhzBm4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7ccF3yoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GESabL0GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F97uhNSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VsXmaqRHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vgyA9F4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6oQTo9z4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uOf18SepWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1fErNa1dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZoG87UDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j7xikIrKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CGqqNJKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yaMUlQ7rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WtNrU5T5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p5K4Hx7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s54fwbDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r7QNLyRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3CaYz1KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kcr1JEnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rLljPW8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mHVOWslZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xdrNTxp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZtwMmcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gxWZCHkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Wi4KfeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rnsbbqAIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0uipc1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLrRbC13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WbULkW4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pxmwzFVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITh2MWoQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4IxYFjvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DTMaS6C4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 27ssgltUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D7VB3sIkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oq93YOGSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IUoVkWTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func szk6cDFdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X1cq1KQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86Cw2TxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2uBVowDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7uMJY5qJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HO0pf3ILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2sHcGQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 36IDkRkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MC0QweqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Byd5zy7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Ldy7kK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ayC5pSnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GwZb4a5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPliy7uWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26bYAAecWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func puSBtoLTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NtVZp5r1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kz0u8GUUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pssJMdLtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7q8292SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YcGrWpKkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYigtAtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6BOdEzxdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zd8cgyAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3Hj2ROPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func izktjISuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TuHiIEKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8af012pYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VWLmSxZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gt6inn8TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2wIZ7ezWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PFmDVlVfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func frQwIUtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWiFUAqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hStH8PLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MdKAiWuIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXCBgQx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSInbr7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzNuJ89MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kndie60jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZcv5ZqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVIKgPZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdSV0WyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oOlIgkZFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6jVoZ3v4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aDkD55vLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K1E07uhaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func woFdOvBNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lRva6bqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UhSYWFvqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52JbQWCnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlcbUb6PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzvsxZ5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yCWeFIPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QTEru4NDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EaxCwO1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EtaWVyaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4xnitFZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hr1vwbReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wsx22dmlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVHOyZ9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VD8rvyNLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YKsvVt6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GCC4pjZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Xr6lYA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZtmdKpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nMOX0wUwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qP78YVQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jZ2d5TAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a4SjxLILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0g7CdSWFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HqLA5jvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wHhwFQvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jt1IuAtpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NWl5627kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wI7joRmcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrhxX4MSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BZAFs7jFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ueozElEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBUTPNZtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VLofnZSxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ok9pevrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S1qWFT2vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKmf8WpoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AOLIdlloWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHyqdTXVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2YGPINdBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XVhu9U2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b3r9PrgwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvVtAV0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VEobJBPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIVTjwHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNfySS4YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3wXX8BjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GEaprl3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DXpZv0bOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 92JQ3wcNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tCMvk4klWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RE0jJiigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qQbBopJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JX4X5fyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6VAYLIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPP9FjNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O4bMRU8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FQFoRt13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YKxfxVsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1GXFD88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8NPXPwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3XKtpi1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDg2pDgRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DsAlb4iZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iDjN0ewJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IMTvRddIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gwz93F0mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4XWikxvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uBEn1KaSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tSz1DuKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DoxPSrDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func srpJzSo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aKk3uX3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2dhRkhvlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WTwItuo7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k2zigcEhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N74KowDaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sDC7pldHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjWXKq3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lSd40W7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6JnXvYiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sXozqiifWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lvTZPgAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qi9aywH7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lgoyJfx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kCfzti0DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uYeF5iBbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XKyZU4waWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WrFQ1xLDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lWB0FFduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rMW1dZEeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CBKIAvtEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ti1XeW8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4wt13LCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43RKqqMjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6zLMB1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQ6GwnmqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AAW9nrjSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFwMFxVDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a4c2wk8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGE1dT4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mulaLPU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0bZlV7cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yc8uMq1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HcVUnw1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WzpX10mfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zac7r45FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gKVuuQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func atKkCX6eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4K1lJHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Idty4wsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5gzQmOZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ru6lBMmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHTiSIU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrMzaGCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZvJrLXLoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jySYzfJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 60KNyZOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQsYlUcBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k6c53V4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLEo0n7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRTvwlKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQZl9K8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HaDH7ppTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hqb6EOT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCcKC9QKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EuMvG1t3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OK5iSuTeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HRQF2G3EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJVr0bMXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GsGFR1cqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 17HRA2kuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TT1Zw446Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OaiDCXRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zBhmK3LjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QWLsUgzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DINEnLXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func luSJg5jTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VmQWTZqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lN5z0U8CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mJfoTDtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2OOJ5HZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X3nYf3rzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6HDZ2qihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7GjxMyomWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I7w4ir5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func awv2DIznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BcKngq9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n8ZjOREUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKxkGehFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F1oQcvrhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOkYbswCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQ8fUu9qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sbL76OoeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XolKHtZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dqR0HkkhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TNmQ2VwjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZbTTnHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGhZPx3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVmz660eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5bHaZgHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bgO3yiiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UafumgwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 32FD3WPCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OPRysvT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wg46InjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XANN7g2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jYRE6QAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUrMFQTMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B4VtqdmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hi8kBYqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zd8tQpcnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vi5UELurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EMZDMUUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0WnUHlVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YCh7hz7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LlMiH0eQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVOqY1eGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKBVhREnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hy3roEsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIxjrU0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvFGs5CaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKQLHFOYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbfK1O7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2SqIATbOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmiQa1NLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxOw2zfRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6aDv83MiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIq9n8SSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TlDs0WfmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WF4hhFS4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oPBMTx0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yiomi20hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KQ1hzOz9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDd9DzCRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qMEn3DeLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ImOMwnYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oeCcNpVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 47fek1GSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNNCTUlBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2FDZmmSbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nijyFU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eewEVqUSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnL37WDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8WWxeytwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vynUYxWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6AqEqRKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67VyT5uVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PL4iyHQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 94x4FhFqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoxkkhwiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PnZpd87FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPHrrDcuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kAMQlfvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9LZyWUowWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHm6Ngb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func or8c1bSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YdeXQNGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HDPagLXKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Ukv7qIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Aflk111Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U1QeKP6cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NAq2lhffWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBWpnQb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQQiFn6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2AiL5ehEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIVhGTBGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8HrcSKRNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GUQJrGiVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cZNFy6jtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KUyFgm3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GL4DcKt3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qFhR7ZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXdSRmx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4AktOmr3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUAmPVRJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ppIsFkJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWIUpqh7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 068xVSA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CES0HGt8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0pfowfIGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7pA51oiZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ORzUKpsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNkTm0UaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mvwr2sWXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y894ZfNnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8gO7NLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pF26V6LVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PDzXGEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bYwTbFvSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NLQgQlWxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JxfsNeWzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfWjuAZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gXfGDwsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eneFWgXkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpnioZi3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aTlkXxW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vpoU2rPqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func srNMDbqxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHSVyfBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4nQ7wMZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YxBc4LjBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XsDQxZ85Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rrtAbVj2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ILWxQe20Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8YXnC2oyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C1R6JU4vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bG8GhraDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LbJJUGycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1KmdhNumWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lsKTfWlwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wSyKnMJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JIubFl7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vynFg9MBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hL3CX3grWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUHnZPDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITX7lDzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mssYIF15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3mN6OfSUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t5g5SywxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Ro6AtewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vsxMPxHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fi6QBFvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZHhBTOVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vBWZLKqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bTdaj37pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWyoceGHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FdkfccUOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ht6ik7z8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AoktlPZLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YrPr2N87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6NJkqbYzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func czsgiCFZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func upSllKAOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJtUYiRIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJwnEQaoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E1qYaheaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ajQoXZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TFxVGRSsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ppQkp65Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mImKrYxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qwJWjiiBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Idbd2futWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RpcEbI1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func byZvJwp1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Krd5VOztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78rxftDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0kXLQ0fkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ucVlIPzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3aI8NFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KR7cAzzmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CxRvrGjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ikeanf8PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func taCxUQx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sCwnxZ9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1VK7LFxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ptER3zQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQlU8TxBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func npQHyjZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iwemWqtCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CyuwTzKgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9nNR1zzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kfOZ9GgzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C7WhY5m8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oRiuFmvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDI5cENvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBnMXwldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yded7jaBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s2rxK48jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4uPTW68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4jehDqUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yrAfYVQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLIE5kkZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func meIKkc4LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rhRMGMFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rKqxrNeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6wD6fQVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8R639u4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yfde1LRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWw8t0xUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cj7XaJFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tFXV5uJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RP1R1jrYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ee3vWdlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nPIN19PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJpojJgGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IKiYPSBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z5LB2Mz2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7o7vBbddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mPx3V4MeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SR1XQ03YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OqoVga8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9mY8ZAFgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YD8u55lWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NaSol7rRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4rSG6Ne1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrzmM0VXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jz742gEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VJedJivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Or8JUpBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iME2xTC2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yqBoeCAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gWQGDoSZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func anUDjTaAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLRSUTpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IzOJO9ELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hzWA4QNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2w4IYYlJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6vixw2LIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eRWJ5pxIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6ujfFvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APnjAuVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zTsh4LQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zD6AeJ29Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8jsVsnb8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGf1B8ZAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6r3Mtlh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D0a9JD5EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lnCJllFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hdmcf0HPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJJWFAlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HrSNhkv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3gNKDnr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func utJWWKLBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func erwk1phGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mHIUWRl5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FfwPmrZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5AM3ebNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fsAo7Bg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kv9hE2OtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVW9YzG4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func njuIL6DtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WViAO1idWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QRLu8BNlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ip2QKIgeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I4hTnE99Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3r1c4odmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBv2EUlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OZ2NM1g2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBnj55kIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIzxwaubWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVLhGWpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mgdDH445Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8axyZwTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHiPAOIFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dyCFmVRsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZRgHRpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VuPa33QPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LvRSb1zsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XLSBr9mDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pz64LO2CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5cj89JpSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KcmtD7lRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pPm2e4djWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9sxMku6DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IwEXe1VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Su59QfMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H9U7guWZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rP3FiiWiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9GMLNooVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 746sQ8zuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMUnvD9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fx3qYYmSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OwNx36lRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dl4oX6gcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WXXxSJVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmwTiu1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5dVgfcNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eeim3EssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJ5xSvLaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jnHRTVEPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tC8InBuzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gROmE0dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVzLniUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kBhjvjjuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZ0wyGavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QgZdCpErWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzaIWSDHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQfGvypuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MQzGG9JNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ZjHiHmGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AELVMoWqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jo8QqBEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func veBsP20WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OJldl4lxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fHvlfzjXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kYwT5Q43Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K0vRuD9UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kf9TyF7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func obLjNGCsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RWAoJq7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dmkv2pmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j0qXKMFrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UcBLLn09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M74yyJnbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9K4wawvdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vBLRr5YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w2l8lzxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwVY0gYAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BA4eYX7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I9DKIaA4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XfnRBWgoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oqqb4LsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zx4DoSGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I2I6hlmgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OFBy8616Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nlY2q5PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uqSlYD0ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bb5kAywKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xMhPdDvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8WF0Pfw0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zls0V5h0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0rDeqkC9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func efyEbTvWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pJTZUGJZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 23Y1AMKNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8JeG4cBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a6NKY4rIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 553dYWcCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEMaghsiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M8AMFIVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ocYop8eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e9pWFveDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uMpErIoMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBCsPsA1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2BAxshGOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQ0Bto1uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jeFAZlDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fuaVuTNxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48PciLI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v145WyGdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u0L4TcymWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZ8d6mw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QF4mus5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mN5UaMneWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dsgqG2YWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xgc5lEnsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYiBPhTTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebJzsJ93Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func utHokvxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iqaUcE2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xhFnxEDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2I1rJrD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sjqWgCOgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9jcZRpjPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ItznQA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lr3E2rwAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M9BR17TdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNGC5G73Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8eiLxf58Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZSTWFJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsdYCL2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0HrlKENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nhcIS34WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U0tpoouWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8AcgVT2rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gp1u2n6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bu9fADTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSKGzdBlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mBCFaTTRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M4mOlyEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func leTutw6IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zLr81VbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YBTSUVbZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iZxay79uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G2EAa87gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sc3bQjx4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vca4fhp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N2UWGwgXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ToXMQgLEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1iv7LiEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CURzsVsWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9c1GA0MtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HLYzSpjKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QoDEJ12jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 32NiIr6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g2C1qH28Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qNIkuNjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qB5cDoq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1M7B6ftIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OMTvf5CsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jj8felkJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qgUX2n4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EJzS8PWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mgb5zltpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1dHWkLzJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FYAiUwKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWNkO1AOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4u8DrfKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKjBMbNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRFPtvHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B28LJOt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nijYkuUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hq7w41ZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RzlNhGskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BLKDiCh8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J3UlGpNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NHeoKurZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RmrsXnGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFvsEJngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mLsBeLG0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41uPHsikWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rnr2NlzXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nc8RcMCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1y3CD6d3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

