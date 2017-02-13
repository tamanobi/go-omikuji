package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"

    "github.com/urfave/cli"
)

func suffle(arr []string) {
    rand.Seed(time.Now().UnixNano())
    n := len(arr)
    for i := n - 1; i > 0; i-- {
        j := rand.Intn(i + 1)
        arr[i], arr[j] = arr[j], arr[i]
    }
}

func main() {
    app := cli.NewApp()
    app.Name = "omikuji"
    app.Usage = "Get today's fortune!"
    app.Version = "1.0"

    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name:  "level",
            Value: "normal",
            Usage: "Fortune level, it is hard, easy and normal",
        },
    }

    app.Action = func(c *cli.Context) error {
        fortunes := []string{
            "大吉",
            "吉",
            "中吉",
            "笑吉",
            "末吉",
            "凶",
            "大凶",
        }

        if c.String("level") == "hard" {
            fortunes = []string{
                "大吉",
                "大凶",
            }
        } else if c.String("level") == "easy" {
            fortunes = []string{
                "大吉",
                "吉",
                "中吉",
                "笑吉",
                "末吉",
            }
        }

        suffle(fortunes)
        fmt.Println(fortunes[0])

        return nil
    }

    app.Run(os.Args)
}
