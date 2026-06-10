package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.8" )

func Sw9aBMJGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sRQ0yzh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mgnwcqhqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLtxK3uLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OKK3O9pKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Or8gorZVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3swhZaohWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XkHjLf9QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GBZb4J4zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tVQE7NJoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 231
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VlBZI8nWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4FHiEb7qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nuz92NBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JWEiRbFEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d5BgZEI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func f8ryOxf7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hf0XKnNsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yrrKCpmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TN4vWmEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mEN1dcg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 38
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tczTOJEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2gfZ5wo3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yiPGQ2sxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func imYthrBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FquDEWxQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mwTgH3YNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FYeii1exWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cAPZmPhoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cxc4zP7BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UupT75kBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4DAVJkveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rwgffdAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tW56anPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oQJaAKhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PxNqvbtkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lsXvfe4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V6o64Fd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E9L7Bjt0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wGAlfue3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sztoZmkVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mLf0TvqgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wBL5i0QAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dppnh6UZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8Xjf4o9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tjdT32zgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func riKl3X9cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MHulHGihWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3AvTn3obWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func woJRl6irWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DW5QtEZGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T4IMj81WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7djcPXh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8N01e6INWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PbXcVLnBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oCNgphJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E2TYLrAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y3UqAa9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z2vF6hPdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func twXxY291Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ljL4ox01Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7xXOUzaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NQFue1L3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func du4Sq1skWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2U4RZomiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Kc0SglLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NCJnfZ2uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PmhPQfE1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T43LepuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AttS5cGuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func udDdGsT0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func avOftKLLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 120hnGCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5XnObvskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QRqEUmVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PhZvR7UnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2rziyTOHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FUOOLMyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1BoOmZfWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VybuzPsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 63
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yRMGoLreWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func brruRrW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iZBgu8MDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GCiHJpPTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func okhaRtT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D0jvatbGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LxZ12NhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oy8WiPcEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lf6sMwTaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J7dV4uu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZTESKTjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BoJWqyfrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HtZsAdNSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AntLhAXIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y6vkVQiLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JC93IkGJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3X9GW21JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hjk6Mei4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func npDyeJPsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nLgIqR8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 70
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vfzAlhqNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L3U3QkPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uAvfn2uzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lEy5Wb9wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qYgoAv8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A3u66d1eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CeKNCPxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CtWtwZjVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EAYruJ6OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sgXWFIIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 50
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z46ofrVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fru4y2iUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qokWQ7JwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DOkFFMM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qex5n2ixWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 179
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6w7Pz3tLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func suwstuMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nr8KroANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8dih5C2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ycc1vXriWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 123
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kZ4BmwkjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 182
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KY1I62fkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bBKUb41XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xz1MgFvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 51zZiXKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LRPawa94Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PiRUlkMpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LtS0PS5lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k9tcOOT4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eoumY4M0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2wUyBWKlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5O1PXXO7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 151
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ybp8VkbmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 38bS6auLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 201
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tsOSJq11Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 10OIyMeAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BvzzIHaJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 277
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NB7EmqKLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 72
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oNbFQm0fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func efO03E1oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m26HevAUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5OPboq5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 239
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JifbrsNtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S0gBMNKwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DoNMPpmFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BCmsAZPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rpyky1YOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lT4Gi9tLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vC4Z8q2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ik5wuPJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sDWmngrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func G5yyHfOqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nccSJDv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LHimEM9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fvx4M8fEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 65
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J3ynTHFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bLOSA7FpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OLTlpU3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MCoDE7aZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hfLYsnmiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 158
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VCnwPfxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U4WmVrykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Uxc86ww2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x0Ca3saqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IoyT28FzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gv5fb4vzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GNMYqINtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9MCBVSqFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LBfl6CkkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 213
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cDj51ZPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lyZlIzduWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 205
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nXblW8CBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func opncRJT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cPHGz1FEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oBrzqzl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cuDW5V27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 219
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lF44mi4WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZGCvme21Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wYvdQPOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZZZmFqYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7iX80wpuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PML6JsNbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dti8yuTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 61
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nMSs8ph0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aNvHKrFhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S8B72eycWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WcEghvP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MWwHJaobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JkRE0KLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L8xrWwLhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ZWPBH8WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U8Q51CJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aXSBbZVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1XlRnKUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MdrpQfWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 84wy0wsIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H41nBl2pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jFP2AWnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Y2aAkSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zZrH3YR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CFwz2KnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DFWvi3H1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rQeO7gDJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2OluLjp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VaeSDCOmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ML7ZI7xFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xtDrHRykWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4HR6Zd6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Rb0uVbcfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OXl3Wdq4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n4Oi0TjPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iG92J0VhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IWe1t3N7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sy440QXRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZpetAg15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5NWyj5NtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kBeDBTd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DPQHSyQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lOp2GYDFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1CP9YmD2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cvAvlE4EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SEVFFUqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 163
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8JTwbEMyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ShQ7wKTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j0jDwt3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nwkS4b1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LJnDivc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZsBZmjDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0tbsohwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X8L2SXVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ePv6g3IaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 115
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x0guqiDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HVjqGs5XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PZAlDaCbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UFl4OBs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g8wZ9gGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 61FUPoZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tiRMe2tyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z5gUzhJPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zy9sY6h8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fARG6vvWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dwvZZ1FGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func loTeFBGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h2AyZOxzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4iThompiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T2gxsBtQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XNfKZsz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ro9vGsgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 280
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z03UICFjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pujL8QUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I5s89KssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IXTBsImLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RzGh5kQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e0ioOppoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HBjEx1kXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hqHlrregWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9Z2nRdMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fUfn9xp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nMVjnce9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t3OJxXRYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bBTvZKflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SodZ94KLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5BwcPQPnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0rCW4QKKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PY5IrSxLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b6PAIkEkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dsKWNwxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MaMJ7ZKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hMhznAfzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qp8PGvEiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y2f8OvJlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 261
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8GBHQXAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T15UQneLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4uYkKtEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mjtAU5XmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jiOHOFJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oA48PkM3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M5giLaY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 188
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1PClKwlpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n9is6DeQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wtY9cb2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZKT0QNqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3yOFslHwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func blmoCHU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uOVJtzPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EswtCeskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GJX2v2ENWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L0Jg5JAAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UuermUkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 78
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tgBCIkcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ePZwZfaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nsxS8ewbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X8zzAcUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BWTDwT9dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0bf79lYcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gb5J37tkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jjm50i2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jCeYeWuwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 285
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o7m3s2yDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nwZ9NpCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p771WayTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T75X5vFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4ucqL2f1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mrLRtOOnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6HEzM0UZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z9GmslpkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o0MWEf9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BPn5bxDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1KTvfRTgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 99
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vQqDKtvdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cdWpUwbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 136
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QOqNaLt6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R2SNrB1bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qUyn1pZoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mVij4fK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x3le2RAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EG6tYiuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EAupRs9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TVVjIctMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XncMwUdYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wJ7getsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vNXIKIEpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wc6O8NROWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pz2ViCZ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oIStR8KRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func btCo5KoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QUL9xcqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B0TfECEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DrxBnu8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KXNj1On3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4gjChZQYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cZgtFFkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func puTrLUgQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uq1GDXxAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NnqrPnpVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func frVxVez8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jG9BJehNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4LRAmk2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GL84iKfwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3KBtbeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HaLRijv8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func i7ZJzqOpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2mcbwU2UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dbAh1weSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0smB4Ci9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nHySDY7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sm31hsd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fNr3ncZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1mZwTHGNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Mp5PHrGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8wu1qncRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K8w2PWOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BrnRJgPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 92
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K4WtTc33Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZJU9xpYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HbM4eqeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RcbvLZ6EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 270
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qrD3IZ9XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func H9enqgAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 74sszqayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nnt26JZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I843pPBUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func v0t8qANtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 74
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ekNWCtaKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 146
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Bk8kmXsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 145
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QWewU6ewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ujg6IwfPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jegL0AvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MFNvVPJeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JVchnyItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SAIRNXM0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 257
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AD8hObTxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func di0vN8qdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YK3jMCDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 07X4UAzLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MWxE3ZkqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3Bu5TrR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T53Oa0BVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fNNY7A1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4dfEYMUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5yy9VGUJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BGMz27E4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N6exx0CuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uZCSJ4mIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vxkCTJssWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 89
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8fQT1ghrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y1iJYOb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YQBqq3NDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qROKRL4sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gwf1DV0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CL7r7ht1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L5JJ2mY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VUCo9X89Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func J2V3WwZiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iHqR4e2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fzbZiv1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1tHHYb79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O0bRqMD3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 97
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kCGQz22EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oh1I1aliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func spL30y4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7uT0eBn4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 186
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IUc9pSkQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K5czeK1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1lVtMyuXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XNHaCHIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UAWo4mabWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LGyGwWYEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MVGOYE72Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eBdcv5d0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 235
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func umwVvp6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sRNFhMA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 49
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 24KyYq7NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5ZTX14Z2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UEucpGo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gsZFJlJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uAEnDQoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func A3JhIWJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yn9wz1ccWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nvamOQDeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GmxPavqRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KwFkAF1GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VR2AT4p0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MWn6k46vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 40vJRxdSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OSUhLhyuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g907shQ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xApB7ovjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bZDWoyprWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wvHyZ3ZPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WJAyQpQvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iN7oI2mVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JBvWimVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 234
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UzglxiQXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ck1u97WLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 91
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zAN5BNmWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 110
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FvwVwE9fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dg1z6y6CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Xrcj4OAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VFWBcPRzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 269
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y85bDXmLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xEIQS2mEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 27
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NjushVacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hByjMYnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wb3bvtbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XZQ4encFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ctWt9d2DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OnIhFB87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WjoGjNlkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sJPKXjRdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w5bhLAJbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iezcg3KxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wjYl3fMAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AhYeNbG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3eHsMELMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pLFiGFXdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 81
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C99CAxiNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func astzv2j9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PGNAvQ17Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GCIHALH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8yd9GEhlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 228
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qVjUG8ICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aADeNWRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nDGGjxxoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uvN3sWUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IA22Xl7CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ltLeqW5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w6icjPk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mtW7m27oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2HDPXWHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xMYoJ5qYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QAWCZvaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QXKh0EOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3W0oXL0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yVc2AbHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 25
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zIArb6nOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KeteT2MFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mMGlZ78sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FtkgukPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aQWWcXrPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JMFnlK7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yWwD5EFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P92uvKYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dUVcDwLvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func on4MwIDdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func V9yNaZfyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kSEp32s0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pSRe8rVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nZFS8MBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 178
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 97UFTFmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kwt8YZDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func p7v2TKeyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZjrbDeGGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ALpbY3UXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eJwwm8JnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DzMQz0GOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sb3CyvapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dXJXBkh4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ObasVJwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UjTq0WGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func apLf7o95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BVjQAHGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wBhou18HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ypj4mVF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uBZTQg79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 49k33kdKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 244
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vAPbBYBVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xISWjspdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4M5LgON0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pVVShPW9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hy1vUDs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 209
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sBB1mp4BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 255
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LiatWEaxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UJOqVLHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ernWAdChWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PtdT4f6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 27Fpp29cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 191
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WJh6mpIcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 13
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Y5ZQVpy7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lXG6SbpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2TP1igTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xy3w0EpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rmkmYQShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sNJDkSnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6bX5rDauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZTwVPpU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IG8UHvVVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Z8f2MIzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FnsULoWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mBQWP824Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 71
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 41lLQwFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TrYVAF1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O2oLB5O9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KPmv2zr8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Kr9u6AtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Q17PMnPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 192
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vrAeP9D2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eqgcZMypWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UhvD7hFiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 267
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FaKtSuytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GP21I2WiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 17
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kmCrtlxDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8hxwABxJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qsH3GbLFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WcQ62k3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qLglFBg3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func agbGniBCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zc0JL85yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5VkC4SHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j5z9nlRqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uwW60KDAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OPJlQPwDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tri1LbvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func d4omTo2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GCRrZnU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FYMhVnIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 283
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bBwTVTjZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OfgSn9JqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8ZafyyVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 251
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VQbmqyjbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 51
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PHDNXRSqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 204
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tz9OC490Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r0ouSB1iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NuEVemYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2swWVrd8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hOXqucYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rXAyG7FfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 118
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ihG31TAdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ziGrXgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8cfvm5ZqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 116
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qQ33IzSWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BOiIUJJ9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 292
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uxqYPLcwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 183
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oCYp4khcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ueW8WnyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aT32LRqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6ZqGHByUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 223
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nS4LKXK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5JVwbC0eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IChluJ9CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 57
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func j3Mui9imWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GZUu03cBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b6ghmRlAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lXeZEdldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UootWWUfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P7er49v3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yCJakNYyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KdRlivivWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 32
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dZ3cyedyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xW1KcNClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Gj4g65kzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MCAdDa2dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QiGdMPFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 200
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 47a2Y9PjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QrIFL0aMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wjsBnQsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AMGG9oxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 181
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4dyfaYUlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 119
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KvaTm8gNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0meuBVsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 135
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UFvvhZyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8OdFRHzkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 22
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZGiA2DNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IVWiTQEsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 302HXH7iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b7vTiyqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y2mZ3ckxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 109
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Pr4rnHjTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DCz1QW2VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 206
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func wqSYgp9tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fQFNXisSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4d60YytoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func R0jxmiPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8R9AN9J0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OvYDuRf1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xBh6rDALWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dtgl1FA7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2lLKlKRCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 137
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VOYKgQRMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hkdZIqYRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5OShABKBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 125
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UCSwRiQtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 111
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MbeEP2sWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 142
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qWn4kSAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QBMcIOEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sjkqwkt9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 150
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EDgGvj8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ou0VX3ZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 248
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dzdIMd7NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 273
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8sYfwRZXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PAhKtaz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uffDJKt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 20
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b0Wp4OMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sj1A4UqvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 263
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VSHMPshuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MHjvdo5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C0Ho02lRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 88
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BPxZsnY6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IMiZtGxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 300
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ecHwNkmyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9WHrqMzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 165
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0OpxhGHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 510NtR9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 59
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yF0qlXhxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ynB20d7vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sJuhQ1ZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nTLgq0VqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oO1E83caWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OtrKopoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WhomRynKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dUlB4YWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 29
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YQJdANtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

