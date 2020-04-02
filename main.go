package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "math/rand"
    "os/exec"
    "time"
    "gopkg.in/yaml.v2"
    "flag"
)

var filename = flag.String("config", "conf.yaml", "Location of the config file")

func main() {

    var c conf
    c.getConf()

    files, err := ioutil.ReadDir(c.Dir)
    if err != nil {
        log.Fatal(err)
    }

    num := numberPicker(len(files))

    count := 1
    for _, f := range files {
        if count == num {
            filename := fmt.Sprintf("%s/%s", c.Dir, f.Name())
            cmd := exec.Command("/usr/bin/feh", "--bg-fill", filename)
            _, err := cmd.Output()
            if err != nil {
                log.Fatal(err)
            }
            break
        }
        count += 1
    }
}

func numberPicker(max int) int {
    rand.Seed(time.Now().UnixNano())
    num := rand.Intn(max) + 1
    return num
}

type conf struct {
    Dir string `yaml:"dir"`
}

func (c *conf) getConf() *conf {
    flag.Parse()
    log.Println(*filename)
    yamlFile, err := ioutil.ReadFile(*filename)
    if err !=  nil {
        log.Fatal(err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatal(err)
    }

    return c
}
